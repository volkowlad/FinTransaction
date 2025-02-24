package FinTransaction

type Wallet struct {
	WalletID int `json:"id" db:"id"`
	UserID   int `json:"user_id" db:"user_id"`
	Balance  int `json:"balance" db:"balance"`
}

type TransferWallet struct {
	Username string `json:"transfer_username" db:"id"`
	Amount   int    `json:"amount" db:"amount"`
}
