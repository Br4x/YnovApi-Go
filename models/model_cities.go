package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

//City
type City struct {
	Id      uint   `gorm:"column:id" form:"id" json:"id" comment:"Id de la ville" sql:"int(11),PRI"`
	Name    string `gorm:"column:name" form:"name" json:"name" comment:"Nom de la ville" sql:"varchar(255)"`
	ZipCode string `gorm:"column:zip_code" form:"zip_code" json:"zip_code" comment:"Code postal de la ville" sql:"varchar(10)"`
}

//TableName
func (m *City) TableName() string {
	return "cities"
}

//One
func (m *City) One() (one *City, err error) {
	one = &City{}
	err = crudOne(m, one)
	return
}

//All
func (m *City) All(q *PaginationQuery) (list *[]City, total uint, err error) {
	list = &[]City{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *City) Update() (err error) {
	where := City{Id: m.Id}
	m.Id = 0

	return crudUpdate(m, where)
}

//Create
func (m *City) Create() (err error) {
	m.Id = 0

	return mysqlDB.Create(m).Error
}

//Delete
func (m *City) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
