package tests

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
)

func NewRequest(url string, form url.Values) []byte {

	client := http.Client{}
	var resp *http.Response
	var err error

	if len(form) == 0 {
		resp, err = client.Get(url);
	} else {
		resp, err = client.PostForm(url, form)
	}

	if (err != nil) {
		panic(err)
	}

	defer resp.Body.Close()
	body, bodyErr := ioutil.ReadAll(resp.Body)

	if bodyErr != nil {
		panic(bodyErr)
	}

	return body
}

func Get(path string) []byte {
	form := url.Values{}
	return NewRequest(path, form)
}

func GetToJson(path string, v interface{}) {
	body := Get(path)
	err := json.Unmarshal(body, v)
	if err != nil {
		panic(err)
	}
}

func PostToJson(path string, form url.Values, v interface{}) {
	body := NewRequest(path, form)
	err := json.Unmarshal(body, v)
	if err != nil {
		panic(err)
	}
}

