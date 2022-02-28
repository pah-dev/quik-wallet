package models

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var AuthModel *Auth

type Auth struct {
	gorm.Model
	Username	string `gorm:"type:varchar(255);unique;not null" json:"username" form:"username"`
	Password  	string `gorm:"type:varchar(255);not null" json:"password" form:"password"`
}

func CheckAuth(db *gorm.DB, username string, password string) (bool, error) {
	auth := Auth{Username : username, Password : password}
	err := db.Select("id").Where(auth).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.Warn(err.Error())
		return false, err
	}
	if auth.ID > 0 {
		return true, nil
	}
	return false, nil
}