package postgres

import (
	"database/sql"
	"log"

	chrononpb "github.com/vflopes/chronon/gen"
)

type Options struct {
	DriverName       string
	ConnectionString string
	Database         *sql.DB
	ScanBatchSize    int
	EventsTable      string
	SnapshotsTable   string
}

func setDefaultOtions(options *Options) {

	if len(options.DriverName) == 0 {
		options.DriverName = "postgres"
	}

	if len(options.EventsTable) == 0 {
		options.EventsTable = "events"
	}

	if len(options.SnapshotsTable) == 0 {
		options.SnapshotsTable = "snapshots"
	}

	if options.ScanBatchSize == 0 {
		options.ScanBatchSize = 50
	}

	if options.Database == nil {

		db, err := sql.Open(options.DriverName, options.ConnectionString)

		if err != nil {
			log.Fatal(err)
		}

		options.Database = db

	}

}

type Server struct {
	chrononpb.EventStoreServer
	chrononpb.SnapshotStoreServer
}

func NewServer(options *Options) *Server {
	s := &Server{}

	setDefaultOtions(options)

	s.EventStoreServer = newEventStore(options.Database, options.ScanBatchSize, options.EventsTable)
	s.SnapshotStoreServer = newSnapshotStore(options.Database, options.SnapshotsTable)

	return s
}
