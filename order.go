package main

type Order struct {
	OrderId  int64              `json:"order_id"`
	ClientId int                `json:"client_id"`
	Orders   []informationOrder `json:"orders"`
}

type informationOrder struct {
	RestaurantId int     `json:"restaurant_id"`
	Items        []int   `json:"items"`
	Priority     int     `json:"priority"`
	MaxWait      float64 `json:"max_wait"`
	CreatedTime  int64   `json:"created_time"`
}

type ManagerOrder struct {
	Orders map[int64]Order
}

var orderId int64
