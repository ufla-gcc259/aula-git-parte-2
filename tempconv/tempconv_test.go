package tempconv

import (
	"math"
	"testing"
)

func TestCToF(t *testing.T) {
	testCases := []struct {
		celsius    Celsius
		fahrenheit Fahrenheit
	}{
		{AbsoluteZeroC, -459.67},
		{FreezingC, 32},
		{BoilingC, 212},
		{100, 212},
		{-40, -40},
	}

	for _, tc := range testCases {
		t.Run(tc.celsius.String(), func(t *testing.T) {
			result := CToF(tc.celsius)
			expected := tc.fahrenheit
			tolerance := 0.01

			if math.Abs(float64(result-expected)) > tolerance {
				t.Errorf("Expected %s to be %s, but got %s", tc.celsius, tc.fahrenheit, result)
			}
		})
	}
}

func TestFToC(t *testing.T) {
	testCases := []struct {
		fahrenheit Fahrenheit
		celsius    Celsius
	}{
		{-459.67, AbsoluteZeroC},
		{32, FreezingC},
		{212, BoilingC},
		{100, 37.77777777777778},
		{-40, -40},
	}

	for _, tc := range testCases {
		t.Run(tc.fahrenheit.String(), func(t *testing.T) {
			result := FToC(tc.fahrenheit)
			if result != tc.celsius {
				t.Errorf("Expected %s to be %s, but got %s", tc.fahrenheit, tc.celsius, result)
			}
		})
	}
}

func TestCelsiusString(t *testing.T) {
	testCases := []struct {
		celsius  Celsius
		expected string
	}{
		{AbsoluteZeroC, "-273.15°C"},
		{FreezingC, "0°C"},
		{BoilingC, "100°C"},
		{37.5, "37.5°C"},
		{-40, "-40°C"},
	}

	for _, tc := range testCases {
		t.Run(tc.celsius.String(), func(t *testing.T) {
			result := tc.celsius.String()
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}

func TestFahrenheitString(t *testing.T) {
	testCases := []struct {
		fahrenheit Fahrenheit
		expected   string
	}{
		{-459.67, "-459.67°F"},
		{32, "32°F"},
		{212, "212°F"},
		{98.6, "98.6°F"},
		{-40, "-40°F"},
	}

	for _, tc := range testCases {
		t.Run(tc.fahrenheit.String(), func(t *testing.T) {
			result := tc.fahrenheit.String()
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}

func TestZeroCelsiusToFahrenheit(t *testing.T) {
	result := CToF(0)
	if result != 32 {
		t.Errorf("Expected 0°C to be 32°F, but got %s", result)
	}
}

func TestZeroFahrenheitToCelsius(t *testing.T) {
	result := FToC(32)
	if result != 0 {
		t.Errorf("Expected 32°F to be 0°C, but got %s", result)
	}
}
