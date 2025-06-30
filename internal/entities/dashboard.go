package entities

import "time"

type Dashboard struct {
	Domain          string     `json:"domain" ch:"domain"`
	Date            *time.Time `json:"date,omitempty" ch:"date"`
	HostName        *string    `json:"hostName,omitempty" ch:"hostname"`
	RegHostName     *string    `json:"regHostName,omitempty" ch:"reg_hostname"`
	Network         *string    `json:"network,omitempty" ch:"network"`
	SubNetwork      *string    `json:"subNetwork,omitempty" ch:"subnetwork"`
	Category        string     `json:"category" ch:"category"`
	Outcome         string     `json:"outcome" ch:"outcome"`
	Count           int32      `json:"count" ch:"count"`
	UnexpectedCount int32      `json:"unexpectedCount" ch:"unexpected_count"`
	Country         string     `json:"country" ch:"country_name"`
	Source          string     `json:"source" ch:"source"`
}

type DashboardFilter struct {
	Country   string    `json:"country"`
	Domains   []string  `json:"domains"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Source    string    `json:"source"`
}

type DashboardFilterCH struct {
	Country   string   `json:"country"`
	Domains   []string `json:"domains"`
	StartDate string   `json:"startDate"`
	EndDate   string   `json:"endDate"`
	Source    string   `json:"source"`
}
