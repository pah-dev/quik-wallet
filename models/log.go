package models

import (
	"gorm.io/gorm"
)

var LogModel *Log

type Log struct {
	gorm.Model
	Hash		string `gorm:"type:varchar(200);not null;" json:"hash"`
	Operation	string `gorm:"type:varchar(20);not null;" json:"operation"`
	Amount		float32 `gorm:"type:decimal(20,6);not null;" json:"amount"`
	Detail		string `gorm:"type:varchar(250);not null;" json:"detail"`
}

func (w *Log) CreateLog(db *gorm.DB, log *Log) (err error) {
	if err = db.Create(log).Error; err != nil {
		return err
	}
	return nil
}