package DbSchema

const LessonCreateReportFormEntry8 = `
CREATE TABLE LessonCreateReportFormEntry8 (
        lesson int references lesson(id) on delete cascade,
        lesson_summetoxi_spoudaston_summetoxi_stin_kanoniki_eksetasi        text default ' ',
        lesson_summetoxi_spoudaston_summetoxi_stin_epanaliptikh_eksetasi    int default 0,
)
`

