package errors

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ParseValidationErrors(err error) []FieldError {
	var ve validator.ValidationErrors
	if !errors.As(err, &ve) {
		return nil
	}

	fields := make([]FieldError, len(ve))
	for i, fe := range ve {
		fields[i] = FieldError{
			Field:   strings.ToLower(fe.Field()),
			Message: validationMessage(fe),
		}
	}
	return fields
}

func validationMessage(fe validator.FieldError) string {
	field := strings.ToLower(fe.Field())
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "email":
		return "invalid email format"
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", field, fe.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", field, fe.Param())
	default:
		return fmt.Sprintf("%s is invalid", field)
	}
}

// AppError represents a structured application error
type AppError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
	Err        error  `json:"-"`
}

// Error implements the error interface
func (e *AppError) Error() string {
	return e.Message
}

// NewValidationError creates a validation error (400)
func NewValidationError(code, message string) *AppError {
	err := &AppError{
		Code:       code,
		Message:    message,
		StatusCode: http.StatusBadRequest,
	}
	log.Printf("[VALIDATION_ERROR] %s: %s\n", code, message)
	return err
}

// NewConflictError creates a conflict error (409)
func NewConflictError(code, message string) *AppError {
	err := &AppError{
		Code:       code,
		Message:    message,
		StatusCode: http.StatusConflict,
	}
	log.Printf("[CONFLICT_ERROR] %s: %s\n", code, message)
	return err
}

// NewNotFoundError creates a not found error (404)
func NewNotFoundError(code, message string) *AppError {
	err := &AppError{
		Code:       code,
		Message:    message,
		StatusCode: http.StatusNotFound,
	}
	log.Printf("[NOT_FOUND_ERROR] %s: %s\n", code, message)
	return err
}

// NewUnauthorizedError creates an unauthorized error (401)
func NewUnauthorizedError(code, message string) *AppError {
	err := &AppError{
		Code:       code,
		Message:    message,
		StatusCode: http.StatusUnauthorized,
	}
	log.Printf("[UNAUTHORIZED_ERROR] %s: %s\n", code, message)
	return err
}

// NewInternalError creates an internal server error (500)
func NewInternalError(code, message string, err error) *AppError {
	appErr := &AppError{
		Code:       code,
		Message:    message,
		StatusCode: http.StatusInternalServerError,
		Err:        err,
	}
	if err != nil {
		log.Printf("[INTERNAL_ERROR] %s: %s (underlying: %v)\n", code, message, err)
	} else {
		log.Printf("[INTERNAL_ERROR] %s: %s\n", code, message)
	}
	return appErr
}

// HandleGormError converts GORM errors to AppError
func HandleGormError(err error, context string) *AppError {
	if err == nil {
		return nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return NewNotFoundError(
			"RECORD_NOT_FOUND",
			fmt.Sprintf("%s not found", context),
		)
	}

	// Check for PostgreSQL unique constraint violation (specific fields first)
	if strings.Contains(err.Error(), "duplicate key value") || strings.Contains(err.Error(), "violates unique constraint") {
		if strings.Contains(err.Error(), "email") {
			return NewConflictError(
				"DUPLICATE_EMAIL",
				"User with this email already exists",
			)
		}
		if strings.Contains(err.Error(), "company_name") {
			return NewConflictError(
				"DUPLICATE_COMPANY_NAME",
				"Client with this company name already exists",
			)
		}
		return NewConflictError(
			"DUPLICATE_ENTRY",
			fmt.Sprintf("%s already exists", context),
		)
	}

	// Generic database error
	return NewInternalError(
		"DATABASE_ERROR",
		fmt.Sprintf("Failed to access database for %s", context),
		err,
	)
}
