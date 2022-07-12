package entity

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Customer struct {
	CustomerId      uuid.UUID `json:"customer_id" gorm:"customer_id"`
	CustomerName    string    `json:"customer_name" gorm:"customer_name"`
	CustomerAddress string    `json:"customer_address" gorm:"customer_address"`
}

func (customer *Customer) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.NewV4().String()
	return scope.SetColumn("customer_id", id)
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name
func (c Customer) TableName() string {
	return "customer"
}

// TODO
// func (c Customer) ValidateInput(map[string]string) {

// }