package models

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var _ = time.Thursday

//User
type User struct {
	Id        uint   `gorm:"column:id" form:"id" json:"id" comment:"Id de l'utilisateur" sql:"int(11),PRI"`
	Email     string `gorm:"column:email" form:"email" json:"email" comment:"Email de l'utilisateur qui sert aussi de login" sql:"varchar(255)"`
	Password  string `gorm:"column:password" form:"password" json:"password" comment:"Mot de passe de l'utilisateur" sql:"varchar(1024)"`
	Avatar    string `gorm:"column:avatar" form:"avatar" json:"avatar" comment:"Url de l'avatar de l'utilisateur" sql:"varchar(1024)"`
	FirstName string `gorm:"column:first_name" form:"first_name" json:"first_name" comment:"Prénom de l'utilisateur" sql:"varchar(255)"`
	LastName  string `gorm:"column:last_name" form:"last_name" json:"last_name" comment:"Nom de l'utilisateur" sql:"varchar(255)"`
	Address   string `gorm:"column:address" form:"address" json:"address" comment:"Adresse" sql:"varchar(255)"`
	ZipCode   string `gorm:"column:zip_code" form:"zip_code" json:"zip_code" comment:"Code postal" sql:"varchar(255)"`
	City      string `gorm:"column:city" form:"city" json:"city" comment:"Ville" sql:"varchar(255)"`
	Latitude  string `gorm:"column:latitude" form:"latitude" json:"latitude" comment:"Latitude" sql:"varchar(255)"`
	Longitude string `gorm:"column:longitude" form:"longitude" json:"longitude" comment:"Longitude" sql:"varchar(255)"`
}

//TableName
func (m *User) TableName() string {
	return "users"
}

//One
func (m *User) One() (one *User, err error) {
	one = &User{}
	err = crudOne(m, one)
	return
}

//All
func (m *User) All(q *PaginationQuery) (list *[]User, total uint, err error) {
	list = &[]User{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *User) Update() (err error) {
	where := User{Id: m.Id}
	m.Id = 0
	m.makePassword()

	return crudUpdate(m, where)
}

//Create
func (m *User) Create() (err error) {
	m.Id = 0
	m.makePassword()

	return mysqlDB.Create(m).Error
}

//Delete
func (m *User) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}

//Login
func (m *User) Login(ip string) (*jwtObj, error) {
	m.Id = 0
	if m.Password == "" {
		return nil, errors.New("password is required")
	}
	inputPassword := m.Password
	m.Password = ""
	loginTryKey := "login:" + ip
	loginRetries, _ := mem.GetUint(loginTryKey)
	if loginRetries > uint(viper.GetInt("app.login_try")) {
		memExpire := viper.GetInt("app.mem_expire_min")
		return nil, fmt.Errorf("for too many wrong login retries the %s will ban for login in %d minitues", ip, memExpire)
	}
	//you can implement more detailed login retry rule
	//for i don't know what your login username i can't implement the ip+username rule in my boilerplate project
	// about username and ip retry rule

	err := mysqlDB.Where(m).First(&m).Error
	if err != nil {
		//username fail ip retries add 5
		loginRetries = loginRetries + 5
		mem.Set(loginTryKey, loginRetries)
		return nil, err
	}
	//password is set to bcrypt check
	if err := bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(inputPassword)); err != nil {
		// when password failed reties will add 1
		loginRetries = loginRetries + 1
		mem.Set(loginTryKey, loginRetries)
		return nil, err
	}
	m.Password = ""
	key := fmt.Sprintf("login:%d", m.Id)

	//save login user  into the memory store

	data, err := jwtGenerateToken(m)
	mem.Set(key, data)
	return data, err
}

func (m *User) makePassword() {
	if m.Password != "" {
		if bytes, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost); err != nil {
			logrus.WithError(err).Error("bcrypt making password is failed")
		} else {
			m.Password = string(bytes)
		}
	}
}
