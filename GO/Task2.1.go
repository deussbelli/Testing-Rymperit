package main

import (
	"fmt"
	"time"
)

func calculateAge(birthDate string) int {
	layout := "2006-01-02"
	bd, _ := time.Parse(layout, birthDate)
	today := time.Now()
	age := today.Year() - bd.Year()
	if today.YearDay() < bd.YearDay() {
		age--
	}
	return age
}
func main() {
	age := calculateAge("1990-01-01")
	fmt.Println("Age:", age)
}
