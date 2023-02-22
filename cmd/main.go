package main

import (
	"log"
	"sync"

	_ "github.com/Jehadsama/daily-attendance/config"
	"github.com/Jehadsama/daily-attendance/internal/juejin"
	"github.com/Jehadsama/daily-attendance/internal/v2free"
)

var wg sync.WaitGroup

var funcsMap = map[string]func(){
	"v2free.SignIn":      v2free.SignIn,
	"juejin.CheckIn":     juejin.CheckIn,
	"juejin.DrawLottery": juejin.DrawLottery,
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
		wg.Add(1)
		go func(funcName string, fun func()) {
			log.Println(funcName, "start")
			defer wg.Done()
			fun()
			log.Println(funcName, "end")
		}(name, f)
	}
	wg.Wait()
	log.Println("main,end")
}
