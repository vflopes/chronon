package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	LockFailErr = status.Error(codes.FailedPrecondition, "Unabled to acquire lock (already locked)")
)

type sourceLocker struct {
	database *sql.DB

	lockStmt *sql.Stmt
}

func newSourceLocker(database *sql.DB) *sourceLocker {

	s := &sourceLocker{
		database: database,
	}

	if err := s.prepareStatements(); err != nil {
		log.Fatalln("prepareStatements() error:", err)
	}

	return s
}

func (s *sourceLocker) prepareStatements() error {

	st, err := s.database.Prepare("SELECT pg_try_advisory_xact_lock($1::bigint)")

	if err != nil {
		return fmt.Errorf("Prepare lockStmt statement error: %v", err)
	}

	s.lockStmt = st

	return nil

}

func (s *sourceLocker) lock(ctx context.Context, numKey int64) (*sql.Tx, error) {

	tx, err := s.database.BeginTx(ctx, nil)

	if err != nil {
		return nil, fmt.Errorf("Source locker beginTx() error: %v", err)
	}

	locked := false

	err = tx.Stmt(s.lockStmt).QueryRow(numKey).Scan(&locked)

	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("Scan locke statement error: %v", err)
	}

	if locked {
		return tx, nil
	}

	tx.Rollback()
	return nil, LockFailErr

}
