package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

//ChatComment
type ChatComment struct {
	Id      uint       `gorm:"column:id" form:"id" json:"id" comment:"Id du commentaire" sql:"int(11),PRI"`
	IdChat  int        `gorm:"column:id_chat" form:"id_chat" json:"id_chat" comment:"Id de la conversation" sql:"int(11)"`
	IdUser  int        `gorm:"column:id_user" form:"id_user" json:"id_user" comment:"Id de l'utilisateur qui a écrit le commentaire" sql:"int(11)"`
	Comment string     `gorm:"column:comment" form:"comment" json:"comment" comment:"Commentaire" sql:"text"`
	Date    *time.Time `gorm:"column:date" form:"date" json:"date,omitempty" comment:"Date du commentaire" sql:"datetime"`
}

//TableName
func (m *ChatComment) TableName() string {
	return "chat_comments"
}

//One
func (m *ChatComment) One() (one *ChatComment, err error) {
	one = &ChatComment{}
	err = crudOne(m, one)
	return
}

//All
func (m *ChatComment) All(q *PaginationQuery) (list *[]ChatComment, total uint, err error) {
	list = &[]ChatComment{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *ChatComment) Update() (err error) {
	where := ChatComment{Id: m.Id}
	m.Id = 0

	return crudUpdate(m, where)
}

//Create
func (m *ChatComment) Create() (err error) {
	m.Id = 0

	return mysqlDB.Create(m).Error
}

//Delete
func (m *ChatComment) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
