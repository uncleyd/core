package utils

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"core/config"
)

var (
	dingdingURL         = "https://oapi.dingtalk.com/robot/send?access_token=d06e56cf9a63ef8e6d6cf6b1ee28c8ec6a78c84038313d798840666a39b97150"
	cur_ding_web_server = ""
)

func init() {
	cur_ding_web_server = config.Get().WebServer
}
func DingPush(content string) error {
	if content != "" {
		formt := `
			{
				"msgtype": "markdown",
				"markdown": {
					"title":"日志出错记录",
					"text": "%s:%s"
				}
			}`
		body := fmt.Sprintf(formt, cur_ding_web_server, content)
		jsonValue := []byte(body)
		//发送消息到钉钉群使用webhook
		//钉钉文档 https://open-doc.dingtalk.com/docs/doc.htm?spm=a219a.7629140.0.0.karFPe&treeId=257&articleId=105735&docType=1
		resp, err := http.Post(dingdingURL, "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			return err
		}
		log.Println(resp)
	}
	return nil
}