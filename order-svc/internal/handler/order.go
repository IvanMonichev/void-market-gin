package handler

import (
	"github.com/IvanMonichev/void-market-gin/order-svc/internal/mapper"
	"github.com/IvanMonichev/void-market-gin/order-svc/internal/model"
	"github.com/IvanMonichev/void-market-gin/order-svc/internal/repository"
	"github.com/IvanMonichev/void-market-gin/order-svc/internal/transport"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type OrderHandler struct {
	repo repository.OrderRepository
}

func NewOrderHandler(repo repository.OrderRepository) *OrderHandler {
	return &OrderHandler{
		repo: repo,
	}
}

func (h *OrderHandler) Create(ctx *gin.Context) {
	var dto transport.OrderDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}

	items := make([]model.OrderItem, len(dto.Items))
	var total float64
	for i, item := range dto.Items {
		items[i] = model.OrderItem{
			Name:      item.Name,
			Quantity:  item.Quantity,
			UnitPrice: item.UnitPrice,
		}
		total += float64(item.Quantity) * item.UnitPrice
	}

	order := model.Order{
		UserID: dto.UserID,
		Status: dto.Status,
		Items:  items,
		Total:  total,
	}

	created, err := h.repo.Create(ctx.Request.Context(), &order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order"})
		return
	}

	ctx.JSON(http.StatusCreated, mapper.ToOrderRDO(*created))
}

func (h *OrderHandler) Find(ctx *gin.Context) {
	idParam := ctx.Param("id")
	idUint, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	id := uint(idUint)

	order, err := h.repo.FindById(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	ctx.JSON(http.StatusOK, mapper.ToOrderRDO(*order))
}

func (h *OrderHandler) GetAll(ctx *gin.Context) {
	offset, _ := strconv.ParseInt(ctx.DefaultQuery("offset", "0"), 10, 64)
	limit, _ := strconv.ParseInt(ctx.DefaultQuery("limit", "10"), 10, 64)

	orders, total, err := h.repo.GetAll(ctx.Request.Context(), offset, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get orders"})
		return
	}

	result := make([]transport.OrderRDO, len(orders))
	for i, o := range orders {
		result[i] = mapper.ToOrderRDO(o)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"total":  total,
		"orders": result,
	})
}

func (h *OrderHandler) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	idUint, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	id := uint(idUint)

	var dto transport.OrderDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}

	items := make([]model.OrderItem, len(dto.Items))
	var total float64
	for i, item := range dto.Items {
		items[i] = model.OrderItem{
			Name:      item.Name,
			Quantity:  item.Quantity,
			UnitPrice: item.UnitPrice,
		}
		total += float64(item.Quantity) * item.UnitPrice
	}

	order := model.Order{
		ID:        id,
		UserID:    dto.UserID,
		Status:    dto.Status,
		Items:     items,
		Total:     total,
		UpdatedAt: time.Now(),
	}

	updated, err := h.repo.Update(ctx.Request.Context(), &order, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update order"})
		return
	}

	ctx.JSON(http.StatusOK, mapper.ToOrderRDO(*updated))
}

func (h *OrderHandler) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	idUint, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	id := uint(idUint)

	if err := h.repo.Delete(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete order"})
		return
	}

	ctx.Status(http.StatusNoContent)
}
