package repository

import (
	fin "FinTransaction"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWallet_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	r := NewWalletDB(db)

	type args struct {
		userID int
		wallet fin.Wallet
	}
	type mockBehavior func(id int, args args)

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		id           int
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				userID: 1,
				wallet: fin.Wallet{
					WalletID: 1,
					Balance:  1,
				},
			},
			id: 1,
			mockBehavior: func(id int, args args) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO wallet").
					WithArgs(args.wallet.WalletID, args.wallet.Balance).WillReturnRows(rows)

				mock.ExpectCommit()
			},
		},
		{
			name: "Empty Fields",
			args: args{
				userID: 1,
				wallet: fin.Wallet{
					Balance: 1,
				},
			},
			mockBehavior: func(id int, args args) {
				mock.ExpectBegin()

				mock.ExpectRollback()
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			testCase.mockBehavior(testCase.id, testCase.args)

			got, err := r.CreateWallet(testCase.args.userID, testCase.args.wallet)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, got, testCase.id)
			}
		})
	}
}
