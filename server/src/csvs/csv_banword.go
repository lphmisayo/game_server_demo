package csvs

import (
	"demoProject_gameServer/server/src/utils"
	"fmt"
)

type ConfigBanWord struct {
	Id  int    `json:"id"`
	Txt string `json:"txt"`
}

var (
	ConfigBanwordSlice []*ConfigBanWord
)

func init() {
	utils.GetCsvUtilMgr().LoadCsv("banword", &ConfigBanwordSlice)
	fmt.Println("csv_banword初始化")
}

func GetBaseBanWord() []string {
	rtnString := make([]string, 0)
	for _, v := range ConfigBanwordSlice {
		rtnString = append(rtnString, v.Txt)
	}
	return rtnString
}
