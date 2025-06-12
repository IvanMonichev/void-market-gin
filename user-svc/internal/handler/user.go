package handler

import (
	"github.com/IvanMonichev/void-market-gin/user-svc/internal/service"
	"github.com/IvanMonichev/void-market-gin/user-svc/internal/transport"
	"github.com/IvanMonichev/void-market-gin/user-svc/pkg/mongo_id"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserHandler struct {
	service service.UserService
}

func New(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Create(c *gin.Context) {
	var req transport.CreateUserDto
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto := transport.CreateUserDto{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	}

	user, err := h.service.Create(c.Request.Context(), dto)
	if err != nil {
		log.Printf("create error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, transport.NewUserRdo(user))
}

func (h *UserHandler) Find(ctx *gin.Context) {
	idStr := ctx.Param("id")

	objectID, err := mongo_id.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	user, err := h.service.Find(ctx.Request.Context(), objectID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	ctx.JSON(http.StatusOK, transport.NewUserRdo(user))
}
