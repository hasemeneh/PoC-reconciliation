package reconciliation

import "github.com/hasemeneh/PoC-reconciliation/svc/reconcileapp/src/repositories"

type reconciliationUsecase struct {
	Transactions  repositories.TransactionsRepo
	Bank          repositories.BankRepositories
	BankStatement repositories.BankStatementRepositories
	Report        repositories.ReportRepositories
}
type Option struct {
	Transactions  repositories.TransactionsRepo
	Bank          repositories.BankRepositories
	BankStatement repositories.BankStatementRepositories
	Report        repositories.ReportRepositories
}

func New(o *Option) *reconciliationUsecase {
	return &reconciliationUsecase{
		Transactions:  o.Transactions,
		Bank:          o.Bank,
		BankStatement: o.BankStatement,
		Report:        o.Report,
	}
}
