package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/censoredplanet/cp-api/internal/api/experimental/model"
	"github.com/censoredplanet/cp-api/internal/database"
	"github.com/censoredplanet/cp-api/internal/entities"
)

type ExperimentalServicePort interface {
	HyperquackByASN(ctx context.Context, filter model.FilterHyperquack) ([]*entities.Hyperquack, error)
	SatelliteByASN(ctx context.Context, filter model.FilterSatellite) ([]*entities.Satellite, error)
}

func (s ServiceRepository) HyperquackByASN(ctx context.Context, filter model.FilterHyperquack) ([]*entities.Hyperquack, error) {
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
		if f == "__typename" {
			continue
		}
		if chCol, ok := database.GQLToCHHyperquack[f]; ok {
			if strings.HasPrefix(chCol, "received_tls") && strings.ToLower(strings.TrimSpace(filter.Protocol)) != "https" {
				continue
			}
			cols = append(cols, chCol)
		}
	}
	if len(cols) == 0 {
		return nil, fmt.Errorf("no valid database columns requested")
	}
	columns := strings.Join(cols, ", ")

	proto := strings.ToLower(strings.TrimSpace(filter.Protocol))
	switch proto {
	case "https", "http", "echo", "discard":
		res, err := s.clickHouseRepository.HyperquackByASN(ctx, entities.HyperquackFilterByASN{
			Protocol:  proto,
			Asn:       filter.Asn,
			StartDate: filter.StartDate.Format("2006-01-02"),
			EndDate:   filter.EndDate.Format("2006-01-02"),
		}, columns, fromMonth, toMonth)

		if err != nil {
			s.slack.Error("service.go", "Hyperquack", "clickHouseRepository.Hyperquack", err.Error())
			return nil, errGeneric
		}
		return res, nil
	default:
		return nil, fmt.Errorf("invalid protocol: %q, must be one of: HTTPS, HTTP, ECHO, DISCARD", filter.Protocol)
	}
}

func (s ServiceRepository) SatelliteByASN(ctx context.Context, filter model.FilterSatellite) ([]*entities.Satellite, error) {
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
		if f == "__typename" {
			continue
		}
		if chCol, ok := database.GQLToCHSatellite[f]; ok {
			cols = append(cols, chCol)
		}
	}
	if len(cols) == 0 {
		return nil, fmt.Errorf("no valid database columns requested")
	}
	columns := strings.Join(cols, ", ")

	res, err := s.clickHouseRepository.SatelliteByASN(ctx, entities.SatelliteFilterByASN{
		Asn:       filter.Asn,
		StartDate: filter.StartDate.Format("2006-01-02"),
		EndDate:   filter.EndDate.Format("2006-01-02"),
	}, columns, fromMonth, toMonth)
	if err != nil {
		s.slack.Error("service.go", "Satellite", "clickHouseRepository.Satellite", err.Error())
		return nil, errGeneric
	}
	return res, nil
}
