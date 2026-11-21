package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	emp1, err := NewEmployee(
		"Vasya",
		[]string{},
		[]string{
			"2026-12-1",
			"2026-12-7",
			"2026-12-15",
		})
	emp2, err := NewEmployee(
		"Petya",
		[]string{},
		[]string{
			"2026-12-5",
			"2026-12-22",
			"2026-12-31",
		})
	emp3, err := NewEmployee(
		"Vanya",
		[]string{},
		[]string{
			"2026-12-11",
			"2026-12-30",
			"2026-12-31",
		})
	emp4, err := NewEmployee(
		"Masha",
		[]string{},
		[]string{
			"2026-12-15",
			"2026-12-8",
			"2026-12-31",
		})

	if err != nil {
		log.Fatal(err)
	}

	schedule, err := NewSchedule(
		time.Date(2026, time.December, 1, 0, 0, 0, 0, time.UTC),  // Дата начала
		time.Date(2026, time.December, 31, 0, 0, 0, 0, time.UTC), // Конечная дата
		2,                                   // Необходимое количество на смене
		[]*Employee{emp1, emp2, emp3, emp4}, // Список сотрудников
		1,                                   // inConsecutive int  Минимум смен подряд
		3,                                   // MaxConsecutive int  Максимум смен подряд
	)
	fmt.Println(schedule)
}
