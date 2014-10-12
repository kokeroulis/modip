package api

import (
	"github.com/kokeroulis/modip/models"
	"github.com/mholt/binding"
	"net/http"
)

////////////////////// Add

type AssetAddForm struct {
	Content   string
	AssetType int
}

func (a *AssetAddForm) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&a.Content: binding.Field{
			Form:     "content",
			Required: true,
		},
		&a.AssetType: binding.Field{
			Form:     "assettype",
			Required: true,
		},
	}
}

func AssetAdd(resp http.ResponseWriter, req *http.Request) {
	form := &AssetAddForm{}
	if binding.Bind(req, form).Handle(resp) {
		return
	}

	a := models.Asset{
		Content:   form.Content,
		AssetType: form.AssetType,
		Teacher:   models.GetTeacherFromSession(req),
	}

	alreadyExists := a.Add()

	if !alreadyExists {
		RenderJson(resp, req, a)
	} else {
		errorJson := models.AlreadyExists()
		RenderErrorJson(resp, req, errorJson, &models.Asset{})
	}
}


////////////////////// Remove

type AssetRemoveForm struct {
	AssetId int
}

func (a *AssetRemoveForm) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&a.AssetId: binding.Field{
			Form:     "assetId",
			Required: true,
		},
	}
}

func AssetRemove(resp http.ResponseWriter, req *http.Request) {
	form := &AssetRemoveForm{}
	if binding.Bind(req, form).Handle(resp) {
		return
	}

	a := models.Asset{
		Id:      form.AssetId,
		Teacher: models.GetTeacherFromSession(req),
	}

	isValid := a.Remove()

	if isValid {
		RenderJson(resp, req, a)
	} else {
		errorJson := models.InvalidAsset()
		RenderErrorJson(resp, req, errorJson, &models.Asset{})
	}
}

////////////////////// Move

type AssetMoveForm struct {
	AssetId     int
	AssetTypeId int
}

func (a *AssetMoveForm) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&a.AssetId: binding.Field{
			Form:     "assetId",
			Required: true,
		},
		&a.AssetTypeId: binding.Field{
			Form:     "assetTypeId",
			Required: true,
		},
	}
}

func AssetMove(resp http.ResponseWriter, req *http.Request) {
	form := &AssetMoveForm{}
	if binding.Bind(req, form).Handle(resp) {
		return
	}

	a := models.Asset{
		Id:        form.AssetId,
		AssetType: form.AssetTypeId,
		Teacher:   models.GetTeacherFromSession(req),
	}

	isValid := a.Move()

	if isValid {
		RenderJson(resp, req, a)
	} else {
		errorJson := models.InvalidAsset()
		RenderErrorJson(resp, req, errorJson, &models.Asset{})
	}
}

////////////////////// Modify

type AssetModifyForm struct {
	AssetId int
	Content string
}

func (a *AssetModifyForm) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&a.AssetId: binding.Field{
			Form:     "assetId",
			Required: true,
		},
		&a.Content: binding.Field{
			Form:     "content",
			Required: true,
		},
	}
}

func AssetModify(resp http.ResponseWriter, req *http.Request) {
	form := &AssetModifyForm{}
	if binding.Bind(req, form).Handle(resp) {
		return
	}

	a := models.Asset{
		Id:      form.AssetId,
		Content: form.Content,
		Teacher: models.GetTeacherFromSession(req),
	}

	isValid := a.Modify()

	if isValid {
		RenderJson(resp, req, a)
	} else {
		errorJson := models.InvalidAsset()
		RenderErrorJson(resp, req, errorJson, &models.Asset{})
	}
}
