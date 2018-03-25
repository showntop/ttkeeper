package models

import ()

type Permission struct {
	ID         int64 `json:"id"`
	Action     int   `json:"action"`
	ResourceId int64 `json:"resource_id"`
}

func AddPermission(r *Permission) (int64, error) {
	ret := dbc.Create(r)
	return r.ID, ret.Error
}
