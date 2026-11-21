package main

import (
	"fmt"
	"log"
	"time"
)

const (
	// Сотрудники ставятся в смену скорингом. Для каждого дня вычислятеся количество баллов и сотрудник с наименьшим баллом
	// попадает на смену. Ниже константы, которые выдают баллы

	// Множитель смен. При отработанных пяти сменах сотрудик получает 500 балолов,
	// а сотрудник с 6 - 600, таким образом сотрудник с 5 сменами будет приоритетнее
	shiftMultiple int = 10
	// Награда за текущую серию:
	// каждый день текущей серии подряд даёт -100. Сортировка по возрастанию,
	// значит длинную серию выгодно продолжать.
	seriesMultiple int = -100
	// Штрафы и бонусы
	needSeriaScoreBonus int = -200 // Бонус за необходимости дработать серию
	preferredDayBonus   int = -50  // Вероятность того, что смена выпадет сотруднику, если он отметил день, как предпочтительный
	FineLongSeries      int = 20   // штраф, если серия дольше, чем минимальная
	FineTooLongSeries   int = 500  // Штраф, если серия дольше, чем максимальая

	// настройки
	maxSeries int = 3 // Максимальное количество дней работы подряд
	minSeries int = 2 // Минимальное количество работы денй подряд
)

func main() {
	emp1, err := NewEmployee(
		"Vasya",
		[]string{},
		[]string{
			"2026-12-01",
			"2026-12-07",
			"2026-12-15",
		},
		maxSeries,
		minSeries,
	)
	emp2, err := NewEmployee(
		"Petya",
		[]string{},
		[]string{
			"2026-12-05",
			"2026-12-22",
			"2026-12-31",
		},
		maxSeries,
		minSeries)
	emp3, err := NewEmployee(
		"Vanya",
		[]string{},
		[]string{
			"2026-12-11",
			"2026-12-30",
			"2026-12-31",
		},
		maxSeries,
		minSeries)
	emp4, err := NewEmployee(
		"Masha",
		[]string{},
		[]string{
			"2026-12-15",
			"2026-12-08",
			"2026-12-31",
		},
		maxSeries,
		minSeries)

	if err != nil {
		log.Fatal(err)
	}

	schedule, err := NewSchedule(
		time.Date(2026, time.December, 01, 0, 0, 0, 0, time.UTC), // Дата начала
		time.Date(2026, time.December, 31, 0, 0, 0, 0, time.UTC), // Конечная дата
		3,                                   // Необходимое количество на смене
		[]*Employee{emp1, emp2, emp3, emp4}, // Список сотрудников
	)

	for _, day := range schedule.Shifts {
		fmt.Printf("Дата: %s, день уомплектован: %t, Список сотрудников:\n", day.Date, !day.IsIncomplete)
		for _, employee := range day.Employees {
			fmt.Println(employee.Name)
		}
	}

	countAndPrintShifts(schedule)

}

func countAndPrintShifts(schedule *Schedule) {
	counts := make(map[string]int)

	// Считаем
	for _, day := range schedule.Shifts {
		for _, emp := range day.Employees {
			// Предполагаем, что у Employee есть поле Name
			counts[emp.Name]++
		}
	}

	// Выводим
	for name, shifts := range counts {
		fmt.Printf("%s: %d смен\n", name, shifts)
	}
}
