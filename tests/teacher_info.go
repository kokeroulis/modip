package tests

import (
	"github.com/kokeroulis/modip/models"
)

func teacherInfoOk() {
	books := []models.BookOrPaperInfo{}
	books = append(books, models.BookOrPaperInfo{1, "My Super book"})

	papers := []models.BookOrPaperInfo{}
	papers = append(papers, models.BookOrPaperInfo{1, "My Super paper"})
	info := models.TeacherInfo{testTeacher(), books, papers}

	expected := models.JsonData{testOkCommonJson(), info}

	result := GetToJsonAsTeacher("http://localhost:3001/teacher/info")

	CompareJson(result, expected)
}
