package handler

import (
	"errors"
	"github.com/IvanMonichev/void-market-gin/user-svc/internal/model"
	"github.com/IvanMonichev/void-market-gin/user-svc/internal/repository"
	"github.com/IvanMonichev/void-market-gin/user-svc/internal/transport"
	"github.com/IvanMonichev/void-market-gin/user-svc/pkg/hash"
	"github.com/IvanMonichev/void-market-gin/user-svc/pkg/mongo_id"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"log"
	"net/http"
	"strconv"
)

type UserHandler struct {
	repository repository.UserRepository
}

func New(repository repository.UserRepository) *UserHandler {
	return &UserHandler{repository: repository}
}

func (h *UserHandler) Create(c *gin.Context) {
	var dto transport.CreateUserDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := hash.HashPassword(dto.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	user := &model.User{
		Email:    dto.Email,
		Name:     dto.Name,
		Password: hashedPassword,
	}

	createdUser, err := h.repository.Create(c.Request.Context(), user)
	if err != nil {
		log.Printf("create error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, transport.NewUserRdo(createdUser))
}

func (h *UserHandler) Find(ctx *gin.Context) {
	idStr := ctx.Param("id")

	objectID, err := mongo_id.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	user, err := h.repository.FindByID(ctx.Request.Context(), objectID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	ctx.JSON(http.StatusOK, transport.NewUserRdo(user))
}

func (h *UserHandler) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")

	objectID, err := mongo_id.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	var dto transport.CreateUserDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	user := &model.User{
		ID:       objectID,
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password, // в боевом коде тут должно быть хеширование!
	}

	updatedUser, err := h.repository.Update(ctx, user, objectID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
		return
	}

	ctx.JSON(http.StatusOK, updatedUser)
}

func (h *UserHandler) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")

	objectID, err := mongo_id.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	err = h.repository.Delete(ctx, objectID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *UserHandler) GetAll(ctx *gin.Context) {

	limit := int64(10)
	offset := int64(0)

	if l := ctx.Query("limit"); l != "" {
		if parsed, err := strconv.ParseInt(l, 10, 64); err == nil {
			limit = parsed
		}
	}

	if o := ctx.Query("offset"); o != "" {
		if parsed, err := strconv.ParseInt(o, 10, 64); err == nil {
			offset = parsed
		}
	}

	users, total, err := h.repository.GetAll(ctx.Request.Context(), offset, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch users"})
		return
	}

	var data []transport.UserRdo
	for _, user := range users {
		data = append(data, transport.NewUserRdo(user))
	}

	response := transport.PaginatedResponse[transport.UserRdo]{
		Data: data,
		Meta: transport.Meta{
			Total:  total,
			Limit:  limit,
			Offset: offset,
		},
	}

	ctx.JSON(http.StatusOK, response)
}
