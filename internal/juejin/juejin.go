package juejin

import (
	"fmt"
	"os"

	"github.com/Jehadsama/daily-attendance/internal/utils"
)

var (
	checkIn     string = os.Getenv("juejinHost") + os.Getenv("juejinCheckin")
	drawLottery string = os.Getenv("juejinHost") + os.Getenv("juejinDrawLottery")
	CK          string = utils.ReadFile("./juejinCk")
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

// 签到
func CheckIn() {
	result := utils.Request("POST", checkIn, CK, &checkInRes{})
	res := result.(*resp)
	if res.Success {
		fmt.Print("签到成功")
	} else {
		fmt.Println("签到失败:", res.Data)
	}
}

// 抽奖
func DrawLottery() {
	result := utils.Request("POST", drawLottery, CK, &drawLotteryRes{})
	res := result.(*resp)
	if res.Success {
		fmt.Println("抽奖成功")
	} else {
		fmt.Println("免费抽奖失败:", res.Data)
	}
}
