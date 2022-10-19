package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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

func hello(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Println("Read failed:", err)
		}
		defer req.Body.Close()
		log.Println("b:", string(b))
		task := PullTask{
			ShardNo: 1,
			File:    "1111111",
			TaskId:  11,
		}
		response := HttpResponse{
			Code:    0,
			Message: "success",
			Data:    task,
		}
		bytes, err := json.Marshal(response)
		w.Write(bytes)
	}

}

func main() {

	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":8091", nil)
}
