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
	_, err := tx.Exec(`
        CREATE TABLE users (
            id UUID PRIMARY KEY,
            created_at TIMESTAMP NOT NULL,
            updated_at TIMESTAMP NOT NULL,
            deleted_at TIMESTAMP,
            username VARCHAR(255) UNIQUE NOT NULL,
            password VARCHAR(1000) NOT NULL
        );
    `)
	if err != nil {
		return err
	}

	return nil
}

func downInit(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE users")
	if err != nil {
		return err
	}
	return nil
}
