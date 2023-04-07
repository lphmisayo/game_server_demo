package csvs

import "demoProject_gameServer/server/src/utils"

type ConfigRole struct {
	RoleId          int    `json:"RoleId"`
	ItemName        string `json:"ItemName"`
	Star            int    `json:"Star"`
	Stuff           int    `json:"Stuff"`           //命座材料对应物品id
	StuffNum        int    `json:"StuffNum"`        //命座材料数量
	StuffItem       int    `json:"StuffItem"`       //材料返还id
	StuffItemNum    int    `json:"StuffItemNum"`    //材料返还数量
	MaxStuffItem    int    `json:"MaxStuffItem"`    //最大材料返还id
	MaxStuffItemNum int    `json:"MaxStuffItemNum"` //最大材料返还数量
}

var (
	ConfigRoleMap map[int]*ConfigRole
)

func init() {
	ConfigRoleMap = make(map[int]*ConfigRole)
	utils.GetCsvUtilMgr().LoadCsv("Role", ConfigRoleMap)
	return
}

func GetRoleConfig(roleId int) *ConfigRole {
	return ConfigRoleMap[roleId]
}
