package tests

import (
	"net/url"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/kokeroulis/modip/types"
	_"fmt"
)

func teacherLogin(username string, password string, result interface{}, expected interface{}) {
	form := url.Values{}
	form.Add("username", username)
	form.Add("password", password)

	PostToJson("http://localhost:3001/teacher/login", form, result)

	So(result, ShouldResemble, expected)
}

func teacherLoginOk() {
	t := types.Teacher{1, "superteacher1", "superteacher1@teilar.gr",
					   types.Department{1, "T.P.T."}}

	expected := &types.TeacherJson{
		types.CommonJson{
			types.AuthJson{1, "superteacher1@teilar.gr"},
			types.ErrorJson{},
		},
		t,
	}

	result := &types.TeacherJson{}
	teacherLogin("superteacher1@teilar.gr", "superteacher1", result, expected)
}

func teacherLoginFail() {
	t := types.Teacher{}

	expected := &types.TeacherJson{
		types.CommonJson{
			types.AuthJson{},
			types.AuthFailed(),
		},
		t,
	}

	result := &types.TeacherJson{}
	teacherLogin("invalid", "invalid", result, expected)
}

