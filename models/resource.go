package models

import (
	"fmt"
)

const (
	RESOURCE_TYPE_INVA = iota
	RESOURCE_TYPE_COMP
	RESOURCE_TYPE_PAGE
	RESOURCE_TYPE_API
)

type Resource struct {
	ID        int64  `json:"id"`
	Rtype     int    `json:"type"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	Extension string `json:"ext"`
	ParentID  int64  `json:"parent_id"`
}

type PermitRes struct {
	Action    string `json:"action"`
	ID        int64  `json:"resource_id"`
	Rtype     int    `json:"type"`
	Name      string `json:"name"`
	Extension string `json:"ext"`
	ParentID  int64  `json:"parent_id"`
}

func AddResource(r *Resource) (int64, error) {
	ret := dbc.Create(r)
	return r.ID, ret.Error
}

func GetResourceByCode() {

}

func GetAllResource(parentID int64, rtype int64, offset, limit int64) ([]Resource, error) {
	var models []Resource

	query := "true"
	if parentID != 0 {
		query = query + " and " + fmt.Sprintf("parent_id=%d", parentID)
	}
	if rtype > RESOURCE_TYPE_INVA {
		query = query + " and " + fmt.Sprintf("rtype=%d", rtype)
	}
	ret := dbc.Where(query).Find(&models).Offset(offset).Limit(limit)
	return models, ret.Error
}
