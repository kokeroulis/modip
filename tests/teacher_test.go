package tests

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
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

		Convey("Add Book", func() {
			Convey("Should succeed", func() { addBookOk() })
			Convey("Should Fail", func() { addBookFail() })
		})

		Convey("Info", func() {
			Convey("Should succeed", func() { teacherInfoOk() })
		})

	})
}
