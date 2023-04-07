package csvs

import "demoProject_gameServer/server/src/utils"

type ConfigIcon struct {
	IconId int `json:"IconId"`
	Check  int `json:"Check"`
}

var (
	ConfigIconMap         map[int]*ConfigIcon
	ConfigIconMapByRoleId map[int]*ConfigIcon //一个临时表，将roleid设为键值，存储icon对象，牺牲一点空间，将时间复杂度n降至常数级
)

func init() {
	ConfigIconMap = make(map[int]*ConfigIcon)
	utils.GetCsvUtilMgr().LoadCsv("Icon", ConfigIconMap)
	ConfigIconMapByRoleId = make(map[int]*ConfigIcon)
	for _, v := range ConfigIconMap {
		ConfigIconMapByRoleId[v.Check] = v
	}
	return
}

func GetIconConfig(iconId int) *ConfigIcon {
	return ConfigIconMap[iconId]
}

func GetIconConfigByRoleId(roleId int) *ConfigIcon {
	return ConfigIconMapByRoleId[roleId]
}
