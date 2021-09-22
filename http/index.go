package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)
var SaveActionLogAns = false
var SaveActionDir = `/runtime`
var SaveActionLog = false


type HttpsClient struct {
	Client *http.Client
	Req   *http.Request
}
func (h *HttpsClient) Set (){
	SaveActionLogAns = true
	SaveActionDir = `/runtime`
	SaveActionLog = true
}
func (h *HttpsClient) Init (proxyUrl string){
	proxy, _ := url.Parse(proxyUrl)
	netTransport := &http.Transport{
		DisableKeepAlives: false,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 3 * time.Minute,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       10 * time.Minute,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		Proxy:                 http.ProxyURL(proxy),
	}
	client := &http.Client{
		Transport: netTransport,
	}
	h.Client = client
}
func (h *HttpsClient) Action(method, urls, action string, postData map[string]string, headers map[string]string, ret string) (by []byte, err error) {
	req:=h.Req
	if strings.Contains(action, "json") {
		buf, _ := json.Marshal(postData)
		req, err = http.NewRequest(method, urls, bytes.NewBuffer(buf))
	} else {
		val := url.Values{}
		for k, v := range postData {
			val.Add(k, v)
		}
		if strings.ToLower(method) == "post" {
			req, err = http.NewRequest(method, urls, strings.NewReader(val.Encode()))
		} else {
			req, err = http.NewRequest(method, urls+"?"+val.Encode(), strings.NewReader(val.Encode()))
		}
	}
	if err != nil {
		err = fmt.Errorf("http.NewRequest is fail: %v", err.Error())
		return
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	if h.Client == nil {
		h.Client = &http.Client{}
	}
	resp, err := h.Client.Do(req)
	if err != nil {
		err = fmt.Errorf("client.Do is fail: %v", err.Error())
		return
	}
	by, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("ioutil.ReadAll is fail: %v", err.Error())
		return
	}
	err = handErr(by)
	if SaveActionLog || err != nil {
		_, err1 := saveFile(urls, method, postData, by, ret, err)
		if err1 != nil {
			fmt.Println(err1.Error(), "err1.Error")
		}
	}
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New("HTTP CODE:" + fmt.Sprint(resp.StatusCode))
		return
	}
	return
}
func handErr(by []byte) (err error) {
	//不同情况下的错误处理
	return
}

func saveFile(urls, method string, postData map[string]string, by []byte, ret string, errs error) (fileName string, err error) {
	now := time.Now()
	logFilePath := ""
	if dir, err0 := os.Getwd(); err0 == nil {
		if err0 == nil {
			logFilePath = dir + SaveActionDir + "/http_log/" + now.Format("200601") + "/" + now.Format("02") + "/"
		} else {
			logFilePath = dir + SaveActionDir + "/http_log/" + now.Format("200601") + "/" + now.Format("02") + "_errs/"
		}
	}
	if _, err2 := os.Stat(logFilePath); os.IsNotExist(err2) {
		os.MkdirAll(logFilePath, 0777)
		os.Chmod(logFilePath, 0777)
	}
	{
		// 请求记录
		var saveObj []byte
		saveObj = append(saveObj, []byte(urls)...)
		saveObj = append(saveObj, '\n')
		saveObj = append(saveObj, []byte(method)...)
		saveObj = append(saveObj, '\n')
		buf, _ := json.Marshal(postData)
		saveObj = append(saveObj, buf...)
		saveObj = append(saveObj, '\n')
		if errs != nil {
			saveObj = append(saveObj, []byte(ret)...)
			saveObj = append(saveObj, '\n')
			saveObj = append(saveObj, []byte(errs.Error())...)
			saveObj = append(saveObj, '\n')
		}
		saveObj = append(saveObj, []byte("http_request_finished")...)
		saveObj = append(saveObj, '\n')

		logFileName := now.Format("1504") + ".log"
		fileName = logFilePath + logFileName
		fileObj, err1 := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err1 != nil {
			err = err1
			return
		}
		defer fileObj.Close()
		fileObj.Write(saveObj)
	}
	//保存文件结束！！
	if ret != "" && errs == nil && SaveActionLogAns {
		logFileName := now.Format("1504") + fmt.Sprintf("_%v", time.Now().UnixNano()) + "." + ret
		fileName = logFilePath + logFileName
		fileObj, err1 := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err1 != nil {
			err = err1
			return
		}
		defer fileObj.Close()
		fileObj.Write(by)
	}
	return
}
