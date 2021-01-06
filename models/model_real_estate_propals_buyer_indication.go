package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

//RealEstatePropalsBuyerIndication
type RealEstatePropalsBuyerIndication struct {
	Id                         uint `gorm:"column:id" form:"id" json:"id" comment:"Id de l'indication" sql:"int(11)"`
	IdPropal                   int  `gorm:"column:id_propal" form:"id_propal" json:"id_propal" comment:"Id de la proposition" sql:"int(11)"`
	IdBuyer                    int  `gorm:"column:id_buyer" form:"id_buyer" json:"id_buyer" comment:"Id de l'acheteur" sql:"int(11)"`
	IdVendor                   int  `gorm:"column:id_vendor" form:"id_vendor" json:"id_vendor" comment:"Id du vendeur" sql:"int(11)"`
	HasTotalInCash             int  `gorm:"column:has_total_in_cash" form:"has_total_in_cash" json:"has_total_in_cash" comment:"a le total en cash ?" sql:"tinyint(1)"`
	HasFinancialSupport        int  `gorm:"column:has_financial_support" form:"has_financial_support" json:"has_financial_support" comment:"a un apport financier ?" sql:"tinyint(1)"`
	HasBigFinancialSupport     int  `gorm:"column:has_big_financial_support" form:"has_big_financial_support" json:"has_big_financial_support" comment:"a un apport financier important ? (>30% du bien)" sql:"tinyint(1)"`
	HasVeryBigFinancialSupport int  `gorm:"column:has_very_big_financial_support" form:"has_very_big_financial_support" json:"has_very_big_financial_support" comment:"a un apport financier très important ? (> 70% du bien)" sql:"tinyint(1)"`
	WantALongSell              int  `gorm:"column:want_a_long_sell" form:"want_a_long_sell" json:"want_a_long_sell" comment:"veut une vente longue ? (+ de 6 mois)" sql:"tinyint(1)"`
}

//TableName
func (m *RealEstatePropalsBuyerIndication) TableName() string {
	return "real_estate_propals_buyer_indication"
}

//One
func (m *RealEstatePropalsBuyerIndication) One() (one *RealEstatePropalsBuyerIndication, err error) {
	one = &RealEstatePropalsBuyerIndication{}
	err = crudOne(m, one)
	return
}

//All
func (m *RealEstatePropalsBuyerIndication) All(q *PaginationQuery) (list *[]RealEstatePropalsBuyerIndication, total uint, err error) {
	list = &[]RealEstatePropalsBuyerIndication{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *RealEstatePropalsBuyerIndication) Update() (err error) {
	where := RealEstatePropalsBuyerIndication{Id: m.Id}
	m.Id = 0

	return crudUpdate(m, where)
}

//Create
func (m *RealEstatePropalsBuyerIndication) Create() (err error) {
	m.Id = 0

	return mysqlDB.Create(m).Error
}

//Delete
func (m *RealEstatePropalsBuyerIndication) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
