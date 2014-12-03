package forms

type TeacherCreateReportForm struct {
	TeacherId      int
	AkademicYearId int
	Entry1         []TeacherCreateReportFormEntry1
	Entry2         TeacherCreateReportFormEntry2
	Entry3         TeacherCreateReportFormEntry3
	Entry4         TeacherCreateReportFormEntry4
	Entry5         TeacherCreateReportFormEntry5
}

func (f *TeacherCreateReportForm) Load() {
	f.Entry1 = ListAllTeacherCreateReportForm1(f.TeacherId, f.AkademicYearId)
	f.Entry2.Load(f.TeacherId, f.AkademicYearId)
	f.Entry3.Load(f.TeacherId, f.AkademicYearId)
	f.Entry4.Load(f.TeacherId, f.AkademicYearId)
	f.Entry5.Load(f.TeacherId, f.AkademicYearId)
}
