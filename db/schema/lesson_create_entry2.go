package DbSchema

const LessonCreateReportFormEntry2 = `
CREATE TABLE LessonCreateReportFormEntry2 (
	lesson                              int  references lesson(id) on delete cascade,
	lesson_didaskalia_typos_ma8iomatos1 text,
	lesson_didaskalia_typos_ma8iomatos2 text,
	lesson_didaskalia_typos_ma8iomatos3 text,
	lesson_didaskalia_typos_ma8iomatos4 text,
	lesson_didaskalia_typos_ma8iomatos5 text
)
`