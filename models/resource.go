package models

import ()

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
