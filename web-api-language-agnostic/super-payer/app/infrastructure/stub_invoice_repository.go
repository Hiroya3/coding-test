package infrastructure

import (
	"context"
	"super-payer/app/domain/entity"
	"super-payer/app/domain/repository"
	"super-payer/pkg/log"
	"time"
)

type stubInvoiceRepository struct {
	logger log.Logger
}

func NewStubInvoiceRepository(logger log.Logger) repository.InvoiceRepository {
	return &stubInvoiceRepository{
		logger: logger,
	}
}

func (i stubInvoiceRepository) Persist(ctx context.Context, invoice entity.Invoice) (entity.Invoice, error) {
	id := 1

	return entity.RestoreInvoice(
		entity.InvoiceID(id), invoice.InvoiceClient.ClientID,
		invoice.InvoiceCompany.CompanyName, invoice.InvoiceCompany.RepresentativeName, invoice.InvoiceCompany.PhoneNumber, invoice.InvoiceCompany.PostalCode, invoice.InvoiceCompany.Address,
		invoice.UserName,
		invoice.InvoiceClient.CompanyName, invoice.InvoiceClient.RepresentativeName, invoice.InvoiceClient.PhoneNumber, invoice.InvoiceClient.PostalCode, invoice.InvoiceClient.Address,
		invoice.InvoiceBankAccount.BankName, invoice.InvoiceBankAccount.BranchName, invoice.InvoiceBankAccount.AccountNumber, invoice.InvoiceBankAccount.AccountName,
		invoice.IssueDate, invoice.PaymentDueDate,
		invoice.PayAmount, invoice.Fee, invoice.ConsumptionTax, invoice.TotalAmount,
		invoice.FeeRate, invoice.ConsumptionTaxRate,
		invoice.Status,
	), nil
}

func (i stubInvoiceRepository) List(ctx context.Context, companyID entity.CompanyID, fromDate, toDate time.Time) ([]entity.Invoice, error) {
	//TODO implement me
	panic("implement me")
}

