package transaction

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hasemeneh/PoC-reconciliation/svc/reconcileapp/src/models"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestGetTransactionByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	transactionDomain := &transactionDomain{DB: sqlxDB}

	ctx := context.Background()
	trxID := "trx123"
	expectedTransaction := &models.TransactionModel{
		TrxID:           trxID,
		Amount:          100.0,
		Type:            "credit",
		TransactionTime: time.Now(),
		BankID:          1,
	}

	rows := sqlmock.NewRows([]string{"trxID", "amount", "type", "transactionTime", "bankID"}).
		AddRow(expectedTransaction.TrxID, expectedTransaction.Amount, expectedTransaction.Type, expectedTransaction.TransactionTime, expectedTransaction.BankID)

	mock.ExpectQuery("SELECT `trxID`, `amount`, `type`, `transactionTime`, `bankID` FROM `Transaction` WHERE `trxID` = ?;").
		WithArgs(trxID).
		WillReturnRows(rows)

	transaction, err := transactionDomain.GetTransactionByID(ctx, trxID)
	assert.NoError(t, err)
	assert.Equal(t, expectedTransaction, transaction)
}

func TestInsertTransaction(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	transactionDomain := &transactionDomain{DB: sqlxDB}

	ctx := context.Background()
	transaction := &models.TransactionModel{
		TrxID:           "trx123",
		Amount:          100.0,
		Type:            "credit",
		TransactionTime: time.Now(),
		BankID:          1,
	}

	mock.ExpectExec("INSERT INTO `Transaction` \\(`trxID`, `amount`, `type`, `transactionTime`, `bankID`\\) VALUES \\(\\?, \\?, \\?, \\?, \\?\\);").
		WithArgs(transaction.TrxID, transaction.Amount, transaction.Type, transaction.TransactionTime, transaction.BankID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = transactionDomain.InsertTransaction(ctx, nil, transaction)
	assert.NoError(t, err)
}

func TestDeleteTransaction(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	transactionDomain := &transactionDomain{DB: sqlxDB}

	ctx := context.Background()
	trxID := "trx123"

	mock.ExpectExec("DELETE FROM `Transaction` WHERE `trxID` = \\?;").
		WithArgs(trxID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = transactionDomain.DeleteTransaction(ctx, nil, trxID)
	assert.NoError(t, err)
}

func TestUpdateTransaction(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	transactionDomain := &transactionDomain{DB: sqlxDB}

	ctx := context.Background()
	transaction := &models.TransactionModel{
		TrxID:           "trx123",
		Amount:          200.0,
		Type:            "debit",
		TransactionTime: time.Now(),
		BankID:          4,
	}

	mock.ExpectExec("UPDATE `Transaction` SET `amount` = \\?, `type` = \\?, `transactionTime` = \\?, `bankID` = \\? WHERE `trxID` = \\?;").
		WithArgs(transaction.Amount, transaction.Type, transaction.TransactionTime, transaction.BankID, transaction.TrxID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = transactionDomain.UpdateTransaction(ctx, nil, transaction)
	assert.NoError(t, err)
}

func TestGetAllTransactions(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	transactionDomain := &transactionDomain{DB: sqlxDB}

	ctx := context.Background()
	expectedTransactions := []*models.TransactionModel{
		{
			TrxID:           "trx123",
			Amount:          100.0,
			Type:            "credit",
			TransactionTime: time.Now(),
			BankID:          1,
		},
		{
			TrxID:           "trx456",
			Amount:          200.0,
			Type:            "debit",
			TransactionTime: time.Now(),
			BankID:          4,
		},
	}

	rows := sqlmock.NewRows([]string{"trxID", "amount", "type", "transactionTime", "bankID"}).
		AddRow(expectedTransactions[0].TrxID, expectedTransactions[0].Amount, expectedTransactions[0].Type, expectedTransactions[0].TransactionTime, expectedTransactions[0].BankID).
		AddRow(expectedTransactions[1].TrxID, expectedTransactions[1].Amount, expectedTransactions[1].Type, expectedTransactions[1].TransactionTime, expectedTransactions[1].BankID)

	mock.ExpectQuery("SELECT `trxID`, `amount`, `type`, `transactionTime`, `bankID` FROM `Transaction`;").
		WillReturnRows(rows)

	transactions, err := transactionDomain.GetAllTransactions(ctx)
	assert.NoError(t, err)
	assert.Equal(t, expectedTransactions, transactions)
}
