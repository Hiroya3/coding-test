package entity

import "time"

type InvoiceID int
type ClientID int
type CompanyID int
type BankIDAccountID int
type UserID int
type Status int

const (
	InvoiceStatusPending = iota
	InvoiceStatusProcessing
	InvoiceStatusPaid
	InvoiceStatusError = 9
)

// 支払い金額
type PayAmount int

// 手数料
type Fee int
type FeeRate float64
type ConsumptionTax int
type ConsumptionTaxRate float64
type TotalAmount int

type UserName string
type CompanyName string
type RepresentativeName string
type Address string
type BankName string
type BranchName string
type AccountName string

type PhoneNumber string
type PostalCode string
type AccountNumber string

type IssueDate time.Time
type PaymentDueDate time.Time
