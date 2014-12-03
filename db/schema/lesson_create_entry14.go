package DbSchema

const LessonCreateReportFormEntry14 = `
CREATE TABLE LessonCreateReportFormEntry14 (
	lesson int references lesson(id) on delete cascade,
	akademic_year int references akademic_year(id) on delete cascade,
	katanomi_ba8mon_spoudaston_0_39                                     int default 0,
	katanomi_ba8mon_spoudaston_4_49                                     int default 0,
	katanomi_ba8mon_spoudaston_5_59                                     int default 0,
	katanomi_ba8mon_spoudaston_6_69                                     int default 0,
	katanomi_ba8mon_spoudaston_7_84                                     int default 0,
	katanomi_ba8mon_spoudaston_85_10                                    int default 0,
	katanomi_ba8mon_spoudaston_mesos_oros_ba8mologias_sunolo_spoudaston int default 0
)
`
