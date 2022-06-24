package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// Config is the global config
	DATABASE_URL string
	PORT         string
	ADMIN        string
	FREE_LIMIT   int
	VIP_LIMIT    int
)

func Load() error {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DATABASE_URL = os.Getenv("DATABASE_URL")
	PORT = os.Getenv("PORT")
	ADMIN = os.Getenv("ADMIN")
	FREE_LIMIT, err = strconv.Atoi(os.Getenv("FREE_LIMIT"))
	if err != nil {
		return err
	}

	VIP_LIMIT, err = strconv.Atoi(os.Getenv("VIP_LIMIT"))
	if err != nil {
		return err
	}

	// print all
	log.Println("DATABASE_URL:", DATABASE_URL)
	log.Println("PORT:", PORT)
	log.Println("ADMIN:", ADMIN)
	log.Println("FREE_LIMIT:", FREE_LIMIT)
	log.Println("VIP_LIMIT:", VIP_LIMIT)

	return nil
}
