package forms

import (
	"github.com/kokeroulis/modip/db"
)

type TeacherCreateReportFormEntry3 struct {
	// Ερευνητικά προγράμματα και έργα
	Field1 int    `schema:"ereunikita_programmata_kai_erga_episthmonikos_ypeu8inos"`
	Field2 int    `schema:"ereunikita_programmata_kai_erga_apli_summetoxi"`
	Field3 string `schema:"ereunikita_programmata_kai_erga_summetoxi_eksoterikon_sunergaton"`
	Field4 string `schema:"ereunikita_programmata_kai_erga_summetoxi_foithton"`
}

func (f *TeacherCreateReportFormEntry3) Update(teacherId int) {
	query := `UPDATE TeacherCreateReportFormEntry3 SET
		ereunikita_programmata_kai_erga_episthmonikos_ypeu8inos = $1,
		ereunikita_programmata_kai_erga_apli_summetoxi = $2,
		ereunikita_programmata_kai_erga_summetoxi_eksoterikon_sunergaton = $3,
		ereunikita_programmata_kai_erga_summetoxi_foithton = $4
		WHERE teacher = $5`

	_, err := Db.Database.Exec(query,
		f.Field1,
		f.Field2,
		f.Field3,
		f.Field4,
		teacherId)

	Db.CheckQueryWithNoRows(err, query)
}

func (f *TeacherCreateReportFormEntry3) Load(teacherId int) {
	query := `SELECT
		ereunikita_programmata_kai_erga_episthmonikos_ypeu8inos,
		ereunikita_programmata_kai_erga_apli_summetoxi,
		ereunikita_programmata_kai_erga_summetoxi_eksoterikon_sunergaton,
		ereunikita_programmata_kai_erga_summetoxi_foithton
		FROM TeacherCreateReportFormEntry3
		WHERE teacher = $1`

	err := Db.Database.QueryRow(query, teacherId).
		Scan(&f.Field1,
		&f.Field2,
		&f.Field3,
		&f.Field4)

	Db.CheckQueryWithNoRows(err, query)
}
