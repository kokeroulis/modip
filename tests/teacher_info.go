package tests

import (
	"github.com/kokeroulis/modip/types"
	. "github.com/smartystreets/goconvey/convey"
)

func teacherInfoOk() {
	books := []types.BookOrPaperInfo{}
	books = append(books, types.BookOrPaperInfo{1, "My Super book"})

	papers := []types.BookOrPaperInfo{}
	papers = append(papers, types.BookOrPaperInfo{1, "My Super paper"})
	info := types.TeacherInfo{testTeacher(), books, papers}

	expected := &types.TeacherInfoJson{testOkCommonJson(), info}

	result := &types.TeacherInfoJson{}
	GetToJsonAsTeacher("http://localhost:3001/teacher/info", result)

	So(result, ShouldResemble, expected)
}
