package DbSchema

const LessonCreateReportFormEntry9 = `
CREATE TABLE LessonCreateReportFormEntry9 (
        lesson int references lesson(id) on delete cascade,
        lesson_aksiologisi_ths_epidosis_grapti_eksetasi                         bool default false,
        lesson_aksiologisi_ths_epidosis_proforiki_eksetasi                      bool default false,
        lesson_aksiologisi_ths_epidosis_proodos                                 bool default false,
        lesson_aksiologisi_ths_epidosis_katoi_kon_ergasia                       bool default false,
        lesson_aksiologisi_ths_epidosis_proforiki_parousiasi_ergasias           bool default false,
        lesson_aksiologisi_ths_epidosis_ergashtirio_praktikes_askiseis          bool default false,
        lesson_aksiologisi_ths_epidosis_alla                                    bool default false,
        lesson_aksiologisi_ths_epidosis_epeksigiseis_alla                       text default ' ',
        lesson_aksiologisi_ths_epidosis_parakolou8ish_spoudaston_sto_erg        bool default false,
        lesson_aksiologisi_ths_epidosis_sxolia_mesa_eksaminou                   bool default false,
        lesson_aksiologisi_ths_epidosis_diafania_aksiologisis_tis_epidosis      text default ' '
)
`
