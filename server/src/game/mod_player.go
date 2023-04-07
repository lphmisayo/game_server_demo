package game

import (
	"demoProject_gameServer/server/src/csvs"
	"fmt"
	"time"
)

type ShowRole struct {
	RoleId    int
	RoleLevel int
}

type ModPlayer struct {
	//显示字段
	UserId             int
	Icon               int
	Card               int
	Name               string
	Sign               string
	PlayerLevel        int
	PlayerExp          int
	WorldLevel         int
	WorldLevelNow      int
	WorldLevelColdTime int64
	Birth              int
	ShowTeam           []*ShowRole //展示阵容
	HideShowTeam       int         //隐藏开关
	ShowCard           []int
	//看不见的字段
	Prohibit int
	IsGM     int
}

// SetIcon 修改头像值，模块交互服务器
func (self *ModPlayer) SetIcon(iconId int, player *Player) {

	if !player.ModIcon.HasIcon(iconId) {
		fmt.Println("修改失败，该用户无此头像")
		return
	}
	player.ModPlayer.Icon = iconId
	fmt.Println("修改成功，当前头像为：", iconId)
}

func (self *ModPlayer) SetCard(cardId int, player *Player) {
	if !player.ModCard.HasCard(cardId) {
		fmt.Println("修改失败，该用户无此卡片")
		return
	}
	player.ModPlayer.Card = cardId
	fmt.Println("修改成功，当前展示的卡片为：", cardId)
}

func (self *ModPlayer) SetName(name string, player *Player) {

	if GetManageBanWord().IsBanWord(name) {
		return
	}
	player.ModPlayer.Name = name
	fmt.Println("当前名字为：", player.ModPlayer.Name)

}

func (self *ModPlayer) SetSign(sign string, player *Player) {

	if GetManageBanWord().IsBanWord(sign) {
		return
	}
	player.ModPlayer.Sign = sign
	fmt.Println("当前签名为：", player.ModPlayer.Sign)
}

func (self *ModPlayer) AddExp(exp int, player *Player) {
	self.PlayerExp += exp

	for {
		config := csvs.GetNowLevelConfig(self.PlayerLevel)
		if config == nil {
			break
		}
		if config.PlayerExp == 0 {
			break
		}
		//世界任务是否完成
		//isFinish := player.ModUniqueQuest.IsQuestFinish(config.ChapterId)
		//fmt.Println(isFinish)
		if config.ChapterId > 0 && !player.ModUniqueQuest.IsQuestFinish(config.ChapterId) {
			break
		}

		if self.PlayerExp >= config.PlayerExp {
			self.PlayerLevel += 1
			self.PlayerExp -= config.PlayerExp
		} else {
			break
		}
	}

	fmt.Println("当前等级：", self.PlayerLevel, "====当前经验值：", self.PlayerExp)

}

func (self *ModPlayer) ReduceWorldLevel(player *Player) {
	if self.WorldLevel < csvs.REDUCE_WORLD_LEVEL_START {
		fmt.Println("操作失败：", "当前等级---", self.PlayerLevel)
		return
	}

	if self.WorldLevel-self.WorldLevelNow >= csvs.REDUCE_WORLD_LEVEL_MAX {
		fmt.Println("操作失败：", "当前世界等级---", self.WorldLevelNow, "原世界等级---", self.WorldLevel)
		return
	}

	if time.Now().Unix() < self.WorldLevelColdTime {
		fmt.Println("操作失败：", "冷却中")
		return
	}

	self.WorldLevelNow -= 1
	self.WorldLevelColdTime = time.Now().Unix() + csvs.REDUCE_WORLD_COlDTIME
	fmt.Println("操作成功：", "当前世界等级---", self.WorldLevelNow, "原世界等级---", self.WorldLevel)
	return
}

func (self *ModPlayer) RestoreWorldLevel(player *Player) {
	if self.WorldLevelNow >= self.WorldLevel {
		fmt.Println("操作失败：", "当前世界等级---", self.WorldLevelNow, "原世界等级---", self.WorldLevel)
		return
	}

	if time.Now().Unix() < self.WorldLevelColdTime {
		fmt.Println("操作失败：", "冷却中")
		return
	}

	self.WorldLevelNow += 1
	self.WorldLevelColdTime = time.Now().Unix() + csvs.REDUCE_WORLD_COlDTIME
	fmt.Println("操作成功：", "当前世界等级---", self.WorldLevelNow, "原世界等级---", self.WorldLevel)
	return
}

