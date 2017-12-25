package common

import (
	"net/http"
	"strings"
	"io/ioutil"
	"github.com/astaxie/beego/logs"
)

func GetRequest(url string, params map[string]string) string {
	//http.NewRequest()
	response, err := http.Get(url)
	if err != nil {
		logs.Error(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	return string(body)
}

func PostRequest(url string, params map[string]string) string {
	response, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	if err != nil {
		logs.Error(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logs.Error(err)
	}

	logs.Error(string(body))
	return string(body)
}
