package game

import (
	"demoProject_gameServer/server/src/csvs"
	"fmt"
)

type Role struct {
	RoleLevel int
	RoleId    int
	GetTimes  int
}

type ModRole struct {
	RoleMap map[int]*Role
}

func (self *ModRole) IsHasRole(roleId int) bool {
	_, ok := self.RoleMap[roleId]
	return ok
}

func (self *ModRole) GetRoleLevel(roleId int) int {
	//return self.RoleLevel
	return 60
}

func (self *ModRole) AddRole(roleId int, num int64, player *Player) {
	config := csvs.GetRoleConfig(roleId)
	if config == nil {
		fmt.Println("配置不存在！")
		return
	}
	for i := int64(0); i < num; i++ {
		_, ok := self.RoleMap[roleId]
		if !ok {
			data := new(Role)
			data.RoleId = roleId
			data.RoleLevel = 1
			data.GetTimes = 1
			self.RoleMap[roleId] = data
		} else {
			//判断实际获得的东西是否转换成材料
			self.RoleMap[roleId].GetTimes++
			if self.RoleMap[roleId].GetTimes >= csvs.ADD_ROLE_TIME_NORMAL_MIN &&
				self.RoleMap[roleId].GetTimes <= csvs.ADD_ROLE_TIME_NORMAL_MAX {
				player.ModBag.AddItem(config.Stuff, int64(config.StuffNum), player)
				player.ModBag.AddItemToBag(config.StuffItem, int64(config.StuffItemNum))
			} else {
				player.ModBag.AddItemToBag(config.StuffItem, int64(config.MaxStuffItemNum))
			}
		}
	}
	itemConfig := csvs.GetItemConfigs(roleId)
	if itemConfig != nil {
		fmt.Println("获得角色：", itemConfig.ItemName, "-------", self.RoleMap[roleId].GetTimes, "次")
	}
	player.ModIcon.CheckGetIcon(roleId)
	player.ModCard.CheckGetCard(roleId, 10)
}
