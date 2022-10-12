package main

import (
	// Package json implements encoding and decoding of JSON.
	// The mapping between JSON and Go values is described in the documentation for the Marshal and Unmarshal functions.

	// Package ioutil implements some I/O utility functions.

	// Package os provides a platform-independent interface to operating system functionality.

	// importing the gin, because is a high-performance HTTP web framework written in Golang (Go).

	"bytes"
	"encoding/json"
	"net/http"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

const TIME_UNIT = 250

func main() {
	router := gin.Default()
	router.POST("/register", getRestaurantRegistration)
	router.POST("/order", getOrderFromClient)

	router.Run(":8090")
}

func getRestaurantRegistration(c *gin.Context) {
	var restaurant *Restaurant
	if err := c.BindJSON(&restaurant); err != nil {
		log.Err(err).Msg("Error!!")
		return
	}
	Restaurants[restaurant.RestaurantId] = *restaurant
	log.Printf("The new restaurant was registered: %+v", restaurant.RestaurantId)
	c.IndentedJSON(http.StatusCreated, restaurant)
}

func getOrderFromClient(c *gin.Context) {
	var order *Order
	var managO = ManagerOrder{
		Orders: make(map[int64]Order),
	}
	if err := c.BindJSON(&order); err != nil {
		log.Err(err).Msg("Error!!")
		return
	}
	log.Printf("The new order was received from :%+v", order)
	managO.ManagerOrderForSendingToRestaurant(*order)
	c.IndentedJSON(http.StatusCreated, order)
}

func (managO *ManagerOrder) ManagerOrderForSendingToRestaurant(order Order) {
	order.OrderId = atomic.AddInt64(&orderId, 1)
	managO.Orders[order.OrderId] = order
	for _, informationOrder := range order.Orders {
		restaurant, exists := Restaurants[informationOrder.RestaurantId]
		if !exists {
			continue
		}
		jsonBody, err := json.Marshal(informationOrder)
		if err != nil {
			log.Err(err).Msg("Error!!")
		}
		contentType := "application/json"
		//_, err = http.Post("http://dining_hall_restaurant:8080/v2/order", contentType, bytes.NewReader(jsonBody))
		_, err = http.Post("http://localhost:8080/v2/order", contentType, bytes.NewReader(jsonBody))
		if err != nil {
			log.Err(err).Msg("Error!!!")
		}
		log.Printf("The new order: %+v was sent to the restaurant havind id: %+v", order.OrderId, restaurant.RestaurantId)
	}
}
