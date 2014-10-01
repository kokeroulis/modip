package tests

import (
	"github.com/kokeroulis/modip/models"
	"net/url"
)

const paperTitle = "My Super paper"

func addPaper(expected models.JsonData) {
	form := url.Values{}
	form.Add("title", paperTitle)

	result := PostToJsonAsTeacher("http://localhost:3001/paper/add", form)

	CompareJson(result, expected)
}

func addPaperOk() {
	p := models.Paper{1, paperTitle, testTeacher()}

	expected := testTeacherOkJson(p)

	addPaper(expected)
}

func addPaperFail() {

	expected := models.JsonData{
		Common: models.CommonJson{testAuthJson(), models.AlreadyExists()},
		Data:   models.Paper{},
	}

	addPaper(expected)
}
