package juejin

import (
	"log"

	"github.com/Jehadsama/daily-attendance/internal/utils"
)

type checkInRes struct {
	Err
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
		log.Println("【juejin sign in】successfully")
	} else {
		log.Println("【juejin sign in】failied")
	}
}
