package msg

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"net/http"
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

func SendMsg(token, title, content string) {
	url := "http://www.pushplus.plus/send"
	bodyData := map[string]interface{}{
		"token":    token,
		"template": "markdown",
		"title":    title,
		"content":  content,
	}
	marshal, err := json.Marshal(bodyData)
	if err != nil {
		log.Fatal(err.Error())
	}
	response := req("post", url, string(marshal))
	fmt.Println(response)
}
