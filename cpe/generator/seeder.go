package main

import (
	"fmt"

	"github.com/proteogenic/eps-cpe-demo-faker/models"

	"github.com/gosuri/uiprogress"
	"github.com/gosuri/uiprogress/util/strutil"
)

type Seeder struct {
	NumCustomers       int
	MaxNumTransactions int
}

func NewSeeder() *Seeder {
	s := &Seeder{
		NumCustomers:       1,
		MaxNumTransactions: 1000,
	}

	return s
}

func (s *Seeder) Run() {

	uiprogress.Start()

	s.seedCustomers()
	s.seedTransactions()

	uiprogress.Stop()

	fmt.Println("Done!")
}

func (s *Seeder) seedCustomers() {
	customer := NewCustomer()

	bar := uiprogress.
		AddBar(s.NumCustomers + 1).
		AppendCompleted().
		PrependElapsed().
		PrependFunc(func(b *uiprogress.Bar) string {
			return strutil.Resize("Generating Customers: ", 30)
		})

	tx := db.Begin()
	if tx.Error != nil {
		panic(tx.Error)
	}

	for i := 0; i < s.NumCustomers; i++ {
		if err := tx.Create(customer.Generate()).Error; err != nil {
			tx.Rollback()
			panic(err)
		}

		bar.Incr()
	}

	if err := tx.Commit().Error; err != nil {
		panic(err)
	}

	bar.Incr()
}

func (s *Seeder) seedTransactions() {
	customers := []*models.Customers{}

	transactions := NewTransactions()
	transactions.MaxNum = s.MaxNumTransactions

	db.Find(&customers)

	bar := uiprogress.
		AddBar(s.NumCustomers).
		AppendCompleted().
		PrependElapsed().
		PrependFunc(func(b *uiprogress.Bar) string {
			return strutil.Resize("Generating Transactions: ", 30)
		})

	count := len(customers)
	for i := 0; i < count; i++ {
		entries := transactions.Generate()

		tx := db.Begin()
		if tx.Error != nil {
			panic(tx.Error)
		}

		for j := range entries {
			entry := entries[j]
			entry.CustomersID = customers[i].ID

			if err := tx.Create(entry).Error; err != nil {
				tx.Rollback()
				panic(err)
			}
		}

		if err := tx.Commit().Error; err != nil {
			panic(err)
		}

		bar.Incr()
	}
}
