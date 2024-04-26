package database

import (
	"database/sql"
	"fmt"
	"graphyy/entity"
	"graphyy/internal/envvar"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GetDatabase returns a database instance.
func GetDatabase() *gorm.DB {

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

	return gormDB
}
