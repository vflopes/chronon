package event

import (
	"context"
	"io"
	"log"
	"sync"

	chrononpb "github.com/vflopes/chronon/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Reducer interface {
	proto.Message

	Clone() Reducer
	Reset()
	Apply(*chrononpb.Event) error
}

type Source[R Reducer] struct {
	key     string
	reducer R

	eventStore    chrononpb.EventStoreClient
	snapshotStore chrononpb.SnapshotStoreClient

	currentSeq int64

	lock sync.RWMutex
}

func NewSource[R Reducer](key string, reducer R, eventStore chrononpb.EventStoreClient, snapshotStore chrononpb.SnapshotStoreClient) *Source[R] {
	return &Source[R]{
		key:           key,
		reducer:       reducer,
		currentSeq:    0,
		eventStore:    eventStore,
		snapshotStore: snapshotStore,
	}
}

func (s *Source[R]) CurrentState() R {
	return s.reducer.Clone().(R)
}

func (s *Source[R]) restoreFromSnapshot(ctx context.Context, stopSeq int64, stopWhen chrononpb.StopWhen) error {

	getSnapshot, err := s.snapshotStore.Get(ctx, &chrononpb.GetRequest{
		Key:          s.key,
		StopSequence: stopSeq,
		StopWhen:     stopWhen,
	})

	if err != nil {

		if status.Code(err) == codes.NotFound {
			return nil
		}

		return err
	}

	if err := getSnapshot.Payload.UnmarshalTo(s.reducer); err != nil {
		log.Fatalln("UnmarshalTo reducer failed on restore from snapshot:", err)
	}

	return nil

}

func (s *Source[R]) restoreFromEventReplay(ctx context.Context, stopSeq int64, stopWhen chrononpb.StopWhen) error {

	scanClient, err := s.eventStore.Scan(ctx, &chrononpb.ScanRequest{
		Key:           s.key,
		StartSequence: s.currentSeq + 1,
	})

	if err != nil {
		return err
	}

	ev := &chrononpb.Event{}

	for {

		if err := scanClient.RecvMsg(ev); err != nil {

			if err == io.EOF {
				return nil
			}

			return status.Errorf(codes.Internal, "Error on recvmsg from event store scan: %v", err)
		}

		ev.Key = s.key

		if err := s.reducer.Apply(ev); err != nil {
			return status.Errorf(codes.Internal, "Reducer.Apply() error on restoreFromEventReplay operation: %v", err)
		}

		s.currentSeq = ev.Sequence

	}

}

func (s *Source[R]) Restore(ctx context.Context, stopSeq int64, stopWhen chrononpb.StopWhen) (int64, error) {

	s.lock.Lock()

	defer s.lock.Unlock()

	switch {

	case stopWhen == chrononpb.StopWhen_STOP_WHEN_LATEST || stopSeq > s.currentSeq:

		if s.snapshotStore != nil {

			if err := s.restoreFromSnapshot(ctx, stopSeq, stopWhen); err != nil {
				return 0, err
			}

		}

		if stopWhen == chrononpb.StopWhen_STOP_WHEN_LATEST || stopSeq > s.currentSeq {
			if err := s.restoreFromEventReplay(ctx, stopSeq, stopWhen); err != nil {
				return 0, err
			}

		}

		return s.currentSeq, nil

	case stopSeq <= s.currentSeq:

		return s.currentSeq, nil

	}

	return 0, status.Error(codes.Internal, "This condition must never be met")

}

func (s *Source[R]) Append(ctx context.Context, events ...*chrononpb.Event) error {

	s.lock.Lock()

	defer s.lock.Unlock()

	for _, ev := range events {

		ev.Key = s.key

		if err := s.reducer.Apply(ev); err != nil {
			return status.Errorf(codes.Internal, "Reducer.Apply() error on append operation: %v", err)
		}

	}

	appendResponse, err := s.eventStore.Append(ctx, &chrononpb.AppendRequest{
		Key:    s.key,
		Events: events,
	})

	if err != nil {
		return err
	}

	s.currentSeq = appendResponse.LatestSequence

	return nil

}

func (s *Source[R]) SaveSnapshot(ctx context.Context) error {

	if s.snapshotStore == nil {
		return nil
	}

	s.lock.RLock()

	payload, err := anypb.New(s.reducer)

	if err != nil {
		log.Fatalln("Failed to create Any proto message from reducer:", err)
	}

	snapshot := &chrononpb.Snapshot{
		Key:       s.key,
		Sequence:  s.currentSeq,
		Timestamp: timestamppb.Now(),
		Payload:   payload,
	}

	s.lock.RUnlock()

	_, err = s.snapshotStore.Save(ctx, snapshot)

	return err

}
