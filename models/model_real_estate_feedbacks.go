package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

//RealEstateFeedback
type RealEstateFeedback struct {
	Id           uint       `gorm:"column:id" form:"id" json:"id" comment:"Id du feedback" sql:"int(11),PRI"`
	IdUser       int        `gorm:"column:id_user" form:"id_user" json:"id_user" comment:"Id de l'utilisateur" sql:"int(11)"`
	IdRealEstate int        `gorm:"column:id_real_estate" form:"id_real_estate" json:"id_real_estate" comment:"Id du bien" sql:"int(11)"`
	Feedback     string     `gorm:"column:feedback" form:"feedback" json:"feedback" comment:"Feedback" sql:"text"`
	Date         *time.Time `gorm:"column:date" form:"date" json:"date,omitempty" comment:"Date (YYYY-mm-dd hh:mm:ss)" sql:"datetime"`
}

//TableName
func (m *RealEstateFeedback) TableName() string {
	return "real_estate_feedbacks"
}

//One
func (m *RealEstateFeedback) One() (one *RealEstateFeedback, err error) {
	one = &RealEstateFeedback{}
	err = crudOne(m, one)
	return
}

//All
func (m *RealEstateFeedback) All(q *PaginationQuery) (list *[]RealEstateFeedback, total uint, err error) {
	list = &[]RealEstateFeedback{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *RealEstateFeedback) Update() (err error) {
	where := RealEstateFeedback{Id: m.Id}
	m.Id = 0

	return crudUpdate(m, where)
}

//Create
func (m *RealEstateFeedback) Create() (err error) {
	m.Id = 0

	return mysqlDB.Create(m).Error
}

//Delete
func (m *RealEstateFeedback) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
