package internal

import (
	"os"
	"strconv"
)

func getStrEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func getBoolEnv(key, defaultValue string) bool {
	env, _ := os.LookupEnv(key)
	value, err := strconv.ParseBool(env)
	if err != nil {
		value = false
	}
	return value
}

var Migrate = getBoolEnv("MIGRATE", "false")

var ServerPort = getStrEnv("PORT", "8081")

var DBName = getStrEnv("MONGODB_DB_NAME", "mongo")

var MongoURL = getStrEnv("MONGODB_URL", "mongodb://localhost:27017")
