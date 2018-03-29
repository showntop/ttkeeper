package models

import (
// "time"

// "github.com/jinzhu/gorm"
)

type User struct {
	Model
	Username string `json:"username"`
	Password string `json:"password"`

	OrgunitID int     `json:"orgunit_id"`
	Orgunit   Orgunit `gorm:"foreignkey:OrgunitID"` // use OrgunitRefer as foreign key
}

type UserRole struct {
	ID int64

	UserID int
	RoleID int
}

type UserProfile struct {
	Model
	Username  string `json:"username"`
	OrgunitID int    `json:"orgunit_id"`
}

// AddUser insert a new User into database and returns
// last inserted Id on success.
func AddUser(m *User) (id int64, err error) {
	ret := dbc.Create(m)
	return m.ID, ret.Error
}

func GetAllUser(offset, limit int64) ([]UserProfile, error) {
	var users []UserProfile
	ret := dbc.Table("users").Find(&users).Offset(offset).Limit(limit)
	return users, ret.Error
}
func GetUserByID(ID string) (*User, error) {
	var user User
	ret := dbc.Select("username, password").First(&user, "username = ?", ID)
	return &user, ret.Error
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	ret := dbc.Select("id,username, password").First(&user, "username = ?", username)
	return &user, ret.Error
}

func GetUserPermissions(userID, parentID int64) ([]PermitRes, error) {
	var permissions []PermitRes
	ret := dbc.Select("c.action, d.*").Table("user_roles").
		Joins("left join permissions as c on user_roles.role_id = c.role_id left join resources as d on c.resource_id = d.id").
		Where("user_roles.user_id = ?", userID).Where("d.parent_id = ?", parentID).Scan(&permissions)
	return permissions, ret.Error
}

func GetUserPermissionsByType(userID, ptype int) ([]PermitRes, error) {
	var permissions []PermitRes
	ret := dbc.Select("c.action, d.*").Table("user_roles").
		Joins("left join permissions as c on user_roles.role_id = c.role_id left join resources as d on c.resource_id = d.id").
		Where("user_roles.user_id = ?", userID).Where("d.rtype = ?", ptype).Scan(&permissions)
	return permissions, ret.Error
}
