package skill

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Skill struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	Category  string         `gorm:"type:varchar(255);not null" json:"category"` // E.g., "Frontend", "Backend", "Data Science", "Tools"
	Order     int            `gorm:"type:int;default:0" json:"order"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (Skill) TableName() string {
	return "skills"
}
