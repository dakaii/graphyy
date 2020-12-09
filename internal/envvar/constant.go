package envvar

import "os"

// GetSecret returns the jwt secret.
func AuthSecret() string {
	secret, exists := os.LookupEnv("AUTH_SECRET")
	if !exists {
		secret = "secret_key"
	}
	return secret
}

func DBUser() string {
	user, exists := os.LookupEnv("POSTGRES_USER")
	if !exists {
		user = "postgres"
	}
	return user
}

func DBPassword() string {
	password, exists := os.LookupEnv("POSTGRES_PASSWORD")
	if !exists {
		password = "postgres"
	}
	return password
}

func DBName() string {
	dbname, exists := os.LookupEnv("POSTGRES_DB_NAME")
	if !exists {
		dbname = "postgres"
	}
	return dbname
}
func DBHost() string {
	host, exists := os.LookupEnv("POSTGRES_HOST")
	if !exists {
		host = "localhost"
	}
	return host
}

func DBPort() string {
	port, exists := os.LookupEnv("POSTGRES_PORT")
	if !exists {
		port = "5432"
	}
	return port
}
