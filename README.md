# PoC-reconciliation

This project is a Proof of Concept (PoC) for a reconciliation system implemented in Go.

## Description

PoC-reconciliation demonstrates a basic reconciliation process that compares and matches records from different data sources to identify discrepancies.

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.18 or higher)
- [Docker](https://www.docker.com/)

### Steps to Run

1. **Clone the repository:**
    ```bash
    git clone https://github.com/hasemeneh/PoC-reconciliation.git
    cd PoC-reconciliation
    ```

2. **Start Docker containers:**
    ```bash
    make docker-start
    ```

3. **Prepare the database:**
    ```bash
    make docker-rebuilddb-reconcile
    ```

4. **Run the application:**
    ```bash
    make docker-run-reconcile
    ```


## Resetting Data

If you need to reset the data and rerun the reconciliation process, repeat steps 3 and 4:

1. **Rebuild the database:**
    ```bash
    make docker-rebuilddb-reconcile
    ```

2. **Run the application again:**
    ```bash
    make docker-run-reconcile
    ```

This will clear existing data, reinitialize the database, and execute the reconciliation process from a clean state.

## API Endpoint Test Cases

The following describes the Postman test cases for the `/api/reconcile` endpoint. These tests cover various scenarios to ensure the reconciliation process works as expected.

### Test Cases

#### 1. POST `/reconcile` - Success, No Discrepancy

**Description:**  
Uploads a system transaction CSV file and two bank statement CSV files with matching data and a date range.  
**Expected Result:**  
The reconciliation process completes successfully with zero discrepancies found.

#### 2. POST `/reconcile` - Transaction Discrepancy

**Description:**  
Uploads a system transaction CSV file and a bank statement CSV file where the system transaction contains discrepancies not present in the bank statement.  
**Expected Result:**  
The response highlights discrepancies found in the system transaction data.

#### 3. POST `/reconcile` - Bank Statement Discrepancy

**Description:**  
Uploads a system transaction CSV file and a bank statement CSV file where the bank statement contains discrepancies not present in the system transaction.  
**Expected Result:**  
The response highlights discrepancies found in the bank statement data.

#### 4. GET `/reconcile` - Query by Date Range

**Description:**  
Sends a GET request with `start_date` and `end_date` as query parameters to retrieve reconciliation results for the specified date range.  
**Expected Result:**  
Returns reconciliation results for the given date range. No files are uploaded in this request.

---

These test cases help verify the correctness and robustness of the reconciliation API under different scenarios.

## How to Test Using Postman


You can use Postman to manually test the `/api/reconcile` endpoint with the following steps:

**please do reset for each testcase**

### 1. Import the Postman Collection

1. Open Postman.
2. Click on **Import**.
3. Select the **Raw Text** tab and paste the provided Postman JSON.
4. Click **Continue** and then **Import**.

### 2. Prepare Sample CSV Files

- Ensure you have the required CSV files as referenced in the test cases:
    - `sample transaction.csv`
    - `sample bankstatement.csv`
    - `sample_bankstatement2.csv`
- Place them in an accessible directory on your machine.

### 3. Configure File Paths

- For each request in the collection, update the file paths in the form-data fields to match the location of your CSV files.

### 4. Run Test Cases

#### a. POST `/reconcile` - Success, No Discrepancy

- Select the request named **POST /reconcile success 0 discrep**.
- Ensure all files and date fields are set.
- Click **Send**.
- Verify the response indicates zero discrepancies.

#### b. POST `/reconcile` - Transaction Discrepancy

- Select **POST /reconcile transcation discrep**.
- Set the appropriate files and date range.
- Click **Send**.
- Check that discrepancies in the system transaction are reported.

#### c. POST `/reconcile` - Bank Statement Discrepancy

- Select **POST /reconcile bankstatement discrep**.
- Set the appropriate files and date range.
- Click **Send**.
- Check that discrepancies in the bank statement are reported.

#### d. GET `/reconcile` - Query by Date Range

- Select **GET /reconcile**.
- Ensure the `start_date` and `end_date` query parameters are set.
- Click **Send**.
- Verify the reconciliation results for the specified date range.

### 5. Review Responses

- Each response should match the expected results described in the test cases section above.

---

By following these steps, you can validate the reconciliation API using the provided Postman collection and sample data.