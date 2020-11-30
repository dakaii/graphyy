package database

import (
	"fmt"
	"graphyy/envvar"
	"graphyy/model"

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
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		user, password, dbname, dbhost, dbport)
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
