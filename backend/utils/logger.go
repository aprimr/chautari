package utils

import (
	"log"
	"os"
)

func LogInfo(message string) {
	log.Println(message)
}

func LogDebug(message string) {
	env := os.Getenv("ENVIRONMENT")
	if env == "development" {
		log.Printf("DEBUG: %s\n", message)
	}
}

func LogError(message string, err error) {
	env := os.Getenv("ENVIRONMENT")
	if env == "development" {
		log.Printf("ERROR: %s -> %v\n", message, err)
	}
}

func LogFatal(message string, err error) {
	env := os.Getenv("ENVIRONMENT")
	if env == "development" {
		log.Fatalf("FATAL: %s -> %v\n", message, err)
	} else {
		log.Fatalf("FATAL: %s\n", message)
	}
}
