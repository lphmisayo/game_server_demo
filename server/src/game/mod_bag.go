package game

import (
	"demoProject_gameServer/server/src/csvs"
	"fmt"
)

type ItemInfo struct {
	ItemId  int
	ItemNum int64
}

type ModBag struct {
	BagInfo map[int]*ItemInfo
}

func (self *ModBag) AddItem(itemId int, num int64, player *Player) {
	itemConfig := csvs.GetItemConfigs(itemId)
	if itemConfig == nil {
		fmt.Println(itemId, "物品不存在")
		return
	}

	switch itemConfig.SortType {
	/*case csvs.ITEMTYPE_NORMAL:
	fmt.Println("普通物品：", itemConfig.ItemName)
	self.AddItemToBag(itemId, num)
	*/
	case csvs.ITEMTYPE_ROLE:
		fmt.Println("角色：", itemConfig.ItemName)
		player.ModRole.AddRole(itemId, num, player)

	case csvs.ITEMTYPE_ICON:
		player.ModIcon.AddItem(itemId)
		//fmt.Println("头像：", itemConfig.ItemName)

	case csvs.ITEMTYPE_CARD:
		player.ModCard.AddItem(itemId, 10)
		//fmt.Println("名片：", itemConfig.ItemName)
	default: //同普通
		//self.AddItemToBag(itemId, 1)
		self.AddItemToBag(itemId, num)
	}
}

func (self *ModBag) RemoveItem(itemId int, player *Player) {
	config := csvs.GetItemConfigs(itemId)
	if config == nil {
		fmt.Println("物品不存在！")
		return
	}
	switch config.SortType {
	case csvs.ITEMTYPE_NORMAL:
		self.RemoveItemFromBagGM(itemId, 2000)
	}
}

func (self *ModBag) AddItemToBag(itemId int, num int64) {
	_, ok := self.BagInfo[itemId]
	if ok {
		self.BagInfo[itemId].ItemNum += num
	} else {
		self.BagInfo[itemId] = &ItemInfo{ItemId: itemId, ItemNum: num}
	}
	config := csvs.GetItemConfigs(itemId)
	if config != nil {
		fmt.Println("获得物品：", config.ItemName, "获得数量：", num, "当前数量：", self.BagInfo[itemId].ItemNum)
	}

}

func (self *ModBag) RemoveItemFromBag(itemId int, num int64, player *Player) {

	if !self.HasEnoughItem(itemId, num) {
		config := csvs.GetItemConfigs(itemId)
		nowNum := int64(0)
		if config != nil {
			_, ok := self.BagInfo[itemId]
			if ok {
				nowNum = self.BagInfo[itemId].ItemNum
			}
		}
		fmt.Println("物品名称：", config.ItemName, "数量不足---", "当前剩余数量：", nowNum)

		return
	}

	_, ok := self.BagInfo[itemId]
	if ok {
		self.BagInfo[itemId].ItemNum -= num
	}
	config := csvs.GetItemConfigs(itemId)
	if config != nil {
		fmt.Println("减少物品名称：", config.ItemName, "减少数量：", num, "当前剩余数量：", self.BagInfo[itemId].ItemNum)
	}

}

func (self *ModBag) RemoveItemFromBagGM(itemId int, num int64) {

	_, ok := self.BagInfo[itemId]
	if ok {
		self.BagInfo[itemId].ItemNum -= num
	} else {
		self.BagInfo[itemId] = &ItemInfo{ItemId: itemId, ItemNum: -num}
	}
	config := csvs.GetItemConfigs(itemId)
	if config != nil {
		fmt.Println("减少物品名称：", config.ItemName, "减少数量：", num, "当前剩余数量：", self.BagInfo[itemId].ItemNum)
	}

}

func (self *ModBag) HasEnoughItem(itemId int, num int64) bool {
	_, ok := self.BagInfo[itemId]
	if !ok {
		return false
	} else if self.BagInfo[itemId].ItemNum < num {
		return false
	}
	return true
}
