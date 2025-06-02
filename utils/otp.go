package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenerateOTP(length int) (string, error) {
	if length < 0 {
        return "", fmt.Errorf("length cannot be negative")
    }
	digits := "0123456789"
	otp := ""

	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", fmt.Errorf("failed to generate OTP: %w", err)
		}
		otp += string(digits[num.Int64()])
	}

	return otp, nil
}

// CheckOTP checks if the provided OTP is valid for the given email.
// This is a placeholder function and should be replaced with actual OTP validation logic.
func CheckOTP(email, otp string) (bool, error) {
	// Placeholder logic: In a real application, you would check the OTP against a database or cache.
	// For this example, we'll assume the OTP is always valid.
	return true, nil
}