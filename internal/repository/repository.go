package repository

import (
	fin "FinTransaction"
	"database/sql"
)

type Authorization interface {
	CreateUser(user fin.User) (int, error)
	GetUser(username, password string) (fin.User, error)
}

type Wallet interface {
	CreateWallet(userID int, wallet fin.Wallet) (int, error)
	GetAllWallets(userID int) ([]fin.Wallet, error)
	GetIDWallet(userID, walletID int) (fin.Wallet, error)
	DeleteIDWallet(userID, walletID int) error
	Transfer(userID, id int, input fin.TransferWallet) (int, error)
}

type Repository struct {
	Authorization
	Wallet
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthDB(db),
		Wallet:        NewWalletDB(db),
	}
}
