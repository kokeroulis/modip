package Db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/kokeroulis/modip/db/api"
	"github.com/kokeroulis/modip/db/schema"
	"github.com/kokeroulis/modip/config"
)

var Database *sql.DB

func Connect() {
	c := config.NewConfig()
	d, err := sql.Open("postgres", c.PgConnectionString())

	if err != nil {
		fmt.Println("Can't connect to postgresql")
		panic(err)
	}

	pingErr := d.Ping()
	if pingErr != nil {
		fmt.Println("Can't connect to postgresql")
		panic(pingErr)
	}

	Database = d
}

func runQuery(list []string) {
	for _, it := range list {
		_, err := Database.Exec(it)
		if err != nil {
			fmt.Println("Error at: ", it)
			panic(err)
		}
	}
}

func CreateDb() {
	schemaList := []string{
		DbSchema.Department,
		DbSchema.Teacher,
		DbSchema.Category,
		DbSchema.CategoryData,
		DbSchema.CategoryGroup,

		DbSchema.AkademicYear,

		DbSchema.Lesson,
		DbSchema.LessonCreateReportForm,
		DbSchema.LessonCreateReportFormEntry1,
		DbSchema.LessonCreateReportFormEntry2,
		DbSchema.LessonCreateReportFormEntry3,
		DbSchema.LessonCreateReportFormEntry4,
		DbSchema.LessonCreateReportFormEntry5,
		DbSchema.LessonCreateReportFormEntry11,
		DbSchema.LessonCreateReportFormEntry12,
		DbSchema.LessonCreateReportFormEntry13,
		DbSchema.LessonCreateReportFormEntry14,
		DbSchema.LessonCreateReportFormEntry15,
		DbSchema.LessonCreateReportFormEntry16,

		DbSchema.TeacherCreateReportFormEntry1,
		DbSchema.TeacherCreateReportFormEntry2,
		DbSchema.TeacherCreateReportFormEntry3,
		DbSchema.TeacherCreateReportFormEntry4,
		DbSchema.TeacherCreateReportFormEntry4Helper,
		DbSchema.TeacherCreateReportFormEntry5,
	}

	apiList := []string{
		DbApi.TeacherLogin,
		DbApi.TeacherCreate,
		DbApi.CategoryAdd,
		DbApi.CategoryGroupAdd,
		DbApi.CategorySave,
		DbApi.LessonCreate,
		DbApi.AkademicYear,
	}

	runQuery(schemaList)
	runQuery(apiList)
}
