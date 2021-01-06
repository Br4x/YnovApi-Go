package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

//RealEstatePropal
type RealEstatePropal struct {
	Id             uint      `gorm:"column:id" form:"id" json:"id" comment:"Id de la proposition" sql:"int(11),PRI"`
	IdBooker       int       `gorm:"column:id_booker" form:"id_booker" json:"id_booker" comment:"Id du booker" sql:"int(11)"`
	IdBuyer        int       `gorm:"column:id_buyer" form:"id_buyer" json:"id_buyer" comment:"Id du visitor" sql:"int(11)"`
	IdRealEstate   int       `gorm:"column:id_real_estate" form:"id_real_estate" json:"id_real_estate" comment:"Id du bien" sql:"int(11)"`
	IdVendor       int       `gorm:"column:id_vendor" form:"id_vendor" json:"id_vendor" comment:"Id du propriétaire" sql:"int(11)"`
	Price          float64   `gorm:"column:price" form:"price" json:"price" comment:"Montant de la proposition" sql:"double"`
	BuyerFirstName string    `gorm:"column:buyer_first_name" form:"buyer_first_name" json:"buyer_first_name" comment:"Prénom de l'acheteur" sql:"varchar(255)"`
	BuyerLastName  string    `gorm:"column:buyer_last_name" form:"buyer_last_name" json:"buyer_last_name" comment:"Nom de l'acheteur" sql:"varchar(255)"`
	BuyerAddress   string    `gorm:"column:buyer_address" form:"buyer_address" json:"buyer_address" comment:"Adresse de l'acheteur" sql:"varchar(255)"`
	BuyerZipCode   string    `gorm:"column:buyer_zip_code" form:"buyer_zip_code" json:"buyer_zip_code" comment:"Code postal de l'acheteur" sql:"varchar(255)"`
	BuyerCity      string    `gorm:"column:buyer_city" form:"buyer_city" json:"buyer_city" comment:"Ville de l'acheteur" sql:"varchar(255)"`
	ExpirationDate time.Time `gorm:"column:expiration_date" form:"expiration_date" json:"expiration_date" comment:"Date d'expiration de l'offre" sql:"date"`
	IsAccepted     int       `gorm:"column:is_accepted" form:"is_accepted" json:"is_accepted" comment:"Est accepté?" sql:"tinyint(1)"`
}

//TableName
func (m *RealEstatePropal) TableName() string {
	return "real_estate_propals"
}

//One
func (m *RealEstatePropal) One() (one *RealEstatePropal, err error) {
	one = &RealEstatePropal{}
	err = crudOne(m, one)
	return
}

//All
func (m *RealEstatePropal) All(q *PaginationQuery) (list *[]RealEstatePropal, total uint, err error) {
	list = &[]RealEstatePropal{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *RealEstatePropal) Update() (err error) {
	where := RealEstatePropal{Id: m.Id}
	m.Id = 0

	return crudUpdate(m, where)
}

//Create
func (m *RealEstatePropal) Create() (err error) {
	m.Id = 0

	return mysqlDB.Create(m).Error
}

//Delete
func (m *RealEstatePropal) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
