package main

import (
	"fmt"
	"strings"
	"time"
)

type Employee struct {
	Name           string
	PreferredDates []time.Time
	BlockedDates   []time.Time
}

func ValidateDates(datesA, datesB []string) error {
	const op = "validateDates"
	seenMap := make(map[string]bool, len(datesA))
	foundDuplicates := []string{}
	for _, element := range datesA {
		seenMap[element] = true
	}
	for _, element := range datesB {
		if seenMap[element] {
			foundDuplicates = append(foundDuplicates, element)
		}
	}
	if len(foundDuplicates) > 0 {
		return fmt.Errorf("%s: Same date found in both lists: %s", op, strings.Join(foundDuplicates, ", "))
	}
	return nil
}

func NewEmployee(name string, preferredDates []string, blockedDates []string) (*Employee, error) {
	const op = "newEmployee"

	err := ValidateDates(preferredDates, blockedDates)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	preferredDatesParsed := make([]time.Time, 0, len(preferredDates))
	blockedDatesParsed := make([]time.Time, 0, len(blockedDates))
	// Validate
	for _, date := range preferredDates {
		parsedDate, err := time.Parse(time.DateOnly, date)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		preferredDatesParsed = append(preferredDatesParsed, parsedDate)
	}
	for _, date := range blockedDates {
		parsedDate, err := time.Parse(time.DateOnly, date)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		blockedDatesParsed = append(blockedDatesParsed, parsedDate)
	}

	return &Employee{
		Name:           name,
		PreferredDates: preferredDatesParsed,
		BlockedDates:   blockedDatesParsed,
	}, nil
}
func main() {
	emp1, err := NewEmployee(
		"Vasya",
		[]string{
			"2026-12-25",
			"2026-12-26",
			"2026-12-27",
		},
		[]string{
			"2026-12-26",
			"2026-12-29",
			"2026-12-30",
		})
	fmt.Println(err, emp1)

}
