package tests

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"fmt"
)

func TestTeacher(t *testing.T) {
    Convey("Teacher", t, func() {
        Convey("Auth", func() {
			fmt.Println("here")
			Convey("Should succeed", func() { teacherLogin() })
		})
    })
}

