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
			  SET author = $1, title = $2, is_magazine $3,
			  publisher = $4,
			  publication_date = $5, type = $5
			  WHERE id = $2`

	err := Db.Database.Exec(query, f.Id)

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

func ListAllTeacherCreateReportForm1(teacherId int) []TeacherCreateReportFormEntry4 {
	var formList []TeacherCreateReportFormEntry4

	query := `SELECT TeacherCreateReportFormEntry4
			  id, author, title, is_magazine,
			  publisher, publication_date
			  FROM TeacherCreateReportFormEntry4
			  WHERE teacher $1`

	rows, err := Db.Database.Query(query, teacherId)

	if err != nil {
		panic(err)
	}

	defer rows.Close()


	for rows.Next() {
		it := TeacherCreateReportFormEntry4{}

		if err := rows.Scan(&it.Id,
							&it.Field1,
							&it.Field2,
							&it.Field3,
							&it.Field4,
							&it.Field5,
							&it.Field6); err != nil {
			panic(err)
		} else {
			formList = append(formList, it)
		}
	}

	if rowsErr := rows.Err(); rowsErr != nil {
		panic(err)
	}

	return formList
}
