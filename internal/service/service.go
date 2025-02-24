package service

import (
	fin "FinTransaction"
	"FinTransaction/internal/repository"
)

type Authorization interface {
	CreateUser(user fin.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Wallet interface {
	CreateWallet(userID int, wallet fin.Wallet) (int, error)
	GetAllWallets(userID int) ([]fin.Wallet, error)
	GetIDWallet(userID, walletID int) (fin.Wallet, error)
	DeleteIDWallet(userID, walletID int) error
	Transfer(userID, id int, input fin.TransferWallet) (int, error)
}

type Service struct {
	Authorization
	Wallet
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Wallet:        NewWalletService(repos.Wallet),
	}
}
