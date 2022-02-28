package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

var WalletModel *Wallet

type Wallet struct {
	gorm.Model
	Hash		string `gorm:"type:varchar(200);not null;" json:"hash"`
	Alias       string  `gorm:"type:varchar(40);not null;" json:"alias"`
	Balance     float32 `gorm:"type:decimal(20,6);not null;" json:"balance"`
}

type UpdateWallet struct {
	Credit		float32 `gorm:"type:decimal(20,6);not null;" json:"credit"`
	Debit 		float32 `gorm:"type:decimal(20,6);not null;" json:"debit"`
}

func (w *Wallet) GetAllWallets(db *gorm.DB, wallet *[]Wallet) (err error) {
	if err = db.Find(wallet).Error; err != nil {
		return err
	}
	return nil
}

func (w *Wallet) CreateWallet(db *gorm.DB, wallet *Wallet) (err error) {
	if err = db.Create(wallet).Error; err != nil {
		return err
	}
	return nil
}

func (w *Wallet) GetWalletByID(db *gorm.DB, wallet *Wallet, id string) (err error) {
	if err = db.Where("id = ?", id).First(wallet).Error; err != nil {
		return err
	}
	return nil
}

func (w *Wallet) UpdateWallet(db *gorm.DB, wallet *Wallet, id string) (err error) {
	fmt.Println(wallet)
	db.Save(wallet)
	return nil
}

func (w *Wallet) DeleteWallet(db *gorm.DB, wallet *Wallet, id string) (err error) {
	db.Where("id = ?", id).Delete(wallet)
	return nil
}

func (w *Wallet) CreditWallet(db *gorm.DB, oldWallet *Wallet, wallet *UpdateWallet) (err error) {
	if wallet.Credit < 0 {
		return errors.New("[!] Credit amount cannot be negative")
	}
	oldWallet.Balance += wallet.Credit
	if err := db.Save(&oldWallet).Error; err != nil {
		return err
	}
	return nil
}

func (w *Wallet) DebitWallet(db *gorm.DB, oldWallet *Wallet, wallet *UpdateWallet) (err error) {
	if wallet.Debit < 0 {
		return errors.New("[!] Debit amount cannot be negative")
	}
	if(oldWallet.Balance - wallet.Debit) < 0 {
		return errors.New("[!] A wallet balance cannot go below 0")
	}
	oldWallet.Balance -= wallet.Debit
	if err := db.Save(&oldWallet).Error; err != nil {
		return err
	}
	return nil
}

