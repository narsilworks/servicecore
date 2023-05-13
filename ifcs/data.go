package ifcs

import (
	"context"
	"database/sql"
	"time"
)

type IData interface {
	ID() string
	Open(ctx context.Context, di IDatabaseInfo) error                  // Open a new connection
	Exec(sql string, args ...interface{}) (int64, error)               // Exec executes a non-returning query
	Exists(sqlWithParams string, args ...interface{}) (bool, error)    // Checks existence of a record
	Query(sql string, args ...interface{}) (sql.Rows, error)           // Query to a database to return one or more records
	QueryArray(sql string, out interface{}, args ...interface{}) error // Query to a database to return one or more records and store to an array
	QueryRow(sql string, args ...interface{}) sql.Row                  // QueryRow to a database and return one record
	Next(serial string, next *int64) error                             // Get next value of a serial
	Now(utc bool) *time.Time                                           // Get time now
	Begin() error                                                      // Begin a transaction. If there is an existing transaction, begin is ignored
	BeginDR() (string, error)                                          // Begin a transaction with transaction id to use when rollback is deferred
	Mark(name string) error                                            // Mark a savepoint
	Discard(name string) error                                         // Discard a savepoint
	Save(name string) error                                            // Save a transaction
	Rollback(...string) error                                          // Rollback a transaction
	Commit(...string) error                                            // Commit the transaction
	DatabaseVersion() string                                           // Get database version
	Escape(fv string) string                                           // Escape a field value (fv) from disruption by single quote
	Close() error                                                      // Close connection
}
