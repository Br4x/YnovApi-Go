package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

//RealEstateVisit
type RealEstateVisit struct {
	Id             uint       `gorm:"column:id" form:"id" json:"id" comment:"Id de la visite" sql:"int(11),PRI"`
	IdRealEstate   int        `gorm:"column:id_real_estate" form:"id_real_estate" json:"id_real_estate" comment:"Id du bien" sql:"int(11)"`
	IdBooker       int        `gorm:"column:id_booker" form:"id_booker" json:"id_booker" comment:"Id du booker" sql:"int(11)"`
	IdVisitor      int        `gorm:"column:id_visitor" form:"id_visitor" json:"id_visitor" comment:"Id du visiteur" sql:"int(11)"`
	StartDate      time.Time  `gorm:"column:start_date" form:"start_date" json:"start_date" comment:"Date de début (YYYY-mm-dd)" sql:"date"`
	EndDate        time.Time  `gorm:"column:end_date" form:"end_date" json:"end_date" comment:"Date de fin (YYYY-mm-dd)" sql:"date"`
	StartTime      *time.Time `gorm:"column:start_time" form:"start_time" json:"start_time,omitempty" comment:"Heure de début (hh:mm)" sql:"time"`
	EndTime        *time.Time `gorm:"column:end_time" form:"end_time" json:"end_time,omitempty" comment:"Heure de fin (hh:mm)" sql:"time"`
	BookerIsReady  int        `gorm:"column:booker_is_ready" form:"booker_is_ready" json:"booker_is_ready" comment:"Le booker est prêt ?" sql:"tinyint(1)"`
	VisitorIsReady int        `gorm:"column:visitor_is_ready" form:"visitor_is_ready" json:"visitor_is_ready" comment:"Les visiteurs sont prêts ?" sql:"tinyint(1)"`
}

//TableName
func (m *RealEstateVisit) TableName() string {
	return "real_estate_visits"
}

//One
func (m *RealEstateVisit) One() (one *RealEstateVisit, err error) {
	one = &RealEstateVisit{}
	err = crudOne(m, one)
	return
}

//All
func (m *RealEstateVisit) All(q *PaginationQuery) (list *[]RealEstateVisit, total uint, err error) {
	list = &[]RealEstateVisit{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *RealEstateVisit) Update() (err error) {
	where := RealEstateVisit{Id: m.Id}
	m.Id = 0

	return crudUpdate(m, where)
}

//Create
func (m *RealEstateVisit) Create() (err error) {
	m.Id = 0

	return mysqlDB.Create(m).Error
}

//Delete
func (m *RealEstateVisit) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
