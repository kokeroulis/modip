package forms

import (
	 _"github.com/kokeroulis/modip/db"
)

type LessonCreateReportFormEntry12 struct {
	// Εκπαιδευτικά Μέσα
	Field1 bool   `schema:"ekpedeutika_mesa_xrhsh_ekpedeutikon_meson"`
	Field2 bool   `schema:"ekpedeutika_mesa_eparkeia_ekpedeutikon_meson"`
	Field3 string `schema:"ekpedeutika_mesa_anafora_elleipseon_ekpedeutikon_meson"`
}


=============== Start create table
CREATE TABLE lessoncreatereportformentry12 (
	lesson int references lesson(id) on delete cascade,
	ekpedeutika_mesa_xrhsh_ekpedeutikon_meson boolean,
	ekpedeutika_mesa_eparkeia_ekpedeutikon_meson boolean,
	ekpedeutika_mesa_anafora_elleipseon_ekpedeutikon_meson text
)
=============== End Create table
=============== Start create update func
func (f *lessoncreatereportformentry12) Update(lessonId int) {
query := `UPDATE lessoncreatereportformentry12 SET
	ekpedeutika_mesa_xrhsh_ekpedeutikon_meson = $1,
	ekpedeutika_mesa_eparkeia_ekpedeutikon_meson = $2,
	ekpedeutika_mesa_anafora_elleipseon_ekpedeutikon_meson = $3
WHERE lesson = $`

_, err := Db.Database.Exec(query,
	field1,
	field2,
	field3
)

Db.CheckQueryWithNoRows(err, query)
}
=============== End Create update func
=============== Start create select func
func (f *lessoncreatereportformentry12) Load(lessonId int) {
query := `SELECT
	ekpedeutika_mesa_xrhsh_ekpedeutikon_meson,
	ekpedeutika_mesa_eparkeia_ekpedeutikon_meson,
	ekpedeutika_mesa_anafora_elleipseon_ekpedeutikon_meson
FROM lessoncreatereportformentry12
WHERE lesson = $`

err := Db.Database.QueryRow(query, TODO).
	Scan(&f.field1,
		&f.field2,
		&f.field3
)

Db.CheckQueryWithNoRows(err, query)
}
=============== End Create select func
