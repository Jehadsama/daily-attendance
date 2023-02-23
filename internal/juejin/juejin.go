package juejin

import (
	"fmt"
	"os"

	"github.com/Jehadsama/daily-attendance/internal/utils"
)

var (
	checkIn          string = os.Getenv("juejinHost") + os.Getenv("juejinCheckin")
	getLotteryConfig string = os.Getenv("juejinHost") + os.Getenv("juejinGetLotteryConfig")
	drawLottery      string = os.Getenv("juejinHost") + os.Getenv("juejinDrawLottery")
	CK               string = utils.ReadFile("juejinCk")
)

type Err struct {
	Err_no  int
	Err_msg string
}

type checkInRes struct {
	Err
	Data bool
}

func (res *checkInRes) ReturnResponse() bool {
	return res.Err_no > 0
}

type drawLotteryRes struct {
	Err
	Data struct {
		Lottery_name      string
		Total_lucky_value int
		Draw_lucky_value  int
	}
}

func (res *drawLotteryRes) ReturnResponse() bool {
	return res.Err_no > 0
}

type lotteryConfigRes struct {
	Err
	Data struct {
		Lottery    interface{}
		Free_count int
		point_cost int
	}
}

func (res *lotteryConfigRes) ReturnResponse() bool {
	return (res.Err_no == 0 && res.Data.Free_count > 0)
}

// 签到
func CheckIn() {
	ok := utils.Request("POST", checkIn, CK, &checkInRes{})
	if ok {
		fmt.Println("【juejin sign in】successfully")
	} else {
		fmt.Println("【juejin sign in failied】")
	}
}

// 获取今日免费抽奖次数
func GetLotteryConfig() bool {
	return utils.Request("GET", getLotteryConfig, CK, &lotteryConfigRes{})
}

// 抽奖
func DrawLottery() {
	ok := GetLotteryConfig()
	if !ok {
		fmt.Println("【juejin draw lottery】no free lottery")
		return
	}
	ok = utils.Request("POST", drawLottery, CK, &drawLotteryRes{})
	if ok {
		fmt.Println("【juejin draw lottery】successfully")
	} else {
		fmt.Println("【juejin draw lottery】failed")
	}
}
