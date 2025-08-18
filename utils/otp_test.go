package utils

import (
	"testing"
)

func TestGenerateOTP(t *testing.T) {
	tests := []struct {
		name      string
		length    int
		expectErr bool
	}{
		{"ValidLength", 6, false},
		{"ZeroLength", 0, false},
		// {"NegativeLength", -1, false}, // Should handle gracefully, though not expected
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateOTP(tt.length)
			if (err != nil) != tt.expectErr {
				t.Errorf("GenerateOTP() error = %v, wantErr %v", err, tt.expectErr)
				return
			}
			if len(got) != tt.length {
				t.Errorf("GenerateOTP() = %v, want %v", got, tt.length)
			}
		})
	}
}
