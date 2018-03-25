package models

import (
	"fmt"
)

type Role struct {
	ID   int64
	Name string
}

type RolePermission struct {
	ID int64

	PermissionID int64
	RoleID       int64
}

func AddRole(r *Role) (int64, error) {
	ret := dbc.Create(r)
	return r.ID, ret.Error
}

func (r *Role) Grant(p *Permission) error {
	rp := RolePermission{RoleID: r.ID, PermissionID: p.ID}
	fmt.Printf("%+v\r\n", rp)
	ret := dbc.Create(&rp)
	return ret.Error
}
