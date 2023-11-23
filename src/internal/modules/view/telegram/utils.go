package telegram

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func popElemFront(a []string) (string, []string) {
	x, a := a[0], a[1:]
	return x, a
}
