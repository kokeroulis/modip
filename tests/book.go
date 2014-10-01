package tests

import (
	"github.com/kokeroulis/modip/models"
	"net/url"
)

const bookTitle = "My Super book"

func addbook(expected models.JsonData) {
	form := url.Values{}
	form.Add("title", bookTitle)

	result := PostToJsonAsTeacher("http://localhost:3001/book/add", form)

	CompareJson(result, expected)
}

func addBookOk() {
	b := models.Book{1, bookTitle, testTeacher()}

	expected := testTeacherOkJson(b)

	addbook(expected)
}

func addBookFail() {

	expected := models.JsonData{
		Common: models.CommonJson{testAuthJson(), models.AlreadyExists()},
		Data:   models.Book{},
	}

	addbook(expected)
}
