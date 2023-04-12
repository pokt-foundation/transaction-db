package postgresdriver

import (
	"context"
	"time"

	"github.com/lib/pq"
	"github.com/pokt-foundation/transaction-db/types"
)

const insertServiceRecords = `
INSERT INTO service_record (
	session_key,
	request_id,
	portal_region_name,
	latency,
	tickets,
	result,
	available,
	successes,
	failures,
	p90_success_latency,
	median_success_latency,
	weighted_success_latency,
	success_rate,
	created_at,
	updated_at
	)
	SELECT * FROM unnest(
	$1::char(44)[],
	$2::varchar[],
	$3::varchar[],
	$4::float[],
	$5::integer[],
	$6::varchar[],
	$7::boolean[],
	$8::integer[],
	$9::integer[],
	$10::float[],
	$11::float[],
	$12::float[],
	$13::float[],
	$14::date[],
	$15::date[]
	) AS t(
	session_key,
	request_id,
	portal_region_name,
	latency,
	tickets,
	result,
	available,
	successes,
	failures,
	p90_success_latency,
	median_success_latency,
	weighted_success_latency,
	success_rate,
	created_at,
	updated_at
	)`

func (d *PostgresDriver) WriteServiceRecord(ctx context.Context, serviceRecord types.ServiceRecord) error {
	now := time.Now()

	return d.InsertServiceRecord(ctx, InsertServiceRecordParams{
		SessionKey:             serviceRecord.SessionKey,
		RequestID:              serviceRecord.RequestID,
		PortalRegionName:       serviceRecord.PortalRegionName,
		Latency:                serviceRecord.Latency,
		Tickets:                int32(serviceRecord.Tickets),
		Result:                 serviceRecord.Result,
		Available:              serviceRecord.Available,
		Successes:              int32(serviceRecord.Successes),
		Failures:               int32(serviceRecord.Failures),
		P90SuccessLatency:      serviceRecord.P90SuccessLatency,
		MedianSuccessLatency:   serviceRecord.MedianSuccessLatency,
		WeightedSuccessLatency: serviceRecord.WeightedSuccessLatency,
		SuccessRate:            serviceRecord.SuccessRate,
		CreatedAt:              now,
		UpdatedAt:              now,
	})
}

func (d *PostgresDriver) WriteServiceRecords(ctx context.Context, serviceRecords []types.ServiceRecord) error {
	now := time.Now()

	var (
		sessionKeys              []string
		requestIDs               []string
		portalRegionNames        []string
		latencies                []float64
		tickets                  []int32
		results                  []string
		availables               []bool
		successes                []int32
		failures                 []int32
		p90SuccessLatencies      []float64
		medianSuccessLatencies   []float64
		weightedSuccessLatencies []float64
		successRates             []float64
		createdTimes             []time.Time
		updatedTimes             []time.Time
	)

	for _, serviceRecord := range serviceRecords {
		sessionKeys = append(sessionKeys, serviceRecord.SessionKey)
		requestIDs = append(requestIDs, serviceRecord.RequestID)
		portalRegionNames = append(portalRegionNames, serviceRecord.PortalRegionName)
		latencies = append(latencies, serviceRecord.Latency)
		tickets = append(tickets, int32(serviceRecord.Tickets))
		results = append(results, serviceRecord.Result)
		availables = append(availables, serviceRecord.Available)
		successes = append(successes, int32(serviceRecord.Successes))
		failures = append(failures, int32(serviceRecord.Failures))
		p90SuccessLatencies = append(p90SuccessLatencies, serviceRecord.P90SuccessLatency)
		medianSuccessLatencies = append(medianSuccessLatencies, serviceRecord.MedianSuccessLatency)
		weightedSuccessLatencies = append(weightedSuccessLatencies, serviceRecord.WeightedSuccessLatency)
		successRates = append(successRates, serviceRecord.SuccessRate)
		createdTimes = append(createdTimes, now)
		updatedTimes = append(updatedTimes, now)
	}

	_, err := d.db.Exec(insertServiceRecords, pq.StringArray(sessionKeys),
		pq.StringArray(requestIDs),
		pq.StringArray(portalRegionNames),
		pq.Float64Array(latencies),
		pq.Int32Array(tickets),
		pq.StringArray(results),
		pq.BoolArray(availables),
		pq.Int32Array(successes),
		pq.Int32Array(failures),
		pq.Float64Array(p90SuccessLatencies),
		pq.Float64Array(medianSuccessLatencies),
		pq.Float64Array(weightedSuccessLatencies),
		pq.Float64Array(successRates),
		pq.Array(createdTimes),
		pq.Array(updatedTimes))
	if err != nil {
		return err
	}

	return nil
}

func (d *PostgresDriver) ReadServiceRecord(ctx context.Context, serviceRecordID int) (types.ServiceRecord, error) {
	serviceRecord, err := d.SelectServiceRecord(ctx, int64(serviceRecordID))
	if err != nil {
		return types.ServiceRecord{}, err
	}

	return types.ServiceRecord{
		ServiceRecordID:        int(serviceRecord.ID),
		SessionKey:             serviceRecord.SessionKey,
		RequestID:              serviceRecord.RequestID,
		PortalRegionName:       serviceRecord.PortalRegionName,
		Latency:                serviceRecord.Latency,
		Tickets:                int(serviceRecord.Tickets),
		Result:                 serviceRecord.Result,
		Available:              serviceRecord.Available,
		Successes:              int(serviceRecord.Successes),
		Failures:               int(serviceRecord.Failures),
		P90SuccessLatency:      serviceRecord.P90SuccessLatency,
		MedianSuccessLatency:   serviceRecord.MedianSuccessLatency,
		WeightedSuccessLatency: serviceRecord.WeightedSuccessLatency,
		SuccessRate:            serviceRecord.SuccessRate,
		CreatedAt:              serviceRecord.CreatedAt,
		UpdatedAt:              serviceRecord.UpdatedAt,
	}, nil
}
