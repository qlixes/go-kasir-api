package configs

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var appInstance sync.Once

func LoadApp() {
	appInstance.Do(func() {
		err := godotenv.Load()

		if err != nil {
			log.Println("Failed load environment")
		}
	})
}

func GetEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	panic("Missing environment key")
}
