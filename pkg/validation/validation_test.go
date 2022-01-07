package validation

import (
	"testing"
)

func TestValidateUID_1(t *testing.T) {
	var uid int64
	uid = 12034

	err := ValidateUID(uid)
	if err != nil {
		t.Error(err)
	}
}

func TestValidateUID_2(t *testing.T) {
	var uid int64
	uid = 0

	err := ValidateUID(uid)
	if err != nil {
		t.Logf("Validation didn't pass for %v", uid)
	}
}

func TestValidateUsername_1(t *testing.T) {
	username := "abcd1234abcd1234~"
	err := ValidateUsername(username)
	if err != nil {
		t.Logf("Validation didn't pass for %v", username)
	}
}

func TestValidateUsername_2(t *testing.T) {
	username := "abcd1234abcd123~"
	err := ValidateUsername(username)
	if err != nil {
		t.Logf("Validation didn't pass for %v", username)
	}
}

func TestValidateUsername_3(t *testing.T) {
	username := "abcd1234abcd1234"
	err := ValidateUsername(username)
	if err != nil {
		t.Error(err)
	}
}

func TestValidateCoordinates_1(t *testing.T) {
	var latitude, longitude float64

	latitude = 39.12355
	longitude = 27.64538

	err := ValidateCoordinates(latitude, longitude)
	if err != nil {
		t.Error(err)
	}
}

func TestValidateCoordinates_2(t *testing.T) {
	var latitude, longitude float64

	latitude = -91.12355
	longitude = 182.64538

	err := ValidateCoordinates(latitude, longitude)
	if err != nil {
		t.Logf("Validation didn't pass for %v, %v", latitude, longitude)
	}
}
