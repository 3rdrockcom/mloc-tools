package main

import (
	"sort"
	"time"

	"github.com/epointpayment/mloc-tools/cpe/generator/models"

	"github.com/jmcvetta/randutil"
)

type Transactions struct {
	StartTime  time.Time
	StopTime   time.Time
	MaxNum     int
	Multiplier float64
	Balance    Balance
	Credit     Credit
	Debit      Debit
}

type Balance struct {
	RangeMin float64
	RangeMax float64
}

type Credit struct {
	Name         string
	Weight       int
	RangeMin     float64
	RangeMax     float64
	Descriptions []string
}

type Debit struct {
	Name         string
	Weight       int
	RangeMin     float64
	RangeMax     float64
	Descriptions []string
}

func NewTransactions() *Transactions {
	t := &Transactions{
		StartTime:  time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC),
		StopTime:   time.Now().UTC(),
		MaxNum:     1000,
		Multiplier: 100,
		Balance: Balance{
			RangeMin: 0,
			RangeMax: 100,
		},
		Credit: Credit{
			Name:     "credit",
			Weight:   2,
			RangeMin: 100,
			RangeMax: 500,
			Descriptions: []string{
				"Bank Deposit",
				"Ebay",
				"Amazon",
				"Gift",
				"Cashed Check",
			},
		},
		Debit: Debit{
			Name:     "debit",
			Weight:   8,
			RangeMin: 0,
			RangeMax: 500,
			Descriptions: []string{
				"Pizza Hut",
				"Ulta",
				"Taco Bell",
				"Walmart",
				"J.C. Penny",
				"Del Taco",
				"Big O Tires",
				"Auto Zone",
				"McDonalds",
			},
		},
	}

	return t
}

func (t *Transactions) Generate() []*models.Transactions {
	transactions := make([]*models.Transactions, 0)

	timestamps := t.generateTimestamps()
	balance := t.generateBalance()

	for i := 0; i < t.MaxNum; i++ {
		transaction := &models.Transactions{}

		transaction.DateTime = time.Unix(int64(timestamps[i]), 0)

		transactionType := t.generateTransactionType()
		switch transactionType {
		case t.Credit.Name:
			transaction.Credit = t.generateAmount(t.Credit.RangeMin, t.Credit.RangeMax, t.Multiplier)
			transaction.Description = t.generateDescription(t.Credit.Descriptions)
		case t.Debit.Name:
			transaction.Debit = t.generateAmount(t.Debit.RangeMin, t.Debit.RangeMax, t.Multiplier)
			transaction.Description = t.generateDescription(t.Debit.Descriptions)
		}

		if transactionType == t.Debit.Name {
			// Cannot debit without positive balance, wait for credit instead
			if balance == 0 {
				continue
			}

			// Replace generated debit amount to empty account
			if balance-transaction.Debit < 0 {
				transaction.Debit = balance
			}
		}

		netChange := transaction.Credit - transaction.Debit
		balance = balance + netChange
		transaction.Balance = balance

		transactions = append(transactions, transaction)
	}

	return transactions
}

func (t *Transactions) generateTimestamps() []int {
	var timestamps []int

	for i := 0; i < t.MaxNum; i++ {
		timestamp, _ := randutil.IntRange(int(t.StartTime.Unix()), int(t.StopTime.Unix()))
		timestamps = append(timestamps, timestamp)
	}

	sort.Ints(timestamps)

	return timestamps
}

func (t *Transactions) generateBalance() float64 {
	return t.generateAmount(t.Balance.RangeMin, t.Balance.RangeMax, t.Multiplier)
}

func (t *Transactions) generateTransactionType() string {
	choices := []randutil.Choice{
		randutil.Choice{
			Item:   t.Credit.Name,
			Weight: t.Credit.Weight,
		},
		{
			Item:   t.Debit.Name,
			Weight: t.Debit.Weight,
		},
	}

	choice, _ := randutil.WeightedChoice(choices)
	return choice.Item.(string)
}

func (t *Transactions) generateAmount(rangeMin, rangeMax, multiplier float64) float64 {
	amount, _ := randutil.IntRange(int(rangeMin*multiplier), int(rangeMax*multiplier))
	return float64(amount) / multiplier
}

func (t *Transactions) generateDescription(descriptions []string) string {
	description, _ := randutil.ChoiceString(descriptions)
	return description
}
