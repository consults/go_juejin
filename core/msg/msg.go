package msg

import (
	"io/ioutil"
	"jj/core/config"
	"log"
	"net/http"
	"net/url"
)

func SendMsg(content string) {
	client := &http.Client{}
	query := url.Values{}
	query.Set("content", content)
	c := config.GetConfig()
	reqUrl := c.MsgUrl + "&" + query.Encode()
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
	log.Printf("%s \n", string(bodyText))
}
