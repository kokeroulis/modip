package forms

type TeacherCreateReportForm struct {
	TeacherId int
	Entry1    TeacherCreateReportFormEntry1
	Entry2    TeacherCreateReportFormEntry2
	Entry3    TeacherCreateReportFormEntry3
	Entry4    TeacherCreateReportFormEntry3
	Entry5    TeacherCreateReportFormEntry5
}

func (f *TeacherCreateReportForm) Load() {
/*	f.Entry1.Load(f.TeacherId)
	f.Entry2.Load(f.TeacherId)
	f.Entry3.Load(f.TeacherId)
	f.Entry4.Load(f.TeacherId)
	f.Entry5.Load(f.TeacherId)
	*/
}
