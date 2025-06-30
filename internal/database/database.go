package database

import (
	"context"
	"time"

	"github.com/censoredplanet/cp-api/internal/entities"
)

type DatabasePort interface {
	Hyperquack(ctx context.Context, filtere entities.HyperquackFilterCH, columns string, fromMonth string, tillMonth string) ([]*entities.Hyperquack, error)
	Satellite(ctx context.Context, filtere entities.SatelliteFilterCH, columns string, fromMonth string, tillMonth string) ([]*entities.Satellite, error)
	DashBoard(ctx context.Context, filtere entities.DashboardFilterCH, columns string, fromMonth string, tillMonth string) ([]*entities.Dashboard, error)
	TotalMeasurementsCount(ctx context.Context) (string, error)
	MeasurementsCountByDate(ctx context.Context, startDate time.Time, endDate time.Time) (string, error)
	InterferenceRateByCountry(ctx context.Context, startDate time.Time, endDate time.Time) ([]*entities.InterferenceRateByCountry, error)
	Domains(ctx context.Context, startDate time.Time, endDate time.Time, protocol string) ([]string, error)
	Countries(ctx context.Context, startDate time.Time, endDate time.Time, protocol string) ([]string, error)
}
