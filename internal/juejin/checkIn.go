package juejin

import (
	"fmt"

	"github.com/Jehadsama/daily-attendance/internal/utils"
)

type checkInRes struct {
	Err
	Data bool
}

func (res *checkInRes) Success() bool {
	return res.Err_no == 0
}

// 签到
func CheckIn() {
	var res checkInRes
	utils.Request("POST", checkIn, CK, nil, &res)
	ok := res.Success()
	if ok {
		fmt.Println("【juejin sign in】successfully")
	} else {
		fmt.Println("【juejin sign in】failied")
	}
}
