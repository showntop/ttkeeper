package models

import ()

type Resource struct {
	ID    int64
	Rtype int
	Name  string
}

func AddResource(r *Resource) (int64, error) {
	ret := dbc.Create(r)
	return r.ID, ret.Error
}
