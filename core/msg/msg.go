package msg

import (
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func req(method string, url string, body string) *gjson.Result {
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

//func SendMsg(token, title, content string) {
//	url := "http://www.pushplus.plus/send"
//	bodyData := map[string]interface{}{
//		"token":    token,
//		"template": "markdown",
//		"title":    title,
//		"content":  content,
//	}
//	marshal, err := json.Marshal(bodyData)
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//	response := req("post", url, string(marshal))
//	fmt.Println(response)
//}

func SendMsg(content string) {
	client := &http.Client{}
	query := url.Values{}
	query.Set("content", content)
	reqUrl := "http://117.50.175.64:8080/send?key=1Wf1ruJexwfE&" + query.Encode()
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("%s \n", bodyText)
}
