package main

import (
	"log"

	_ "github.com/Jehadsama/daily-attendance/config"
	"github.com/Jehadsama/daily-attendance/internal/juejin"
	"github.com/Jehadsama/daily-attendance/internal/v2free"
)

var funcsMap = map[string]func(){
	"v2free.SignIn":      v2free.SignIn,
	"juejin.DrawLottery": juejin.DrawLottery,
	"juejin.DipLucky":    juejin.DipLucky,
}

func main() {

	// if len(os.Args) > 1 {
	// 	for _, f := range funcs {
	// 		wg.Add(1)
	// 		go func(fun func()) {
	// 			defer wg.Done()
	// 			fun()
	// 		}(f)
	// 	}
	// 	wg.Wait()
	// } else {
	// 	cron.Run(v2free.SignIn)
	// }

	log.Println("main,start")
	for name, f := range funcsMap {
		log.Println(name, "start")
		f()
		log.Println(name, "end")
	}
	log.Println("main,end")
}
