package forms

import (
	"github.com/kokeroulis/modip/db"
)

type TeacherCreateReportFormEntry4Helper struct {
	Content string
}

type TeacherCreateReportFormEntry4 struct {
	// Αριθμός Δημοσιεύσεων
	Id      int    `schema:-`
	Field1  string `schema:"ereunitikes_ypodomes_arithmos_kai_xwritikotita_ereunitikon_ergasthrion_pou_xrisimopoieitai"`
	Field2  string `schema:"ereunitikes_ypodomes_eparkeia_katalilotita_kai_poiotita_ton_ereunitikon_ergasthrion"`
	Field3  string `schema:"ereunitikes_ypodomes_eparkeia_katallilotita_kai_poithta_tou_ergasthriakou_eksoplismou"`
	Field4  string `schema:"ereunitikes_ypodomes_kalyptoun_oi_dia8esimes_ypodomes_tis_anagkes_tis_ereunhtikis_diadikasias"`
	Field5  string `schema:"ereunitikes_ypodomes_poia_apo_ta_ereunitika_sas_antikeimena_den_kalyptontai_apo_tis_ypodomes"`
	Field6  string `schema:"ereunitikes_ypodomes_poso_entatiki_xrish_kanete_ton_sugkrekrimenon_ereunitikon_ypodomon"`
	Field7  string `schema:"ereunitikes_ypodomes_ananeosi_ereunitikon_ypodomon"`
	Field8  string `schema:"ereunitikes_ypodomes_pws_epidiokete_th_xrimatodothsh_gia_promi8eia"`
	Field9  string `schema:"ereunitikes_ypodomes_praktiki_akiopoihsh_ton_ereunitikon_apotelesmaton"`
	Helpers []TeacherCreateReportFormEntry4Helper
}

func (f *TeacherCreateReportFormEntry4) Update() {
	query := `UPDATE TeacherCreateReportFormEntry4
			  SET
				ereunitikes_ypodomes_arithmos_kai_xwritikotita_ereunitikon_ergasthrion_pou_xrisimopoieitai = $1,
				ereunitikes_ypodomes_eparkeia_katalilotita_kai_poiotita_ton_ereunitikon_ergasthrion = $2,
				ereunitikes_ypodomes_eparkeia_katallilotita_kai_poithta_tou_ergasthriakou_eksoplismou = $3,
				ereunitikes_ypodomes_kalyptoun_oi_dia8esimes_ypodomes_tis_anagkes_tis_ereunhtikis_diadikasias = $4,
				ereunitikes_ypodomes_poia_apo_ta_ereunitika_sas_antikeimena_den_kalyptontai_apo_tis_ypodomes = $5,
				ereunitikes_ypodomes_poso_entatiki_xrish_kanete_ton_sugkrekrimenon_ereunitikon_ypodomon = $6,
				ereunitikes_ypodomes_ananeosi_ereunitikon_ypodomon = $7,
				ereunitikes_ypodomes_pws_epidiokete_th_xrimatodothsh_gia_promi8eia = $8,
				ereunitikes_ypodomes_praktiki_akiopoihsh_ton_ereunitikon_apotelesmaton = $9,
			 WHERE id = $10`

	err := Db.Database.Exec(query,
							f.Field1,
							f.Field2,
							f.Field3,
							f.Field4,
							f.Field5,
							f.Field6,
							f.Field7,
							f.Field8,
							f.Field9,
							f.Id)

	Db.CheckQueryWithNoRows(err, query)

	if len(f.Helpers) > 0 {
		// we have helpers lets add them
		for _, it := range f.Helpers {
			query = `INSERT INTO TeacherCreateReportFormEntry4Helper
					 (form, content) VALUES ($1, $2)`

			err := Db.Database.Exec(query, f.Id, it.Content)

			Db.CheckQueryWithNoRows(err, query)
		}
	}
}

func (f *TeacherCreateReportFormEntry4) Load(teacherId int) {
	query := `SELECT
				ereunitikes_ypodomes_arithmos_kai_xwritikotita_ereunitikon_ergasthrion_pou_xrisimopoieitai,
				ereunitikes_ypodomes_eparkeia_katalilotita_kai_poiotita_ton_ereunitikon_ergasthrion,
				ereunitikes_ypodomes_eparkeia_katallilotita_kai_poithta_tou_ergasthriakou_eksoplismou,
				ereunitikes_ypodomes_kalyptoun_oi_dia8esimes_ypodomes_tis_anagkes_tis_ereunhtikis_diadikasias,
				ereunitikes_ypodomes_poia_apo_ta_ereunitika_sas_antikeimena_den_kalyptontai_apo_tis_ypodomes,
				ereunitikes_ypodomes_poso_entatiki_xrish_kanete_ton_sugkrekrimenon_ereunitikon_ypodomon,
				ereunitikes_ypodomes_ananeosi_ereunitikon_ypodomon,
				ereunitikes_ypodomes_pws_epidiokete_th_xrimatodothsh_gia_promi8eia,
				ereunitikes_ypodomes_praktiki_akiopoihsh_ton_ereunitikon_apotelesmaton,
			  FROM TeacherCreateReportFormEntry4
			  WHERE id = $1 AND teacher = $2`

	err := Db.Database.QueryRow(query, f.Id, teacherId).
			Scan(&f.Field1,
				&f.Field2,
				&f.Field3,
				&f.Field4,
				&f.Field5,
				&f.Field6,
				&f.Field7,
				&f.Field8,
				&f.Field9)

	Db.CheckQuery(err, query)

	var helpers []TeacherCreateReportFormEntry4Helper

	query := `SELECT content
			  FROM TeacherCreateReportFormEntry4Helper
			  WHERE form = $1`

	rows, err = Db.Database.Query(query, f.Id)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		it := TeacherCreateReportFormEntry4Helper{}

		if err = rows.Scan(&it.Content); err != nil {
			panic(err)
		} else {
			helpers = append(helpers, it)
		}
	}

	if rowsErr := rows.Err(); rowsErr != nil {
		panic(err)
	}

	f.Helpers = helpers
}
