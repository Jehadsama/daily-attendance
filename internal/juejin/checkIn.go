package juejin

import (
	"github.com/Jehadsama/daily-attendance/internal/utils"
)

type checkInRes struct {
	Err
}

func (res *checkInRes) Success() bool {
	return res.Err_no == 0
}

// 签到
func CheckIn() string {
	var res checkInRes
	utils.Request("POST", checkIn, CK, nil, &res)
	ok := res.Success()
	if ok {
		return "【juejin sign in】successfully"
	} else {
		return "【juejin sign in】failied"
	}
}
