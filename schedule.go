package main

import (
	"slices"
	"time"
)

type DaySchedule struct {
	Date         time.Time   `json:"date"`
	Employees    []*Employee `json:"employees"`
	IsIncomplete bool        `json:"isIncomplete"`
}

type Schedule struct {
	DateFrom time.Time
	DateTo   time.Time
	Shifts   []*DaySchedule
}

type EmployeeState struct {
	employee    *Employee
	blocked     map[time.Time]bool
	preferred   map[time.Time]bool
	totalShifts int
	shiftSeries int // Текущее количество смен подряд
	minSeries   int
	maxSeries   int
}

func NewSchedule(dateFrom time.Time, dateTo time.Time, necessaryCount int, employees []*Employee) (*Schedule, error) {
	const op = "newSchedule"

	states := createStates(employees)
	currentDate := dateFrom

	scheduleDaysSlice := make([]*DaySchedule, 0)

	for !currentDate.After(dateTo) {
		necessaryCountDay := necessaryCount
		// Смотрим, кто физически подходит под этот день:
		// У кого серия не превышает день подряд
		// У кого не стоит заблокированного дня
		// Собираем их стейты в слайс available
		var available []*EmployeeState
		for _, state := range states {
			// TODO: Переделать на тоггл.
			// День не выбран, как заблокированный
			if state.blocked[currentDate] {
				continue
			}
			available = append(available, state)
		}

		// Сортируем кандидатов по приоритету
		// это базовая функция, которая считает скоринг
		slices.SortFunc(available, func(a, b *EmployeeState) int {
			return calculateScore(a, currentDate) - calculateScore(b, currentDate)
		})

		// Набираем смену на день
		assignedEmployes := make([]*Employee, 0, 2)
		if len(available) < necessaryCountDay {
			necessaryCountDay = len(available)
		}
		for i := 0; i < necessaryCountDay; i++ {
			chosen := available[i]
			assignedEmployes = append(assignedEmployes, chosen.employee)
		}
		// Проходимся по ВСЕМ статусам, обновляем (по всем - если серия оборвалась - прервать)
		for _, ws := range states {
			if slices.Contains(assignedEmployes, ws.employee) {
				ws.totalShifts++
				ws.shiftSeries++
			} else {
				ws.shiftSeries = 0
			}

		}

		day := DaySchedule{
			Date:         currentDate,
			Employees:    assignedEmployes,
			IsIncomplete: len(assignedEmployes) < necessaryCount,
		}

		scheduleDaysSlice = append(scheduleDaysSlice, &day)
		currentDate = currentDate.AddDate(0, 0, 1)
	}

	return &Schedule{
		DateFrom: dateFrom,
		DateTo:   dateTo,
		Shifts:   scheduleDaysSlice,
	}, nil
}

func createStates(employees []*Employee) []*EmployeeState {
	states := make([]*EmployeeState, len(employees))
	for index, employee := range employees {
		ws := &EmployeeState{
			employee:  employee,
			blocked:   make(map[time.Time]bool),
			preferred: make(map[time.Time]bool),
			minSeries: employee.minSeries,
			maxSeries: employee.maxSeries,
		}
		for _, d := range employee.BlockedDates {
			ws.blocked[d] = true
		}
		for _, d := range employee.PreferredDates {
			ws.preferred[d] = true
		}
		states[index] = ws
	}

	return states
}

func calculateScore(empState EmployeeState, currentDate time.Time) int {
	// Базовое - количество смен
	score := empState.totalShifts * shiftMultiple
	// Убавляем баллы, если не хватает до минимальной серии
	if empState.shiftSeries < empState.minSeries {
		score = score + needSeriaScoreBonus
	}
	// Добавляем, если превысили максимальную серию
	if empState.shiftSeries > empState.maxSeries {
		score = score + FineTooLongSeries
	}
	// Штраф за превышение минимальной серии
	if empState.shiftSeries > empState.minSeries {
		score = score + FineLongSeries
	}
	// Бонус за предпочтительный день
	if empState.preferred[currentDate] {
		score = score + preferredDayBonus
	}
	return score
}
