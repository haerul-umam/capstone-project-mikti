package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role string

const (
	Admin Role = "ADMIN"
	Buyer Role = "BUYER"
)

type User struct {
	UserID    string `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name      string
	Email     string
	Password  string
	Role      Role `gorm:"type:enum('ADMIN', 'BUYER')"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	user.UserID = uuid.NewString()
	return
}

func (user *User) TableName() string {
	return "user"
}
