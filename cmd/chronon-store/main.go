package main

import (
	"log"
	"net"

	_ "github.com/lib/pq"

	"github.com/vflopes/chronon/event/store/postgres"
	chrononpb "github.com/vflopes/chronon/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var config = &chrononpb.ServerConfiguration{
	Address: "localhost:53351",
	Name:    "Chronon Server",
	EventStore: &chrononpb.ServerConfiguration_PostgresEvents{
		PostgresEvents: &chrononpb.ServerConfiguration_Postgres{
			ConnectionString: "postgres://postgres:example@192.168.0.101:5432/chronon?sslmode=disable",
		},
	},
	SnapshotStore: &chrononpb.ServerConfiguration_PostgresSnapshots{
		PostgresSnapshots: &chrononpb.ServerConfiguration_Postgres{
			ConnectionString: "postgres://postgres:example@192.168.0.101:5432/chronon?sslmode=disable",
		},
	},
}

func main() {

	server := grpc.NewServer()

	switch store := config.EventStore.(type) {

	case *chrononpb.ServerConfiguration_PostgresEvents:

		eventStore := postgres.NewServer(&postgres.Options{
			ConnectionString: store.PostgresEvents.ConnectionString,
		})

		chrononpb.RegisterEventStoreServer(server, eventStore)

		log.Printf("Postgres event store initalized")

	}

	switch store := config.SnapshotStore.(type) {

	case *chrononpb.ServerConfiguration_PostgresSnapshots:

		snapshotStore := postgres.NewServer(&postgres.Options{
			ConnectionString: store.PostgresSnapshots.ConnectionString,
		})

		chrononpb.RegisterSnapshotStoreServer(server, snapshotStore)

		log.Printf("Postgres snapshot store initalized")

	}

	lis, err := net.Listen("tcp", config.Address)

	if err != nil {
		log.Fatalf("Failed to listen server address: %v", err)
	}

	reflection.Register(server)

	log.Printf("Server listening at: %v", lis.Addr())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
