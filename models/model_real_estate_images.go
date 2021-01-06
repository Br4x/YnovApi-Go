package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

//RealEstateImage
type RealEstateImage struct {
	Id           uint   `gorm:"column:id" form:"id" json:"id" comment:"Id de l'image" sql:"int(11),PRI"`
	IdRealEstate int    `gorm:"column:id_real_estate" form:"id_real_estate" json:"id_real_estate" comment:"Id du bien" sql:"int(11)"`
	Url          string `gorm:"column:url" form:"url" json:"url" comment:"Url de l'image" sql:"varchar(1024)"`
}

//TableName
func (m *RealEstateImage) TableName() string {
	return "real_estate_images"
}

//One
func (m *RealEstateImage) One() (one *RealEstateImage, err error) {
	one = &RealEstateImage{}
	err = crudOne(m, one)
	return
}

//All
func (m *RealEstateImage) All(q *PaginationQuery) (list *[]RealEstateImage, total uint, err error) {
	list = &[]RealEstateImage{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *RealEstateImage) Update() (err error) {
	where := RealEstateImage{Id: m.Id}
	m.Id = 0

	return crudUpdate(m, where)
}

//Create
func (m *RealEstateImage) Create() (err error) {
	m.Id = 0

	return mysqlDB.Create(m).Error
}

//Delete
func (m *RealEstateImage) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
