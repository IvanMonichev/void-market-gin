package model

type OrderStatus string

const (
	StatusPending   OrderStatus = "pending"
	StatusPaid      OrderStatus = "paid"
	StatusShipped   OrderStatus = "shipped"
	StatusDelivery  OrderStatus = "delivery"
	StatusCancelled OrderStatus = "cancelled"
)
