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

