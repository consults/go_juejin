package requests

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"net/http"
	"strings"
)

func req(method string, url string, cookie string, body string) *gjson.Result {
	method = strings.ToUpper(method)
	client := &http.Client{}
	var req *http.Request
	var err error
	if body != "" {
		req, err = http.NewRequest(method, url, strings.NewReader(body))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		log.Fatal(err.Error())
	}
	req.Header.Set("authority", "api.juejin.cn")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cookie", cookie)
	req.Header.Set("origin", "https://juejin.cn")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("referer", "https://juejin.cn/")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	response := gjson.ParseBytes(bodyText)
	return &response
}

func GetUser(cookie string) string {
	url := "https://api.juejin.cn/user_api/v1/user/get"
	response := req("get", url, cookie, "")
	userName := response.Get("data").Get("user_name")
	return userName.String()
}
func GetTodyStatus(cookie string) bool {
	url := "https://api.juejin.cn/growth_api/v1/get_today_status"
	response := req("get", url, cookie, "")
	checkStatus := response.Get("data").Bool()
	return checkStatus
}
func GetCounts(cookie string) *gjson.Result {
	// 签到天数
	url := "https://api.juejin.cn/growth_api/v1/get_counts"
	response := req("get", url, cookie, "")
	return response
}
func CheckIn(cookie string) int64 {
	url := "https://api.juejin.cn/growth_api/v1/check_in"
	response := req("post", url, cookie, "")
	incr_point := response.Get("data").Get("incr_point").Int()
	return incr_point
}
func GetLotteryHistory(cookie string) *gjson.Result {
	// 粘喜气
	url := "https://api.juejin.cn/growth_api/v1/lottery_history/global_big"
	body := `{"page_no":1,"page_size":5}`
	response := req("post", url, cookie, body)
	lotteries := response.Get("data").Get("lotteries")
	return &lotteries
}
func DipLucky(cookie string, history_id string) *gjson.Result {
	// 粘喜气
	url := "https://api.juejin.cn/growth_api/v1/lottery_lucky/dip_lucky"
	body := map[string]interface{}{"lottery_history_id": history_id}
	marshal, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err.Error())
	}
	response := req("post", url, cookie, string(marshal))
	data := response.Get("data")
	return &data
}
func GetLotteryConfig(cookie string) int64 {
	// 免费抽奖次数
	url := "https://api.juejin.cn/growth_api/v1/lottery_config/get"
	response := req("get", url, cookie, "")
	return response.Get("data").Get("free_count").Int()
}

func DrawLottery(cookie string) string {
	// 抽奖
	url := "https://api.juejin.cn/growth_api/v1/lottery/draw"
	response := req("post", url, cookie, "")
	return response.Get("data").Get("lottery_name").String()
}

func GetCurrentPoint(cookie string) int64 {
	// 当前矿石
	url := "https://api.juejin.cn/growth_api/v1/get_cur_point"
	response := req("get", url, cookie, "")
	return response.Get("data").Int()
}
func GetNotCollectBug(cookie string) *gjson.Result {
	// 当前矿石
	url := "https://api.juejin.cn/user_api/v1/bugfix/not_collect"
	response := req("post", url, cookie, "{}")
	data := response.Get("data")
	return &data
}
