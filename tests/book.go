package tests

import (
	"github.com/kokeroulis/modip/types"
	"net/url"
)

const bookTitle = "My Super book"

func addbook(expected types.JsonData) {
	form := url.Values{}
	form.Add("title", bookTitle)

	result := PostToJsonAsTeacher("http://localhost:3001/book/add", form)

	CompareJson(result, expected)
}

func addBookOk() {
	b := types.Book{1, bookTitle, testTeacher()}

	expected := testTeacherOkJson(b)

	addbook(expected)
}

func addBookFail() {

	expected := types.JsonData{
		Common: types.CommonJson{testAuthJson(), types.AlreadyExists()},
		Data:   types.Book{},
	}

	addbook(expected)
}
