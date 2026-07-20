package experience

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Experience struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Company     string         `gorm:"type:varchar(255);not null" json:"company"`
	Role        string         `gorm:"type:varchar(255);not null" json:"role"`
	Location    string         `gorm:"type:varchar(255)" json:"location"`
	StartDate   string         `gorm:"type:varchar(50);not null" json:"start_date"` // E.g., "July 2023"
	EndDate     string         `gorm:"type:varchar(50)" json:"end_date"`            // E.g., "Present" or "Dec 2023"
	CurrentJob  bool           `gorm:"type:boolean;default:false" json:"current_job"`
	Description string         `gorm:"type:text;not null" json:"description"` // Supports Markdown description
	Order       int            `gorm:"type:int;default:0" json:"order"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (Experience) TableName() string {
	return "experiences"
}
