package game

import "sync"

type QuestInfo struct { //任务状态
	QuestId int
	State   int
}

type ModUniqueQuest struct {
	QuestInfo map[int]*QuestInfo
	Locker    *sync.RWMutex
}

func (self *ModUniqueQuest) IsQuestFinish(questId int) bool {

	//测试用
	if questId == 10001 || questId == 10002 || questId == 10003 {
		return true
	}

	quest, ok := self.QuestInfo[questId]
	if !ok {
		return false
	}
	return quest.State == QUEST_STATE_FINISH //对比任务状态是否完成
}
