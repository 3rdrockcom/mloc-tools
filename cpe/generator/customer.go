package main

import (
	randomdata "github.com/Pallinder/go-randomdata"
	"github.com/proteogenic/eps-cpe-demo-faker/models"
)

type Customer struct{}

func NewCustomer() *Customer {
	c := &Customer{}
	return c
}

func (c *Customer) Generate() *models.Customers {
	profile := randomdata.GenerateProfile(randomdata.RandomGender)

	return &models.Customers{
		Email:     profile.Email,
		Gender:    profile.Gender,
		FirstName: profile.Name.First,
		LastName:  profile.Name.Last,
	}
}
