package utils

import (
	"io/ioutil"
	"log"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func ReadFile(filepath string) string {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic("read file failed!")
	}
	return string(content)
}
