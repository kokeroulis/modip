package models

import "github.com/kokeroulis/modip/db"

type Teacher struct {
	Id         int        `json:"id"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Department Department `json:"department"`
}

func (t *Teacher) Login(username string, password string) bool {
	var auth bool
	query := `SELECT id, name, email, departmentId, departmentName, auth
			  FROM teacher_auth($1, $2)`
	err := Db.Database.QueryRow(query, username, password).
		Scan(&t.Id, &t.Name, &t.Email, &t.Department.Id, &t.Department.Name, &auth)

	Db.CheckQuery(err, query)

	return auth
}

type TeacherInfo struct {
	Teacher    Teacher     `json:"teacher"`
	AssetTypes []AssetType `json:"assetTypes"`
}

func (t *TeacherInfo) takeAssets() {
	queryAssets := `
		SELECT id, content, assetType
		FROM asset
		WHERE teacher = $1 AND assetType = $2
	 `

	 for index, it := range t.AssetTypes {
		assets := []Asset{}
		rows, err := Db.Database.Query(queryAssets, t.Teacher.Id, it.Id)

		if err != nil {
			panic(err)
		}

		defer rows.Close()

		for rows.Next() {
			asset := Asset{}

			if err := rows.Scan(&asset.Id, &asset.Content, &asset.AssetType); err != nil {
				panic(err)
			} else {
				assets = append(assets, asset)
			}
		}

		if err := rows.Err(); err != nil {
			panic(err)
		}

		t.AssetTypes[index].Assets = assets
	 }
}

func (t *Teacher) takeAssetTypes() []AssetType {
	assetTypes := []AssetType{}

	queryAllAssetTypes := `SELECT id, name FROM assetTypes`
	rows, err := Db.Database.Query(queryAllAssetTypes)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		assetType := AssetType{}

		if err := rows.Scan(&assetType.Id, &assetType.Name); err != nil {
			panic(err)
		} else {
			assetTypes = append(assetTypes, assetType)
		}
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	return assetTypes
}

func (t *Teacher) Info() TeacherInfo {
	assetTypes := t.takeAssetTypes()

	info := TeacherInfo{
		Teacher:    *t,
		AssetTypes: assetTypes,
	}

	info.takeAssets()

	return info
}
