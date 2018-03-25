package models

import ()

type Orgunit struct {
	Id   int64
	Name string

	UpunitID int64
	Subunits []Orgunit `gorm:"ForeignKey:UpunitID"`

	Users []User `gorm:"ForeignKey:OrgunitID"`
}
