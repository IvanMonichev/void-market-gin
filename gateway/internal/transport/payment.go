package transport

type OrderStatusUpdateRequest struct {
	Status string `json:"status" binding:"required,oneof=pending paid shipped delivery cancelled"`
}
