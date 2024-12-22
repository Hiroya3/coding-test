package usecase

import (
	"context"
	"super-payer/app/domain/entity"
	"super-payer/app/domain/repository"
	"super-payer/pkg/log"
	"time"
)

type InvoiceUseCase interface {
	ListByUserIDAndDate(ctx context.Context, userID entity.UserID, fromDate, toDate time.Time) ([]entity.Invoice, error)
	Persist(ctx context.Context, input PersistInvoiceInput) (entity.Invoice, error)
}

type invoiceUseCase struct {
	logger            log.Logger
	invoiceRepository repository.InvoiceRepository
	companyRepository repository.CompanyRepository
}

func NewInvoiceUseCase(logger log.Logger, invoiceRepository repository.InvoiceRepository, companyRepository repository.CompanyRepository) InvoiceUseCase {
	return &invoiceUseCase{
		logger:            logger,
		invoiceRepository: invoiceRepository,
		companyRepository: companyRepository,
	}
}

type PersistInvoiceInput struct {
	ClientID           int
	IssueDate          time.Time
	PayAmount          int
	Fee                int
	FeeRate            float64
	ConsumptionTax     int
	ConsumptionTaxRate float64
	PaymentDueDate     string
}
