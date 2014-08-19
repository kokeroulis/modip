package tests

import (
	"net/http"
	"net/url"
	"net/http/cookiejar"
	"io/ioutil"
	"encoding/json"
	"types"
)

func testTeacher() types.Teacher {
	return types.Teacher{1, "superteacher1", "superteacher1@teilar.gr",
						 types.Department{1, "T.P.T."}}
}

func testAuthJson() types.AuthJson {
	return types.AuthJson{1, "superteacher1@teilar.gr"}
}

func testOkCommonJson() types.CommonJson {
	return types.CommonJson{testAuthJson(), types.ErrorJson{}}
}

func NewRequest(client *http.Client, path string, form url.Values) []byte {
	var resp *http.Response
	var err error

	if len(form) == 0 {
		resp, err = client.Get(path);
	} else {
		resp, err = client.PostForm(path, form)
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
	return NewRequest(&http.Client{}, path, form)
}

func GetToJson(path string, v interface{}) {
	body := Get(path)
	err := json.Unmarshal(body, v)
	if err != nil {
		panic(err)
	}
}

func PostToJson(path string, form url.Values, v interface{}) {
	body := NewRequest(&http.Client{}, path, form)
	err := json.Unmarshal(body, v)
	if err != nil {
		panic(err)
	}
}

func loginAsTeacher(client *http.Client) {
	f := url.Values{}
	f.Add("username", "superteacher1@teilar.gr")
	f.Add("password", "superteacher1")

	_ = NewRequest(client, "http://localhost:3001/teacher/login", f)
}

func PostToJsonAsTeacher(path string, form url.Values, v interface{}) {
	cookieJar, _ := cookiejar.New(nil)
    client := &http.Client{Jar: cookieJar}

	loginAsTeacher(client)

	body := NewRequest(client, path, form)
	err := json.Unmarshal(body, v)
	if err != nil {
		panic(err)
	}
}

func GetToJsonAsTeacher(path string, v interface{}) {
	cookieJar, _ := cookiejar.New(nil)
    client := &http.Client{Jar: cookieJar}

	loginAsTeacher(client)

	body := NewRequest(client, path, url.Values{})
	err := json.Unmarshal(body, v)
	if err != nil {
		panic(err)
	}
}

