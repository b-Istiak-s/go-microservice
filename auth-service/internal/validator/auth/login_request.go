package validator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// LoginRequest represents the expected request body for user login.
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func BindAndValidateLogin(c *gin.Context, obj interface{}) (bool, map[string]string) {
	if err := c.ShouldBindJSON(obj); err != nil {
		// try to extract validator.ValidationErrors
		var verrs validator.ValidationErrors
		errs := make(map[string]string)

		if errors.As(err, &verrs) {
			// validation‐error path
			for _, fe := range verrs {
				// use the struct field name (or JSON tag) as the key
				field := strings.ToLower(fe.Field())
				var msg string
				switch fe.Tag() {
				case "required":
					msg = fmt.Sprintf("%s is required", fe.Field())
				case "email":
					msg = fmt.Sprintf("%s must be a valid email", fe.Field())
				default:
					msg = fmt.Sprintf("%s is invalid", fe.Field())
				}
				errs[field] = msg
			}
			return false, errs
		}

		// non‐validation error (e.g. malformed JSON)
		errs["error"] = err.Error()
		return false, errs
	}

	// success
	return true, nil
}
