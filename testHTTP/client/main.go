package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"time"
)

// ResponseStatus 响应公共字段
type HttpResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// @Description 任务拉取数据
type PullTask struct {
	ShardNo int    `json:"shard_no" valid:"Required;"` // 文件分片ID
	File    string `json:"file" valid:"Required;"`     // 所需计算的文件，是个压缩包
	TaskId  int    `json:"task_id" valid:"Required;"`  // 任务id
}

type postMsg struct {
	Ip string `json:"ip"`
}

const (
	ContentType      = "Content-Type"
	ApplicationJson  = "application/json"
	MutipartFormData = "multipart/form-data; boundary=<calculated when request is sent>"
)

var httpClient *resty.Client = nil

func NewHttpClient() *resty.Client {
	if httpClient != nil {
		return httpClient
	}
	httpClient = resty.New()
	httpClient.SetTimeout(5 * time.Second)
	httpClient.SetRetryCount(3)

	//debug := config.GetConfig("httpclient.debug")
	//if len(debug) > 0 && debug == "ON" {
	//	httpClient.SetDebug(true)
	//}

	return httpClient
}

func Post(baseURL string, url string, body any, data interface{}) (resp *resty.Response) {
	client := NewHttpClient()
	client.SetBaseURL(baseURL)
	respBody := HttpResponse{}
	req := client.R()
	resp, err := req.
		SetHeader(ContentType, ApplicationJson).
		SetBody(body).
		SetResult(respBody).
		Post(url)
	if err != nil {
		logrus.Errorf("request failed: %v \n", errors.WithStack(err))
		return
	}
	if resp.StatusCode() == 200 {
		logrus.Infof("request success, req: %+v, resp: %+v", resp.Request, resp)
		//json.Unmarshal(resp.Body(), &respBody)
		if respBody.Code == 0 {
			logrus.Infof("response success, resp body: %+v", respBody)

			bytes, _ := json.Marshal(respBody.Data)
			json.Unmarshal(bytes, &data)
			fmt.Printf("%T \n", data)
			fmt.Printf("%+v \n", data)
			// do something
			return

		} else {
			logrus.Errorf("error response: %+v", respBody)
			return
		}
	}
	logrus.Errorf("request failed, http resp: %+v", resp)
	return
}

func main() {
	reqBody := postMsg{Ip: "1.1.1.1"}
	task := PullTask{}
	resp := Post("http://localhost:8080", "/hello", reqBody, task)

	//resp, err := http.Get("https://www.baidu.com")
	//if err != nil {
	//	panic(err)
	//}
	fmt.Printf("%+v", resp)
}
