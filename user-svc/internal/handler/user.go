package handler

import (
	"github.com/IvanMonichev/void-market-gin/user-svc/internal/service"
	"github.com/IvanMonichev/void-market-gin/user-svc/internal/transport"
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

	input := service.CreateUserInput{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	}

	user, err := h.service.Create(c.Request.Context(), input)
	if err != nil {
		log.Printf("create error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, transport.NewUserRdo(user))
}
