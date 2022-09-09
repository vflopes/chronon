package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	chrononpb "github.com/vflopes/chronon/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	VersionMismatchErr = status.Error(codes.Aborted, "Event not inserted 'cause specified sequence is out of sync")
)

type eventStore struct {
	chrononpb.UnimplementedEventStoreServer

	database *sql.DB

	scanBatchSize int
	eventsTable   string

	sourceLocker *sourceLocker

	appendStmt *sql.Stmt

	scanExclusiveToLatestStmt *sql.Stmt
	scanInclusiveToLatestStmt *sql.Stmt

	scanExclusiveToExactOrPreviousStmt *sql.Stmt
	scanInclusiveToExactOrPreviousStmt *sql.Stmt

	seekTimestampStmt *sql.Stmt
}

func newEventStore(database *sql.DB, scanBatchSize int, eventsTable string) *eventStore {

	e := &eventStore{
		database:      database,
		scanBatchSize: scanBatchSize,
		eventsTable:   eventsTable,
		sourceLocker:  newSourceLocker(database),
	}

	if err := e.prepareStatements(); err != nil {
		log.Fatalln("prepareStatements() error:", err)
	}

	return e

}

func (e *eventStore) prepareStatements() error {

	var sb strings.Builder

	sb.WriteString("INSERT INTO ")
	sb.WriteString(e.eventsTable)
	sb.WriteString(" (source_key, source_sequence, emit_timestamp, payload_type_url, payload_value) SELECT $1, $2, $3, $4, $5 WHERE COALESCE((SELECT MAX(ev.source_sequence) FROM ")
	sb.WriteString(e.eventsTable)
	sb.WriteString(" AS ev WHERE ev.source_key = $1), 0) + 1 = $2")

	st, err := e.database.Prepare(sb.String())

	if err != nil {
		_, err = fmt.Println("Prepare appendStmt statement error:", err)
		return err
	}

	e.appendStmt = st

	var selectSb, orderBySb strings.Builder

	selectSb.WriteString("SELECT ev.source_sequence, ev.emit_timestamp, ev.payload_type_url, ev.payload_value FROM ")
	selectSb.WriteString(e.eventsTable)
	selectSb.WriteString(" AS ev WHERE ev.source_key = $1 AND ev.source_sequence ")

	orderBySb.WriteString(" ORDER BY ev.source_sequence ASC LIMIT ")
	orderBySb.WriteString(strconv.Itoa(e.scanBatchSize))

	sb.Reset()

	sb.WriteString(selectSb.String())
	sb.WriteString(" > $2 ")
	sb.WriteString(orderBySb.String())

	st, err = e.database.Prepare(sb.String())

	if err != nil {
		_, err = fmt.Println("Prepare scanExclusiveToLatestStmt statement error:", err)
		return err
	}

	e.scanExclusiveToLatestStmt = st

	sb.Reset()

	sb.WriteString(selectSb.String())
	sb.WriteString(" >= $2 ")
	sb.WriteString(orderBySb.String())

	st, err = e.database.Prepare(sb.String())

	if err != nil {
		_, err = fmt.Println("Prepare scanInclusiveToLatestStmt statement error:", err)
		return err
	}

	e.scanInclusiveToLatestStmt = st

	sb.Reset()

	sb.WriteString(selectSb.String())
	sb.WriteString(" > $2 AND ev.source_sequence <= $3 ")
	sb.WriteString(orderBySb.String())

	st, err = e.database.Prepare(sb.String())

	if err != nil {
		_, err = fmt.Println("Prepare scanExclusiveToExactOrPreviousStmt statement error:", err)
		return err
	}

	e.scanExclusiveToExactOrPreviousStmt = st

	sb.Reset()

	sb.WriteString(selectSb.String())
	sb.WriteString(" >= $2 AND ev.source_sequence <= $3 ")
	sb.WriteString(orderBySb.String())

	st, err = e.database.Prepare(sb.String())

	if err != nil {
		_, err = fmt.Println("Prepare scanInclusiveToExactOrPreviousStmt statement error:", err)
		return err
	}

	e.scanInclusiveToExactOrPreviousStmt = st

	sb.Reset()

	sb.WriteString("WITH max_seq AS (SELECT MAX(ev.source_sequence) AS source_sequence, MAX(ev.emit_timestamp) AS emit_timestamp FROM ")
	sb.WriteString(e.eventsTable)
	sb.WriteString(" AS ev WHERE ev.source_key = $1) SELECT ev.source_sequence, ev.emit_timestamp, max_seq.source_sequence, max_seq.emit_timestamp FROM ")
	sb.WriteString(e.eventsTable)
	sb.WriteString(" AS ev CROSS JOIN max_seq WHERE ev.source_key = $1 AND ev.emit_timestamp <= $2 ORDER BY ev.source_sequence DESC LIMIT 1")

	st, err = e.database.Prepare(sb.String())

	if err != nil {
		_, err = fmt.Println("Prepare seekTimestampStmt statement error:", err)
		return err
	}

	e.seekTimestampStmt = st

	return nil

}

