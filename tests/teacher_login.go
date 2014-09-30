package tests

import (
	"github.com/kokeroulis/modip/types"
	"net/url"
)

func teacherLogin(username string, password string, expected types.JsonData) {
	form := url.Values{}
	form.Add("username", username)
	form.Add("password", password)

	result := PostToJson("http://localhost:3001/teacher/login", form)

	CompareJson(result, expected)
}

func teacherLoginOk() {
	t := types.Teacher{1, "superteacher1", "superteacher1@teilar.gr",
		types.Department{1, "T.P.T."}}

	expected := types.JsonData{
		Common: types.CommonJson{types.AuthJson{1, "superteacher1@teilar.gr"}, types.ErrorJson{}},
		Data:   t,
	}

	teacherLogin("superteacher1@teilar.gr", "superteacher1", expected)
}

func teacherLoginFail() {
	t := types.Teacher{}

	expected := types.JsonData{
		Common: types.CommonJson{types.AuthJson{},types.AuthFailed()},
		Data:   t,
	}

	teacherLogin("invalid", "invalid", expected)
}
