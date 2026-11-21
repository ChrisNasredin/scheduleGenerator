package main

import "time"

type Shift struct {
	Date      time.Time   `json:"date"`
	Employees []*Employee `json:"employees"`
}

type Schedule struct {
	DateFrom time.Time
	DateTo   time.Time
	Shifts   []*Shift
}

func NewSchedule(
	dateFrom time.Time,
	dateTo time.Time,
	necessaryCount int,
	employees []*Employee,
	minSequence int,
	maxSequence int,
) (*Schedule, error) {

	return nil, nil
}
