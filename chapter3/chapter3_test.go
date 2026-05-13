package chapter3

import "testing"

func TestToFahrentheit(t *testing.T) {
	tests := []struct {
		input    Celsius
		expected Fahrenheit
	}{
		{0, 32},
		{100, 212},
		{-40, -40},
		{37, 98.6},
	}

	for _, tt := range tests {
		result := tt.input.ToFahrenheit()
		if result != tt.expected {
			t.Errorf("Celsius(%v).ToFahrenheit() = %v, want %v",
				tt.input, result, tt.expected)
		}
	}
}