func (self *ModPlayer) SetBirth(birth int, player *Player) {

	if self.Birth > 0 {
		fmt.Println("操作错误：", "当前账号生日已设置")
		return
	}

	month := birth / 100
	day := birth % 100

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		if day <= 0 || day > 31 {
			fmt.Println("操作错误：", "当前设置天数为--", day, "月数为---", month)
			return
		}
	case 4, 6, 9, 11:
		if day <= 0 || day > 30 {
			fmt.Println("操作错误：", "当前设置天数为--", day, "月数为---", month)
			return
		}
	case 2:
		if day <= 0 || day > 29 {
			fmt.Println("操作错误：", "当前设置天数为--", day, "月数为---", month)
			return
		}
	default:
		fmt.Println("操作错误：", "当前设置天数为--", day, "月数为---", month)
		return
	}

	self.Birth = birth
	fmt.Println("设置成功：", "当前生日为---", month, "月", day, "日")

	if self.IsBirthDay() {
		fmt.Println("今天是你的生日，生日快乐！")
	} else {
		fmt.Println("期待你生日的到来！")
	}
}

// go语言 time.Now() 取出来的时间，跟设备的时区是直接对应的
func (self *ModPlayer) IsBirthDay() bool {
	month := time.Now().Month()
	day := time.Now().Day()
	if int(month) == self.Birth/100 && day == self.Birth%100 {
		return true
	}
	return false
}

func (self *ModPlayer) SetShowCard(showCards []int, player *Player) {
	if len(showCards) > csvs.SHOW_SIZE || len(showCards) < csvs.SHOW_ERR_SIZE {
		return
	}

	existCards := make(map[int]int)
	newCardList := make([]int, 0)
	for _, cardId := range showCards {
		_, ok := existCards[cardId]
		if ok { //检查是否有重复值
			continue
		}
		if !player.ModCard.HasCard(cardId) { //检查用户是否拥有这张卡片
			return
		}
		newCardList = append(newCardList, cardId)
		existCards[cardId] = 1
	}
	self.ShowCard = newCardList
	fmt.Println("设置已完成，当前卡片为：", self.ShowCard)
}

func (self *ModPlayer) SetShowTeam(showRoleIds []int, player *Player) {

	if len(showRoleIds) > csvs.SHOW_SIZE || len(showRoleIds) < csvs.SHOW_ERR_SIZE {
		return
	}

	existRoles := make(map[int]int)
	newRoleList := make([]*ShowRole, 0)
	for _, roleId := range showRoleIds {
		_, ok := existRoles[roleId]
		if ok { //检查是否有重复值
			continue
		}
		if !player.ModRole.IsHasRole(roleId) { //检查用户是否拥有这张卡片
			return
		}
		showRole := new(ShowRole)
		showRole.RoleId = roleId
		showRole.RoleLevel = player.ModRole.GetRoleLevel(roleId)
		newRoleList = append(newRoleList, showRole)
		existRoles[roleId] = 1
	}
	self.ShowTeam = newRoleList
	fmt.Println("设置已完成，当前阵容为：")
	for _, role := range self.ShowTeam {
		fmt.Printf("id为：%d -- ", role.RoleId)
	}
	fmt.Println()
}

func (self *ModPlayer) SetShowHideTeam(isHide int, player *Player) {

	if isHide != csvs.LOGIC_TRUE && isHide != csvs.LOGIC_FALSE {
		fmt.Println("设置不展示阵容失败 ", "--当前值不规范，值为：", isHide)
		return
	}
	self.HideShowTeam = isHide
	fmt.Println("设置不展示阵容完成")
}

func (self *ModPlayer) SetProhibit(prohibit int) {
	self.Prohibit = prohibit
}

func (self *ModPlayer) SetIsGM(isGM int) {
	self.IsGM = isGM
}

func (self *ModPlayer) IsCanEnter() bool {
	return int64(self.Prohibit) < time.Now().Unix()
}
