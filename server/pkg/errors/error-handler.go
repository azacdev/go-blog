package errors

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm" // Import GORM for potential database error checks
)

type ErrorResponse struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors,omitempty"` // Optional field
}

// ValidationErrorResponse sends an appropriate HTTP response based on the error type
func ValidationErrorResponse(c *gin.Context, err error) {
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		formattedErrors := make(map[string]string)
		for _, fieldError := range validationErrors {
			formattedErrors[fieldError.Field()] = getErrorMessage(fieldError.Tag())
		}
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: "Validation failed",
			Errors:  formattedErrors,
		})
		return
	}

	// Check for common database errors (using GORM's error type if applicable)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Status:  "error",
			Message: "Resource not found",
		})
		return
	}

	// Check for the specific "relation does not exist" error (or similar database errors)
	if strings.Contains(err.Error(), "relation \"articles\" does not exist") {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  "error",
			Message: "Database table 'articles' does not exist",
		})
		return
	}

	// Fallback for other internal server errors
	c.JSON(http.StatusInternalServerError, ErrorResponse{
		Status:  "error",
		Message: "Internal server error",
	})
}

func getErrorMessage(tag string) string {
	errorMessages := map[string]string{
		"required": "This field is required",
		"email":    "This field must be a valid email address",
		"min":      "Should be more than the limit",
		"max":      "Should be less than the limit",
		// Add other validation tags and their corresponding messages
	}
	if msg, ok := errorMessages[tag]; ok {
		return msg
	}
	return "Validation failed" // Default message if tag is not found
}
