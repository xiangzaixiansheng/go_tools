package http_util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	XMLHeader  = "xml"
	JSONHeader = "json"
)

const (
	POST = "POST"
	GET  = "GET"
)

type Curl interface {
	SetHeader(key, val string) //设置header头
	Do() ([]byte, error)       //执行请求
}

//请求对象
type ReqParams struct {
	Url    string                 //地址
	Method string                 //请求方法
	Header string                 //请求头 JSON或者XML
	Params map[string]interface{} //请求参数 需要decode
}

type reqObj struct {
	req *http.Request
}

//初始请求参数
func (p *ReqParams) InitRequest() (req Curl, err error) {
	var reqParams *bytes.Reader
	obj := new(reqObj)

	if p.Params != nil {
		if p.Method == GET {
			obj.req, err = http.NewRequest(p.Method, p.Url, nil)
			q := obj.req.URL.Query()
			for k, v := range p.Params {
				q.Add(k, fmt.Sprintf("%v", v))
			}
			obj.req.URL.RawQuery = q.Encode()
			fmt.Println(obj.req.URL.String())
			fmt.Println("come in===== get")
		} else {
			//post 请求.
			body, _ := json.Marshal(p.Params)
			reqParams = bytes.NewReader(body)
			obj.req, err = http.NewRequest(p.Method, p.Url, reqParams)
			fmt.Println("come in===== post")
		}
	} else {
		obj.req, err = http.NewRequest(p.Method, p.Url, nil)
	}

	if err != nil {
		return nil, err
	}
	if p.Method == POST {
		switch p.Header {
		case JSONHeader:
			obj.req.Header.Set("Content-Type", "application/json;charset=UTF-8")
			break
		case XMLHeader:
			obj.req.Header.Set("Accept", "application/xml")
			obj.req.Header.Set("Content-Status", "application/xml;charset=utf-8")
			break
		default:
			obj.req.Header.Set("Content-Type", "application/json;charset=UTF-8")
		}
	}

	return obj, nil
}

//设置header头
func (obj *reqObj) SetHeader(key, val string) {
	obj.req.Header.Set(key, val)
}

//执行请求
func (obj *reqObj) Do() ([]byte, error) {
	defer func() {
		if er := recover(); er != nil {
			fmt.Print(fmt.Errorf("%v", er))
		}
	}()
	c := http.Client{}
	resp, err := c.Do(obj.req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
