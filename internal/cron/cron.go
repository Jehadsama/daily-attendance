package cron

import (
	"fmt"
	"os"

	"github.com/robfig/cron/v3"
)

var crontab string = os.Getenv("cronInterval")

var c = cron.New(cron.WithSeconds())

func Run(f func() string) {

	c.AddFunc(crontab, func() {
		r := f()
		fmt.Println(r)
	})

	c.Start()

	select {}
}
