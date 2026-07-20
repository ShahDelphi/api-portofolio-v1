package profile

import (
	"gorm.io/gorm"
)

type Repository interface {
	Get() (*Profile, error)
	Update(profile *Profile) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Get() (*Profile, error) {
	var prof Profile
	if err := r.db.First(&prof).Error; err != nil {
		return nil, err
	}
	return &prof, nil
}

func (r *repository) Update(profile *Profile) error {
	return r.db.Save(profile).Error
}
