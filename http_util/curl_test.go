package http_util

import (
	"fmt"
	"testing"
)

func TestFileCache(t *testing.T) {
	//get 请求
	params := map[string]interface{}{
		"page": 1,
		"size": 4,
		"wb":   "测试数据",
	}
	reqParam := &ReqParams{
		Url:    "http://localhost:3000/api/testArray",
		Method: "GET",
		Header: "json",
		Params: params,
	}
	req, err := reqParam.InitRequest()
	if err != nil {
		fmt.Println("request err", err)
		return
	}
	byteData, _ := req.Do()
	fmt.Printf("result is %v", string(byteData))
}
