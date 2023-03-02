package juejin

import (
	"log"
	"strings"

	"github.com/Jehadsama/daily-attendance/internal/utils"
)

type getLuckyUserListRes struct {
	Err
	Data struct {
		Lotteries []struct {
			History_id string
		}
	}
}

func (res *getLuckyUserListRes) Success() bool {
	return res.Err_no == 0
}

type dipLuckyRes struct {
	Err
	Data struct {
		Total_value int
	}
}

func (res *dipLuckyRes) Success() bool {
	return res.Err_no == 0
}

func GetLuckyUserList() string {
	var res getLuckyUserListRes
	utils.Request("POST", getLuckyUserList, "", nil, &res)
	return res.Data.Lotteries[0].History_id
}

func DipLucky() {
	historyId := GetLuckyUserList()
	var res dipLuckyRes
	utils.Request("POST", dipLucky, CK, strings.NewReader("{\"History_id\":"+historyId+"}"), &res)
	ok := res.Success()
	if ok {
		log.Println("【juejin dip lucky】successfully")
	} else {
		log.Println("【juejin dip lucky】failied")
	}
}
