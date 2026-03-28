package auth

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

func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		appErr := errors.NewValidationError("VALIDATION_ERROR", "Validation failed")
		c.JSON(appErr.StatusCode, gin.H{
			"error": gin.H{
				"code":    appErr.Code,
				"message": appErr.Message,
				"fields":  errors.ParseValidationErrors(err),
			},
		})
		return
	}

	user, token, err := h.service.Register(c.Request.Context(), &req)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			c.JSON(appErr.StatusCode, gin.H{
				"error": gin.H{
					"code":    appErr.Code,
					"message": appErr.Message,
				},
			})
			return
		}

		genericErr := errors.NewInternalError("UNKNOWN_ERROR", "An unexpected error occured", err)
		c.JSON(genericErr.StatusCode, gin.H{
			"error": gin.H{
				"code":    genericErr.Code,
				"message": genericErr.Message,
			},
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "successfully registered",
		"data": gin.H{
			"uuid":  user.UUID,
			"email": user.Email,
			"name":  user.Name,
		},
		"token": token,
	})
}

func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest

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

	user, token, err := h.service.Login(c.Request.Context(), req.Email, req.Password)
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
		"message": "login berhasil",
		"user": gin.H{
			"uuid":  user.UUID,
			"email": user.Email,
			"name":  user.Name,
		},
		"token": token,
	})
}
