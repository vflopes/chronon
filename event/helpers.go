package event

import (
	"log"

	chrononpb "github.com/vflopes/chronon/gen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func New(seq int64, ts *timestamppb.Timestamp, m proto.Message) *chrononpb.Event {

	payload, err := anypb.New(m)

	if err != nil {
		log.Fatalln("Error while creating Any proto message from proto message", err)
	}

	if ts == nil {
		ts = timestamppb.Now()
	}

	return &chrononpb.Event{
		Sequence:  seq,
		Timestamp: ts,
		Payload:   payload,
	}

}
