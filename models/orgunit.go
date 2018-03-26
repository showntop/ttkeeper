package models

import ()

type Orgunit struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`

	ParentID int64     `json:"parent_id"`
	Subunits []Orgunit `json:"-"`

	Users []User
}

func GetAllOrgunit(parentID int64) ([]Orgunit, error) {
	var units []Orgunit
	if parentID <= 0 {
		ret := dbc.Find(&units)
		return units, ret.Error
	} else {
		ret := dbc.Where("parent_id =?", parentID).Find(&units)
		return units, ret.Error
	}

}

func AddOrgunit(r *Orgunit) (int64, error) {
	ret := dbc.Create(r)
	return r.ID, ret.Error
}
