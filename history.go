package FinTransaction

type History struct {
	OpID   int    `json:"id" db:"id"`
	UserID int    `json:"user_id" db:"user_id"`
	Action string `json:"action" db:"action"`
	Money  int    `json:"money" db:"money"`
}
