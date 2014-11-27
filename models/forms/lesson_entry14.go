package forms

import (
	 _"github.com/kokeroulis/modip/db"
)

type LessonCreateReportFormEntry14 struct {
	// Κατανομή Βαθμών Σπουδαστών (%)
	Field1 int `schema:"katanomi_ba8mon_spoudaston_0_39"`
	Field2 int `schema:"katanomi_ba8mon_spoudaston_4_49"`
	Field3 int `schema:"katanomi_ba8mon_spoudaston_5_59"`
    Field4 int `schema:"katanomi_ba8mon_spoudaston_6_69"`
    Field5 int `schema:"katanomi_ba8mon_spoudaston_7_84"`
    Field6 int `schema:"katanomi_ba8mon_spoudaston_85_10"`
    Field7 int `schema:"katanomi_ba8mon_spoudaston_mesos_oros_ba8mologias_sunolo_spoudaston"`
}


=============== Start create update func
func (f *LessonCreateReportFormEntry14) Update(lessonId int) {
query := `UPDATE LessonCreateReportFormEntry14 SET
	katanomi_ba8mon_spoudaston_0_39 = $1,
	katanomi_ba8mon_spoudaston_4_49 = $2,
	katanomi_ba8mon_spoudaston_5_59 = $3,
	katanomi_ba8mon_spoudaston_6_69 = $4,
	katanomi_ba8mon_spoudaston_7_84 = $5,
	katanomi_ba8mon_spoudaston_85_10 = $6,
	katanomi_ba8mon_spoudaston_mesos_oros_ba8mologias_sunolo_spoudaston = $7
WHERE lesson = $8`

_, err := Db.Database.Exec(query,
	Field1,
	Field2,
	Field3,
	Field4,
	Field5,
	Field6,
	Field7
    lessonId)

Db.CheckQueryWithNoRows(err, query)
}
=============== End Create update func
=============== Start create select func
func (f *LessonCreateReportFormEntry14) Load(lessonId int) {
query := `SELECT
	katanomi_ba8mon_spoudaston_0_39,
	katanomi_ba8mon_spoudaston_4_49,
	katanomi_ba8mon_spoudaston_5_59,
	katanomi_ba8mon_spoudaston_6_69,
	katanomi_ba8mon_spoudaston_7_84,
	katanomi_ba8mon_spoudaston_85_10,
	katanomi_ba8mon_spoudaston_mesos_oros_ba8mologias_sunolo_spoudaston
FROM LessonCreateReportFormEntry14
WHERE lesson = $1`

err := Db.Database.QueryRow(query, lessonId).
	Scan(&f.Field1,
		&f.Field2,
		&f.Field3,
		&f.Field4,
		&f.Field5,
		&f.Field6,
		&f.Field7)

Db.CheckQueryWithNoRows(err, query)
}
=============== End Create select func
