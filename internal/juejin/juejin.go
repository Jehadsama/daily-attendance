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

type resp struct {
	Success bool
	Data    string
}

type Err struct {
	Err_no  int
	Err_msg string
}

type checkInRes struct {
	Err
	Data bool
}

func (res *checkInRes) ReturnResponse() interface{} {
	r := &resp{
		Success: true,
		Data:    "",
	}
	if res.Err_no > 0 {
		r.Success = false
		r.Data = res.Err_msg
	}
	return r
}

type drawLotteryRes struct {
	Err
	Data struct {
		Lottery_name      string
		Total_lucky_value int
		Draw_lucky_value  int
	}
}

func (res *drawLotteryRes) ReturnResponse() interface{} {
	r := &resp{
		Success: true,
		Data:    fmt.Sprintf("免费抽奖成功,奖品:%v,增加幸运值:%v,总幸运值:%v", res.Data.Lottery_name, res.Data.Draw_lucky_value, res.Data.Total_lucky_value),
	}
	if res.Err_no > 0 {
		r.Success = false
		r.Data = res.Err_msg
	}
	return r
}

type lotteryConfigRes struct {
	Err
	Data struct {
		Lottery    interface{}
		Free_count int
		point_cost int
	}
}

func (res *lotteryConfigRes) ReturnResponse() interface{} {
	r := &resp{
		Success: false,
	}
	if res.Err_no == 0 && res.Data.Free_count > 0 {
		r.Success = true
	}
	return r
}

// 签到
func CheckIn() {
	result := utils.Request("POST", checkIn, CK, &checkInRes{})
	res := result.(*resp)
	if res.Success {
		fmt.Println("juejin sign in successfully")
	} else {
		fmt.Println("juejin sign in failied:", res.Data)
	}
}

// 获取今日免费抽奖次数
func GetLotteryConfig() bool {
	result := utils.Request("GET", getLotteryConfig, CK, &lotteryConfigRes{})
	res := result.(*resp)
	return res.Success
}

// 抽奖
func DrawLottery() {
	ok := GetLotteryConfig()
	if !ok {
		fmt.Println("no free lottery")
		return
	}
	result := utils.Request("POST", drawLottery, CK, &drawLotteryRes{})
	res := result.(*resp)
	if res.Success {
		fmt.Println("juejin draw lottery successfully")
	} else {
		fmt.Println("juejin draw lottery failed:", res.Data)
	}
}
