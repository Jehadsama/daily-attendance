package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func ReadFile(filepath string) string {
	content, err := os.ReadFile(filepath)
	if err != nil {
		panic("read file failed!")
	}
	return string(content)
}

type Result interface {
	Response() string
}

func Request(method string, url string, cookie string, result Result) (string, error) {
	var req *http.Request
	req, err := http.NewRequest(method, url, nil)
	CheckError(err)
	req.Header.Set("Cookie", cookie)
	resp, err := http.DefaultClient.Do(req)
	CheckError(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	CheckError(err)
	err = json.Unmarshal(body, result)
	CheckError(err)
	return result.Response(), nil

}
