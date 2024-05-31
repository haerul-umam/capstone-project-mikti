package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Status string
type Payment string

const (
	Menunggu Status = "MENUNGGU"
	Diterima Status = "DITERIMA"
	Ditolak  Status = "DITOLAK"
)

const (
	Credit  Payment = "CREDIT"
	Debit   Payment = "DEBIT"
	Va      Payment = "VA"
	Ewallet Payment = "EWALLET"
)

type Order struct {
	OrderID       string  `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"`
	EventID       int     `gorm:"column:event_id"`
	UserID        string  `gorm:"column:user_id"`
	NameEvent     string  `gorm:"column:name_event"`
	DateEvent     string  `gorm:"column:date_event"`
	PriceEvent    int     `gorm:"column:price_event"`
	IsFree        bool    `gorm:"column:is_free"`
	Description   string  `gorm:"column:description"`
	City          string  `gorm:"column:city"`
	Quantity      int     `gorm:"column:quantity"`
	PaymentMethod Payment `gorm:"column:payment_method;type:enum('CREDIT', 'DEBIT', 'VA', 'EWALLET')"`
	Amount        int     `gorm:"column:amount"`
	Status        Status  `gorm:"column:status;type:enum('MENUNGGU', 'DITERIMA', 'DITOLAK')"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (order *Order) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	order.OrderID = uuid.NewString()
	return
}

func (order *Order) TableName() string {
	return "order"
}
