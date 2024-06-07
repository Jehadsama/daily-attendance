package cron

import (
	"os"

	_ "github.com/Jehadsama/daily-attendance/config"
	"github.com/robfig/cron/v3"
)

var crontab string = os.Getenv("cronInterval")

func Run(f func()) {
	c := cron.New()
	c.AddFunc(crontab, f)
	c.Start()
	select {}
}
