package model

type Result struct {
	UserName    string // 用户名
	CheckedIn   bool   // 是否签到
	IncrPoint   int64  // 签到结果
	SumPoint    int64  //矿石总量
	ContCount   int64  // 连续签到天数
	SumCount    int64  // 累计签到天数
	DippedLucky bool   // 是否粘喜气
	DipValue    int64  // 幸运值
	LuckyValue  int64  // 总幸运值
	FreeCount   int    // 免费签到次数
	FreeDrawed  bool   // 是否免费签到
	LotteryName string // 奖品名称
}
