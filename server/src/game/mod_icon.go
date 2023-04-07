package game

import (
	"demoProject_gameServer/server/src/csvs"
	"fmt"
)

type Icon struct {
	IconId int
}

type ModIcon struct {
	IconInfo map[int]*Icon
}

func (self *ModIcon) HasIcon(iconId int) bool {
	_, ok := self.IconInfo[iconId]
	if !ok {
		return false
	}
	return true
}

func (self *ModIcon) AddItem(itemId int) {
	_, ok := self.IconInfo[itemId]
	if ok {
		fmt.Println("当前头像已存在, 请勿重复添加！ 头像ID：", itemId)
		return
	}
	config := csvs.GetIconConfig(itemId)
	if config == nil {
		fmt.Println("非法头像，Id：", itemId)
		return
	}
	//赋值对象
	self.IconInfo[itemId] = &Icon{IconId: itemId}
}

func (self *ModIcon) CheckGetIcon(roleId int) {
	config := csvs.GetIconConfigByRoleId(roleId)
	if config == nil {
		return
	}
	self.AddItem(config.IconId)
}
