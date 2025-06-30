package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/censoredplanet/cp-api/internal/api/graphql/model"
	"github.com/censoredplanet/cp-api/internal/database"
	"github.com/censoredplanet/cp-api/internal/entities"
	"github.com/censoredplanet/cp-api/internal/slack"
)

type ServicePort interface {
	Hyperquack(ctx context.Context, filter model.FilterHyperquack) ([]*entities.Hyperquack, error)
	Satellite(ctx context.Context, filter model.FilterSatellite) ([]*entities.Satellite, error)
	Dashboard(ctx context.Context, filter model.FilterDashboard) ([]*entities.Dashboard, error)
	TotalMeasurementsCount(ctx context.Context) (string, error)
	MeasurementsCountByDate(ctx context.Context, filter model.DateRange) (string, error)
	InterferenceRateByCountry(ctx context.Context, filter model.DateRange) ([]*entities.InterferenceRateByCountry, error)
	Domains(ctx context.Context, filter model.DateRange, protocol string) ([]string, error)
	Countries(ctx context.Context, filter model.DateRange, protocol string) ([]string, error)
}

type ServiceRepository struct {
	slack                slack.SlackPort
	clickHouseRepository database.DatabasePort
}

func NewService(slack slack.SlackPort, repository database.DatabasePort) (*ServiceRepository, error) {
	return &ServiceRepository{
		slack:                slack,
		clickHouseRepository: repository,
	}, nil
}

func (s ServiceRepository) Hyperquack(ctx context.Context, filter model.FilterHyperquack) ([]*entities.Hyperquack, error) {
	if len(filter.Country) != 2 {
		return nil, fmt.Errorf("country must be a 2-character ISO country code")
	}
	if filter.EndDate.Before(filter.StartDate) {
		return nil, fmt.Errorf("end date cannot be earlier than start date")
	}
	maxEndDate := filter.StartDate.AddDate(0, 3, 0)
	if filter.EndDate.After(maxEndDate) {
		return nil, fmt.Errorf("date range cannot exceed 3 months")
	}

	fromMonth := filter.StartDate.Format("200601")
	toMonth := filter.EndDate.Format("200601")

	requested := graphql.CollectAllFields(ctx)
	cols := make([]string, 0, len(requested))
	for _, f := range requested {
		if chCol, ok := database.GQLToCHHyperquack[f]; ok {
			if strings.HasPrefix(chCol, "received_tls") && strings.ToUpper(strings.TrimSpace(filter.Protocol)) != "HTTPS" {
				continue
			}
			cols = append(cols, chCol)
		}
	}
	columns := strings.Join(cols, ", ")

	switch strings.ToUpper(strings.TrimSpace(filter.Protocol)) {
	case "HTTPS", "HTTP", "ECHO", "DISCARD":
		res, err := s.clickHouseRepository.Hyperquack(ctx, entities.HyperquackFilterCH{
			Protocol:  filter.Protocol,
			Domain:    filter.Domain,
			Country:   filter.Country,
			StartDate: filter.StartDate.Format("2006-01-02"),
			EndDate:   filter.EndDate.Format("2006-01-02"),
		}, columns, fromMonth, toMonth)

		if err != nil {
			s.slack.Error("service.go", "Hyperquack", "clickHouseRepository.Hyperquack", err.Error())
			return nil, fmt.Errorf("something went wrong :(")
		}
		return res, nil
	default:
		return nil, fmt.Errorf("invalid protocol: %q, must be one of: HTTPS, HTTP, ECHO, DISCARD", filter.Protocol)
	}
}

func (s ServiceRepository) Satellite(ctx context.Context, filter model.FilterSatellite) ([]*entities.Satellite, error) {
	if len(filter.Country) != 2 {
		return nil, fmt.Errorf("country must be a 2-character ISO country code")
	}
	if filter.EndDate.Before(filter.StartDate) {
		return nil, fmt.Errorf("end date cannot be earlier than start date")
	}
	maxEndDate := filter.StartDate.AddDate(0, 3, 0)
	if filter.EndDate.After(maxEndDate) {
		return nil, fmt.Errorf("date range cannot exceed 3 months")
	}

	fromMonth := filter.StartDate.Format("200601")
	toMonth := filter.EndDate.Format("200601")

	requested := graphql.CollectAllFields(ctx)
	cols := make([]string, 0, len(requested))
	for _, f := range requested {
		if chCol, ok := database.GQLToCHSatellite[f]; ok {
			cols = append(cols, chCol)
		}
	}
	columns := strings.Join(cols, ", ")

	res, err := s.clickHouseRepository.Satellite(ctx, entities.SatelliteFilterCH{
		Domain:    filter.Domain,
		Country:   filter.Country,
		StartDate: filter.StartDate.Format("2006-01-02"),
		EndDate:   filter.EndDate.Format("2006-01-02"),
	}, columns, fromMonth, toMonth)
	if err != nil {
		s.slack.Error("service.go", "Satellite", "clickHouseRepository.Satellite", err.Error())
		return nil, fmt.Errorf("something went wrong :(")
	}
	return res, nil
}

