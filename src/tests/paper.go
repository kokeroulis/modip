package tests

import (
	"net/url"
	. "github.com/smartystreets/goconvey/convey"
	"types"
)

const paperTitle = "My Super paper"

func addPaper(title string, result interface{}, expected interface{}) {
	form := url.Values{}
	form.Add("title", title)

	PostToJsonAsTeacher("http://localhost:3001/paper/add", form, result)

	So(result, ShouldResemble, expected)
}

func addPaperOk() {
	p := types.Paper{1, paperTitle, testTeacher()}

	expected := &types.PaperJson{testOkCommonJson(), p}

	result := &types.PaperJson{}
	addPaper(paperTitle, result, expected)
}

func addPaperFail() {

	expected := &types.PaperJson{
		types.CommonJson{testAuthJson(), types.AlreadyExists()},
		types.Paper{},
	}

	result := &types.PaperJson{}
	addPaper(paperTitle, result, expected)
}

