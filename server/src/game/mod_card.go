package game

import (
	"demoProject_gameServer/server/src/csvs"
	"fmt"
)

type Card struct {
	cardId int
}

type ModCard struct {
	CardInfo map[int]*Card
}

func (self *ModCard) HasCard(cardId int) bool {
	_, ok := self.CardInfo[cardId]
	return ok
}

func (self *ModCard) AddItem(cardId int, friendliness int) {
	ok := self.HasCard(cardId)
	if ok {
		fmt.Println("已存在名片，ID：", cardId)
		return
	}
	config := csvs.GetConfigCard(cardId)
	if config == nil {
		fmt.Println("非法名片，ID：", cardId)
		return
	}
	if friendliness < config.Friendliness {
		fmt.Println("当前角色好感度不足，不可增加名片，名片ID：", cardId)
		return
	}

	self.CardInfo[cardId] = &Card{cardId: cardId}
	fmt.Println("增加名片成功，当前名片ID：", cardId)
}

func (self *ModCard) CheckGetCard(roleId int, friendliness int) {
	config := csvs.GetConfigCardByRoleId(roleId)
	if config == nil {
		fmt.Println("当前用户已拥有此名片")
		return
	}
	fmt.Println(config.CardId, ":::", config.Check, "::: ", friendliness)
	self.AddItem(config.CardId, friendliness)
}
