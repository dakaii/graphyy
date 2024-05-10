package testing

import (
	"fmt"
	"graphyy/internal/envvar"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TruncateAllTables() {
	password := envvar.DBPassword()
	dbname := envvar.DBName()
	dbhost := envvar.DBHost()
	dbport := envvar.DBPort()
	user := envvar.DBUser()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", dbhost, user, password, dbname, dbport)
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		db, _ := gormDB.DB()
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// Delete all rows from all tables
	err = gormDB.Exec(`
			DO $$
			DECLARE
				r RECORD;
			BEGIN
				FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = current_schema()) LOOP
					EXECUTE 'DELETE FROM ' || quote_ident(r.tablename) || ' CASCADE';
				END LOOP;
			END $$;
		`).Error
	if err != nil {
		log.Println("Failed to delete all rows from all tables:", err)
	}
}
