package auth

import (
	"portfolio-backend/internal/modules/admin"

	"gorm.io/gorm"
)

type Repository interface {
	FindByUsername(username string) (*admin.Admin, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindByUsername(username string) (*admin.Admin, error) {
	var adm admin.Admin
	if err := r.db.Where("username = ?", username).First(&adm).Error; err != nil {
		return nil, err
	}
	return &adm, nil
}
