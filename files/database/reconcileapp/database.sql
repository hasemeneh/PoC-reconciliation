-- phpMyAdmin SQL Dump
-- version 4.9.0.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Apr 02, 2021 at 03:37 AM
-- Server version: 10.3.15-MariaDB
-- PHP Version: 7.3.6
USE reconcile_db;
SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `reconcile_db`
--


-- Drop tables if they exist (for rerun safety)
DROP TABLE IF EXISTS Bank;
DROP TABLE IF EXISTS Transaction;
DROP TABLE IF EXISTS BankStatement;


-- Table: Bank
CREATE TABLE Bank (
    bankID INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    code VARCHAR(50) UNIQUE NOT NULL
);

-- Table: Transaction
CREATE TABLE Transaction (
    trxID VARCHAR(64) PRIMARY KEY,
    amount DECIMAL(15, 2) NOT NULL,
    type ENUM('DEBIT', 'CREDIT') NOT NULL,
    transactionTime DATETIME NOT NULL,
    isReconciled BOOLEAN DEFAULT FALSE,
    -- Foreign key to Bank table
    bankID INT NOT NULL,
    FOREIGN KEY (bankID) REFERENCES Bank(bankID)
);

-- Table: BankStatement
CREATE TABLE BankStatement (
    unique_identifier VARCHAR(100) NOT NULL,
    amount DECIMAL(15, 2) NOT NULL,
    statement_date DATE NOT NULL,
    bankID INT NOT NULL,
    isReconciled BOOLEAN DEFAULT FALSE,
    PRIMARY KEY (unique_identifier, bankID),
    FOREIGN KEY (bankID) REFERENCES Bank(bankID)
);


-- Main Reconciliation Report
CREATE TABLE ReconciliationReport (
    id INT AUTO_INCREMENT PRIMARY KEY,
    upload_date DATETIME NOT NULL,
    report_date_start DATE NOT NULL,
    report_date_end DATE NOT NULL,
    total_transactions INT NOT NULL,
    matched_transactions INT NOT NULL,
    unmatched_transactions INT NOT NULL,
    total_discrepancies DECIMAL(15, 2) NOT NULL
);

-- Unmatched system transactions (present in system but not in bank statement)
CREATE TABLE UnmatchedSystemTransaction (
    id INT AUTO_INCREMENT PRIMARY KEY,
    report_id INT NOT NULL,
    trxID VARCHAR(64),
    amount DECIMAL(15, 2),
    transactionTime DATETIME,
    type ENUM('DEBIT', 'CREDIT'),
    bankID INT,
    FOREIGN KEY (report_id) REFERENCES ReconciliationReport(id),
    FOREIGN KEY (bankID) REFERENCES Bank(bankID)
);

-- Unmatched bank statement records (present in bank statement but not in system)
CREATE TABLE UnmatchedBankStatement (
    id INT AUTO_INCREMENT PRIMARY KEY,
    report_id INT NOT NULL,
    unique_identifier VARCHAR(100),
    amount DECIMAL(15, 2),
    date DATE,
    bankID INT,
    FOREIGN KEY (report_id) REFERENCES ReconciliationReport(id),
    FOREIGN KEY (bankID) REFERENCES Bank(bankID)
);


