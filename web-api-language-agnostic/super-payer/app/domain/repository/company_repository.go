package repository

import "super-payer/app/domain/entity"

type CompanyRepository interface {
	GetByUserID(userID entity.UserID) (entity.Company, error)
}
