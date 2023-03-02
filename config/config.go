package config

import (
	"log"
	"os"

	"github.com/Jehadsama/daily-attendance/internal/utils"
	"github.com/joho/godotenv"
)

func init() {
	env := os.Getenv("DAILY_ENV")
	if env == "" {
		env = "dev"
	}
	log.Println("env:", env)
	err := godotenv.Load("config/"+env, "config/dev")
	utils.CheckError("config,init", err)
}