func (s ServiceRepository) Dashboard(ctx context.Context, filter model.FilterDashboard) ([]*entities.Dashboard, error) {
	if len(filter.Domains) == 0 {
		return nil, fmt.Errorf("at least one domain needs to be provided")
	}
	if len(filter.Domains) > 10 {
		return nil, fmt.Errorf("sorry, but you can only query up to 10 domains at once")
	}
	if filter.EndDate.Before(filter.StartDate) {
		return nil, fmt.Errorf("end date cannot be earlier than start date")
	}
	maxEndDate := filter.StartDate.AddDate(0, 6, 0)
	if filter.EndDate.After(maxEndDate) {
		return nil, fmt.Errorf("date range cannot exceed 6 months")
	}
	fromMonth := filter.StartDate.Format("200601")
	toMonth := filter.EndDate.Format("200601")
	requested := graphql.CollectAllFields(ctx)
	cols := make([]string, 0, len(requested))
	for _, f := range requested {
		if chCol, ok := database.GQLToCHDashboard[f]; ok {
			cols = append(cols, chCol)
		}
	}
	columns := strings.Join(cols, ", ")

	switch strings.ToUpper(strings.TrimSpace(filter.Source)) {
	case "HTTPS", "HTTP", "ECHO", "DISCARD", "DNS":
		res, err := s.clickHouseRepository.DashBoard(ctx, entities.DashboardFilterCH{
			Source:    filter.Source,
			Domains:   filter.Domains,
			Country:   filter.Country,
			StartDate: filter.StartDate.Format("2006-01-02"),
			EndDate:   filter.EndDate.Format("2006-01-02"),
		}, columns, fromMonth, toMonth)

		if err != nil {
			s.slack.Error("service.go", "Dashboard", "clickHouseRepository.DashBoard", err.Error())
			return nil, fmt.Errorf("something went wrong :(")
		}
		return res, nil
	default:
		return nil, fmt.Errorf("invalid protocol: %q, must be one of: DNS, HTTPS, HTTP, ECHO, DISCARD", filter.Source)
	}
}

func (s ServiceRepository) TotalMeasurementsCount(ctx context.Context) (string, error) {
	res, err := s.clickHouseRepository.TotalMeasurementsCount(ctx)
	if err != nil {
		s.slack.Error("service.go", "TotalMeasurementsCount", "clickHouseRepository.TotalMeasurementsCount", err.Error())
		return "", fmt.Errorf("something went wrong :(")
	}
	return res, nil
}

func (s ServiceRepository) MeasurementsCountByDate(ctx context.Context, filter model.DateRange) (string, error) {
	if filter.EndDate.Before(filter.StartDate) {
		return "", fmt.Errorf("end date cannot be earlier than start date")
	}

	res, err := s.clickHouseRepository.MeasurementsCountByDate(ctx, filter.StartDate, filter.EndDate)
	if err != nil {
		s.slack.Error("service.go", "MeasurementsCountByDate", "clickHouseRepository.MeasurementsCountByDate", err.Error())
		return "", fmt.Errorf("something went wrong :(")
	}
	return res, nil
}

func (s ServiceRepository) InterferenceRateByCountry(ctx context.Context, filter model.DateRange) ([]*entities.InterferenceRateByCountry, error) {
	if filter.EndDate.Before(filter.StartDate) {
		return nil, fmt.Errorf("end date cannot be earlier than start date")
	}

	res, err := s.clickHouseRepository.InterferenceRateByCountry(ctx, filter.StartDate, filter.EndDate)
	if err != nil {
		s.slack.Error("service.go", "MeasurementsCountByDate", "clickHouseRepository.MeasurementsCountByDate", err.Error())
		return nil, fmt.Errorf("something went wrong :(")
	}
	return res, nil
}

func (s ServiceRepository) Domains(ctx context.Context, filter model.DateRange, protocol string) ([]string, error) {
	if filter.EndDate.Before(filter.StartDate) {
		return nil, fmt.Errorf("end date cannot be earlier than start date")
	}
	switch strings.ToUpper(strings.TrimSpace(protocol)) {
	case "DNS", "HTTPS", "HTTP", "ECHO", "DISCARD":
		res, err := s.clickHouseRepository.Domains(ctx, filter.StartDate, filter.EndDate, strings.ToLower(protocol))
		if err != nil {
			s.slack.Error("service.go", "Domains", "clickHouseRepository.Domains", err.Error())
			return nil, fmt.Errorf("something went wrong :(")
		}
		return res, nil
	default:
		return nil, fmt.Errorf("invalid protocol: %q, must be one of: DNS, HTTPS, HTTP, ECHO, DISCARD", protocol)
	}
}

func (s ServiceRepository) Countries(ctx context.Context, filter model.DateRange, protocol string) ([]string, error) {
	if filter.EndDate.Before(filter.StartDate) {
		return nil, fmt.Errorf("end date cannot be earlier than start date")
	}
	switch strings.ToUpper(strings.TrimSpace(protocol)) {
	case "DNS", "HTTPS", "HTTP", "ECHO", "DISCARD":
		res, err := s.clickHouseRepository.Countries(ctx, filter.StartDate, filter.EndDate, strings.ToLower(protocol))
		if err != nil {
			s.slack.Error("service.go", "Countries", "clickHouseRepository.Countries", err.Error())
			return nil, fmt.Errorf("something went wrong :(")
		}
		return res, nil
	default:
		return nil, fmt.Errorf("invalid protocol: %q, must be one of: DNS, HTTPS, HTTP, ECHO, DISCARD", protocol)
	}
}
