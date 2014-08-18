package tests

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTeacher(t *testing.T) {
    Convey("Teacher", t, func() {
        Convey("Auth", func() {
			Convey("Should succeed", func() { teacherLoginOk() })
			Convey("Should Fail", func() { teacherLoginFail() })
		})
    })
}

