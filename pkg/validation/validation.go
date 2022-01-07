package validation

import (
	"errors"
	appErrors "github.com/kvendingoldo/gu-common/pkg/errors"
	"regexp"
)

func ValidateUID(uid int64) error {
	if uid <= 0 {
		return appErrors.NewAppError(errors.New("ID should be >= 0"), appErrors.ValidationError)
	}
	return nil
}

func ValidateUsername(username string) error {
	rgx := regexp.MustCompile(`^[A-Za-z0-9]+$`).MatchString

	if !rgx(username) {
		return appErrors.NewAppError(errors.New("only (a-zA-Z0-9) symbols are acceptable for username"), appErrors.ValidationError)
	}

	length := len(username)
	if length < 4 || length > 16 {
		return appErrors.NewAppError(errors.New("username should contain 4-16 symbols"), appErrors.ValidationError)

	}

	return nil
}

func ValidateCoordinates(latitude, longitude float64) error {
	if !(latitude < 90.0 && latitude > -90.0) {
		return errors.New("wrong latitude coordinate; It should be between -90 and 90")
	}

	if !(longitude < 180.0 && longitude > -180.0) {
		return errors.New("wrong latitude coordinate; It should be between -180 and 180")
	}

	return nil
}

func ValidateDate() error {
	//dates - use ISO 8601 date format (2021-09-02T11:26:18+00:00)
	return nil
}
