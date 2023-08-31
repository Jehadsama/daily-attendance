package main

import (
	"log"
	"os"
	"strings"

	_ "github.com/Jehadsama/daily-attendance/config"
	"github.com/Jehadsama/daily-attendance/internal/cron"
	"github.com/Jehadsama/daily-attendance/internal/email"
	"github.com/Jehadsama/daily-attendance/internal/juejin"
	"github.com/Jehadsama/daily-attendance/internal/v2free"
)

var funcsMap = map[string]func() string{
	"v2free.SignIn":      v2free.SignIn,
	"juejin.DrawLottery": juejin.DrawLottery,
	"juejin.DipLucky":    juejin.DipLucky,
}

var run = func() {
	log.Println("cron start")
	messages := []string{}
	for name, f := range funcsMap {
		log.Println(name, "start")
		msg := f()
		messages = append(messages, msg)
		log.Println(name, "end")
	}

	ed := &email.EmailData{
		From:    "ayumijehad@qq.com",
		To:      []string{"ayumijehad@qq.com"},
		Subject: "daily-attendance",
		Content: strings.Join(messages, "\n"),
	}

	ed.SendMail()
	log.Println("main end")
}

func main() {
	if len(os.Args) > 1 {
		cron.Run(run)
	} else {
		run()
	}
}
