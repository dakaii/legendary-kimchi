package envvar

import (
	"os"
	"strconv"
)

func Migrate() bool {
	env, _ := os.LookupEnv("MIGRATE")
	migrate, err := strconv.ParseBool(env)
	if err != nil {
		migrate = false
	}
	return migrate
}

func ServerPort() string {
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

func DBName() string {
	dbName, exists := os.LookupEnv("MONGODB_DB_NAME")
	if !exists {
		dbName = "mongo"
	}
	return dbName
}

func PointCollection() string {
	url, exists := os.LookupEnv("MONGODB_COLLECTION_SCOOTER")
	if !exists {
		url = "collection"
	}
	return url
}

func MongoURL() string {
	url, exists := os.LookupEnv("MONGODB_URL")
	if !exists {
		url = "mongodb://localhost:27017"
	}
	return url
}
