package DbSchema

const LessonCreateReportFormEntry5 = `
CREATE TABLE LessonCreateReportFormEntry5 (
	lesson int references lesson(id) on delete cascade,
	akademic_year int references akademic_year(id) on delete cascade,
	lesson_didaktika_boi8imata_boi8imata_pou_dianemontai_stous_foithtes text default ' ',
	lesson_didaktika_boi8imata_epikairopoihsh_ton_boh8imaton text default ' ',
	lesson_didaktika_boi8imata_pososto_didaskomenis_ylis_pou_kaliptete text default ' ',
	lesson_didaktika_boi8imata_prostheti_bibliografia text default ' ',
	lesson_didaktika_boi8imata_epikairopoihsh_ylis text default ' ',
	lesson_didaktika_boi8imata_gnostopoihsh_ylis_stous_foithtes text default ' '
)
`
