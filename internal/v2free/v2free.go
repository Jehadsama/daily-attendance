package v2free

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	utils "github.com/Jehadsama/daily-attendance/internal/utils"
)

var url string = os.Getenv("V2FREE_HOST")
var CK string = utils.ReadFile("./v2freeCk")

type V2freeRes struct {
	Ret int    `json:"ret"`
	Msg string `json:"msg"`
}

func SignIn() string {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	utils.CheckError(err)
	req.Header.Add("ContentType", "application/json")
	req.Header.Add("Cookie", CK)
	res, err := client.Do(req)
	utils.CheckError(err)

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	utils.CheckError(err)

	result := &V2freeRes{}
	err = json.Unmarshal(body, result)
	utils.CheckError(err)

	var msg string = "v2free sign in failed"
	if result.Ret == 1 {
		msg = "v2free sign in successfully"
	}
	return msg
}
