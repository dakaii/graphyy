package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upInit, downInit)
}

func upInit(ctx context.Context, tx *sql.Tx) error {
	// dbname := envvar.DBName()
	// _, err := tx.Exec("CREATE TABLE " + dbname)
	// if err != nil {
	// 	return err
	// }
	return nil
}

func downInit(ctx context.Context, tx *sql.Tx) error {
	// dbname := envvar.DBName()
	// _, err := tx.Exec("DROP TABLE " + dbname)
	// if err != nil {
	// 	return err
	// }
	return nil
}
