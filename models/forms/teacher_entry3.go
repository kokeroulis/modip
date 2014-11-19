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

