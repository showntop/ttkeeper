package models

import (
	// "time"
	"fmt"
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
	ID int64 `json:"-"`

	UserID int `json:"user_id"`
	RoleID int `json:"role_id"`
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

// AddUser insert a new User into database and returns
// last inserted Id on success.
func AddUserRole(m *UserRole) (id int64, err error) {
	ret := dbc.Create(m)
	return m.ID, ret.Error
}

func GetAllUser(offset, limit int64) ([]UserProfile, error) {
	var users []UserProfile
	ret := dbc.Table("users").Find(&users).Offset(offset).Limit(limit)
	return users, ret.Error
}
func GetAllUserRole(userID, offset, limit int64) ([]UserRole, error) {
	var models []UserRole
	ret := dbc.Table("user_roles").Where("user_id = ?", userID).Find(&models).Offset(offset).Limit(limit)
	return models, ret.Error
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
	query := "true"
	if userID > 0 {
		query += " and " + fmt.Sprintf("user_roles.user_id = %d", userID)
	}
	if parentID > 0 {
		query += " and " + fmt.Sprintf("d.parent_id = %d", parentID)
	}
	ret := dbc.Select("c.action, d.*").Table("user_roles").
		Joins("left join permissions as c on user_roles.role_id = c.role_id left join resources as d on c.resource_id = d.id").
		Where(query).Scan(&permissions)
	return permissions, ret.Error
}

func GetUserPermissionsByType(userID, ptype int) ([]PermitRes, error) {
	var permissions []PermitRes
	ret := dbc.Select("c.action, d.*").Table("user_roles").
		Joins("left join permissions as c on user_roles.role_id = c.role_id left join resources as d on c.resource_id = d.id").
		Where("user_roles.user_id = ?", userID).Where("d.rtype = ?", ptype).Scan(&permissions)
	return permissions, ret.Error
}
