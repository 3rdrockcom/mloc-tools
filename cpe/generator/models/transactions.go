package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transactions struct {
	ID          uint `gorm:"PRIMARY_KEY"`
	CustomersID uint `sql:"type:integer REFERENCES customers(id) ON DELETE CASCADE ON UPDATE CASCADE"`
	DateTime    time.Time
	Description string
	Credit      decimal.Decimal `sql:"type:decimal(10,2);"`
	Debit       decimal.Decimal `sql:"type:decimal(10,2);"`
	Balance     decimal.Decimal `sql:"type:decimal(10,2);"`
}
