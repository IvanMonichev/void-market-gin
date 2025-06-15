package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"io"
	"net/http"
)

type UserHandler struct {
	client *resty.Client
}

func NewUserHandler(client *resty.Client) *UserHandler {
	return &UserHandler{client: client}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	resp, err := h.client.R().
		SetBody(body).
		Post("/users")

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "user service unavailable"})
		return
	}

	c.Data(resp.StatusCode(), resp.Header().Get("Content-Type"), resp.Body())
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.client.R().
		Get(fmt.Sprintf("/users/%s", id))

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "user service unavailable"})
		return
	}

	c.Data(resp.StatusCode(), resp.Header().Get("Content-Type"), resp.Body())
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	resp, err := h.client.R().
		SetBody(body).
		Put(fmt.Sprintf("/users/%s", id))

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "user service unavailable"})
		return
	}

	c.Data(resp.StatusCode(), resp.Header().Get("Content-Type"), resp.Body())
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.client.R().
		Delete(fmt.Sprintf("/users/%s", id))

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "user service unavailable"})
		return
	}

	c.Status(resp.StatusCode())
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	resp, err := h.client.R().
		SetQueryParamsFromValues(c.Request.URL.Query()).
		Get("/users/all")

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "user service unavailable"})
		return
	}

	c.Data(resp.StatusCode(), resp.Header().Get("Content-Type"), resp.Body())
}
