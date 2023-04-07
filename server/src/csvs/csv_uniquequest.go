package csvs

import (
	"demoProject_gameServer/server/src/utils"
	"fmt"
)

type ConfigUniqueQuest struct {
	QuestId   int `json:"QuestId"`
	SortType  int `json:"SortType"`
	OpenLevel int `json:"OpenLevel"`
	QuestType int `json:"QuestType"`
	Condition int `json:"Condition"`
}

var (
	ConfigUniqueQuestSlice map[int]*ConfigUniqueQuest
)

func init() {
	ConfigUniqueQuestSlice := make(map[int]*ConfigUniqueQuest)
	utils.GetCsvUtilMgr().LoadCsv("uniqueQuest", &ConfigUniqueQuestSlice)
	fmt.Println("世界等级任务模块csv加载完成...")
	return
}
