package game

import (
	"demoProject_gameServer/server/src/csvs"
	"fmt"
	"sync"
)

const (
	QUEST_STATE_INIT   = 0
	QUEST_STATE_DOING  = 1
	QUEST_STATE_FINISH = 2
)

type Player struct {
	ModPlayer      *ModPlayer
	ModIcon        *ModIcon
	ModCard        *ModCard
	ModUniqueQuest *ModUniqueQuest
	ModRole        *ModRole
	ModBag         *ModBag
}

func NewTestPlayer() *Player {
	//*****************************
	//模块初始化
	player := new(Player)
	player.ModPlayer = new(ModPlayer)
	player.ModIcon = new(ModIcon)
	player.ModIcon.IconInfo = make(map[int]*Icon)
	player.ModCard = new(ModCard)
	player.ModCard.CardInfo = make(map[int]*Card)
	player.ModUniqueQuest = new(ModUniqueQuest)
	player.ModUniqueQuest.QuestInfo = make(map[int]*QuestInfo)
	player.ModUniqueQuest.Locker = new(sync.RWMutex)
	player.ModRole = new(ModRole)
	player.ModRole.RoleMap = make(map[int]*Role)
	player.ModBag = new(ModBag)
	player.ModBag.BagInfo = make(map[int]*ItemInfo)

	//******************************
	//参数初始化
	player.ModPlayer.PlayerLevel = 1
	player.ModPlayer.WorldLevelNow = 4
	player.ModPlayer.WorldLevel = 5

	//******************************
	return player
}

// RecvSetIcon 客户端交互模块，修改头像值
func (self *Player) RecvSetIcon(iconId int) {
	self.ModPlayer.SetIcon(iconId, self)
}

// RecvSetCard 客户端交互模块，修改卡片值
func (self *Player) RecvSetCard(cardId int) {
	self.ModPlayer.SetCard(cardId, self)
}

func (self *Player) RecvSetName(name string) {
	self.ModPlayer.SetName(name, self)
}

func (self *Player) RecvSetSign(sign string) {
	self.ModPlayer.SetSign(sign, self)
}

func (self *Player) ReduceWorldLevel() {
	self.ModPlayer.ReduceWorldLevel(self)
}

func (self *Player) RestoreWorldLevel() {
	self.ModPlayer.RestoreWorldLevel(self)
}

func (self *Player) RecvSetBirth(birth int) {
	self.ModPlayer.SetBirth(birth, self)
}

func (self *Player) RecvSetShowCard(showCard []int) {
	self.ModPlayer.SetShowCard(showCard, self)
}

func (self *Player) RecvSetShowTeam(showRole []int) {
	self.ModPlayer.SetShowTeam(showRole, self)
}

func (self *Player) RecvSetHideShowTeam(isHide int) {
	self.ModPlayer.SetShowHideTeam(isHide, self)
}

func (self *Player) Run() {
	fmt.Println("测试工具v0.1")
	fmt.Println("模拟用户创建成功OK------开始测试")
	fmt.Println("↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓")
	for {
		fmt.Println(self.ModPlayer.Name, ",欢迎来到提瓦特大陆,请选择功能：1基础信息2背包3地图(未开放)")
		var modChoose int
		fmt.Scan(&modChoose)
		switch modChoose {
		case 1:
			self.HandleBase()
		case 2:
			self.HandleBag()
		case 3:
			self.HandleMap()
		}
	}
}

// 基础信息
func (self *Player) HandleBase() {
	for {
		fmt.Println("当前处于基础信息界面,请选择操作：0返回1查询信息2设置名字3设置签名4头像5名片6设置生日")
		var action int
		fmt.Scan(&action)
		switch action {
		case 0:
			return
		case 1:
			self.HandleBaseGetInfo()
		case 2:
			self.HandleBagSetName()
		case 3:
			self.HandleBagSetSign()
		case 4:
			self.HandleBagSetIcon()
		case 5:
			self.HandleBagSetCard()
		case 6:
			self.HandleBagSetBirth()
		}
	}
}

