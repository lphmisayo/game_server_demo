package game

import (
	"demoProject_gameServer/server/src/csvs"
	"fmt"
	"regexp"
	"time"
)

var manageBanWord *ManageBanWord

type ManageBanWord struct {
	BanWordBase  []string //基础配置
	BanWordExtra []string //后期更新升级
}

func GetManageBanWord() *ManageBanWord {
	if nil == manageBanWord {
		manageBanWord = new(ManageBanWord)
		manageBanWord.BanWordBase = []string{"草", "外挂", "工具", "+v"}
		manageBanWord.BanWordExtra = []string{"❤"}
	}
	return manageBanWord
}

func (self *ManageBanWord) IsBanWord(txt string) bool {
	for _, v := range self.BanWordBase {
		match, _ := regexp.MatchString(v, txt)
		if match {
			fmt.Println("基础词库中包含违禁词，该违禁词为：", v)
			return match
		}
	}
	for _, v := range self.BanWordExtra {
		match, _ := regexp.MatchString(v, txt)
		if match {
			fmt.Println("拓展词库中包含违禁词，该违禁词为：", v)
			return match
		}
	}

	return false
}

func (self *ManageBanWord) Run() {
	self.BanWordBase = csvs.GetBaseBanWord()
	ticker := time.NewTicker(time.Second * 20)
	for {
		select {
		case <-ticker.C:
			if time.Now().Unix()%5 == 0 {
				fmt.Println("更新词库")
			} else {
				//fmt.Println("待机")
			}
		}
	}
}
