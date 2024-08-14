package main

import (
	"fmt"
	"math"
)

type PersonAnalysis struct {
	Name   string
	Age    int
	Height float64
	Weight float64
}

func (p *PersonAnalysis) CalculateBMI() float64 {
	return p.Weight / math.Pow(p.Height/100, 2)
}

func (p *PersonAnalysis) GetHealthStatus() string {
	bmi := p.CalculateBMI()
	switch {
	case bmi < 18.5:
		return "Underweight"
	case bmi >= 18.5 && bmi < 24.9:
		return "Normal weight"
	case bmi >= 25 && bmi < 29.9:
		return "Overweight"
	default:
		return "Obesity"
	}
}

func (p *PersonAnalysis) String() string {
	return fmt.Sprintf("Name: %s, Age: %d, Height: %.2f, Weight: %.2f, BMI: %.2f, Health Status: %s",
		p.Name, p.Age, p.Height, p.Weight, p.CalculateBMI(), p.GetHealthStatus())
}

func main() {
	person := PersonAnalysis{
		Name:   "John Doe",
		Age:    30,
		Height: 175,
		Weight: 70,
	}

	fmt.Println(person.String())
}
