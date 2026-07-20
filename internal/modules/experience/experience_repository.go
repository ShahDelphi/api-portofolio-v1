package experience

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Experience, error)
	FindByID(id string) (*Experience, error)
	Create(exp *Experience) error
	Update(exp *Experience) error
	Delete(id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindAll() ([]Experience, error) {
	var experiences []Experience
	if err := r.db.Order("\"order\" asc, created_at desc").Find(&experiences).Error; err != nil {
		return nil, err
	}
	return experiences, nil
}

func (r *repository) FindByID(id string) (*Experience, error) {
	var exp Experience
	if err := r.db.Where("id = ?", id).First(&exp).Error; err != nil {
		return nil, err
	}
	return &exp, nil
}

func (r *repository) Create(exp *Experience) error {
	return r.db.Create(exp).Error
}

func (r *repository) Update(exp *Experience) error {
	return r.db.Save(exp).Error
}

func (r *repository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&Experience{}).Error
}