func (self *Player) HandleBaseGetInfo() {
	fmt.Println("名字:", self.ModPlayer.Name)
	fmt.Println("等级:", self.ModPlayer.PlayerLevel)
	fmt.Println("大世界等级:", self.ModPlayer.WorldLevelNow)
	if self.ModPlayer.Sign == "" {
		fmt.Println("签名:", "未设置")
	} else {
		fmt.Println("签名:", self.ModPlayer.Sign)
	}

	if self.ModPlayer.Icon == 0 {
		fmt.Println("头像:", "未设置")
	} else {
		fmt.Println("头像:", csvs.GetItemConfigs(self.ModPlayer.Icon), self.ModPlayer.Icon)
	}

	if self.ModPlayer.Card == 0 {
		fmt.Println("名片:", "未设置")
	} else {
		fmt.Println("名片:", csvs.GetItemConfigs(self.ModPlayer.Card), self.ModPlayer.Card)
	}

	if self.ModPlayer.Birth == 0 {
		fmt.Println("生日:", "未设置")
	} else {
		fmt.Println("生日:", self.ModPlayer.Birth/100, "月", self.ModPlayer.Birth%100, "日")
	}
}

func (self *Player) HandleBagSetName() {
	fmt.Println("请输入名字:")
	var name string
	fmt.Scan(&name)
	self.RecvSetName(name)
}

func (self *Player) HandleBagSetSign() {
	fmt.Println("请输入签名:")
	var sign string
	fmt.Scan(&sign)
	self.RecvSetSign(sign)
}

func (self *Player) HandleBagSetIcon() {
	for {
		fmt.Println("当前处于基础信息--头像界面,请选择操作：0返回1查询头像背包2设置头像")
		var action int
		fmt.Scan(&action)
		switch action {
		case 0:
			return
		case 1:
			self.HandleBagSetIconGetInfo()
		case 2:
			self.HandleBagSetIconSet()
		}
	}
}

func (self *Player) HandleBagSetIconGetInfo() {
	fmt.Println("当前拥有头像如下:")
	for _, v := range self.ModIcon.IconInfo {
		config := csvs.GetItemConfigs(v.IconId)
		if config != nil {
			fmt.Println(config.ItemName, ":", config.ItemId)
		}
	}
}

func (self *Player) HandleBagSetIconSet() {
	fmt.Println("请输入头像id:")
	var icon int
	fmt.Scan(&icon)
	self.RecvSetIcon(icon)
}

func (self *Player) HandleBagSetCard() {
	for {
		fmt.Println("当前处于基础信息--名片界面,请选择操作：0返回1查询名片背包2设置名片")
		var action int
		fmt.Scan(&action)
		switch action {
		case 0:
			return
		case 1:
			self.HandleBagSetCardGetInfo()
		case 2:
			self.HandleBagSetCardSet()
		}
	}
}

func (self *Player) HandleBagSetCardGetInfo() {
	fmt.Println("当前拥有名片如下:")
	for _, v := range self.ModCard.CardInfo {
		//fmt.Println(v)
		config := csvs.GetItemConfigs(v.cardId)
		if config != nil {
			fmt.Println(config.ItemName, ":", config.ItemId)
		}
	}
}

func (self *Player) HandleBagSetCardSet() {
	fmt.Println("请输入名片id:")
	var card int
	fmt.Scan(&card)
	self.RecvSetCard(card)
}

func (self *Player) HandleBagSetBirth() {
	if self.ModPlayer.Birth > 0 {
		fmt.Println("已设置过生日!")
		return
	}
	fmt.Println("生日只能设置一次，请慎重填写,输入月:")
	var month, day int
	fmt.Scan(&month)
	fmt.Println("请输入日:")
	fmt.Scan(&day)
	self.ModPlayer.SetBirth(month*100+day, self)
}

// 背包
func (self *Player) HandleBag() {
	for {
		fmt.Println("当前处于基础信息界面,请选择操作：0返回1增加物品2扣除物品")
		var action int
		fmt.Scan(&action)
		switch action {
		case 0:
			return
		case 1:
			self.HandleBagAddItem()
		case 2:
			self.HandleBagRemoveItem()
		}
	}
}

func (self *Player) HandleBagAddItem() {
	itemId := 0
	itemNum := 0
	fmt.Println("物品ID")
	fmt.Scan(&itemId)
	fmt.Println("物品数量")
	fmt.Scan(&itemNum)
	self.ModBag.AddItem(itemId, int64(itemNum), self)
}

func (self *Player) HandleBagRemoveItem() {
	itemId := 0
	itemNum := 0
	fmt.Println("物品ID")
	fmt.Scan(&itemId)
	fmt.Println("物品数量")
	fmt.Scan(&itemNum)
	self.ModBag.RemoveItemFromBag(itemId, int64(itemNum), self)
}

// 地图
func (self *Player) HandleMap() {
	fmt.Println("向着星辰与深渊,欢迎来到冒险家协会！")
	fmt.Println("当前位置:", "蒙德城")
	fmt.Println("地图模块还没写到......")
}
