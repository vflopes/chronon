package chronon

import (
	_ "github.com/lib/pq"

	"github.com/spf13/cobra"
	"github.com/vflopes/chronon/event/store/postgres"
	chrononpb "github.com/vflopes/chronon/gen"
)

var (
	storeEventsPostgres    bool
	storeSnapshotsPostgres bool

	storeConfigPostgres = &chrononpb.StoreConfiguration_Postgres{}

	storePostgresCmd = &cobra.Command{
		Use:   "postgres",
		Short: "Use Postgres as event and/or snapshot store",
		Run: func(cmd *cobra.Command, args []string) {

			postgresServer := postgres.NewServer(&postgres.Options{
				DriverName:       "postgres",
				ConnectionString: storeConfigPostgres.ConnectionString,
				EventsTable:      storeConfigPostgres.EventsTable,
				SnapshotsTable:   storeConfigPostgres.SnapshotsTable,
				ScanBatchSize:    int(storeConfigPostgres.ScanBatchSize),
			})

			if storeEventsPostgres {
				chrononpb.RegisterEventStoreServer(storeServer, postgresServer)
			}

			if storeSnapshotsPostgres {
				chrononpb.RegisterSnapshotStoreServer(storeServer, postgresServer)
			}

			serveStoreServer()

		},
	}
)

func init() {

	storeCmd.AddCommand(storePostgresCmd)

	storePostgresCmd.PersistentFlags().StringVar(&storeConfigPostgres.ConnectionString, "postgresConnectionString", "", "Postgres connection string")

	storePostgresCmd.PersistentFlags().Int64Var(&storeConfigPostgres.ScanBatchSize, "postgresScanBatchSize", 500, "Scan batch size for pagination strategy")

	storePostgresCmd.PersistentFlags().BoolVar(&storeEventsPostgres, "storeEventsPostgres", true, "Use postgres as event store")
	storePostgresCmd.PersistentFlags().StringVar(&storeConfigPostgres.EventsTable, "postgresEventsTable", "events", "Postgres events table name")

	storePostgresCmd.PersistentFlags().BoolVar(&storeSnapshotsPostgres, "storeSnapshotsPostgres", true, "Use postgres as snapshot store")
	storePostgresCmd.PersistentFlags().StringVar(&storeConfigPostgres.SnapshotsTable, "postgresSnapshotsTable", "snapshots", "Postgres snapshots table name")

}
