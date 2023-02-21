package main

import (
	"fmt"
	"os"

	_ "github.com/Jehadsama/daily-attendance/config"
	"github.com/Jehadsama/daily-attendance/internal/cron"
	"github.com/Jehadsama/daily-attendance/internal/v2free"
)

func main() {

	if len(os.Args) > 1 {
		result := v2free.SignIn()
		fmt.Println(result)
	} else {
		cron.Run(v2free.SignIn)
	}

}
