package forms

type LessonCreateReportForm struct {
	LessonId     int
	Entry1       LessonCreateReportFormEntry1
	Entry2       LessonCreateReportFormEntry2
	Entry3       LessonCreateReportFormEntry3
	Entry4       LessonCreateReportFormEntry4
	Entry5       LessonCreateReportFormEntry5
}

func (f *LessonCreateReportForm) Load() {
	f.Entry1.Load(f.LessonId)
	f.Entry2.Load(f.LessonId)
	f.Entry3.Load(f.LessonId)
	f.Entry4.Load(f.LessonId)
	f.Entry5.Load(f.LessonId)
}
