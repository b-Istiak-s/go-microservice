package validator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BindAndValidateNote(c *gin.Context, obj interface{}) (bool, map[string]string) {
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
				case "min":
					msg = fmt.Sprintf("%s must be at least %s characters", fe.Field(), fe.Param())
				case "max":
					msg = fmt.Sprintf("%s must be at most %s characters", fe.Field(), fe.Param())
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
