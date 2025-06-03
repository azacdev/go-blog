package errors

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type FormattedErrorResponse struct {
	Errors map[string]string `json:"errors"`
}

func ValidationErrorResponse(c *gin.Context, err error) {
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		formattedErrors := make(map[string]string)
		for _, fieldError := range validationErrors {
			formattedErrors[fieldError.Field()] = getErrorMessage(fieldError.Tag())
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"error":   "Bad Request",
			"message": formattedErrors,
		})
		return
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"error":   "Not Found",
			"message": "Resource not found",
		})
		return
	}

	if strings.Contains(err.Error(), "relation \"articles\" does not exist") {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"error":   "Internal Server Error",
			"message": "Articles' does not exist.",
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  http.StatusInternalServerError,
		"error":   "Internal Server Error",
		"message": err.Error(),
	})
}

func getErrorMessage(tag string) string {
	errorMessages := map[string]string{
		"required": "This field is required",
		"email":    "This field must be a valid email address",
		"min":      "Should be more than the limit",
		"max":      "Should be less than the limit",
	}
	if msg, ok := errorMessages[tag]; ok {
		return msg
	}
	return "Validation failed"
}

func FieldErrorResponse(c *gin.Context, statusCode int, message string) {
	errStatus := ""
	switch statusCode {
	case http.StatusBadRequest:
		errStatus = "Bad Request"
	case http.StatusConflict:
		errStatus = "Conflict"
	default:
		errStatus = http.StatusText(statusCode)
	}

	c.JSON(statusCode, gin.H{
		"status":  statusCode,
		"error":   errStatus,
		"message": message,
	})
}
