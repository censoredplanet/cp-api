package database

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/censoredplanet/cp-api/internal/entities"

	"github.com/ClickHouse/clickhouse-go/v2"
)

type clickHouseRepository struct {
	client *driver.Conn
}

func NewClickHouse(client *driver.Conn) (*clickHouseRepository, error) {
	ch := &clickHouseRepository{
		client: client,
	}
	return ch, nil
}

func ClickHouseConnect() (driver.Conn, error) {
	ctx := context.Background()
	client, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{os.Getenv("CLICKHOUSE_URL")},
		Auth: clickhouse.Auth{
			Database: os.Getenv("DATABASE"),
			Username: os.Getenv("USERNAME"),
			Password: os.Getenv("PASSWORD"),
		},
	})
	if err != nil {
		return nil, err
	}
	if err = client.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return nil, err
	}
	return client, nil
}

func (ch clickHouseRepository) Hyperquack(ctx context.Context, filter entities.HyperquackFilterCH, columns string, fromMonth string, tillMonth string) ([]*entities.Hyperquack, error) {
	query := fmt.Sprintf(
		"SELECT %s FROM base.%s WHERE server_country = ? AND domain = ? AND yyyymm BETWEEN ? AND ? AND date BETWEEN ? AND ?",
		columns,
		filter.Protocol,
	)
	rows, err := (*ch.client).Query(ctx, query, filter.Country, filter.Domain, fromMonth, tillMonth, filter.StartDate, filter.EndDate)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var results []*entities.Hyperquack

	for rows.Next() {
		var res entities.Hyperquack
		if err := rows.ScanStruct(&res); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		results = append(results, &res)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	return results, nil
}

func (ch clickHouseRepository) Satellite(ctx context.Context, filter entities.SatelliteFilterCH, columns string, fromMonth string, tillMonth string) ([]*entities.Satellite, error) {
	query := fmt.Sprintf(
		"SELECT %s FROM base.satellite WHERE resolver_country = ? AND domain = ? AND yyyymm BETWEEN ? AND ? AND date BETWEEN ? AND ?",
		columns,
	)
	rows, err := (*ch.client).Query(ctx, query, filter.Country, filter.Domain, fromMonth, tillMonth, filter.StartDate, filter.EndDate)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var results []*entities.Satellite

	for rows.Next() {
		var res entities.Satellite
		if err := rows.ScanStruct(&res); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		results = append(results, &res)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	return results, nil
}

func (ch clickHouseRepository) DashBoard(ctx context.Context, filter entities.DashboardFilterCH, columns string, fromMonth string, tillMonth string) ([]*entities.Dashboard, error) {
	placeholders := make([]string, len(filter.Domains))
	for i := range placeholders {
		placeholders[i] = "?"
	}
	inClause := strings.Join(placeholders, ",")

	query := fmt.Sprintf(
		"SELECT %s FROM cp.derived WHERE country_name = ? AND domain IN (%s) AND yyyymm BETWEEN ? AND ? AND date BETWEEN ? AND ? AND source = ?",
		columns,
		inClause,
	)

	params := make([]interface{}, 0, len(filter.Domains)+5)
	params = append(params, filter.Country)

	for _, domain := range filter.Domains {
		params = append(params, domain)
	}

	params = append(params, fromMonth, tillMonth, filter.StartDate, filter.EndDate, filter.Source)

	rows, err := (*ch.client).Query(ctx, query, params...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var results []*entities.Dashboard

	for rows.Next() {
		var res entities.Dashboard
		if err := rows.ScanStruct(&res); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		results = append(results, &res)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	return results, nil
}

func (ch clickHouseRepository) TotalMeasurementsCount(ctx context.Context) (string, error) {
	var total uint64
	err := (*ch.client).QueryRow(ctx, "SELECT COALESCE(sum(rows_per_day), 0) FROM base.daily_rows_all").Scan(&total)
	if err != nil {
		return "", fmt.Errorf("failed to execute query: %w", err)
	}

	return strconv.FormatUint(total, 10), nil
}

func (ch clickHouseRepository) MeasurementsCountByDate(ctx context.Context, startDate time.Time, endDate time.Time) (string, error) {
	var total uint64
	query := "SELECT COALESCE(sum(rows_per_day), 0) FROM base.daily_rows_all WHERE date BETWEEN ? AND ?"
	err := (*ch.client).QueryRow(ctx, query, startDate, endDate).Scan(&total)
	if err != nil {
		return "", fmt.Errorf("failed to execute query: %w", err)
	}

	return strconv.FormatUint(total, 10), nil
}

func (ch clickHouseRepository) InterferenceRateByCountry(ctx context.Context, startDate time.Time, endDate time.Time) ([]*entities.InterferenceRateByCountry, error) {
	query := `WITH
    GlobalPriorForPeriod AS (
        SELECT
            sum(total_count) / count(DISTINCT country_name) AS prior_strength_C,
            sum(flag_count) / sum(total_count) AS global_rate
        FROM
            cp.country_interference_counts
        WHERE
            date BETWEEN ? AND ?
		)
	SELECT
		country_name,
		round(100.0 *
			(sum(flag_count) + ((SELECT prior_strength_C FROM GlobalPriorForPeriod) * (SELECT global_rate FROM GlobalPriorForPeriod))) /
			(sum(total_count) + (SELECT prior_strength_C FROM GlobalPriorForPeriod)),
		2) AS unexpected_rate
	FROM
		cp.country_interference_counts
	WHERE
		date BETWEEN ? AND ?
	GROUP BY
		country_name
	ORDER BY
		unexpected_rate DESC`
	rows, err := (*ch.client).Query(ctx, query, startDate, endDate, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to execute interference rate query: %w", err)
	}
	defer rows.Close()

	var results []*entities.InterferenceRateByCountry
	for rows.Next() {
		var rec entities.InterferenceRateByCountry
		if err := rows.ScanStruct(&rec); err != nil {
			return nil, fmt.Errorf("failed to scan interference rate row: %w", err)
		}
		results = append(results, &rec)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over interference rate rows: %w", err)
	}

	return results, nil
}

func (ch clickHouseRepository) Domains(ctx context.Context, startDate time.Time, endDate time.Time, protocol string) ([]string, error) {
	query := "SELECT DISTINCT domain FROM base.domain_lookup WHERE table_name = ? AND date BETWEEN ? AND ? ORDER BY domain"
	rows, err := (*ch.client).Query(ctx, query, protocol, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to execute domains query: %w", err)
	}
	defer rows.Close()

	var results []string
	for rows.Next() {
		var domain string
		if err := rows.Scan(&domain); err != nil {
			return nil, fmt.Errorf("failed to scan domains row: %w", err)
		}
		results = append(results, domain)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over domains rows: %w", err)
	}

	return results, nil
}

func (ch clickHouseRepository) Countries(ctx context.Context, startDate time.Time, endDate time.Time, protocol string) ([]string, error) {
	query := "SELECT DISTINCT country FROM base.country_lookup WHERE table_name = ? AND date BETWEEN ? AND ? ORDER BY country"
	rows, err := (*ch.client).Query(ctx, query, protocol, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to execute countries query: %w", err)
	}
	defer rows.Close()

	var results []string
	for rows.Next() {
		var country string
		if err := rows.Scan(&country); err != nil {
			return nil, fmt.Errorf("failed to scan domains row: %w", err)
		}
		results = append(results, country)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over countries rows: %w", err)
	}

	return results, nil
}
