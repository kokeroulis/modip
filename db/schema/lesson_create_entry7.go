package DbSchema

const LessonCreateReportFormEntry7 = `
CREATE TABLE LessonCreateReportFormEntry7 (
        lesson int references lesson(id) on delete cascade,
        akademic_year int references akademic_year(id) on delete cascade,
        lesson_epikoinonia_wres_grafeiou                        text default ' ',
        lesson_epikoinonia_ekepedeusi_methodos                  text default ' ',
        lesson_epikoinonia_organosi_ekpedeutikon_episkepseon    text default ' '
)
`

