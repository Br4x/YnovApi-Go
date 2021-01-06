package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

//RealEstateAvailability
type RealEstateAvailability struct {
	Id        uint       `gorm:"column:id" form:"id" json:"id" comment:"Id de la disponibilité" sql:"int(11),PRI"`
	Price     float64    `gorm:"column:price" form:"price" json:"price" comment:"Prix de la visite" sql:"double"`
	StartDate time.Time  `gorm:"column:start_date" form:"start_date" json:"start_date" comment:"Date de début" sql:"date"`
	EndDate   time.Time  `gorm:"column:end_date" form:"end_date" json:"end_date" comment:"Date de fin" sql:"date"`
	StartTime *time.Time `gorm:"column:start_time" form:"start_time" json:"start_time,omitempty" comment:"De " sql:"time"`
	EndTime   *time.Time `gorm:"column:end_time" form:"end_time" json:"end_time,omitempty" comment:"à" sql:"time"`
}

//TableName
func (m *RealEstateAvailability) TableName() string {
	return "real_estate_availability"
}

//One
func (m *RealEstateAvailability) One() (one *RealEstateAvailability, err error) {
	one = &RealEstateAvailability{}
	err = crudOne(m, one)
	return
}

//All
func (m *RealEstateAvailability) All(q *PaginationQuery) (list *[]RealEstateAvailability, total uint, err error) {
	list = &[]RealEstateAvailability{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *RealEstateAvailability) Update() (err error) {
	where := RealEstateAvailability{Id: m.Id}
	m.Id = 0

	return crudUpdate(m, where)
}

//Create
func (m *RealEstateAvailability) Create() (err error) {
	m.Id = 0

	return mysqlDB.Create(m).Error
}

//Delete
func (m *RealEstateAvailability) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
