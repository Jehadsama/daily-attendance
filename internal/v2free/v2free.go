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

func (result *V2freeRes) Response() string {
	var msg string = "v2free sign in failed"
	if result.Ret == 1 {
		msg = "v2free sign in successfully"
	}
	return msg
}

func SignIn() string {
	msg, err := utils.Request("POST", url, CK, &V2freeRes{})
	utils.CheckError(err)
	return msg
}
