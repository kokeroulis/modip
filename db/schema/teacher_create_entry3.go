package DbSchema

const TeacherCreateReportFormEntry3 = `
CREATE TABLE TeacherCreateReportFormEntry3 (
	teacher int references teacher(id) on delete cascade,
	akademic_year int references akademic_year(id) on delete cascade,
	ereunikita_programmata_kai_erga_episthmonikos_ypeu8inos             int default 0,
	ereunikita_programmata_kai_erga_apli_summetoxi                      int default 0,
	ereunikita_programmata_kai_erga_summetoxi_eksoterikon_sunergaton text default ' ',
	ereunikita_programmata_kai_erga_summetoxi_foithton               text default ' '
)
`
