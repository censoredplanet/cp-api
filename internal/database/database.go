package database

import (
	"context"
	"time"

	"github.com/censoredplanet/cp-api/internal/entities"
)

type DatabasePort interface {
	Hyperquack(ctx context.Context, filter entities.HyperquackFilterCH, columns string, fromMonth string, tillMonth string) ([]*entities.Hyperquack, error)
	Satellite(ctx context.Context, filter entities.SatelliteFilterCH, columns string, fromMonth string, tillMonth string) ([]*entities.Satellite, error)
	DashBoard(ctx context.Context, filter entities.DashboardFilterCH, columns string, fromMonth string, tillMonth string) ([]*entities.Dashboard, error)
	TotalMeasurementsCount(ctx context.Context) (string, error)
	MeasurementsCountByDate(ctx context.Context, startDate time.Time, endDate time.Time) (string, error)
	InterferenceRateByCountry(ctx context.Context, startDate time.Time, endDate time.Time) ([]*entities.InterferenceRateByCountry, error)
	Domains(ctx context.Context, startDate time.Time, endDate time.Time, protocol string) ([]string, error)
	Countries(ctx context.Context, startDate time.Time, endDate time.Time, protocol string) ([]string, error)
	CenAlertTimeSeries(ctx context.Context, columns string, startDate, endDate string, country string) ([]*entities.CenAlertTimeSeries, error)
	CenAlertCountries(ctx context.Context) ([]string, error)
	CenAlertEvents(ctx context.Context, columns string, startDate, endDate string, country *string) ([]*entities.CenAlertEvents, error)
}
