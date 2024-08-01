package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	APIURL   = ""
	Port     = 0
	HashKey  []byte
	BlockKey []byte
)

func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading.env file: ", err)
	}

	if Port, err = strconv.Atoi(os.Getenv("APP_PORT")); err != nil {
		Port = 3000
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
