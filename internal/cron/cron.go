package cron

import (
	"os"

	"github.com/robfig/cron/v3"
)

var crontab string = os.Getenv("cronInterval")

var c = cron.New(cron.WithSeconds())

func Run(f func()) {

	c.AddFunc(crontab, func() {
		f()
	})

	c.Start()

	select {}
}
