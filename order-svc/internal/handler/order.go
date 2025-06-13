package handler

import (
	"github.com/IvanMonichev/void-market-gin/order-svc/internal/model"
	"github.com/IvanMonichev/void-market-gin/order-svc/internal/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Order interface {
	Find(ctx *gin.Context, id uint) (*model.Order, error)
	Create(ctx *gin.Context, order model.Order) (*model.Order, error)
	Update(ctx *gin.Context, id uint, order model.Order) (*model.Order, error)
	Delete(ctx *gin.Context, id uint) error
	GetAll(ctx *gin.Context, offset int64, limit int64) ([]model.Order, int64, error)
}
type OrderHandler struct {
	repository repository.OrderRepository
}

func NewOrderHandler(repository repository.OrderRepository) *OrderHandler {
	return &OrderHandler{repository: repository}
}

func (h *OrderHandler) Find(ctx *gin.Context, id uint) (*model.Order, error) {
	id := ctx.Param("id")

	user, err := h.repository.FindById(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})

}
