package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

//CitiesDistrict
type CitiesDistrict struct {
	Id     uint   `gorm:"column:id" form:"id" json:"id" comment:"Id du quartier" sql:"int(11),PRI"`
	IdCity int    `gorm:"column:id_city" form:"id_city" json:"id_city" comment:"Id de la ville" sql:"int(11)"`
	Name   string `gorm:"column:name" form:"name" json:"name" comment:"Nom du quartier" sql:"varchar(255)"`
}

//TableName
func (m *CitiesDistrict) TableName() string {
	return "cities_districts"
}

//One
func (m *CitiesDistrict) One() (one *CitiesDistrict, err error) {
	one = &CitiesDistrict{}
	err = crudOne(m, one)
	return
}

//All
func (m *CitiesDistrict) All(q *PaginationQuery) (list *[]CitiesDistrict, total uint, err error) {
	list = &[]CitiesDistrict{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *CitiesDistrict) Update() (err error) {
	where := CitiesDistrict{Id: m.Id}
	m.Id = 0

	return crudUpdate(m, where)
}

//Create
func (m *CitiesDistrict) Create() (err error) {
	m.Id = 0

	return mysqlDB.Create(m).Error
}

//Delete
func (m *CitiesDistrict) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
