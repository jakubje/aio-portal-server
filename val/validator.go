package val

import (
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math"
	"regexp"
)

var (
	isValidEmail = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`).MatchString
	isValidName  = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
)

type Numeric interface {
	int32 | int64
}

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLength, maxLength)
	}
	return nil
}

func ValidateNumberLimit[T Numeric](value T, minLength T, maxLength T) error {
	if value < minLength || value > maxLength {
		return fmt.Errorf("must be in the range %d-%d", minLength, maxLength)
	}
	return nil
}

func ValidateSocialMediaLinks(links []string) error {
	// Check if the array is empty
	if len(links) == 0 {
		return fmt.Errorf("at least one social media link is required")
	}

	// Regular expression pattern to match valid URLs
	urlPattern := `^(https?|ftp)://[^\s/$.?#].[^\s]*$`
	regex := regexp.MustCompile(urlPattern)

	// Validate each link in the array
	for _, link := range links {
		if !regex.MatchString(link) {
			return fmt.Errorf("invalid social media link: %s", link)
		}
	}

	return nil
}

func ValidateID(value int64) error {
	if value <= 0 {
		return fmt.Errorf("invalid user id")
	}
	return nil
}

func ValidateInt[T Numeric](value T) error {
	if value <= 0 {
		return fmt.Errorf("invalid input, must be an integer")
	}
	return nil
}

func ValidateFloat(value float64) error {
	if math.IsNaN(value) && math.IsInf(value, 0) {
		fmt.Errorf("must be a valid number")
	}
	return nil
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidEmail(value) {
		return fmt.Errorf("must be a valid email address")
	}
	return nil
}

func ValidatePassword(value string) error {
	return ValidateString(value, 6, 100)
}

func ValidateName(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidName(value) {
		return fmt.Errorf("must contain only letters")
	}
	return nil
}

func ValidateEmailId(value int64) error {
	if value <= 0 {
		return fmt.Errorf("must be a positive integer")
	}
	return nil
}

func ValidateSecretCode(value string) error {
	return ValidateString(value, 32, 128)
}

func ValidateType(value int32) error {
	if value != 0 && value != 1 {
		return fmt.Errorf("type can only be a 0 (buy) or 1 (sell)")
	}
	return nil
}

func ValidateUUID(value string) error {
	_, err := uuid.Parse(value)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "error parsing UUID")
	}
	return nil
}
