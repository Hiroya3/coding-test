package entity

import "time"

type Invoice struct {
	InvoiceID          InvoiceID
	InvoiceCompany     Company
	UserName           UserName
	InvoiceClient      Client
	InvoiceBankAccount BankAccount
	IssueDate          IssueDate
	PayAmount          PayAmount // 支払い金額
	Fee                Fee
	FeeRate            FeeRate
	ConsumptionTax     ConsumptionTax
	ConsumptionTaxRate ConsumptionTaxRate
	TotalAmount        TotalAmount
	PaymentDueDate     PaymentDueDate
	Status             Status
}

func NewInvoice(
	clientID int,
	companyName, representativeName, phoneNumber, postalCode, address string,
	userName string,
	clientCompanyName, clientRepresentativeName, clientPhoneNumber, clientPostalCode, clientAddress string,
	bankName, branchName, accountNumber, accountName string,
	issueDate, paymentDueDate time.Time,
	payAmount, fee, consumptionTax int,
	feeRate, consumptionTaxRate float64,
) (Invoice, error) {
	// TODO domain作成時にvalidationする
	return Invoice{
		InvoiceCompany: Company{
			CompanyName:        CompanyName(companyName),
			RepresentativeName: RepresentativeName(representativeName),
			PhoneNumber:        PhoneNumber(phoneNumber),
			PostalCode:         PostalCode(postalCode),
			Address:            Address(address),
		},
		UserName: UserName(userName),
		InvoiceClient: Client{
			ClientID:           ClientID(clientID),
			CompanyName:        CompanyName(clientCompanyName),
			RepresentativeName: RepresentativeName(clientRepresentativeName),
			PhoneNumber:        PhoneNumber(clientPhoneNumber),
			PostalCode:         PostalCode(clientPostalCode),
			Address:            Address(clientAddress),
		},
		InvoiceBankAccount: BankAccount{
			BankName:      BankName(bankName),
			BranchName:    BranchName(branchName),
			AccountNumber: AccountNumber(accountNumber),
			AccountName:   AccountName(accountName),
		},
		IssueDate:          IssueDate(issueDate),
		PayAmount:          PayAmount(payAmount),
		Fee:                Fee(fee),
		FeeRate:            FeeRate(feeRate),
		ConsumptionTax:     ConsumptionTax(consumptionTax),
		ConsumptionTaxRate: ConsumptionTaxRate(consumptionTaxRate),
		TotalAmount:        TotalAmount(calcTotalAmount(float64(payAmount), feeRate, consumptionTaxRate)),
		PaymentDueDate:     PaymentDueDate(paymentDueDate),
		Status:             InvoiceStatusPending,
	}, nil
}

func RestoreInvoice(
	invoiceID InvoiceID,
	clientID ClientID,
	companyName CompanyName,
	representativeName RepresentativeName,
	phoneNumber PhoneNumber,
	postalCode PostalCode,
	address Address,
	userName UserName,
	clientCompanyName CompanyName,
	clientRepresentativeName RepresentativeName,
	clientPhoneNumber PhoneNumber,
	clientPostalCode PostalCode,
	clientAddress Address,
	bankName BankName,
	branchName BranchName,
	accountNumber AccountNumber,
	accountName AccountName,
	issueDate IssueDate,
	paymentDueDate PaymentDueDate,
	payAmount PayAmount,
	fee Fee,
	consumptionTax ConsumptionTax,
	totalAmount TotalAmount,
	feeRate FeeRate,
	consumptionTaxRate ConsumptionTaxRate,
	status Status,
) Invoice {
	return Invoice{
		InvoiceID: invoiceID,
		InvoiceCompany: Company{
			CompanyName:        companyName,
			RepresentativeName: representativeName,
			PhoneNumber:        phoneNumber,
			PostalCode:         postalCode,
			Address:            address,
		},
		UserName: userName,
		InvoiceClient: Client{
			ClientID:           clientID,
			CompanyName:        clientCompanyName,
			RepresentativeName: clientRepresentativeName,
			PhoneNumber:        clientPhoneNumber,
			PostalCode:         clientPostalCode,
			Address:            clientAddress,
		},
		InvoiceBankAccount: BankAccount{
			BankName:      bankName,
			BranchName:    branchName,
			AccountNumber: accountNumber,
			AccountName:   accountName,
		},
		IssueDate:          issueDate,
		PayAmount:          payAmount,
		Fee:                fee,
		FeeRate:            feeRate,
		ConsumptionTax:     consumptionTax,
		ConsumptionTaxRate: consumptionTaxRate,
		TotalAmount:        totalAmount,
		PaymentDueDate:     paymentDueDate,
		Status:             status,
	}
}

// calcTotalAmount : 支払い金額を計算するが、小数点は切り捨てられることに注意
func calcTotalAmount(payAmount, feeRate, consumptionTaxRate float64) int {
	return int(payAmount + (payAmount * feeRate * consumptionTaxRate))
}
