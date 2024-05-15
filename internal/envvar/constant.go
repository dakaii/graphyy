package envvar

import (
	"os"
	"strconv"
)

func Port() string {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8081"
	}
	return port
}

// GetSecret returns the jwt secret.
func AuthSecret() string {
	secret, exists := os.LookupEnv("AUTH_SECRET")
	if !exists {
		secret = "secret_key"
	}
	return secret
}

func DBHost() string {
	host, exists := os.LookupEnv("DB_HOST")
	if !exists {
		host = "localhost"
	}
	return host
}

func DBUser() string {
	user, exists := os.LookupEnv("DB_USER")
	if !exists {
		user = "postgres"
	}
	return user
}

func DBPassword() string {
	password, exists := os.LookupEnv("DB_PASS")
	if !exists {
		password = "postgres"
	}
	return password
}

func DBName() string {
	dbname, exists := os.LookupEnv("DB_NAME")
	if !exists {
		dbname = "postgres"
	}
	return dbname
}

func DBPort() string {
	port, exists := os.LookupEnv("DB_PORT")
	if !exists {
		port = "5432"
	}
	return port
}

func HashCost() int {
	costString, _ := os.LookupEnv("HASH_COST")
	res, err := strconv.Atoi(costString)
	if err != nil {
		return 8
	}

	return res
}
