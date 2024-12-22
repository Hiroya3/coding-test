package repository

import (
	"context"
	"super-payer/app/domain/entity"
)

type CompanyRepository interface {
	GetByUserID(ctx context.Context, userID entity.UserID) (entity.Company, error)
}
