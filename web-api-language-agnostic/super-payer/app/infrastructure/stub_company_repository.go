package infrastructure

import (
	"context"
	"super-payer/app/domain/entity"
	"super-payer/app/domain/repository"
	"super-payer/pkg/log"
)

type stubCompanyRepository struct {
	logger log.Logger
}

func NewStubCompanyRepository(logger log.Logger) repository.CompanyRepository {
	return &stubCompanyRepository{
		logger: logger,
	}
}

func (i stubCompanyRepository) GetByUserID(_ context.Context, _ entity.UserID) (entity.Company, error) {
	return entity.Company{
		CompanyID:          entity.CompanyID(1),
		CompanyName:        "株式会社A",
		RepresentativeName: "山田太郎",
		PhoneNumber:        "03-1234-5678",
		PostalCode:         "100-0001",
		Address:            "東京都千代田区千代田1-1-1",
	}, nil
}
