package csvs

import (
	"demoProject_gameServer/server/src/utils"
	"fmt"
)

type ConfigPlayerLevel struct {
	PlayerLevel int `json:"PlayerLevel"`
	PlayerExp   int `json:"PlayerExp"`
	WorldLevel  int `json:"WorldLevel"`
	ChapterId   int `json:"ChapterID"`
}

var (
	ConfigPlayerLevelSlice []*ConfigPlayerLevel
)

func init() {
	utils.GetCsvUtilMgr().LoadCsv("PlayerLevel", &ConfigPlayerLevelSlice)
	fmt.Println("初始化经验csv完成")
}

func GetNowLevelConfig(level int) *ConfigPlayerLevel {
	if level < 0 || level >= len(ConfigPlayerLevelSlice) {
		return nil
	}
	return ConfigPlayerLevelSlice[level-1]
}
