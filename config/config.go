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
	dirname, _ := os.Getwd()
	err := godotenv.Load(dirname+"/config/"+env, dirname+"/config/dev")
	utils.CheckError(err)
}
