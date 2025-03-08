package service

import (
	fin "FinTransaction"
	"FinTransaction/internal/repository"
)

//go:generate mockgen -source=service.go -destination=mock/mock.go

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

type History interface {
	HistoryWallet(userID int) ([]fin.History, error)
}

type Service struct {
	Authorization
	Wallet
	History
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Wallet:        NewWalletService(repos.Wallet),
		History:       NewHistoryService(repos.History),
	}
}
