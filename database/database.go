package database

import (
	"fmt"
	"graphyy/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GetDatabase returns a database instance.
func GetDatabase() *gorm.DB {

	user, exists := os.LookupEnv("POSTGRES_USER")
	if !exists {
		user = "postgres"
	}

	password, exists := os.LookupEnv("POSTGRES_PASSWORD")
	if !exists {
		password = "postgres"
	}

	dbname, exists := os.LookupEnv("POSTGRES_DB_NAME")
	if !exists {
		dbname = "postgres"
	}

	host, exists := os.LookupEnv("POSTGRES_HOST")
	if !exists {
		host = "localhost"
	}

	port, exists := os.LookupEnv("POSTGRES_PORT")
	if !exists {
		port = "5432"
	}

	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		user, password, dbname, host, port)
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&model.User{})
	return db
}

// func createPgDb(port string, host string, user string, password string, dbname string) {
// 	cmd := exec.Command("createdb", "-p", port, "-h", host, "-U", user, "-W", password, "-e", dbname)
// 	var out bytes.Buffer
// 	cmd.Stdout = &out
// 	if err := cmd.Run(); err != nil {
// 		log.Printf("Error: %v", err)
// 	}
// 	log.Printf("Output: %q\n", out.String())
// }
