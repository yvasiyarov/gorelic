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
