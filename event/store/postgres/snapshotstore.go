package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	chrononpb "github.com/vflopes/chronon/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type snapshotStore struct {
	chrononpb.UnimplementedSnapshotStoreServer

	database       *sql.DB
	snapshotsTable string

	saveStmt *sql.Stmt

	getLatestStmt *sql.Stmt

	getExactOrPreviousStmt *sql.Stmt
}

func newSnapshotStore(database *sql.DB, snapshotsTable string) *snapshotStore {

	s := &snapshotStore{
		database:       database,
		snapshotsTable: snapshotsTable,
	}

	if err := s.prepareStatements(); err != nil {
		log.Fatalf("prepareStatements() error: %v", err)
	}

	return s

}

func (s *snapshotStore) prepareStatements() error {

	var sb strings.Builder

	sb.WriteString("INSERT INTO ")
	sb.WriteString(s.snapshotsTable)
	sb.WriteString(" (source_key, source_sequence, emit_timestamp, payload_type_url, payload_value) VALUES ($1, $2, $3, $4, $5) ON CONFLICT (source_key, source_sequence) DO UPDATE SET emit_timestamp = $3, payload_type_url = $4, payload_value = $5")

	st, err := s.database.Prepare(sb.String())

	if err != nil {
		return fmt.Errorf("Prepare saveStmt statement error: %v", err)
	}

	s.saveStmt = st

	var selectSb, orderBySb strings.Builder

	selectSb.WriteString("SELECT sn.source_sequence, sn.emit_timestamp, sn.payload_type_url, sn.payload_value FROM ")
	selectSb.WriteString(s.snapshotsTable)
	selectSb.WriteString(" AS sn WHERE sn.source_key = $1 ")

	orderBySb.WriteString(" ORDER BY sn.source_sequence DESC LIMIT 1")

	sb.Reset()

	sb.WriteString(selectSb.String())
	sb.WriteString(orderBySb.String())

	st, err = s.database.Prepare(sb.String())

	if err != nil {
		return fmt.Errorf("Prepare getLatestStmt statement error: %v", err)
	}

	s.getLatestStmt = st

	sb.Reset()

	sb.WriteString(selectSb.String())
	sb.WriteString(" AND sn.source_sequence <= $2 ")
	sb.WriteString(orderBySb.String())

	st, err = s.database.Prepare(sb.String())

	if err != nil {
		return fmt.Errorf("Prepare getExactOrPreviousStmt statement error: %v", err)
	}

	s.getExactOrPreviousStmt = st

	return nil

}

func (s *snapshotStore) Save(ctx context.Context, snapshot *chrononpb.Snapshot) (*chrononpb.SaveResponse, error) {

	_, err := s.saveStmt.ExecContext(
		ctx,
		snapshot.Key,
		snapshot.Sequence,
		snapshot.Timestamp.AsTime(),
		snapshot.Payload.TypeUrl,
		snapshot.Payload.Value,
	)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error on saveStmt: %v", err)
	}

	response := &chrononpb.SaveResponse{}

	return response, nil
}

func (s *snapshotStore) Get(ctx context.Context, request *chrononpb.GetRequest) (*chrononpb.Snapshot, error) {

	getStmt := s.getLatestStmt
	args := []any{request.Key}

	if request.StopWhen == chrononpb.StopWhen_STOP_WHEN_EXACT_OR_PREVIOUS {

		getStmt = s.getExactOrPreviousStmt

		args = append(args, request.StopSequence)

	}

	emitTimestamp := time.Time{}

	snapshot := &chrononpb.Snapshot{
		Payload: &anypb.Any{},
	}

	err := getStmt.QueryRowContext(ctx, args...).
		Scan(&snapshot.Sequence, &emitTimestamp, &snapshot.Payload.TypeUrl, &snapshot.Payload.Value)

	switch {

	case err == sql.ErrNoRows:
		return nil, status.Error(codes.NotFound, "No snapshot found for this source key.")

	case err != nil:
		return nil, status.Errorf(codes.Internal, "Internal error: %v", err)

	}

	snapshot.Timestamp = timestamppb.New(emitTimestamp)

	return snapshot, nil
}
