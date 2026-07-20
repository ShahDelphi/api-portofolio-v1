package certificate

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Certificate, error)
	FindByID(id string) (*Certificate, error)
	Create(cert *Certificate) error
	Update(cert *Certificate) error
	Delete(id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindAll() ([]Certificate, error) {
	var certs []Certificate
	if err := r.db.Order("\"order\" asc, created_at desc").Find(&certs).Error; err != nil {
		return nil, err
	}
	return certs, nil
}

func (r *repository) FindByID(id string) (*Certificate, error) {
	var cert Certificate
	if err := r.db.Where("id = ?", id).First(&cert).Error; err != nil {
		return nil, err
	}
	return &cert, nil
}

func (r *repository) Create(cert *Certificate) error {
	return r.db.Create(cert).Error
}

func (r *repository) Update(cert *Certificate) error {
	return r.db.Save(cert).Error
}

func (r *repository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&Certificate{}).Error
}
