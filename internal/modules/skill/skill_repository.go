package skill

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Skill, error)
	FindByID(id string) (*Skill, error)
	Create(skill *Skill) error
	Update(skill *Skill) error
	Delete(id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindAll() ([]Skill, error) {
	var skills []Skill
	if err := r.db.Order("category asc, \"order\" asc").Find(&skills).Error; err != nil {
		return nil, err
	}
	return skills, nil
}

func (r *repository) FindByID(id string) (*Skill, error) {
	var sk Skill
	if err := r.db.Where("id = ?", id).First(&sk).Error; err != nil {
		return nil, err
	}
	return &sk, nil
}

func (r *repository) Create(skill *Skill) error {
	return r.db.Create(skill).Error
}

func (r *repository) Update(skill *Skill) error {
	return r.db.Save(skill).Error
}

func (r *repository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&Skill{}).Error
}
