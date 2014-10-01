package tests

import (
	"github.com/kokeroulis/modip/models"
	"net/url"
)

func teacherLogin(username string, password string, expected models.JsonData) {
	form := url.Values{}
	form.Add("username", username)
	form.Add("password", password)

	result := PostToJson("http://localhost:3001/teacher/login", form)

	CompareJson(result, expected)
}

func teacherLoginOk() {
	t := models.Teacher{1, "superteacher1", "superteacher1@teilar.gr",
		models.Department{1, "T.P.T."}}

	expected := models.JsonData{
		Common: models.CommonJson{models.AuthJson{1, "superteacher1@teilar.gr"}, models.ErrorJson{}},
		Data:   t,
	}

	teacherLogin("superteacher1@teilar.gr", "superteacher1", expected)
}

func teacherLoginFail() {
	t := models.Teacher{}

	expected := models.JsonData{
		Common: models.CommonJson{models.AuthJson{},models.AuthFailed()},
		Data:   t,
	}

	teacherLogin("invalid", "invalid", expected)
}
