package models

import (
	"fmt"
)

type Role struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type RolePermission struct {
	ID int64

	PermissionID int64
	RoleID       int64
}

func GetAllRole(offset, limit int64) ([]Role, error) {
	var roles []Role
	ret := dbc.Find(&roles).Offset(offset).Limit(limit)
	return roles, ret.Error
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
