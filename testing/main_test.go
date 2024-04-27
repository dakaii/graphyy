package testing

import (
	"database/sql"
	"fmt"
	"graphyy/entity"
	"graphyy/internal/envvar"
	"log"
	"os"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	setup()

	// runs all tests in the package
	exitCode := m.Run()

	teardown()

	os.Exit(exitCode)
}

func setup() {
	user := envvar.DBUser()
	password := envvar.DBPassword()
	dbname := envvar.DBName()
	dbhost := envvar.DBHost()
	dbport := envvar.DBPort()

	isTestMode := dbname == "graphyy_test"
	if isTestMode {
		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", dbhost, dbport, user, password)
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		// Create the database if it doesn't exist
		_, err = db.Exec("CREATE DATABASE " + dbname)
		if err != nil {
			log.Println("Database already exists:", err)
		}
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", dbhost, user, password, dbname, dbport)
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if isTestMode {
		gormDB.AutoMigrate(&entity.User{})
	}
}

func teardown() {
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
