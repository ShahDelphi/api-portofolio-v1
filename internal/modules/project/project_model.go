package project

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Project struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Title       string         `gorm:"type:varchar(255);not null" json:"title"`
	Description string         `gorm:"type:text;not null" json:"description"`
	Thumbnail   string         `gorm:"type:varchar(255)" json:"thumbnail"`
	GithubURL   string         `gorm:"type:varchar(255)" json:"github_url"`
	DemoURL     string         `gorm:"type:varchar(255)" json:"demo_url"`
	TechStack   string         `gorm:"type:varchar(255)" json:"tech_stack"` // Comma-separated stack (e.g., "React,Golang,PostgreSQL")
	Order       int            `gorm:"type:int;default:0" json:"order"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (Project) TableName() string {
	return "projects"
}
