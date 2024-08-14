package main

import (
	"math"
	"testing"
)

func TestCalculateBMI(t *testing.T) {
	tests := []struct {
		name   string
		person PersonAnalysis
		want   float64
	}{
		{"Normal BMI", PersonAnalysis{Height: 175, Weight: 70}, 22.86},
		{"Underweight BMI", PersonAnalysis{Height: 175, Weight: 50}, 16.33},
		{"Overweight BMI", PersonAnalysis{Height: 175, Weight: 90}, 29.39},
		{"Obesity BMI", PersonAnalysis{Height: 175, Weight: 100}, 32.65},

		{"Normal BMI Faulty", PersonAnalysis{Height: 175, Weight: 70}, 25.00},      // Хибний тест 22.857142857142858
		{"Underweight BMI Faulty", PersonAnalysis{Height: 175, Weight: 50}, 18.00}, // Хибний тест 16.3265306122449
		{"Overweight BMI Faulty", PersonAnalysis{Height: 175, Weight: 90}, 35.00},  // Хибний тест 29.387755102040817
		{"Obesity BMI Faulty", PersonAnalysis{Height: 175, Weight: 100}, 40.00},    // Хибний тест 32.6530612244898
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.person.CalculateBMI()
			if math.Abs(got-tt.want) > 0.01 {
				t.Errorf("CalculateBMI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetHealthStatus(t *testing.T) {
	tests := []struct {
		name   string
		person PersonAnalysis
		want   string
	}{
		{"Underweight", PersonAnalysis{Height: 175, Weight: 50}, "Underweight"},
		{"Normal weight", PersonAnalysis{Height: 175, Weight: 70}, "Normal weight"},
		{"Overweight", PersonAnalysis{Height: 175, Weight: 90}, "Overweight"},
		{"Obesity", PersonAnalysis{Height: 175, Weight: 100}, "Obesity"},

		{"Underweight Faulty", PersonAnalysis{Height: 175, Weight: 50}, "Normal weight"}, // Хибний тест Underweight
		{"Normal weight Faulty", PersonAnalysis{Height: 175, Weight: 70}, "Obesity"},     // Хибний тест Normal weigh
		{"Overweight Faulty", PersonAnalysis{Height: 175, Weight: 90}, "Underweight"},    // Хибний тест Overweight
		{"Obesity Faulty", PersonAnalysis{Height: 175, Weight: 100}, "Overweight"},       // Хибний тест Obesity
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.person.GetHealthStatus()
			if got != tt.want {
				t.Errorf("GetHealthStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name   string
		person PersonAnalysis
		want   string
	}{
		{"Normal", PersonAnalysis{Name: "John Doe", Age: 30, Height: 175, Weight: 70},
			"Name: John Doe, Age: 30, Height: 175.00, Weight: 70.00, BMI: 22.86, Health Status: Normal weight"},

		{"Normal Faulty", PersonAnalysis{Name: "John Doe", Age: 30, Height: 175, Weight: 70},
			"Name: John Doe, Age: 30, Height: 175.00, Weight: 70.00, BMI: 25.00, Health Status: Obesity"}, // Хибний тест Name: John Doe, Age: 30, Height: 175.00, Weight: 70.00, BMI: 22.86, Health Status: Normal weight
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.person.String()
			if got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
