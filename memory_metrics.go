package main

import (
	"fmt"
	metrics "github.com/yvasiyarov/go-metrics"
	"github.com/yvasiyarov/newrelic_platform_go"
	"time"
)

type MemoryMetricaDataSource struct {
	*metrics.StandardRegistry
}

func NewMemoryMetricaDataSource() MemoryMetricaDataSource {
	r := metrics.NewRegistry()

	metrics.RegisterRuntimeMemStats(r)
	metrics.CaptureRuntimeMemStatsOnce(r)
	go metrics.CaptureRuntimeMemStats(r, time.Duration(MEMORY_ALLOCATOR_POLL_INTERVAL_IN_SECONDS)*time.Second)
	return MemoryMetricaDataSource{r}
}

func (ds MemoryMetricaDataSource) GetValue(key string, statFunction string) (float64, error) {
	if valueContainer := ds.Get(key); valueContainer == nil {
		return 0, fmt.Errorf("Metrica with name %s is not registered\n", key)
	} else {
		switch valueContainer := valueContainer.(type) {
		default:
			return 0, fmt.Errorf("Metrica container has unexpected type: %T\n", valueContainer)
		case *metrics.StandardGauge:
			return float64(valueContainer.Value()), nil
		case *metrics.StandardHistogram:
			switch statFunction {
			default:
				return 0, fmt.Errorf("Unsupported stat function for histogram: %s\n", statFunction)
			case "Max":
				return float64(valueContainer.Max()), nil
			case "Min":
				return float64(valueContainer.Min()), nil
			case "Mean":
				return float64(valueContainer.Mean()), nil
			case "StdDev":
				return float64(valueContainer.StdDev()), nil
			case "Percentile95":
				return float64(valueContainer.Percentile(0.95)), nil
			}
		}
	}
}

type MemoryMetrica struct {
	dataSource    MemoryMetricaDataSource
	basePath      string
	name          string
	units         string
	dataSourceKey string
	statFunction  string
}

func (metrica *MemoryMetrica) GetName() string {
	name := metrica.basePath + metrica.name
	if metrica.statFunction != "" {
		name += "/" + metrica.statFunction
	}
	return name
}
func (metrica *MemoryMetrica) GetUnits() string {
	return metrica.units
}
func (metrica *MemoryMetrica) GetValue() (float64, error) {
	return metrica.dataSource.GetValue(metrica.dataSourceKey, metrica.statFunction)
}

type MemoryIncMetrica struct {
	*MemoryMetrica
	previousValue float64
}

func (metrica *MemoryIncMetrica) GetValue() (float64, error) {
	if currentValue, err := metrica.dataSource.GetValue(metrica.dataSourceKey, metrica.statFunction); err != nil {
		return 0, err
	} else {
		value := currentValue - metrica.previousValue
		metrica.previousValue = currentValue
		return value, nil
	}
}

func addMemoryMericsToComponent(component newrelic_platform_go.IComponent) {
	metrics := []*MemoryMetrica{
		&MemoryMetrica{
			name:          "MemoryAllocated",
			units:         "bytes",
			dataSourceKey: "runtime.MemStats.Alloc",
		},
	}

	ds := NewMemoryMetricaDataSource()
	for _, m := range metrics {
		m.basePath = "Runtime/Memory/"
		m.dataSource = ds
		component.AddMetrica(m)
	}
}
