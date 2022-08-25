package validatorutils

import (
	"errors"
	"fmt"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

// Validate - To validate store attributes
func Validate(myStruct interface{}) error {
	v := validator.New()
	err := v.Struct(myStruct)
	if err == nil {
		return nil
	}

	var result []string
	for _, e := range err.(validator.ValidationErrors) {
		result = append(result, fmt.Sprintf("%v", e))
	}

	if len(result) > 0 {
		return errors.New(strings.Join(result, "\n"))
	}

	return nil
}
