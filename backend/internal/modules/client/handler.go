package client

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/internal/errors"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) CreateClient(c *gin.Context) {
	var req CreateClientRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		appErr := errors.NewValidationError("VALIDATION_ERROR", err.Error())
		c.JSON(appErr.StatusCode, gin.H{
			"error": gin.H{
				"code":    appErr.Code,
				"message": appErr.Message,
			},
		})
		return
	}

	client, err := h.service.CreateClient(c.Request.Context(), req)
	if err != nil {
		// Check if it's an AppError
		if appErr, ok := err.(*errors.AppError); ok {
			c.JSON(appErr.StatusCode, gin.H{
				"error": gin.H{
					"code":    appErr.Code,
					"message": appErr.Message,
				},
			})
			return
		}

		// Fallback for unexpected errors
		genericErr := errors.NewInternalError("UNKNOWN_ERROR", "An unexpected error occurred", err)
		c.JSON(genericErr.StatusCode, gin.H{
			"error": gin.H{
				"code":    genericErr.Code,
				"message": genericErr.Message,
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "client created successfully",
		"data": gin.H{
			"id":            client.ID,
			"uuid":          client.UUID,
			"company_name":  client.CompanyName,
			"contact_name":  client.ContactName,
			"contact_email": client.ContactEmail,
			"contact_phone": client.ContactPhone,
			"notes":         client.Notes,
		},
	})
}

func (h *Handler) FindAllClients(c *gin.Context) {
	clients, err := h.service.FindAllClients(c.Request.Context())
	if err != nil {
		// Check if it's an AppError
		if appErr, ok := err.(*errors.AppError); ok {
			c.JSON(appErr.StatusCode, gin.H{
				"error": gin.H{
					"code":    appErr.Code,
					"message": appErr.Message,
				},
			})
			return
		}

		// Fallback for unexpected errors
		genericErr := errors.NewInternalError("UNKNOWN_ERROR", "An unexpected error occurred", err)
		c.JSON(genericErr.StatusCode, gin.H{
			"error": gin.H{
				"code":    genericErr.Code,
				"message": genericErr.Message,
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "clients retrieved successfully",
		"data":    clients,
	})
}
