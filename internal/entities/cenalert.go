package entities

import "time"

type CenAlertTimeSeries struct {
	Value   float64   `json:"value" ch:"value"`
	Date    time.Time `json:"date"  ch:"date"`
	Country string    `json:"country" ch:"country"`
}

type CenAlertEvents struct {
	Country    string    `json:"country" ch:"country"`
	StartDate  time.Time `json:"startDate"  ch:"start"`
	EndDate    time.Time `json:"endDate"  ch:"end"`
	Peak       time.Time `json:"peak"  ch:"peak"`
	Impact     float64   `json:"impact" ch:"impact"`
	Cause      string    `json:"cause" ch:"cause"`
	ReportedBy string    `json:"reportedBy" ch:"reported_by"`
}
