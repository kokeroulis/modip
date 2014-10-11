package models

type AssetType struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Assets []Asset `json:"assets"`
}

