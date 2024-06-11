package cron

import (
	"os"
	"time"

	_ "github.com/Jehadsama/daily-attendance/config"
	"github.com/robfig/cron/v3"
)

var timezone, _ = time.LoadLocation("Asia/Shanghai")

func Run(crontabKey string, f func()) {
	c := cron.New(cron.WithLocation(timezone))
	crontab := os.Getenv(crontabKey)
	c.AddFunc(crontab, f)
	c.Start()
	select {}
}
