package usecase

import (
	"context"
	"super-payer/app/domain/entity"
)

func (p invoiceUseCase) Persist(ctx context.Context, input PersistInvoiceInput) (entity.Invoice, error) {
	// TODO implement
	return entity.Invoice{}, nil
}
