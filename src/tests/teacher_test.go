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

		Convey("Add Paper", func() {
			Convey("Should succeed", func() { addPaperOk() })
			Convey("Should Fail", func() { addPaperFail() })
		})
    })
}