func (e *eventStore) Append(ctx context.Context, request *chrononpb.AppendRequest) (*chrononpb.AppendResponse, error) {

	response := &chrononpb.AppendResponse{
		AppendedCount: 0,
	}

	numKey := uuidToInt64(request.Key)

	tx, err := e.sourceLocker.lock(ctx, numKey)

	switch {

	case err == LockFailErr:
		return nil, status.Errorf(codes.Aborted, "Aborted append operation: %v", err)

	case err != nil:
		return nil, status.Errorf(codes.Internal, "Internal store error: %v", err)

	}

	defer tx.Rollback()

	appendStmt := tx.Stmt(e.appendStmt)

	for _, event := range request.Events {

		result, err := appendStmt.ExecContext(
			ctx,
			request.Key,
			event.Sequence,
			event.Timestamp.AsTime(),
			event.Payload.TypeUrl,
			event.Payload.Value,
		)

		if err != nil {
			return nil, status.Errorf(codes.Internal, "appendStmt exec error: %v", err)
		}

		rowsAffected, err := result.RowsAffected()

		if err != nil {
			return nil, status.Errorf(codes.Internal, "RowsAffected() error: %v", err)
		}

		if rowsAffected == 0 {
			return nil, VersionMismatchErr
		}

		response.AppendedCount++

		response.LatestSequence = event.Sequence
		response.LatestTimestamp = event.Timestamp

	}

	err = tx.Commit()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Append transaction commit error: %v", err)
	}

	return response, nil

}

func (e *eventStore) Scan(request *chrononpb.ScanRequest, stream chrononpb.EventStore_ScanServer) error {

	exclusiveStmt := e.scanExclusiveToLatestStmt
	inclusiveStmt := e.scanInclusiveToLatestStmt

	args := []any{request.Key, request.StartSequence}

	if request.StopWhen == chrononpb.StopWhen_STOP_WHEN_EXACT_OR_PREVIOUS {

		exclusiveStmt = e.scanExclusiveToExactOrPreviousStmt
		inclusiveStmt = e.scanInclusiveToExactOrPreviousStmt

		args = append(args, request.StopSequence)

	}

	ctx := stream.Context()

	rows, err := inclusiveStmt.QueryContext(ctx, args...)

	var ev *chrononpb.Event

	emitTimestamp := time.Time{}

	for {

		if err != nil {
			return status.Errorf(codes.Internal, "inclusive/execlusiveStmt QueryContext: %v", err)
		}

		sentCount := 0

		for rows.Next() {

			ev = &chrononpb.Event{
				Payload: &anypb.Any{},
			}

			if err := rows.Scan(&ev.Sequence, &emitTimestamp, &ev.Payload.TypeUrl, &ev.Payload.Value); err != nil {
				rows.Close()
				return status.Errorf(codes.Internal, "Event row scan error: %v", err)
			}

			ev.Timestamp = timestamppb.New(emitTimestamp)

			if err := stream.Send(ev); err != nil {
				rows.Close()
				return status.Errorf(codes.Internal, "Failed to send event through grpc server stream: %v", err)
			}

			sentCount++

		}

		if err := rows.Close(); err != nil {
			return status.Errorf(codes.Internal, "rows.Close() error: %v", err)
		}

		if sentCount < e.scanBatchSize {
			break
		}

		args[1] = ev.Sequence

		rows, err = exclusiveStmt.QueryContext(ctx, args...)

	}

	return nil
}

func (e *eventStore) SeekTimestamp(ctx context.Context, request *chrononpb.SeekTimestampRequest) (*chrononpb.SeekTimestampResponse, error) {

	timestamp := time.Now()

	if request.ExactOrPreviousTimestamp != nil && request.ExactOrPreviousTimestamp.IsValid() {
		timestamp = request.ExactOrPreviousTimestamp.AsTime()
	}

	response := &chrononpb.SeekTimestampResponse{}

	var foundTimestamp, latestTimestamp time.Time

	err := e.seekTimestampStmt.QueryRowContext(ctx, request.Key, timestamp).Scan(
		&response.FoundSequence,
		&foundTimestamp,
		&response.LatestSequence,
		&latestTimestamp,
	)

	switch {

	case err == sql.ErrNoRows:
		return nil, status.Error(codes.NotFound, "No source found for this source key.")

	case err != nil:
		return nil, status.Errorf(codes.Internal, "Internal error: %v", err)

	}

	response.FoundTimestamp = timestamppb.New(foundTimestamp)
	response.LatestTimestamp = timestamppb.New(latestTimestamp)

	return response, nil

}
