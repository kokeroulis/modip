package models

import (
	"github.com/kokeroulis/modip/db"
)

type ResearchProgramCreateReportForm struct {
    //Α.Δ. Ερευν. Προγραμ.
    Id      int
	Field1  string `schema:"analutika_stoixeia_programmatos_titlos_programmatos"`
	Field2  bool   `schema:"analutika_stoixeia_programmatos_ylopoihsh"`
	Field3  string `schema:"analutika_stoixeia_programmatos_foreas_xrimatodotisis"`
	Field4  int    `schema:"analutika_stoixeia_programmatos_proupologismos"`
	Field5  int    `schema:"analutika_stoixeia_programmatos_diarkia"`
	Field6  string `schema:"analutika_stoixeia_programmatos_hmerominia"`
	Field7  string `schema:"analutika_stoixeia_programmatos_sxolia"`
}

func (f *ResearchProgramCreateReportForm) Update(researchProgramId int) {
	query := `UPDATE ResearchProgram SET
              analutika_stoixeia_programmatos_titlos_programmatos = $1,
              analutika_stoixeia_programmatos_ylopoihsh = $2,
              analutika_stoixeia_programmatos_foreas_xrimatodotisis = $3,
              analutika_stoixeia_programmatos_proupologismos = $4,
              analutika_stoixeia_programmatos_diarkia = $5,
              analutika_stoixeia_programmatos_hmerominia = $6,
              analutika_stoixeia_programmatos_sxolia = $7
			  WHERE id = $8`

	_, err := Db.Database.Exec(query,
								f.Field1,
								f.Field2,
								f.Field3,
								f.Field4,
								f.Field5,
								f.Field6,
								f.Field7,
								researchProgramId)

	Db.CheckQueryWithNoRows(err, query)
}

func (f *ResearchProgramCreateReportForm) Load(researchProgramId int) {
	query := `SELECT
              analutika_stoixeia_programmatos_titlos_programmatos,
              analutika_stoixeia_programmatos_ylopoihsh,
              analutika_stoixeia_programmatos_foreas_xrimatodotisis,
              analutika_stoixeia_programmatos_proupologismos,
              analutika_stoixeia_programmatos_diarkia,
              analutika_stoixeia_programmatos_hmerominia,
              analutika_stoixeia_programmatos_sxolia,
              id
			  FROM ResearchProgram
			  WHERE id = $1`

	err := Db.Database.QueryRow(query, researchProgramId).
		Scan(&f.Field1,
			 &f.Field2,
             &f.Field3,
             &f.Field4,
             &f.Field5,
             &f.Field6,
             &f.Field7,
             &f.Id)

	Db.CheckQueryWithNoRows(err, query)
}

func ListAllResearchPrograms(teacherId int) []ResearchProgramCreateReportForm {
	var formList []ResearchProgramCreateReportForm

	query := `SELECT
              id
			  FROM ResearchProgram
			  WHERE teacher = $1 order by id`

	rows, err := Db.Database.Query(query, teacherId)

	if err != nil {
		panic(err)
	}

	defer rows.Close()


	for rows.Next() {
        var id int
		if err := rows.Scan(&id); err != nil {
			panic(err)
		} else {
            currentResearchProgram := ResearchProgramCreateReportForm{}
            currentResearchProgram.Load(id)
			formList = append(formList, currentResearchProgram)
		}
	}

	if rowsErr := rows.Err(); rowsErr != nil {
		panic(err)
	}

	return formList
}

func (r *ResearchProgramCreateReportForm) Create(teacherId int) {
	var alreadyExists bool
	query := `SELECT alreadyExists
			  FROM research_program_create($1, $2, $3, $4, $5, $6, $7, $8)`

	err := Db.Database.QueryRow(query,
              r.Field1,
              r.Field2,
              r.Field3,
              r.Field4,
              r.Field5,
              r.Field6,
              r.Field7,
              teacherId).
		Scan(&alreadyExists)

	Db.CheckQuery(err, query)

	if alreadyExists {
		panic("The research program already exists!!")
	}
}
