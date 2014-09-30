package tests

import (
	"github.com/kokeroulis/modip/types"
)

func teacherInfoOk() {
	books := []types.BookOrPaperInfo{}
	books = append(books, types.BookOrPaperInfo{1, "My Super book"})

	papers := []types.BookOrPaperInfo{}
	papers = append(papers, types.BookOrPaperInfo{1, "My Super paper"})
	info := types.TeacherInfo{testTeacher(), books, papers}

	expected := types.JsonData{testOkCommonJson(), info}

	result := GetToJsonAsTeacher("http://localhost:3001/teacher/info")

	CompareJson(result, expected)
}
