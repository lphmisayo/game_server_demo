package csvs

import (
	"demoProject_gameServer/server/src/utils"
	"fmt"
)

type ConfigCard struct {
	CardId       int `json:"CardId"`
	Friendliness int `json:"Friendliness"`
	Check        int `json:"Check"`
}

var (
	ConfigCardMap         map[int]*ConfigCard
	ConfigCardMapByRoleId map[int]*ConfigCard
)

func init() {
	ConfigCardMap = make(map[int]*ConfigCard)
	utils.GetCsvUtilMgr().LoadCsv("Card", ConfigCardMap)
	ConfigCardMapByRoleId = make(map[int]*ConfigCard)
	for _, v := range ConfigCardMap {
		ConfigCardMapByRoleId[v.Check] = v
	}
	fmt.Println("加载Card...")
}

func GetConfigCard(cardId int) *ConfigCard {
	return ConfigCardMap[cardId]
}

func GetConfigCardByRoleId(roleId int) *ConfigCard {
	return ConfigCardMapByRoleId[roleId]
}
