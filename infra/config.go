package infra

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	ServerName         string
	Environment        string
	TotalParkingSpaces int
}

func NewConfig() Config {
	//Environment it is not loaded
	if os.Getenv("ENVIRONMENT") == "" {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatalln("Error loading env file")
		}
	}

	number, err := strconv.Atoi(os.Getenv("TOTAL_PARKING_SPACES"))
	if err != nil {
		log.Fatalln("Error convert string to int, total parking spaces env.")
	}
	return Config{
		ServerName:         os.Getenv("SERVER_NAME"),
		Environment:        os.Getenv("ENVIRONMENT"),
		TotalParkingSpaces: number,
	}
}
