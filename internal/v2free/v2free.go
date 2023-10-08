package v2free

import (
	"os"

	utils "github.com/Jehadsama/daily-attendance/internal/utils"
)

var (
	url string = os.Getenv("v2freeHost") + os.Getenv("v2freeCheckin")
	CK  string = utils.ReadFile("v2freeCk")
)

type V2freeRes struct {
	Ret           int64       `json:"ret"`
	Msg           string      `json:"msg"`
	Unflowtraffic int64       `json:"unflowtraffic"`
	Traffic       string      `json:"traffic"`
	TrafficInfo   interface{} `json:"trafficInfo"`
}

func (res *V2freeRes) Success() bool {
	return res.Ret == 1
}

func SignIn() string {
	var res V2freeRes
	utils.Request("POST", url, CK, nil, &res)
	msg := "【v2free sign in】failed"
	ok := res.Success()
	if ok {
		return "【v2free sign in】successfully"
	}
	return msg
}
