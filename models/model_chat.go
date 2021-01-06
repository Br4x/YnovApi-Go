package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

//Chat
type Chat struct {
	Id      uint `gorm:"column:id" form:"id" json:"id" comment:"Id de la conversation" sql:"int(11),PRI"`
	IdUser1 int  `gorm:"column:id_user1" form:"id_user1" json:"id_user1" comment:"Id utilisateur 1" sql:"int(11)"`
	IdUser2 int  `gorm:"column:id_user2" form:"id_user2" json:"id_user2" comment:"Id utilisateur 2" sql:"int(11)"`
}

//TableName
func (m *Chat) TableName() string {
	return "chat"
}

//One
func (m *Chat) One() (one *Chat, err error) {
	one = &Chat{}
	err = crudOne(m, one)
	return
}

//All
func (m *Chat) All(q *PaginationQuery) (list *[]Chat, total uint, err error) {
	list = &[]Chat{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *Chat) Update() (err error) {
	where := Chat{Id: m.Id}
	m.Id = 0

	return crudUpdate(m, where)
}

//Create
func (m *Chat) Create() (err error) {
	m.Id = 0

	return mysqlDB.Create(m).Error
}

//Delete
func (m *Chat) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
