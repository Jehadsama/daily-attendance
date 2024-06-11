package healthz

import (
	"fmt"
	"time"
)

var loc, _ = time.LoadLocation("Asia/Shanghai")

func Run() {
	fmt.Println("Hello healthz,", time.Now().In(loc))
}
