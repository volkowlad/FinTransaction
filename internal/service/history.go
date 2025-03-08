package service

import (
	fin "FinTransaction"
	"FinTransaction/internal/repository"
)

type HistoryService struct {
	repo repository.History
}

func NewHistoryService(repo repository.History) *HistoryService {
	return &HistoryService{
		repo: repo,
	}
}

func (s *HistoryService) HistoryWallet(userID int) ([]fin.History, error) {
	return s.repo.History(userID)
}