var stubInvoices = []entity.Invoice{
	{
		InvoiceID: entity.InvoiceID(1),
		InvoiceCompany: entity.Company{
			CompanyID:          entity.CompanyID(1),
			CompanyName:        entity.CompanyName("株式会社A"),
			RepresentativeName: entity.RepresentativeName("山田太郎"),
			PhoneNumber:        entity.PhoneNumber("03-1234-5678"),
			PostalCode:         entity.PostalCode("100-0001"),
			Address:            entity.Address("東京都千代田区千代田1-1-1"),
		},
		UserName: entity.UserName("user1"),
		InvoiceClient: entity.Client{
			ClientID:           entity.ClientID(101),
			CompanyName:        entity.CompanyName("株式会社X"),
			RepresentativeName: entity.RepresentativeName("佐藤次郎"),
			PhoneNumber:        entity.PhoneNumber("03-8765-4321"),
			PostalCode:         entity.PostalCode("150-0002"),
			Address:            entity.Address("東京都渋谷区渋谷2-2-2"),
		},
		InvoiceBankAccount: entity.BankAccount{
			BankName:      entity.BankName("みずほ銀行"),
			BranchName:    entity.BranchName("東京支店"),
			AccountNumber: entity.AccountNumber("1234567"),
			AccountName:   entity.AccountName("カ）エー"),
		},
		IssueDate:          entity.IssueDate(time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC)),
		PayAmount:          entity.PayAmount(100000),
		Fee:                entity.Fee(4000),
		FeeRate:            entity.FeeRate(0.04),
		ConsumptionTax:     entity.ConsumptionTax(10400),
		ConsumptionTaxRate: entity.ConsumptionTaxRate(0.10),
		TotalAmount:        entity.TotalAmount(114400),
		PaymentDueDate:     entity.PaymentDueDate(time.Date(2023, 5, 31, 0, 0, 0, 0, time.UTC)),
		Status:             entity.InvoiceStatusPending,
	},
	{
		InvoiceID: entity.InvoiceID(2),
		InvoiceCompany: entity.Company{
			CompanyID:          entity.CompanyID(1),
			CompanyName:        entity.CompanyName("株式会社A"),
			RepresentativeName: entity.RepresentativeName("山田太郎"),
			PhoneNumber:        entity.PhoneNumber("03-1234-5678"),
			PostalCode:         entity.PostalCode("100-0001"),
			Address:            entity.Address("東京都千代田区千代田1-1-1"),
		},
		UserName: entity.UserName("user1"),
		InvoiceClient: entity.Client{
			ClientID:           entity.ClientID(102),
			CompanyName:        entity.CompanyName("株式会社Y"),
			RepresentativeName: entity.RepresentativeName("鈴木花子"),
			PhoneNumber:        entity.PhoneNumber("03-2345-6789"),
			PostalCode:         entity.PostalCode("160-0003"),
			Address:            entity.Address("東京都新宿区新宿3-3-3"),
		},
		InvoiceBankAccount: entity.BankAccount{
			BankName:      entity.BankName("三菱UFJ銀行"),
			BranchName:    entity.BranchName("新宿支店"),
			AccountNumber: entity.AccountNumber("7654321"),
			AccountName:   entity.AccountName("カ）エー"),
		},
		IssueDate:          entity.IssueDate(time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC)),
		PayAmount:          entity.PayAmount(200000),
		Fee:                entity.Fee(8000),
		FeeRate:            entity.FeeRate(0.04),
		ConsumptionTax:     entity.ConsumptionTax(20800),
		ConsumptionTaxRate: entity.ConsumptionTaxRate(0.10),
		TotalAmount:        entity.TotalAmount(228800),
		PaymentDueDate:     entity.PaymentDueDate(time.Date(2023, 6, 15, 0, 0, 0, 0, time.UTC)),
		Status:             entity.InvoiceStatusProcessing,
	},
	{
		InvoiceID: entity.InvoiceID(3),
		InvoiceCompany: entity.Company{
			CompanyID:          entity.CompanyID(2),
			CompanyName:        entity.CompanyName("株式会社B"),
			RepresentativeName: entity.RepresentativeName("田中三郎"),
			PhoneNumber:        entity.PhoneNumber("06-9876-5432"),
			PostalCode:         entity.PostalCode("530-0001"),
			Address:            entity.Address("大阪府大阪市北区梅田1-1-1"),
		},
		UserName: entity.UserName("user2"),
		InvoiceClient: entity.Client{
			ClientID:           entity.ClientID(103),
			CompanyName:        entity.CompanyName("株式会社Z"),
			RepresentativeName: entity.RepresentativeName("高橋幸子"),
			PhoneNumber:        entity.PhoneNumber("06-3456-7890"),
			PostalCode:         entity.PostalCode("540-0002"),
			Address:            entity.Address("大阪府大阪市中央区大阪城2-2-2"),
		},
		InvoiceBankAccount: entity.BankAccount{
			BankName:      entity.BankName("三井住友銀行"),
			BranchName:    entity.BranchName("大阪支店"),
			AccountNumber: entity.AccountNumber("2468135"),
			AccountName:   entity.AccountName("カ）ビー"),
		},
		IssueDate:          entity.IssueDate(time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC)),
		PayAmount:          entity.PayAmount(150000),
		Fee:                entity.Fee(6000),
		FeeRate:            entity.FeeRate(0.04),
		ConsumptionTax:     entity.ConsumptionTax(15600),
		ConsumptionTaxRate: entity.ConsumptionTaxRate(0.10),
		TotalAmount:        entity.TotalAmount(171600),
		PaymentDueDate:     entity.PaymentDueDate(time.Date(2023, 6, 30, 0, 0, 0, 0, time.UTC)),
		Status:             entity.InvoiceStatusPending,
	},
	{
		InvoiceID: entity.InvoiceID(4),
		InvoiceCompany: entity.Company{
			CompanyID:          entity.CompanyID(1),
			CompanyName:        entity.CompanyName("株式会社A"),
			RepresentativeName: entity.RepresentativeName("山田太郎"),
			PhoneNumber:        entity.PhoneNumber("03-1234-5678"),
			PostalCode:         entity.PostalCode("100-0001"),
			Address:            entity.Address("東京都千代田区千代田1-1-1"),
		},
		UserName: entity.UserName("user1"),
		InvoiceClient: entity.Client{
			ClientID:           entity.ClientID(104),
			CompanyName:        entity.CompanyName("株式会社W"),
			RepresentativeName: entity.RepresentativeName("中村四郎"),
			PhoneNumber:        entity.PhoneNumber("03-4567-8901"),
			PostalCode:         entity.PostalCode("170-0004"),
			Address:            entity.Address("東京都豊島区池袋4-4-4"),
		},
		InvoiceBankAccount: entity.BankAccount{
			BankName:      entity.BankName("りそな銀行"),
			BranchName:    entity.BranchName("池袋支店"),
			AccountNumber: entity.AccountNumber("1357924"),
			AccountName:   entity.AccountName("カ）エー"),
		},
		IssueDate:          entity.IssueDate(time.Date(2023, 6, 15, 0, 0, 0, 0, time.UTC)),
		PayAmount:          entity.PayAmount(300000),
		Fee:                entity.Fee(12000),
		FeeRate:            entity.FeeRate(0.04),
		ConsumptionTax:     entity.ConsumptionTax(31200),
		ConsumptionTaxRate: entity.ConsumptionTaxRate(0.10),
		TotalAmount:        entity.TotalAmount(343200),
		PaymentDueDate:     entity.PaymentDueDate(time.Date(2023, 7, 15, 0, 0, 0, 0, time.UTC)),
		Status:             entity.InvoiceStatusPaid,
	},
	{
		InvoiceID: entity.InvoiceID(5),
		InvoiceCompany: entity.Company{
			CompanyID:          entity.CompanyID(2),
			CompanyName:        entity.CompanyName("株式会社B"),
			RepresentativeName: entity.RepresentativeName("田中三郎"),
			PhoneNumber:        entity.PhoneNumber("06-9876-5432"),
			PostalCode:         entity.PostalCode("530-0001"),
			Address:            entity.Address("大阪府大阪市北区梅田1-1-1"),
		},
		UserName: entity.UserName("user2"),
		InvoiceClient: entity.Client{
			ClientID:           entity.ClientID(105),
			CompanyName:        entity.CompanyName("株式会社V"),
			RepresentativeName: entity.RepresentativeName("小林五郎"),
			PhoneNumber:        entity.PhoneNumber("06-5678-9012"),
			PostalCode:         entity.PostalCode("550-0003"),
			Address:            entity.Address("大阪府大阪市西区西本町3-3-3"),
		},
		InvoiceBankAccount: entity.BankAccount{
			BankName:      entity.BankName("関西みらい銀行"),
			BranchName:    entity.BranchName("本町支店"),
			AccountNumber: entity.AccountNumber("8642097"),
			AccountName:   entity.AccountName("カ）ビー"),
		},
		IssueDate:          entity.IssueDate(time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC)),
		PayAmount:          entity.PayAmount(250000),
		Fee:                entity.Fee(10000),
		FeeRate:            entity.FeeRate(0.04),
		ConsumptionTax:     entity.ConsumptionTax(26000),
		ConsumptionTaxRate: entity.ConsumptionTaxRate(0.10),
		TotalAmount:        entity.TotalAmount(286000),
		PaymentDueDate:     entity.PaymentDueDate(time.Date(2023, 7, 31, 0, 0, 0, 0, time.UTC)),
		Status:             entity.InvoiceStatusProcessing,
	},
}
