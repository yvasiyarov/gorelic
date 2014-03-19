package gorelic

import (
	"fmt"
	metrics "github.com/yvasiyarov/go-metrics"
)

const (
	HISTOGRAM_MIN = iota
	HISTOGRAM_MAX
	HISTOGRAM_MEAN
	HISTOGRAM_PERCENTILE
	HISTOGRAM_STD_DEV
	HISTOGRAM_VARIANCE
	NO_HISTOGRAM_FUNCTIONS
)

type GoMetricaDataSource struct {
	metrics.Registry
}

func (ds GoMetricaDataSource) GetGaugeValue(key string) (float64, error) {
	if valueContainer := ds.Get(key); valueContainer == nil {
		return 0, fmt.Errorf("Metrica with name %s is not registered\n", key)
	} else if gauge, ok := valueContainer.(metrics.Gauge); ok {
		return float64(gauge.Value()), nil
	} else {
		return 0, fmt.Errorf("Metrica container has unexpected type: %T\n", valueContainer)
	}
}

func (ds GoMetricaDataSource) GetHistogramValue(key string, statFunction int, percentile float64) (float64, error) {
	if valueContainer := ds.Get(key); valueContainer == nil {
		return 0, fmt.Errorf("Metrica with name %s is not registered\n", key)
	} else if histogram, ok := valueContainer.(metrics.Histogram); ok {
		switch statFunction {
		default:
			return 0, fmt.Errorf("Unsupported stat function for histogram: %s\n", statFunction)
		case HISTOGRAM_MAX:
			return float64(histogram.Max()), nil
		case HISTOGRAM_MIN:
			return float64(histogram.Min()), nil
		case HISTOGRAM_MEAN:
			return float64(histogram.Mean()), nil
		case HISTOGRAM_STD_DEV:
			return float64(histogram.StdDev()), nil
		case HISTOGRAM_VARIANCE:
			return float64(histogram.Variance()), nil
		case HISTOGRAM_PERCENTILE:
			return float64(histogram.Percentile(percentile)), nil
		}
	} else {
		return 0, fmt.Errorf("Metrica container has unexpected type: %T\n", valueContainer)
	}
}

type BaseGoMetrica struct {
	dataSource    GoMetricaDataSource
	basePath      string
	name          string
	units         string
	dataSourceKey string
}

func (metrica *BaseGoMetrica) GetName() string {
	return metrica.basePath + metrica.name
}

func (metrica *BaseGoMetrica) GetUnits() string {
	return metrica.units
}

type GaugeMetrica struct {
	*BaseGoMetrica
}

func (metrica *GaugeMetrica) GetValue() (float64, error) {
	return metrica.dataSource.GetGaugeValue(metrica.dataSourceKey)
}

type GaugeIncMetrica struct {
	*BaseGoMetrica
	previousValue float64
}

func (metrica *GaugeIncMetrica) GetValue() (float64, error) {
	if currentValue, err := metrica.dataSource.GetGaugeValue(metrica.dataSourceKey); err != nil {
		return 0, err
	} else {
		value := currentValue - metrica.previousValue
		metrica.previousValue = currentValue
		return value, nil
	}
}

type HistogramMetrica struct {
	*BaseGoMetrica
	statFunction    int
	percentileValue float64
}

func (metrica *HistogramMetrica) GetValue() (float64, error) {
	return metrica.dataSource.GetHistogramValue(metrica.dataSourceKey, metrica.statFunction, metrica.percentileValue)
}

type BaseTimerMetrica struct {
	dataSource metrics.Timer
	name       string
	units      string
}

func (metrica *BaseTimerMetrica) GetName() string {
	return metrica.name
}

func (metrica *BaseTimerMetrica) GetUnits() string {
	return metrica.units
}

type TimerRate1Metrica struct {
	*BaseTimerMetrica
}

func (metrica *TimerRate1Metrica) GetValue() (float64, error) {
	return metrica.dataSource.Rate1(), nil
}

type TimerRateMeanMetrica struct {
	*BaseTimerMetrica
}

func (metrica *TimerRateMeanMetrica) GetValue() (float64, error) {
	return metrica.dataSource.RateMean(), nil
}

type TimerMeanMetrica struct {
	*BaseTimerMetrica
}

func (metrica *TimerMeanMetrica) GetValue() (float64, error) {
	return metrica.dataSource.Mean(), nil
}

type TimerMinMetrica struct {
	*BaseTimerMetrica
}

func (metrica *TimerMinMetrica) GetValue() (float64, error) {
	return float64(metrica.dataSource.Min()), nil
}

type TimerMaxMetrica struct {
	*BaseTimerMetrica
}

func (metrica *TimerMaxMetrica) GetValue() (float64, error) {
	return float64(metrica.dataSource.Max()), nil
}

type TimerPercentile75Metrica struct {
	*BaseTimerMetrica
}

func (metrica *TimerPercentile75Metrica) GetValue() (float64, error) {
	return metrica.dataSource.Percentile(75), nil
}

type TimerPercentile90Metrica struct {
	*BaseTimerMetrica
}

func (metrica *TimerPercentile90Metrica) GetValue() (float64, error) {
	return metrica.dataSource.Percentile(90), nil
}

type TimerPercentile95Metrica struct {
	*BaseTimerMetrica
}

func (metrica *TimerPercentile95Metrica) GetValue() (float64, error) {
	return metrica.dataSource.Percentile(95), nil
}
