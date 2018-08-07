package main

import (
	"strconv"
	"time"

	"github.com/epointpayment/mloc-tools/cpe/generator/models"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/customers/list", getCustomersList)
	r.GET("/customer/:id/info", getCustomerInfo)
	r.GET("/customer/:id/transactions", getCustomerTransactions)

	return r
}

func getCustomersList(c *gin.Context) {
	var err error

	start, err := strconv.Atoi(c.DefaultQuery("start", "0"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	customers := []models.Customers{}
	err = db.Offset(start).Limit(limit).Find(&customers).Error
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, customers)
}

func getCustomerInfo(c *gin.Context) {
	var err error

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	customer := models.Customers{}
	err = db.Where("id = ?", id).Find(&customer).Error
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, customer)
}

func getCustomerTransactions(c *gin.Context) {
	var err error

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	customer := models.Customers{}
	err = db.Where("id = ?", id).Find(&customer).Error
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	// loc, _ := time.LoadLocation("America/Los_Angeles")

	query := db.Where("customers_id = ?", id)

	t1, err := time.ParseInLocation(
		"20060102",
		startDate, time.UTC)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	query = query.Where("date_time >= ?", t1)

	t2, err := time.ParseInLocation(
		"20060102",
		endDate, time.UTC)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	query = query.Where("date_time < ?", t2)

	transactions := []models.Transactions{}
	err = query.Find(&transactions).Error
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return

	}

	/*
		err = db.Where("customers_id = ?", id).
			Where("date_time >= ?", t1).
			Where("date_time < ?", t2).
			Find(&transactions).Error
	*/

	c.JSON(200, transactions)
}
