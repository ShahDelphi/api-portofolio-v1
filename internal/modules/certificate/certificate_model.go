package certificate

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Certificate struct {
	ID            uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Title         string         `gorm:"type:varchar(255);not null" json:"title"`
	Issuer        string         `gorm:"type:varchar(255);not null" json:"issuer"`
	IssueDate     string         `gorm:"type:varchar(50);not null" json:"issue_date"` // E.g., "June 2024"
	CredentialURL string         `gorm:"type:varchar(255)" json:"credential_url"`
	Thumbnail     string         `gorm:"type:varchar(255)" json:"thumbnail"`
	Order         int            `gorm:"type:int;default:0" json:"order"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (Certificate) TableName() string {
	return "certificates"
}
