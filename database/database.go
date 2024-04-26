package database

import (
	"fmt"
	"graphyy/internal/envvar"

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

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", dbhost, user, password, dbname, dbport)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
