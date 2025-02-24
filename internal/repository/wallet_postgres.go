package repository

import (
	fin "FinTransaction"
	"database/sql"
	"fmt"
	"log/slog"
)

type WalletDB struct {
	db *sql.DB
}

func NewWalletDB(db *sql.DB) *WalletDB {
	return &WalletDB{db: db}
}

func (w *WalletDB) CreateWallet(userID int, wallet fin.Wallet) (int, error) {
	tx, err := w.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createWalletQuery := fmt.Sprintf(`
INSERT INTO %s (user_id, balance)
VALUES ($1, $2)
RETURNING id`, walletTable)
	err = tx.QueryRow(createWalletQuery, userID, wallet.Balance).Scan(&id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (w *WalletDB) GetAllWallets(userID int) ([]fin.Wallet, error) {

	var wallets []fin.Wallet
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", walletTable)
	rows, err := w.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch wallets: %w", err)
	}

	for rows.Next() {
		var wallet fin.Wallet
		if err := rows.Scan(&wallet.WalletID, &wallet.UserID, &wallet.Balance); err != nil {
			return wallets, fmt.Errorf("failed to fetch wallet: %w", err)
		}
		wallets = append(wallets, wallet)
	}
	if err := rows.Err(); err != nil {
		return wallets, fmt.Errorf("failed to fetch wallets: %w", err)
	}

	return wallets, nil
}

func (w *WalletDB) GetIDWallet(userID, walletID int) (fin.Wallet, error) {
	var wallet fin.Wallet
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1 AND id = $2", walletTable)
	err := w.db.QueryRow(query, userID, walletID).Scan(&wallet.WalletID, &wallet.UserID, &wallet.Balance)
	if err != nil {
		return fin.Wallet{}, err
	}

	return wallet, nil
}

func (w *WalletDB) DeleteIDWallet(userID int, walletID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND id = $2", walletTable)
	_, err := w.db.Exec(query, userID, walletID)
	if err != nil {
		return err
	}

	return nil
}

func (w *WalletDB) Transfer(userID, id int, wallet fin.TransferWallet) (int, error) {
	tx, err := w.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("failed to begin transaction: %w", err)
	}

	queryMyWallet := fmt.Sprintf(`
								SELECT balance FROM %s
								WHERE id=$1 AND user_id=$2`, walletTable)

	var myBalance int
	err = tx.QueryRow(queryMyWallet, id, userID).Scan(&myBalance)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to transfer wallet query1: %w", err)
	}

	slog.Info(wallet.Username, wallet.Amount)
	queryTransferUserID := fmt.Sprintf(`
								SELECT id FROM %s
								WHERE username=$1`, usersTable)

	var transferUserID int
	err = tx.QueryRow(queryTransferUserID, wallet.Username).Scan(&transferUserID)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to transfer wallet query2: %w", err)
	}

	//queryUserID := fmt.Sprintf(`
	//							SELECT id FROM %s
	//							WHERE username=$1`, walletTable)
	//
	//var TransferUserID int
	//err = tx.QueryRow(queryUserID, wallet.Username).Scan(&TransferUserID)
	//if err != nil {
	//	tx.Rollback()
	//	return 0, fmt.Errorf("failed to transfer wallet query2: %w", err)
	//}

	queryTransferBalance := fmt.Sprintf(`
								SELECT balance FROM %s
								WHERE user_id=$1`, walletTable)

	var transferBalance int
	err = tx.QueryRow(queryTransferBalance, transferUserID).Scan(&transferBalance)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to transfer wallet query3: %w", err)
	}

	newTransferBalance := transferBalance + wallet.Amount
	newMyBalance := myBalance - wallet.Amount

	queryMyUpdate := fmt.Sprintf(`
								UPDATE %s 
								SET balance=$1
								WHERE user_id=$2`, walletTable)

	_, err = tx.Exec(queryMyUpdate, newMyBalance, userID)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to transfer wallet query4: %w", err)
	}

	queryTransferUpdate := fmt.Sprintf(`
								UPDATE %s 
								SET balance=$1
								WHERE user_id=$2`, walletTable)

	_, err = tx.Exec(queryTransferUpdate, newTransferBalance, transferUserID)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to transfer wallet query5: %w", err)
	}

	return newMyBalance, tx.Commit()
}
