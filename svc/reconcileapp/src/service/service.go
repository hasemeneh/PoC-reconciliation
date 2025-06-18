package service

import (
	"github.com/hasemeneh/PoC-OnlineStore/helper/database"
	"github.com/hasemeneh/PoC-reconciliation/svc/reconcileapp/src/definitions"
	"github.com/hasemeneh/PoC-reconciliation/svc/reconcileapp/src/domain/bank"
	"github.com/hasemeneh/PoC-reconciliation/svc/reconcileapp/src/domain/bankstatement"
	"github.com/hasemeneh/PoC-reconciliation/svc/reconcileapp/src/domain/report"
	"github.com/hasemeneh/PoC-reconciliation/svc/reconcileapp/src/domain/transaction"
	"github.com/hasemeneh/PoC-reconciliation/svc/reconcileapp/src/models"
	"github.com/hasemeneh/PoC-reconciliation/svc/reconcileapp/src/usecase/reconciliation"
)

type Service struct {
	cfg       *models.MainConfig
	Interface struct {
	}
	Usecase struct {
		Reconciliation definitions.ReconcileDefinition
	}
}

func New(cfg *models.MainConfig) *Service {
	db := database.New().Connect(cfg.DBConnectionString)
	transactionRepo := transaction.New(db)
	bankstatementRepo := bankstatement.New(db)
	bankRepo := bank.New(db)
	reportRepo := report.New(db)
	serviceObj := Service{
		cfg: cfg,
	}
	serviceObj.Usecase.Reconciliation = reconciliation.New(&reconciliation.Option{
		Transactions:  transactionRepo,
		Bank:          bankRepo,
		BankStatement: bankstatementRepo,
		Report:        reportRepo,
	})

	return &serviceObj
}
