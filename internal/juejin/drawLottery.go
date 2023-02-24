package juejin

import (
	"fmt"

	"github.com/Jehadsama/daily-attendance/internal/utils"
)

type drawLotteryRes struct {
	Err
	Data struct {
		Lottery_name      string
		Total_lucky_value int
		Draw_lucky_value  int
	}
}

func (res *drawLotteryRes) Success() bool {
	return res.Err_no == 0
}

type lotteryConfigRes struct {
	Err
	Data struct {
		Lottery    interface{}
		Free_count int
		point_cost int
	}
}

func (res *lotteryConfigRes) Success() bool {
	return (res.Err_no == 0 && res.Data.Free_count > 0)
}

// 获取今日免费抽奖次数
func GetLotteryConfig() bool {
	var res lotteryConfigRes
	utils.Request("GET", getLotteryConfig, CK, nil, &res)
	return res.Success()
}

// 抽奖
func DrawLottery() {
	CheckIn()
	ok := GetLotteryConfig()
	if !ok {
		fmt.Println("【juejin draw lottery】no free lottery")
		return
	}
	var res drawLotteryRes
	utils.Request("POST", drawLottery, CK, nil, &res)
	ok = res.Success()
	if ok {
		fmt.Println("【juejin draw lottery】successfully")
	} else {
		fmt.Println("【juejin draw lottery】failed")
	}
}
