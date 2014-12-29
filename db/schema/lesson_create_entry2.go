package DbSchema

const LessonCreateReportFormEntry2 = `
CREATE TABLE LessonCreateReportFormEntry2 (
	lesson                              int  references lesson(id) on delete cascade,
	akademic_year int references akademic_year(id) on delete cascade,
	lesson_didaskalia_typos_ma8iomatos1 text default ' ',
	lesson_didaskalia_typos_ma8iomatos2 text default ' ',
	lesson_didaskalia_typos_ma8iomatos3 text default ' ',
	lesson_didaskalia_typos_ma8iomatos4 text default ' ',
	lesson_didaskalia_typos_ma8iomatos5 text default ' ',
	lesson_didaskalia_wres_didaskaliasi_8  text default ' ',
	lesson_didaskalia_wres_didaskalias_AP text default ' ',
	lesson_didaskalia_wres_didaskalias_E text default ' ',
	lesson_didaskalia_monadesECTS_8AP text default ' ',
	lesson_didaskalia_monadesECTS_E text default ' ',
	lesson_didaskalia_wres_didaskalias_ana_bdomada text default ' ',
	lesson_didaskalia_wres_ergasthriou_askhshs boolean default false,
	lesson_didaskalia_wres_ergasthriou_ana_bdomada text default ' ',
	lesson_didaskalia_problepomenes_wres_ergasthriou_ana_eksamino text default ' ',
	lesson_didaskalia_problepomenes_wres_didaskalias_mikron_omadon_ana_eksamino text default ' ',
	lesson_didaskalia_problepomenes_wres_didaskalias_ana_eksamino_alli text default ' ',
	lesson_didaskalia_pollapli_bibiografia boolean default false,
	lesson_didaskalia_diati8ete_ergasia_proodos boolean default false,
	lesson_didaskalia_ypoxreotiki_ergasia_proodos boolean default false
)
`
