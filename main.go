package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"jj/core/config"
	"jj/core/model"
	"jj/core/msg"
	"jj/core/requests"
	"strings"
	"time"
)

func Task() {
	for _, user := range config.GetConfig().User {
		var result model.Result
		result.UserName = requests.GetUser(user.Cookie)
		result.CheckedIn = requests.GetTodyStatus(user.Cookie)
		if result.CheckedIn == false {
			incr_point := requests.CheckIn(user.Cookie)
			result.CheckedIn = true
			result.IncrPoint = incr_point
		}
		counts := requests.GetCounts(user.Cookie)
		result.ContCount = counts.Get("data").Get("cont_count").Int()
		result.SumCount = counts.Get("data").Get("sum_count").Int()
		//lottery_history_id 幸运值
		lotteries := requests.GetLotteryHistory(user.Cookie)
		history_id := lotteries.Array()[0].Get("history_id").String()
		dipData := requests.DipLucky(user.Cookie, history_id)
		result.DippedLucky = dipData.Get("has_dip").Bool()
		result.DipValue = dipData.Get("dip_value").Int()
		result.LuckyValue = dipData.Get("total_value").Int()
		//freeCount
		freeCount := requests.GetLotteryConfig(user.Cookie)
		if freeCount > 0 {
			lotteryName := requests.DrawLottery(user.Cookie)
			result.FreeDrawed = true
			result.LotteryName = lotteryName
		}
		//get cur sumpoint
		sumPoint := requests.GetCurrentPoint(user.Cookie)
		result.SumPoint = sumPoint
		// send msg
		content := make([]string, 0)
		content = append(content, fmt.Sprintf("hello %s \n", result.UserName))
		if result.CheckedIn {
			content = append(content, fmt.Sprintf("今日签到 + %d 矿石 \n", result.IncrPoint))
		} else {
			content = append(content, "已经签到 \n")
		}
		content = append(content, fmt.Sprintf("当前矿石数量 %d \n", result.SumPoint))
		content = append(content, fmt.Sprintf("连续签到天数 %d \n", result.ContCount))
		content = append(content, fmt.Sprintf("累计签到天数 %d \n", result.SumCount))
		if result.DippedLucky {
			content = append(content, fmt.Sprintf("沾喜气 + %d \n", result.DipValue))
		} else {
			content = append(content, "今日已经沾过喜气 \n")
		}
		content = append(content, fmt.Sprintf("当前幸运值 %d \n", result.LuckyValue))
		content = append(content, fmt.Sprintf("免费抽奖次数 %d \n", result.FreeCount))
		if result.FreeDrawed {
			content = append(content, fmt.Sprintf("恭喜抽中 %s \n", result.LotteryName))
		} else {
			content = append(content, "今日已免费抽奖 \n")
		}
		msg.SendMsg(strings.Join(content, "\n"))
	}
}
func main() {
	config.Init()
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	s := gocron.NewScheduler(timezone)
	s.Every(1).Days().At(config.GetConfig().Time).Do(Task)
	s.StartAsync()
	s.StartBlocking()
}
