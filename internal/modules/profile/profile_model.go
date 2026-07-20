package profile

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Profile struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name      string         `gorm:"type:varchar(100);not null" json:"name"`
	Title     string         `gorm:"type:varchar(255);not null" json:"title"`     // e.g. "Brittany Chiang."
	Subtitle  string         `gorm:"type:varchar(255);not null" json:"subtitle"`  // e.g. "I build things for the web."
	Intro     string         `gorm:"type:text" json:"intro"`                     // The paragraph in Hero section
	Bio       string         `gorm:"type:text" json:"bio"`                       // The paragraph in About section (HTML/markdown supported)
	Avatar    string         `gorm:"type:varchar(255)" json:"avatar"`            // The profile picture URL
	ResumeURL string         `gorm:"type:varchar(255)" json:"resume_url"`        // The resume URL / file path
	Github    string         `gorm:"type:varchar(255)" json:"github"`
	Instagram string         `gorm:"type:varchar(255)" json:"instagram"`
	Linkedin  string         `gorm:"type:varchar(255)" json:"linkedin"`
	Email     string         `gorm:"type:varchar(255)" json:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (Profile) TableName() string {
	return "profiles"
}
