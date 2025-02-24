package service

import (
	fin "FinTransaction"
	"FinTransaction/internal/repository"
)

type WalletService struct {
	repo repository.Wallet
}

func NewWalletService(repo repository.Wallet) *WalletService {
	return &WalletService{
		repo: repo,
	}
}
func (s *WalletService) CreateWallet(userID int, wallet fin.Wallet) (int, error) {
	return s.repo.CreateWallet(userID, wallet)
}

func (s *WalletService) GetAllWallets(userID int) ([]fin.Wallet, error) {
	return s.repo.GetAllWallets(userID)
}

func (s *WalletService) GetIDWallet(userID, walletID int) (fin.Wallet, error) {
	return s.repo.GetIDWallet(userID, walletID)
}

func (s *WalletService) DeleteIDWallet(userID, walletID int) error {
	return s.repo.DeleteIDWallet(userID, walletID)
}

func (s *WalletService) Transfer(userID, id int, input fin.TransferWallet) (int, error) {
	return s.repo.Transfer(userID, id, input)
}
