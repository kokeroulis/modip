package forms

import (
	 _"github.com/kokeroulis/modip/db"
)

type LessonCreateReportFormEntry9 struct {
	// Αξιολόγηση της επίδοσης των μαθητών στο μάθημα
	Field1 bool `schema:"lesson_aksiologisi_ths_epidosis_troipoi_aksiologisis"`
	Field2 string `schema:"lesson_aksiologisi_ths_epidosis_epeksigiseis_alla"`
	Field3 bool `schema:"lesson_aksiologisi_ths_epidosis_parakolou8ish_spoudaston_sto_erg"`
	Field4 bool `schema:"lesson_aksiologisi_ths_epidosis_sxolia_mesa_eksaminou"`
	Field5 string `schema:"lesson_aksiologisi_ths_epidosis_diafania_aksiologisis_tis_epidosis"`
}
