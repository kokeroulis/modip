package forms

import (
	 _"github.com/kokeroulis/modip/db"
)

type TeacherCreateReportFormEntry3 struct {
	// Ερευνητικά προγράμματα και έργα
	Field1 int        `schema:"ereunikita_programmata_kai_erga_episthmonikos_ypeu8inos"`
    Field2 int        `schema:"ereunikita_programmata_kai_erga_apli_summetoxi"`
    Field3 string     `schema:"ereunikita_programmata_kai_erga_summetoxi_eksoterikon_sunergaton"`
    Field4 string     `schema:"ereunikita_programmata_kai_erga_summetoxi_foithton"`
}


=============== Start create table
CREATE TABLE TeacherCreateReportFormEntry3 (
	lesson int references lesson(id) on delete cascade,
	ereunikita_programmata_kai_erga_episthmonikos_ypeu8inos int,
	ereunikita_programmata_kai_erga_apli_summetoxi int,
	ereunikita_programmata_kai_erga_summetoxi_eksoterikon_sunergaton text,
	ereunikita_programmata_kai_erga_summetoxi_foithton text
)
=============== End Create table
=============== Start create update func
func (f *TeacherCreateReportFormEntry3) Update(lessonId int) {
query := `UPDATE TeacherCreateReportFormEntry3 SET
	ereunikita_programmata_kai_erga_episthmonikos_ypeu8inos = $1,
	ereunikita_programmata_kai_erga_apli_summetoxi = $2,
	ereunikita_programmata_kai_erga_summetoxi_eksoterikon_sunergaton = $3,
	ereunikita_programmata_kai_erga_summetoxi_foithton = $4
WHERE lesson = $`

_, err := Db.Database.Exec(query,
	Field1,
	Field2,
	Field3,
	Field4
)

Db.CheckQueryWithNoRows(err, query)
}
=============== End Create update func
=============== Start create select func
func (f *TeacherCreateReportFormEntry3) Load(lessonId int) {
query := `SELECT
	ereunikita_programmata_kai_erga_episthmonikos_ypeu8inos,
	ereunikita_programmata_kai_erga_apli_summetoxi,
	ereunikita_programmata_kai_erga_summetoxi_eksoterikon_sunergaton,
	ereunikita_programmata_kai_erga_summetoxi_foithton
FROM TeacherCreateReportFormEntry3
WHERE lesson = $`

err := Db.Database.QueryRow(query, TODO).
	Scan(&f.Field1,
		&f.Field2,
		&f.Field3,
		&f.Field4
)

Db.CheckQueryWithNoRows(err, query)
}
=============== End Create select func
