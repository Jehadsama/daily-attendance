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

func (res *V2freeRes) ReturnResponse() (data interface{}) {
	if res.Ret == 1 {
		return "v2free sign in successfully"
	}
	return "v2free sign in failed"
}

func SignIn() {
	msg := utils.Request("POST", url, CK, &V2freeRes{})
	fmt.Println(msg.(string))
}
