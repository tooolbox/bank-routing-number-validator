package validator

import (
	"testing"
)

func TestValidator(t *testing.T) {

	t.Run("Should fail an alphanumeric value", func(t *testing.T) {
		err := ABARoutingNumberIsValid("abcdacbde")
		if err != ErrNotMatch {
			t.Errorf("Expected %v but found: %v", ErrNotMatch, err)
		}
	})

	t.Run("Should fail a number greater than 9 digits", func(t *testing.T) {
		err := ABARoutingNumberIsValid("1234567890")
		if err != ErrTooLong {
			t.Errorf("Expected %v but found: %v", ErrTooLong, err)
		}
	})

	t.Run("Should pass this real bank routing number", func(t *testing.T) {
		err := ABARoutingNumberIsValid("021000021")
		if err != nil {
			t.Errorf("Expected no error but found: %v", err)
		}
	})

	t.Run("Should pass this real bank routing number", func(t *testing.T) {
		err := ABARoutingNumberIsValid("322271627")
		if err != nil {
			t.Errorf("Expected no error but found: %v", err)
		}
	})

	t.Run("Should pass this real bank routing number", func(t *testing.T) {
		err := ABARoutingNumberIsValid("061102400")
		if err != nil {
			t.Errorf("Expected no error but found: %v", err)
		}
	})

	t.Run("Should handle omitted leading zero", func(t *testing.T) {
		err := ABARoutingNumberIsValid("21000021")
		if err != nil {
			t.Errorf("Expected no error but found: %v", err)
		}
	})

	t.Run("Should fail this real-looking number (first two digits)", func(t *testing.T) {
		err := ABARoutingNumberIsValid("131000021")
		if err != ErrInvalidFirstTwo {
			t.Errorf("Expected %v but found: %v", ErrInvalidFirstTwo, err)
		}
	})

	t.Run("Should fail this real-looking number (checksum)", func(t *testing.T) {
		err := ABARoutingNumberIsValid("322271628")
		if err != ErrInvalidChecksum {
			t.Errorf("Expected %v but found: %v", ErrInvalidChecksum, err)
		}
	})

	t.Run("Should fail an empty string", func(t *testing.T) {
		err := ABARoutingNumberIsValid("")
		if err != ErrInactive {
			t.Errorf("Expected %v but found: %v", ErrInactive, err)
		}
	})

}
