package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	mongodb "./external/mongodb"
	handler "./handler"
)

func main() {

	var port = getEnv("WEB_PORT", "8081")
	var mongohost = getEnv("MONGO_HOST", "localhost")
	var mongoport = getEnvInt("MONGO_PORT", 27017)
	var username = getEnv("MONGO_USERNAME", "")
	var password = getEnv("MONGO_PASSWORD", "")
	var database = getEnv("MONGO_DATABASE", "todo")
	var collection = getEnv("MONGO_COLLECTION", "todo")

	repository := mongodb.NewMyMongoRepository(mongohost, mongoport, username, password, database, collection)
	r := handler.CreateRoutes(repository)

	fmt.Printf("Server started on port %v...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
	fmt.Printf("Done!\n")
}

// Returns the environment variable or a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Returns the environment variable or a default value
func getEnvInt(key string, defaultVal int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}
