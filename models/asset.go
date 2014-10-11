package models

import "github.com/kokeroulis/modip/db"

type Asset struct {
	Id        int     `json:"id"`
	Content   string  `json:"content"`
	Teacher   Teacher `json:"teacher"`
	AssetType int     `json:"assetType"`
}

func (a *Asset) Add() bool {
	var alreadyExists bool
	query := `SELECT id, alreadyExists
			  FROM asset_add($1, $2, $3)`

	err := Db.Database.QueryRow(query, a.Content, a.Teacher.Id, a.AssetType).
		Scan(&a.Id, &alreadyExists)

	Db.CheckQuery(err, query)

	return alreadyExists
}

func (a *Asset) Remove() bool {
	var isValid bool
	query := `SELECT content, assetTypeId, isvalid
			  FROM asset_remove($1)`

	err := Db.Database.QueryRow(query, a.Id).
		Scan(&a.Content, &a.AssetType, &isValid)

	Db.CheckQuery(err, query)

	return isValid
}

func (a *Asset) Move() bool {
	var isValid bool
	query := `SELECT content, assetTypeId, isvalid
			  FROM asset_move($1, $2)`

	err := Db.Database.QueryRow(query, a.Id, a.AssetType).
		Scan(&a.Content, &a.AssetType, &isValid)

	Db.CheckQuery(err, query)

	return isValid
}

func (a *Asset) Modify() bool {
	var isValid bool
	query := `SELECT assetTypeId, isvalid
			  FROM asset_move($1, $2)`

	err := Db.Database.QueryRow(query, a.Id, a.Content).
		Scan(&a.AssetType, &isValid)

	Db.CheckQuery(err, query)

	return isValid
}
