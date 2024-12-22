package usecase

import (
	"context"
	"super-payer/app/domain/entity"
	"time"
)

func (p invoiceUseCase) ListByUserIDAndDate(ctx context.Context, userID entity.UserID, fromDate, toDate time.Time) ([]entity.Invoice, error) {

	// userIDからcompanyを取得
	company, err := p.companyRepository.GetByUserID(ctx, userID)
	if err != nil {
		p.logger.Warnf(ctx, "failed to get company by userID: %v,err : %v", userID, err)
		return nil, err
	}

	invoices, err := p.invoiceRepository.ListByDuration(ctx, company.CompanyID, fromDate, toDate)
	if err != nil {
		p.logger.Warnf(ctx, "failed to get invoices by duration: %v ~ %v,companyID : %v,,err : %v", fromDate.Format("2006-01-02"), toDate.Format("2006-01-02"), company.CompanyID, err)
		return nil, err
	}
	return invoices, nil
}
