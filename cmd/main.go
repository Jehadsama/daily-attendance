package main

import (
	"log"
	"os"
	"strings"

	_ "github.com/Jehadsama/daily-attendance/config"
	"github.com/Jehadsama/daily-attendance/internal/cron"
	"github.com/Jehadsama/daily-attendance/internal/email"
	"github.com/Jehadsama/daily-attendance/internal/healthz"
	"github.com/Jehadsama/daily-attendance/internal/juejin"
)

var funcsMap = map[string]func() string{
	// TODO: v2free增加机器人验证，后面再想想怎么处理..
	// "v2free.SignIn":      v2free.SignIn,
	"juejin.DrawLottery": juejin.DrawLottery,
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
		go func() {
			cron.Run("everyMinute", healthz.Run)
		}()
		cron.Run("everyDay", run)
	} else {
		run()
	}
}
