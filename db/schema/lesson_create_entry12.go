package DbSchema

const LessonCreateReportFormEntry12 = `
CREATE TABLE lessoncreatereportformentry12 (
	lesson                      int references lesson(id) on delete cascade,
	ekpedeutika_mesa_xrhsh_ekpedeutikon_meson               bool default false,
	ekpedeutika_mesa_eparkeia_ekpedeutikon_meson            bool default false,
	ekpedeutika_mesa_anafora_elleipseon_ekpedeutikon_meson    text default ' '
)
`
