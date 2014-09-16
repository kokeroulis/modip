package tests

import (
	"net/url"
	. "github.com/smartystreets/goconvey/convey"
	"types"
)

const bookTitle = "My Super book"

func addbook(title string, result interface{}, expected interface{}) {
	form := url.Values{}
	form.Add("title", title)

	PostToJsonAsTeacher("http://localhost:3001/book/add", form, result)

	So(result, ShouldResemble, expected)
}

func addBookOk() {
	b := types.Book{1, bookTitle, testTeacher()}

	expected := &types.BookJson{testOkCommonJson(), b}

	result := &types.BookJson{}
	addbook(bookTitle, result, expected)
}

func addBookFail() {

	expected := &types.BookJson{
		types.CommonJson{testAuthJson(), types.AlreadyExists()},
		types.Book{},
	}

	result := &types.BookJson{}
	addbook(bookTitle, result, expected)
}

