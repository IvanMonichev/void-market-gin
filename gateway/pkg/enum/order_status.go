package enum

type OrderStatus string

const (
	StatusPending   OrderStatus = "pending"
	StatusPaid      OrderStatus = "paid"
	StatusShipped   OrderStatus = "shipped"
	StatusDelivery  OrderStatus = "delivery"
	StatusCancelled OrderStatus = "cancelled"
)

func (s OrderStatus) IsValid() bool {
	switch s {
	case StatusPending, StatusPaid, StatusShipped, StatusDelivery, StatusCancelled:
		return true
	}
	return false
}
