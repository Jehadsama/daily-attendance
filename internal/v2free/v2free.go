package v2free

import (
	"fmt"
	"os"

	utils "github.com/Jehadsama/daily-attendance/internal/utils"
)

var (
	url string = os.Getenv("v2freeHost") + os.Getenv("v2freeCheckin")
	CK  string = utils.ReadFile("v2freeCk")
)

type V2freeRes struct {
	Ret int    `json:"ret"`
	Msg string `json:"msg"`
}

func (res *V2freeRes) Success() bool {
	return res.Ret == 1
}

func SignIn() {
	var res V2freeRes
	utils.Request("POST", url, CK, nil, &res)
	msg := "【v2free sign in】failed"
	ok := res.Success()
	if ok {
		msg = "【v2free sign in】successfully"
	}
	fmt.Println(msg)
}
