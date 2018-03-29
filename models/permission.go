package models

import (
	"fmt"
)

type Permission struct {
	ID         int64 `json:"id"`
	RoleID     int64 `json:"role_id"`
	Could      bool  `json:"could"`
	Action     int   `json:"action"`
	ResourceId int64 `json:"resource_id"`
}

func AddPermission(r *Permission) (int64, error) {
	ret := dbc.Create(r)
	return r.ID, ret.Error
}

func GetAllPermission(roleID, offset, limit int64) ([]Permission, error) {
	var models []Permission
	query := "true"
	if roleID >= 0 {
		query += " and " + fmt.Sprintf("role_id = %d", roleID)
	}
	ret := dbc.Where(query).Find(&models).Offset(offset).Limit(limit)
	return models, ret.Error
}

func GrantedPermission(userID int64, rtype int, code string) bool {
	var permissions []PermitRes
	ret := dbc.Select("d.id").Table("user_roles").
		Joins("left join permissions as c on user_roles.role_id = c.role_id left join resources as d on c.resource_id = d.id").
		Where("user_roles.user_id = ?", userID).Where("d.rtype = ? and code = ?", rtype, code).Scan(&permissions)
	fmt.Println(permissions)
	if ret.Error != nil || len(permissions) <= 0 {
		return false
	}
	return true
}
