package forms

type LessonCreateReportForm struct {
	LessonId       int
	AkademicYearId int
	Entry1  LessonCreateReportFormEntry1
	Entry2  LessonCreateReportFormEntry2
	Entry3  LessonCreateReportFormEntry3
	Entry4  LessonCreateReportFormEntry4
	Entry5  LessonCreateReportFormEntry5

	Entry6  LessonCreateReportFormEntry6
	Entry7  LessonCreateReportFormEntry7
	Entry8  LessonCreateReportFormEntry8
	Entry9  LessonCreateReportFormEntry9
	Entry10 LessonCreateReportFormEntry10

	Entry11 LessonCreateReportFormEntry11
	Entry12 LessonCreateReportFormEntry12
	Entry13 LessonCreateReportFormEntry13
	Entry14 LessonCreateReportFormEntry14
	Entry15 LessonCreateReportFormEntry15
	Entry16 LessonCreateReportFormEntry16

}

func (f *LessonCreateReportForm) Load() {
	f.Entry1.Load(f.LessonId, f.AkademicYearId)
	f.Entry2.Load(f.LessonId, f.AkademicYearId)
	f.Entry3.Load(f.LessonId, f.AkademicYearId)
	f.Entry4.Load(f.LessonId, f.AkademicYearId)
	f.Entry5.Load(f.LessonId, f.AkademicYearId)

	f.Entry6.Load(f.LessonId, f.AkademicYearId)
	f.Entry7.Load(f.LessonId, f.AkademicYearId)
	f.Entry8.Load(f.LessonId, f.AkademicYearId)
	f.Entry9.Load(f.LessonId, f.AkademicYearId)
	f.Entry10.Load(f.LessonId, f.AkademicYearId)

    f.Entry11.Load(f.LessonId, f.AkademicYearId)
	f.Entry12.Load(f.LessonId, f.AkademicYearId)
	f.Entry13.Load(f.LessonId, f.AkademicYearId)
	f.Entry14.Load(f.LessonId, f.AkademicYearId)
	f.Entry15.Load(f.LessonId, f.AkademicYearId)
	f.Entry16.Load(f.LessonId, f.AkademicYearId)
}
