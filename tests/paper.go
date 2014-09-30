package tests

import (
	"github.com/kokeroulis/modip/types"
	"net/url"
)

const paperTitle = "My Super paper"

func addPaper(expected types.JsonData) {
	form := url.Values{}
	form.Add("title", paperTitle)

	result := PostToJsonAsTeacher("http://localhost:3001/paper/add", form)

	CompareJson(result, expected)
}

func addPaperOk() {
	p := types.Paper{1, paperTitle, testTeacher()}

	expected := testTeacherOkJson(p)

	addPaper(expected)
}

func addPaperFail() {

	expected := types.JsonData{
		Common: types.CommonJson{testAuthJson(), types.AlreadyExists()},
		Data:   types.Paper{},
	}

	addPaper(expected)
}
