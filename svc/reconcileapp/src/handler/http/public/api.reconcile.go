package public

import (
	"encoding/csv"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/hasemeneh/PoC-OnlineStore/helper/response"
	"github.com/hasemeneh/PoC-reconciliation/svc/reconcileapp/src/models"
)

func (p *Public) HandleReconcile(r *http.Request) response.HttpResponse {
	err := r.ParseMultipartForm(10 << 20) // Limit file size to 10 MB
	if err != nil {
		return response.NewJsonResponse().SetError(response.NewResponseError(err.Error(), http.StatusBadRequest))
	}

	systemTransactionFile, _, err := r.FormFile("system_transaction_csv_file")
	if err != nil {
		return response.NewJsonResponse().SetError(response.NewResponseError("missing system_transaction_csv_file", http.StatusBadRequest))
	}
	defer systemTransactionFile.Close()

	bankStatementFiles := r.MultipartForm.File["bank_statement_csv_files"]
	if len(bankStatementFiles) == 0 {
		return response.NewJsonResponse().SetError(response.NewResponseError("missing bank_statement_csv_files", http.StatusBadRequest))
	}

	startDate := r.FormValue("start_date")
	endDate := r.FormValue("end_date")
	if startDate == "" || endDate == "" {
		return response.NewJsonResponse().SetError(response.NewResponseError("missing start_date or end_date", http.StatusBadRequest))
	}

	// Validate date format
	validate := validator.New()
	err = validate.Var(startDate, "required,datetime=2006-01-02")
	if err != nil {
		return response.NewJsonResponse().SetError(response.NewResponseError("invalid start_date format", http.StatusBadRequest))
	}
	err = validate.Var(endDate, "required,datetime=2006-01-02")
	if err != nil {
		return response.NewJsonResponse().SetError(response.NewResponseError("invalid end_date format", http.StatusBadRequest))
	}

	// Read bank statement files
	var bankStatementData []*models.CSVData
	for _, fileHeader := range bankStatementFiles {
		bankStatementFile, err := fileHeader.Open()
		if err != nil {
			return response.NewJsonResponse().SetError(response.NewResponseError("failed to open bank_statement_csv_file", http.StatusInternalServerError))
		}
		reader := csv.NewReader(bankStatementFile)
		bankStatementCSV, err := reader.ReadAll()
		if err != nil {
			return response.NewJsonResponse().SetError(response.NewResponseError("failed to parse system_transaction_csv_file", http.StatusInternalServerError))
		}
		bankstatementdatum := models.NewCSVData(bankStatementCSV)
		bankStatementData = append(bankStatementData, &bankstatementdatum)
		bankStatementFile.Close()
	}

	// Convert system transaction data to [][]string
	reader := csv.NewReader(systemTransactionFile)
	systemTransactions, err := reader.ReadAll()
	if err != nil {
		return response.NewJsonResponse().SetError(response.NewResponseError("failed to parse system_transaction_csv_file", http.StatusBadRequest))
	}
	systemTransactionFile.Close()
	// Convert startDate and endDate to time.Time
	startTime, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return response.NewJsonResponse().SetError(response.NewResponseError("invalid start_date format", http.StatusBadRequest))
	}

	endTime, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return response.NewJsonResponse().SetError(response.NewResponseError("invalid end_date format", http.StatusBadRequest))
	}

	// Validate that startTime is before endTime
	if !startTime.Before(endTime) {
		return response.NewJsonResponse().SetError(response.NewResponseError("start_date must be before end_date", http.StatusBadRequest))
	}

	endTime = endTime.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second) // Set end time to the end of the day
	err = p.Service.Usecase.Reconciliation.ReconcileTransactions(
		models.NewCSVData(systemTransactions),
		bankStatementData,
		startTime,
		endTime,
	)
	if err != nil {
		return response.NewJsonResponse().SetError(err).SetMessage(err.Error())
	}
	return response.NewJsonResponse().SetData("successfully reconciled transactions").SetMessage("Reconciliation completed successfully")
}

func (p *Public) HandleGetReconcileReport(r *http.Request) response.HttpResponse {
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")
	if startDate == "" || endDate == "" {
		return response.NewJsonResponse().SetError(response.NewResponseError("missing start_date or end_date", http.StatusBadRequest))
	}

	validate := validator.New()
	if err := validate.Var(startDate, "required,datetime=2006-01-02"); err != nil {
		return response.NewJsonResponse().SetError(response.NewResponseError("invalid start_date format", http.StatusBadRequest))
	}
	if err := validate.Var(endDate, "required,datetime=2006-01-02"); err != nil {
		return response.NewJsonResponse().SetError(response.NewResponseError("invalid end_date format", http.StatusBadRequest))
	}

	startTime, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return response.NewJsonResponse().SetError(response.NewResponseError("invalid start_date format", http.StatusBadRequest))
	}
	endTime, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return response.NewJsonResponse().SetError(response.NewResponseError("invalid end_date format", http.StatusBadRequest))
	}
	if !startTime.Before(endTime) {
		return response.NewJsonResponse().SetError(response.NewResponseError("start_date must be before end_date", http.StatusBadRequest))
	}
	endTime = endTime.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

	report, err := p.Service.Usecase.Reconciliation.GetUnmatchedReportByDateRange(r.Context(), startTime, endTime)
	if err != nil {
		return response.NewJsonResponse().SetError(err).SetMessage(err.Error())
	}
	return response.NewJsonResponse().SetData(report).SetMessage("Report fetched successfully")
}
