package transport

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	authdomain "github.com/warmdev17/hapi/internal/auth/domain"
)

func MapError(err error) (int, string, string) {
	var (
		validationErrs validator.ValidationErrors
		syntaxErr      *json.SyntaxError
		typeErr        *json.UnmarshalTypeError
	)

	switch {
	case errors.Is(err, authdomain.ErrEmailTaken):
		return http.StatusBadRequest, "EMAIL_TAKEN", "email already taken"

	case errors.Is(err, authdomain.ErrUsernameTaken):
		return http.StatusBadRequest, "USERNAME_TAKEN", "username already taken"

	case errors.Is(err, io.EOF):
		return http.StatusBadRequest,
			"INVALID_REQUEST",
			"request body is empty"

	case errors.As(err, &syntaxErr):
		return http.StatusBadRequest,
			"INVALID_JSON",
			"request body contains invalid JSON"

	case errors.As(err, &typeErr):
		return http.StatusBadRequest,
			"INVALID_JSON",
			"request contains an invalid field type"

	case errors.As(err, &validationErrs):
		return http.StatusBadRequest, "VALIDATION_ERROR", "request validation failed"
	}

	return http.StatusInternalServerError, "INTERNAL_SERVER", "an error occured while processing your request"
}
