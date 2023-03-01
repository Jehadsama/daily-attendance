package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func CheckError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
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

type Responser interface {
	Success() bool
}

func Request(method string, url string, cookie string, data io.Reader, response Responser) {
	log.Println("utils,Request", method, url)
	var req *http.Request
	req, err := http.NewRequest(method, url, data)
	CheckError("utils,Request,NewRequest", err)
	req.Header.Set("Cookie", cookie)
	resp, err := http.DefaultClient.Do(req)
	CheckError("utils,Request,Do", err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	CheckError("utils,Request,ReadAll", err)
	err = json.Unmarshal(body, response)
	fmt.Printf("======\n%#v\n=========\n", response)
	CheckError("utils,Request,Unmarshal", err)
}
