package common

import (
	"net/http"
	"strings"
	"io/ioutil"
	"github.com/astaxie/beego/logs"
)

func GetRequest(url string) string {
	//http.NewRequest()
	response, err := http.Get(url)
	if err != nil {
		logs.Error(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	return string(body)
}

func GetRequestWithParams(url string, params map[string]string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logs.Error(err)
		//os.Exit(1)
		return nil, err
	}
	q := request.URL.Query()
	for k, v := range params {
		logs.Info("paramKey=%v, paramValue=%v\n", k, v)
		q.Set(k, v)
	}

	request.URL.RawQuery = q.Encode()

	logs.Info(request.URL.String())
	// Output:
	// http://api.themoviedb.org/3/tv/popular?another_thing=foo+%26+bar&api_key=key_from_environment_or_flag
	//var resp *http.Response
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	//logs.Info(json.NewDecoder(response.Body))
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return body, nil
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
