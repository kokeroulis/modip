package forms

import (
	"github.com/kokeroulis/modip/db"
)

type LessonCreateReportFormEntry12 struct {
	// Εκπαιδευτικά Μέσα
	Field1 bool   `schema:"ekpedeutika_mesa_xrhsh_ekpedeutikon_meson"`
	Field2 bool   `schema:"ekpedeutika_mesa_eparkeia_ekpedeutikon_meson"`
	Field3 string `schema:"ekpedeutika_mesa_anafora_elleipseon_ekpedeutikon_meson"`
}

func (f *LessonCreateReportFormEntry12) Update(lessonId int, akademicYearId int) {
	query := `UPDATE lessoncreatereportformentry12 SET
	ekpedeutika_mesa_xrhsh_ekpedeutikon_meson = $1,
	ekpedeutika_mesa_eparkeia_ekpedeutikon_meson = $2,
	ekpedeutika_mesa_anafora_elleipseon_ekpedeutikon_meson = $3
	 WHERE lesson = $4 AND akademic_year = $5`

	_, err := Db.Database.Exec(query,
		f.Field1,
		f.Field2,
		f.Field3,
		lessonId,
		akademicYearId)

	Db.CheckQueryWithNoRows(err, query)
}

func (f *LessonCreateReportFormEntry12) Load(lessonId int, akademicYearId int) {
	query := `SELECT
			ekpedeutika_mesa_xrhsh_ekpedeutikon_meson,
			ekpedeutika_mesa_eparkeia_ekpedeutikon_meson,
			ekpedeutika_mesa_anafora_elleipseon_ekpedeutikon_meson
		FROM lessoncreatereportformentry12
		WHERE lesson = $1`

	err := Db.Database.QueryRow(query, lessonId,
		akademicYearId).
		Scan(&f.Field1,
		&f.Field2,
		&f.Field3)

	Db.CheckQueryWithNoRows(err, query)
}
