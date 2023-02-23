package juejin

import (
	"fmt"

	"github.com/Jehadsama/daily-attendance/internal/utils"
)

type checkInRes struct {
	Err
	Data bool
}

func (res *checkInRes) ReturnResponse() bool {
	return res.Err_no == 0
}

// 签到
func CheckIn() {
	ok := utils.Request("POST", checkIn, CK, &checkInRes{})
	if ok {
		fmt.Println("【juejin sign in】successfully")
	} else {
		fmt.Println("【juejin sign in】failied")
	}
}
