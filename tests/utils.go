package tests

import (
	"encoding/json"
	"github.com/kokeroulis/modip/types"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
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

func testTeacherOkJson(bar interface{}) types.JsonData {

	common := types.CommonJson{testAuthJson(), types.ErrorJson{}}
	jsonData := types.JsonData{
		Common: common,
		Data:   bar,
	}

	return jsonData
}

func NewRequest(client *http.Client, path string, form url.Values) []byte {
	var resp *http.Response
	var err error

	if len(form) == 0 {
		resp, err = client.Get(path)
	} else {
		resp, err = client.PostForm(path, form)
	}

	if err != nil {
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

func GetToJson(path string) string {
	body := Get(path)
	return string(body)
}

func PostToJson(path string, form url.Values) string {
	body := NewRequest(&http.Client{}, path, form)
	return string(body)
}

func loginAsTeacher(client *http.Client) {
	f := url.Values{}
	f.Add("username", "superteacher1@teilar.gr")
	f.Add("password", "superteacher1")

	_ = NewRequest(client, "http://localhost:3001/teacher/login", f)
}

func PostToJsonAsTeacher(path string, form url.Values) string {
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: cookieJar}

	loginAsTeacher(client)

	body := NewRequest(client, path, form)

	return string(body)
}

func GetToJsonAsTeacher(path string) string {
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: cookieJar}

	loginAsTeacher(client)

	body := NewRequest(client, path, url.Values{})
	return string(body)
}

func CompareJson(result string, expected types.JsonData) {
	expectedJson, err := json.Marshal(expected)

	if err != nil {
		panic(err)
	}

	So(result, ShouldResemble, string(expectedJson))
}

