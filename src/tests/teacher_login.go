package tests

import (
	"net/url"
	. "github.com/smartystreets/goconvey/convey"
	"types"
)

func teacherLogin() {
	form := url.Values{}
	form.Add("username", "superteacher1@teilar.gr")
	form.Add("password", "superteacher1")

	result := types.TeacherJson{}

	PostToJson("http://localhost:3000/teacher/login", form, &result)

	expected := types.TeacherJson{1,
								  "superteacher1",
								  "superteacher1@teilar.gr",
								  types.DepartmentJson{1, "T.P.T."}}

	So(result, ShouldResemble, expected)
}

