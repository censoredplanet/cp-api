package entities

type InterferenceRateByCountry struct {
	Country         string  `json:"country" ch:"country_name"`
	UnexpectedCount float64 `json:"unexpectedRate" ch:"unexpected_rate"`
}
