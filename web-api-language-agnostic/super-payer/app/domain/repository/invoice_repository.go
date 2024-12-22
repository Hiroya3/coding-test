package repository

import (
	"context"
	"super-payer/app/domain/entity"
	"time"
)

type InvoiceRepository interface {
	Persist(ctx context.Context, invoice entity.Invoice) error
	List(ctx context.Context, companyID entity.CompanyID, fromDate, toDate time.Time) ([]entity.Invoice, error)
}
