package forms

import (
	"github.com/kokeroulis/modip/db"
)

type LessonCreateReportFormEntry14 struct {
	// Κατανομή Βαθμών Σπουδαστών (%)
	Field1 string `schema:"katanomi_ba8mon_spoudaston_0_39"`
	Field2 string `schema:"katanomi_ba8mon_spoudaston_4_49"`
	Field3 string `schema:"katanomi_ba8mon_spoudaston_5_59"`
	Field4 string `schema:"katanomi_ba8mon_spoudaston_6_69"`
	Field5 string `schema:"katanomi_ba8mon_spoudaston_7_84"`
	Field6 string `schema:"katanomi_ba8mon_spoudaston_85_10"`
	Field7 string `schema:"katanomi_ba8mon_spoudaston_mesos_oros_ba8mologias_sunolo_spoudaston"`
}

func (f *LessonCreateReportFormEntry14) Update(lessonId int, akademicYearId int) {
	query := `UPDATE LessonCreateReportFormEntry14 SET
			katanomi_ba8mon_spoudaston_0_39 = $1,
			katanomi_ba8mon_spoudaston_4_49 = $2,
			katanomi_ba8mon_spoudaston_5_59 = $3,
			katanomi_ba8mon_spoudaston_6_69 = $4,
			katanomi_ba8mon_spoudaston_7_84 = $5,
			katanomi_ba8mon_spoudaston_85_10 = $6,
			katanomi_ba8mon_spoudaston_mesos_oros_ba8mologias_sunolo_spoudaston = $7
			 WHERE lesson = $8 AND akademic_year = $9`

	_, err := Db.Database.Exec(query,
		f.Field1,
		f.Field2,
		f.Field3,
		f.Field4,
		f.Field5,
		f.Field6,
		f.Field7,
		lessonId,
		akademicYearId)

	Db.CheckQueryWithNoRows(err, query)
}

func (f *LessonCreateReportFormEntry14) Load(lessonId int, akademicYearId int) {
	query := `SELECT
	katanomi_ba8mon_spoudaston_0_39,
	katanomi_ba8mon_spoudaston_4_49,
	katanomi_ba8mon_spoudaston_5_59,
	katanomi_ba8mon_spoudaston_6_69,
	katanomi_ba8mon_spoudaston_7_84,
	katanomi_ba8mon_spoudaston_85_10,
	katanomi_ba8mon_spoudaston_mesos_oros_ba8mologias_sunolo_spoudaston
	FROM LessonCreateReportFormEntry14
	WHERE lesson = $1 AND akademic_year = $2`

	err := Db.Database.QueryRow(query, lessonId,
		akademicYearId).
		Scan(&f.Field1,
		&f.Field2,
		&f.Field3,
		&f.Field4,
		&f.Field5,
		&f.Field6,
		&f.Field7)

	Db.CheckQueryWithNoRows(err, query)
}
