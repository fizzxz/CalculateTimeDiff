package main

import (
	"fmt"
	"time"
)

func main() {
	originDateInput := time.Date(2000, 01, 01, 00, 00, 00, 00, time.UTC)

	ageInYears := yearsCountSince(originDateInput)
	ageInMonths := monthsCountSince(originDateInput)
	ageInDays := daysCountSince(originDateInput)
	ageInMinutes := minutesCountSince(originDateInput)
	ageInSeconds := secsCountSince(originDateInput)

	fmt.Println("Age in Years:")
	fmt.Println(ageInYears)

	fmt.Println("Age in months:")
	fmt.Println(ageInMonths)

	fmt.Println("Age in Days: ")
	fmt.Println(ageInDays)

	fmt.Println("Age in Minutes: ")
	fmt.Println(ageInMinutes)

	fmt.Println("Age in Seconds: ")
	fmt.Println(ageInSeconds)

}

func yearsCountSince(originDate time.Time) int {
	now := time.Now()

	day := now.Sub(originDate)
	years := (day.Hours() / 24) / 365

	return int(years)
}

func monthsCountSince(originDate time.Time) int {
	now := time.Now()

	day := now.Sub(originDate)
	months := (day.Hours() / 24) / 365 * 12

	return int(months)
}

func daysCountSince(originDate time.Time) int {
	now := time.Now()

	day := now.Sub(originDate)
	days := (day.Hours() / 24)

	return int(days)
}

func minutesCountSince(originDate time.Time) int {
	now := time.Now()
	mins := now.Sub(originDate).Minutes()
	return int(mins)
}

func secsCountSince(originDate time.Time) int {
	now := time.Now()
	secs := now.Sub(originDate).Seconds()
	return int(secs)
}