INSERT INTO Bank (bankID, name, code) VALUES (1, 'Bank Alpha', 'BANK001');
INSERT INTO Bank (bankID, name, code) VALUES (2, 'Bank Beta', 'BANK002');
INSERT INTO Bank (bankID, name, code) VALUES (3, 'Bank Gamma', 'BANK003');
INSERT INTO Bank (bankID, name, code) VALUES (4, 'Bank Delta', 'BANK004');
INSERT INTO Bank (bankID, name, code) VALUES (5, 'Bank Epsilon', 'BANK005');
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('d790e576-a18f-46dc-9a74-6e1fd154cf54', 6397.87, 'DEBIT', '2025-01-23 20:26:50', 2);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('a6108c60-b19d-497c-9b01-f481c6dc5baf', 1403.98, 'DEBIT', '2024-07-01 04:26:50', 5);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('b022cbb3-2661-4bdd-ad86-dcc2371d7a3d', 878.52, 'CREDIT', '2025-05-28 03:26:50', 1);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('c1089a39-e972-4f4e-82e3-73bbdb7dfa8b', 2194.19, 'DEBIT', '2024-08-29 21:26:50', 5);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('ba6ddafe-99bd-4cf6-8bb2-351110ab9191', 4201.0, 'CREDIT', '2024-08-15 19:26:50', 1);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('341a6b56-1fca-47be-9852-5c737858f8f7', 7590.49, 'DEBIT', '2024-06-20 14:26:50', 3);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('44bc6b29-85ff-401a-938e-6c47b4b8d2d4', 2785.93, 'DEBIT', '2024-12-23 00:26:50', 1);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('d04fc90f-41f6-4ac3-9834-4f457c906b3d', 3805.47, 'CREDIT', '2024-12-18 08:26:50', 3);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('fbe2f73c-57ab-4e58-95bc-52a16cadfad8', 8073.21, 'CREDIT', '2024-09-12 00:26:50', 4);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('295b0498-2311-452a-afe5-a71cc32207b6', 797.21, 'CREDIT', '2024-07-26 08:26:50', 3);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('84c82acf-4261-4725-beba-f0c703fb18d1', 5777.75, 'DEBIT', '2025-05-20 06:26:50', 2);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('6c445034-7dcc-4d40-b2d4-8fa883b5cc07', 7732.95, 'DEBIT', '2025-02-14 00:26:50', 4);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('57b592a2-0921-49e6-8587-af6da7953fe6', 2786.96, 'CREDIT', '2025-03-21 16:26:50', 3);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('e4ee3b76-caf6-4eb6-b85d-781857e0c59c', 2102.98, 'CREDIT', '2024-06-18 06:26:50', 1);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('0f0a8ac2-1d3e-4f94-ada5-6febf2a51b2c', 6095.22, 'DEBIT', '2024-09-12 04:26:50', 2);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('c8e2577c-43a4-4786-803b-57f2e1ba50d8', 1642.39, 'CREDIT', '2025-01-25 07:26:50', 5);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('fecc4340-0720-499c-b09e-8e8d656ec64e', 2203.96, 'CREDIT', '2025-05-15 20:26:50', 1);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('1adefe82-c567-46f0-9c5f-de8144385fbc', 8052.41, 'CREDIT', '2025-01-27 01:26:50', 2);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('772a9b39-0e7e-41cc-8a25-7e28b3dc0015', 9132.23, 'CREDIT', '2025-02-24 07:26:50', 4);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('32ed4cce-bee9-4ba0-8ca6-7966b234121f', 3962.36, 'CREDIT', '2025-03-31 19:26:50', 2);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('397d2e52-6215-4a21-8fa0-aec35d924e85', 2473.81, 'CREDIT', '2024-08-17 14:26:50', 5);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('4ba4a894-c464-4b45-a6c1-dc2187b60ecc', 4000.01, 'DEBIT', '2025-04-03 11:26:50', 4);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('a182cb09-4543-4c22-bd90-28f48709c12d', 918.19, 'DEBIT', '2025-04-17 23:26:50', 2);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('34bb9062-e303-4bb3-a9cb-8470b792599b', 7922.87, 'CREDIT', '2024-08-12 01:26:50', 4);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('383914ff-79fb-4d7b-b3d5-32ca19b3835a', 3822.38, 'CREDIT', '2024-09-15 19:26:50', 5);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('cdb415c5-8a20-4a16-9cbb-24957f76c480', 8609.19, 'DEBIT', '2024-06-29 04:26:50', 1);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('cc61736e-8004-4eca-8b9f-a0b6422dc647', 6820.29, 'CREDIT', '2024-07-19 17:26:50', 1);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('1de26231-c104-4398-8f16-0ddc260f4816', 2942.07, 'DEBIT', '2024-10-24 03:26:50', 3);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('279378e4-e8b6-422d-8777-e8f2e6c46001', 9719.16, 'DEBIT', '2024-09-27 00:26:50', 3);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('28bfe736-8062-4640-9d12-110108858786', 8418.28, 'DEBIT', '2025-03-26 16:26:50', 2);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('2c2d3c42-4cce-4c65-8839-6d469c6501b9', 5398.4, 'DEBIT', '2024-08-10 17:26:50', 4);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('d242315f-7143-416e-8999-a9ce838320ae', 204.57, 'CREDIT', '2025-01-06 20:26:50', 1);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('1e337c34-92eb-43de-9b5b-6df91183991a', 2416.31, 'DEBIT', '2025-04-30 04:26:50', 4);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('fb9fd2af-6a6d-4ce2-8f87-11f010e8f4b9', 8162.07, 'DEBIT', '2025-04-08 06:26:50', 4);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('dc051178-2cba-4b2d-84ee-ceed1fb2a929', 9469.02, 'DEBIT', '2025-01-28 11:26:50', 5);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('156d6930-02b4-4cae-b31b-86cf7cab4b77', 4237.15, 'DEBIT', '2024-09-09 04:26:50', 2);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('4dc5ae8f-c3ea-48a9-afa7-c27795c4bde9', 7132.36, 'CREDIT', '2024-07-04 07:26:50', 3);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('bc95f928-b0f5-4a4e-bdbe-33e063478a36', 4386.62, 'CREDIT', '2025-04-12 20:26:50', 2);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('8bf0fbc5-ad7a-4c43-82b9-12e7e6058ca0', 649.62, 'DEBIT', '2024-08-15 10:26:50', 2);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('648b5767-d1a0-4463-b6cb-4143648a828f', 5888.52, 'DEBIT', '2025-05-07 05:26:50', 1);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('f5c8ed3e-5a10-4cc0-96e4-3f7b6f6186f4', 2297.13, 'DEBIT', '2024-12-26 01:26:50', 5);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('6cf02134-a032-4072-8cf1-5e249b29812e', 2387.67, 'CREDIT', '2025-02-23 10:26:50', 2);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('56ad1c66-5b8d-4054-bdf2-c45216bae796', 7236.29, 'CREDIT', '2025-02-08 12:26:50', 4);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('2e402b37-e595-4461-8957-5a846e18c2a7', 1912.2, 'DEBIT', '2024-07-10 14:26:50', 3);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('78539b74-7a5a-4282-8c8a-43c18439afc0', 4241.55, 'CREDIT', '2025-05-16 06:26:50', 1);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('79ebc753-ebae-4b47-9e15-456a18c597c4', 615.52, 'CREDIT', '2025-04-18 20:26:50', 2);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('b78e00d1-d1d5-4c02-b975-8fd972a406e2', 1910.19, 'CREDIT', '2025-04-02 14:26:50', 2);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('df4fe9f8-6e03-4898-b8ee-e58db5672366', 2792.67, 'DEBIT', '2025-05-05 13:26:50', 5);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('f9ffe918-1bea-4112-b8fc-5c9d08d3fa1e', 988.11, 'DEBIT', '2025-04-26 20:26:50', 2);
INSERT INTO Transaction (trxID, amount, type, transactionTime, bankID) VALUES ('fad05b93-16e8-455c-92c6-7e6de372b267', 4070.08, 'CREDIT', '2025-02-23 15:26:50', 1);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('bbf27357-4734-4e25-85c6-8ae335ef636a', -2530.44, '2025-06-12', 4);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('9e0e1ac0-b9dc-405c-8b3a-fcb6e6fb423a', -1021.95, '2024-10-24', 3);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('7094ea01-bcf9-4e02-ac15-8419a840d244', 1345.11, '2024-09-02', 4);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('98ca81af-d377-4641-99b2-fae32f18b98d', -2678.05, '2025-01-13', 2);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('96414e74-a4a9-40eb-a0de-913abcf56e26', 9530.64, '2024-08-21', 5);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('c643475a-bf44-4f89-93b7-74ae7db86276', -4085.64, '2025-01-04', 1);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('4f297f3e-2370-461d-bf6a-b2d850e77140', -4247.86, '2024-10-12', 5);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('a02b47d9-d624-4452-bbd5-091146c7b6ce', 8790.81, '2024-09-15', 2);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('c5406d5e-42bb-4468-ba10-e715e4c2782d', -4146.81, '2024-09-26', 1);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('77dae9ed-6569-48b2-bb0f-2eb27427cf1a', 7770.14, '2025-05-09', 5);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('b8c75ee8-9057-4834-a905-3209972b9b54', -3980.59, '2025-02-13', 4);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('0f64cd90-6797-4f7b-b06c-dada64895d0f', -3201.7, '2024-08-26', 2);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('f056a14b-7f06-459f-a451-1fc36d4dd92d', 3683.79, '2025-05-24', 5);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('e5eb5969-b3e9-4c89-8615-a44e23b9c252', -3770.25, '2024-07-12', 5);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('03943c05-8265-4296-8c37-4fb587849fe6', 3478.4, '2025-01-03', 3);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('e6fb458e-da59-4edd-98e2-8bc66d8c49dd', -1936.11, '2025-01-04', 2);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('f412eebf-55b4-428e-ad06-fc123d2fe14a', -1015.78, '2025-04-07', 3);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('7b492d2d-f4fa-4eb8-b821-78159c58cdbe', 1858.37, '2025-05-07', 1);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('56868667-c189-44e9-8765-c4f2df3a61f1', 1874.28, '2024-08-29', 1);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('53a17b64-590f-4112-b36d-a64ccd2dbdff', -3901.09, '2025-02-24', 5);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('76ef92b4-09e2-43a0-b6f6-4dcf9bf8b43c', -1021.99, '2024-12-17', 1);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('5fddb748-cb18-42d4-ad70-eec459ae271e', 8189.05, '2024-12-06', 3);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('fcf3cd19-c40e-400c-b3aa-a6934f00266e', -2633.8, '2024-09-08', 3);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('b22ec048-d4d1-4d14-a209-4bd5d576f823', 4175.17, '2024-07-14', 5);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('53b9e493-8ed4-43ec-b80d-46fe91a92064', -4882.65, '2024-09-03', 3);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('0ff9bd62-15d7-463b-b080-06e9b6b4a376', 8976.34, '2025-04-21', 2);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('e3595b80-b805-41dc-a6db-433b33b47fcc', -1033.01, '2025-04-20', 5);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('bd8a67ad-f704-4303-9aab-21ff9976c0b8', -2668.28, '2025-01-20', 5);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('211519bb-d52d-4225-96e4-8560f8ec63cc', -1840.56, '2024-12-20', 2);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('043b53b1-20f1-42a9-a418-e3ed4564de52', 5312.46, '2025-01-29', 5);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('9889854d-0f62-42f8-857d-54aacbcc2cc3', 2327.98, '2025-05-18', 1);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('8d1a2a35-6786-461e-8415-d0870b10093f', 4514.27, '2025-01-23', 1);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('9c976b44-8303-4af2-aae5-a76ebcb7a5ed', -4946.81, '2025-04-08', 3);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('ad4ed111-5890-47f4-8606-11e46cf1e80d', -2576.27, '2024-10-30', 5);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('ffd57140-60d3-4465-922e-f1b21bd945b6', 5585.01, '2024-08-30', 1);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('e02f7c37-397c-4021-8b4a-8c1911bd819f', -3321.89, '2024-06-25', 2);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('c38373e4-17e6-435a-931c-06f520bb88c6', 3183.85, '2024-12-06', 5);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('d9238bb8-12f4-4ffe-9133-5a88a16b1ea0', 3287.88, '2024-11-05', 2);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('174e184e-3e95-4682-9de4-a659595dc654', -4372.57, '2024-12-09', 1);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('bd06aefe-f452-44a8-b7b7-811880a712b7', 8483.87, '2025-02-26', 2);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('6ab61c7c-2ce8-4784-bf05-7e8789a83da2', 5004.0, '2024-12-14', 5);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('5a22b5d9-aa12-4d0f-a72e-b093a7432689', 8262.02, '2024-11-17', 5);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('6ca0c477-f880-44c1-b659-50daee7efbe6', 6242.16, '2025-02-12', 2);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('c7dd2a89-22a5-4660-8bd1-0e6c92296db5', 9643.09, '2025-03-15', 4);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('5891ef21-45f2-4271-ab80-3d811ab90fd5', -4628.2, '2024-12-25', 4);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('7c1cc7ae-29ef-4b97-8d0d-cd9ed1a05694', 7033.53, '2025-02-06', 3);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('168d6590-063f-4f39-b340-82e4c47ecc7b', -2612.03, '2024-06-19', 1);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('97e5438e-9429-4cb1-83ed-6569deb53369', 738.13, '2025-05-25', 4);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('f2e24163-b69e-4329-83f5-d92938a6cbce', -1663.49, '2024-10-21', 3);
INSERT INTO BankStatement (unique_identifier, amount, statement_date, bankID) VALUES ('0e43b48d-22fb-4e1a-bc2b-f946bd4041c0', -422.14, '2025-02-17', 2);

COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
