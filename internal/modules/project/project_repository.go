package project

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Project, error)
	FindByID(id string) (*Project, error)
	Create(project *Project) error
	Update(project *Project) error
	Delete(id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindAll() ([]Project, error) {
	var projects []Project
	if err := r.db.Order("\"order\" asc, created_at desc").Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *repository) FindByID(id string) (*Project, error) {
	var project Project
	if err := r.db.Where("id = ?", id).First(&project).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *repository) Create(project *Project) error {
	return r.db.Create(project).Error
}

func (r *repository) Update(project *Project) error {
	return r.db.Save(project).Error
}

func (r *repository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&Project{}).Error
}
