package main

import (
	"demoProject_gameServer/server/src/csvs"
	"demoProject_gameServer/server/src/game"
	"fmt"
)

type Player game.Player

func main() {

	/**
	需求分析：

	//1 UID
	//2 头像，名片
	//3 签名
	//4 名字
	//5 冒险等级，冒险阅历
	//6 世界等级，恢复时间
	//7 生日
	//8 展示阵容，展示名片

	当前模块：背包 增加角色模块
	//1 物品识别
	//2 物品增加
	//3 物品消耗
	//4 物品使用
	//5 角色模块 --》头像模块
	*/

	/**
	设计顺序思路：
		客户端 --> 实体模块 --> 模组模块 --> 服务端
	将客户端和服务端之间尽量分离

	*/
	//******************************************************
	//加载配置
	csvs.CheckLodingCsv()

	//******************************************************
	//启用公共模块
	fmt.Println("数据测试------start")
	go game.GetManageBanWord().Run()

	playerTest := game.NewTestPlayer()
	go playerTest.Run()
	for {

	}
	return
	//playerTest.ModBag.AddItem(1000003, playerTest)
	//playerTest.ModBag.RemoveItem(1000003, playerTest)
	//playerTest.ModBag.RemoveItem(1000003, playerTest)
	//playerTest.ModBag.RemoveItem(1000003, playerTest)
	//playerTest.ModBag.RemoveItem(1000003, playerTest)
	//playerTest.ModBag.RemoveItem(1000003, playerTest)
	//playerTest.ModPlayer.SetCard(4000014, playerTest)
	//playerTest.ModBag.AddItem(4000014, playerTest)
	//playerTest.ModPlayer.SetCard(4000014, playerTest)

	/*playerTest.ModPlayer.SetIcon(3000001, playerTest)
	playerTest.ModBag.AddItem(3000001, playerTest)
	playerTest.ModPlayer.SetIcon(3000001, playerTest)*/
	//playerTest.ModBag.AddItem(1000003, playerTest)
	//playerTest.ModBag.AddItem(13000003)
	//playerTest.ModBag.AddItem(2000004)
	//playerTest.ModBag.AddItem(3000004)
	//playerTest.ModBag.AddItem(4000004)

	/*ticker := time.NewTicker(time.Second * 10)

	for {
		select {
		case <-ticker.C:
			playerTest := game.NewTestPlayer()
			go playerTest.Run()
		}
	}
	return*/

	//playerGM := game.NewTestPlayer()
	/*playerGM.RecvSetShowTeam([]int{1001, 1002, 1003, 1001})
	playerGM.RecvSetHideShowTeam(0)
	playerGM.RecvSetHideShowTeam(10)*/
	//playerGM.RecvSetShowCard([]int{101, 102, 103, 14, 101})
	//playerGM.RecvSetShowCard([]int{})
	//playerGM.RecvSetShowCard([]int{101})
	//playerGM.RecvSetIcon(1)
	//playerGM.ModPlayer.AddExp(10000000, playerGM)
	//playerGM.RecvSetBirth(3001)
	//playerGM.RecvSetBirth(1235)
	//playerGM.RecvSetBirth(230)
	//playerGM.RecvSetBirth(104)
	//playerGM.RecvSetBirth(105)
	//ticker := time.NewTicker(time.Second * 1)
	/*tickerIn := time.NewTicker(time.Second * 3)
	tickerOut := time.NewTicker(time.Second * 3)*/

	/*for {
		select {
		case <-ticker.C:
			if time.Now().Unix()%3 == 0 {
				playerGM.ReduceWorldLevel()
			}
			if time.Now().Unix()%5 == 0 {
				playerGM.RestoreWorldLevel()
			}
			//playerGM.ModPlayer.AddExp(5000, playerGM)
			//case <-tickerIn.C:
			//	player.RecvSetIcon(int(time.Now().Unix()))
			//
			//case <-tickerOut.C:
			//	player.RecvSetName("加微信买外挂")
		}
	}*/
	//******************************************************
	//监听客户端窗口

	//测试学习协程锁知识，要能理解map的读写安全问题，同时也要能理解
	//go playerSet(playerGM)
	//go playerGet(playerGM)

}

// 以下两种加锁的方式，在实际业务中会导致时间损耗高达2至3倍，若每个协程中都加一把锁，100个人就100把锁，效率损耗十分恐怖
// 解决方法可以是，通过设计理念来规避同时读写的操作发生，比如在主线程中限定读协程和写协程的操作，
// 而不是放到协程中去处理，这样可以极大降低协程锁的使用率，从而达到用最小的代价去完成数据安全
func playerSet(player *game.Player) {
	for i := 0; i < 1000000; i++ {
		player.ModUniqueQuest.Locker.Lock()
		player.ModUniqueQuest.QuestInfo[10001] = new(game.QuestInfo)
		player.ModUniqueQuest.Locker.Unlock()
	}
}

func playerGet(player *game.Player) {
	for i := 0; i < 1000000; i++ {
		player.ModUniqueQuest.Locker.RLock()
		_, ok := player.ModUniqueQuest.QuestInfo[10001]
		if ok {

		}
		player.ModUniqueQuest.Locker.RUnlock()
	}
}
