package juejin

import (
	"os"

	"github.com/Jehadsama/daily-attendance/internal/utils"
)

var (
	checkIn          string = os.Getenv("juejinHost") + os.Getenv("juejinCheckin")
	getLotteryConfig string = os.Getenv("juejinHost") + os.Getenv("juejinGetLotteryConfig")
	drawLottery      string = os.Getenv("juejinHost") + os.Getenv("juejinDrawLottery")
	getLuckyUserList string = os.Getenv("juejinHost") + os.Getenv("juejinGetLuckyUserList")
	dipLucky         string = os.Getenv("juejinHost") + os.Getenv("juejinDipLucky")
	CK               string = utils.ReadFile("juejinCk")
)

type Err struct {
	Err_no  int
	Err_msg string
}
