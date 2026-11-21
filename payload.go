package main

import (
	"time"
)

type scheduleRequest struct {
	NecessaryCount int        `json:"necessary_count"`
	DateFrom       time.Time  `json:"dateFrom"`
	DateTo         time.Time  `json:"dateTo"`
	Employees      []Employee `json:"employees"`
}
type scheduleResponse struct {
	Month  time.Time     `json:"month"`
	Shifts []DaySchedule `json:"shifts"`
}
