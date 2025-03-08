package repository

import (
	fin "FinTransaction"
	"database/sql"
	"fmt"
)

type HistoryDB struct {
	db *sql.DB
}

func NewHistoryDB(db *sql.DB) *HistoryDB {
	return &HistoryDB{db: db}
}

func (h *HistoryDB) History(userID int) ([]fin.History, error) {
	var historyAll []fin.History
	query := fmt.Sprintf("SELECT * FROM history WHERE user_id = $1 ORDER BY id DESC LIMIT 10")
	rows, err := h.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get history wallet: %w", err)
	}

	for rows.Next() {
		var history fin.History
		if err := rows.Scan(&history.OpID, &history.UserID, &history.Action, &history.Money); err != nil {
			return historyAll, fmt.Errorf("failed to get history wallet: %w", err)
		}
		historyAll = append(historyAll, history)
	}
	if err := rows.Err(); err != nil {
		return historyAll, fmt.Errorf("failed to get history wallets: %w", err)
	}

	return historyAll, nil
}
