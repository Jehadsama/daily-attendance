package v2free

import (
	"os"

	utils "github.com/Jehadsama/daily-attendance/internal/utils"
)

var url string = os.Getenv("V2FREE_HOST")
var CK string = utils.ReadFile("./v2freeCk")

type V2freeRes struct {
	Ret int    `json:"ret"`
	Msg string `json:"msg"`
}

func (result *V2freeRes) Response() (msg string) {
	if result.Ret == 1 {
		return "v2free sign in successfully"
	}
	return "v2free sign in failed"
}

func SignIn() string {
	msg, err := utils.Request("POST", url, CK, &V2freeRes{})
	utils.CheckError(err)
	return msg
}
