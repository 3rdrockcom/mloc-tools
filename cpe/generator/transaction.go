package main

import (
	"math"
	"sort"
	"time"

	"github.com/epointpayment/mloc-tools/cpe/generator/models"

	"github.com/jmcvetta/randutil"
	"github.com/shopspring/decimal"
)

type Transactions struct {
	StartTime time.Time
	StopTime  time.Time
	MaxNum    int
	Precision int
	Balance   Balance
	Credit    Credit
	Debit     Debit
}

type Balance struct {
	RangeMin int
	RangeMax int
}

type Credit struct {
	Name         string
	Weight       int
	RangeMin     int
	RangeMax     int
	Descriptions []string
}

type Debit struct {
	Name         string
	Weight       int
	RangeMin     int
	RangeMax     int
	Descriptions []string
}

func NewTransactions() *Transactions {
	t := &Transactions{
		StartTime: time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC),
		StopTime:  time.Now().UTC(),
		MaxNum:    1000,
		Precision: 2,
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
			transaction.Credit = t.generateAmount(t.Credit.RangeMin, t.Credit.RangeMax, t.Precision)
			transaction.Description = t.generateDescription(t.Credit.Descriptions)
		case t.Debit.Name:
			transaction.Debit = t.generateAmount(t.Debit.RangeMin, t.Debit.RangeMax, t.Precision)
			transaction.Description = t.generateDescription(t.Debit.Descriptions)
		}

		if transactionType == t.Debit.Name {
			// Cannot debit without positive balance, wait for credit instead
			if balance.Equal(decimal.Zero) {
				continue
			}

			// Replace generated debit amount to empty account
			if balance.Sub(transaction.Debit).LessThan(decimal.Zero) {
				transaction.Debit = balance
			}
		}

		netChange := transaction.Credit.Sub(transaction.Debit)
		balance = balance.Add(netChange)
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

func (t *Transactions) generateBalance() decimal.Decimal {
	return t.generateAmount(t.Balance.RangeMin, t.Balance.RangeMax, t.Precision)
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

func (t *Transactions) generateAmount(rangeMin int, rangeMax int, precision int) decimal.Decimal {
	amount, _ := randutil.IntRange(int(rangeMin*int(math.Pow10(precision))), int(rangeMax*int(math.Pow10(precision))))

	return decimal.New(int64(amount), 0).Div(decimal.New(1, int32(precision)))
}

func (t *Transactions) generateDescription(descriptions []string) string {
	description, _ := randutil.ChoiceString(descriptions)
	return description
}
