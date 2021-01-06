package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

//RealEstate
type RealEstate struct {
	Id              uint   `gorm:"column:id" form:"id" json:"id" comment:"Id du bien" sql:"int(11),PRI"`
	IdUser          int    `gorm:"column:id_user" form:"id_user" json:"id_user" comment:"Id du proprietaire" sql:"int(11)"`
	Accroche        string `gorm:"column:accroche" form:"accroche" json:"accroche" comment:"Phrase d'accroche (max 100 caractères)" sql:"varchar(100)"`
	Type            string `gorm:"column:type" form:"type" json:"type" comment:"Type de bien ('apartment','house','vacant_lot','parking','loft','castle','building')" sql:"enum('apartment','house','vacant_lot','parking','loft','castle','building')"`
	NbRooms         int    `gorm:"column:nb_rooms" form:"nb_rooms" json:"nb_rooms" comment:"Nombre de pièce" sql:"int(11)"`
	NbBedroom       int    `gorm:"column:nb_bedroom" form:"nb_bedroom" json:"nb_bedroom" comment:"Nombre de chambre" sql:"int(11)"`
	Description     string `gorm:"column:description" form:"description" json:"description" comment:"Description" sql:"text"`
	Size            int    `gorm:"column:size" form:"size" json:"size" comment:"Superficie (en entier)" sql:"int(11)"`
	Price           int    `gorm:"column:price" form:"price" json:"price" comment:"Prix du bien" sql:"int(11)"`
	Address         string `gorm:"column:address" form:"address" json:"address" comment:"Adresse" sql:"varchar(255)"`
	ZipCode         string `gorm:"column:zip_code" form:"zip_code" json:"zip_code" comment:"Code postal" sql:"varchar(255)"`
	City            string `gorm:"column:city" form:"city" json:"city" comment:"Ville" sql:"varchar(255)"`
	Latitude        string `gorm:"column:latitude" form:"latitude" json:"latitude" comment:"Latitude" sql:"varchar(255)"`
	Longitude       string `gorm:"column:longitude" form:"longitude" json:"longitude" comment:"Longitude" sql:"varchar(255)"`
	EnergyClass     string `gorm:"column:energy_class" form:"energy_class" json:"energy_class" comment:"Classe d'énergie" sql:"varchar(2)"`
	GesClass        string `gorm:"column:ges_class" form:"ges_class" json:"ges_class" comment:"Classe GES" sql:"varchar(2)"`
	HasGarden       int    `gorm:"column:has_garden" form:"has_garden" json:"has_garden" comment:"A un jardin" sql:"tinyint(1)"`
	HasExposedStone int    `gorm:"column:has_exposed_stone" form:"has_exposed_stone" json:"has_exposed_stone" comment:"A des pierres apparentes" sql:"tinyint(1)"`
	HasCimentTiles  int    `gorm:"column:has_ciment_tiles" form:"has_ciment_tiles" json:"has_ciment_tiles" comment:"A des carreaux de ciments" sql:"tinyint(1)"`
	HasParquetFloor int    `gorm:"column:has_parquet_floor" form:"has_parquet_floor" json:"has_parquet_floor" comment:"A du parquet au sol" sql:"tinyint(1)"`
}

//TableName
func (m *RealEstate) TableName() string {
	return "real_estate"
}

//One
func (m *RealEstate) One() (one *RealEstate, err error) {
	one = &RealEstate{}
	err = crudOne(m, one)
	return
}

//All
func (m *RealEstate) All(q *PaginationQuery) (list *[]RealEstate, total uint, err error) {
	list = &[]RealEstate{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *RealEstate) Update() (err error) {
	where := RealEstate{Id: m.Id}
	m.Id = 0

	return crudUpdate(m, where)
}

//Create
func (m *RealEstate) Create() (err error) {
	m.Id = 0

	return mysqlDB.Create(m).Error
}

//Delete
func (m *RealEstate) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
