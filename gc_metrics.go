package main

import (
	"fmt"
	metrics "github.com/yvasiyarov/go-metrics"
	"github.com/yvasiyarov/newrelic_platform_go"
	"time"
)

//metrics:
//Total durations of all GC calls
//Quantiles of GC call durations 

type GCMetricaDataSource struct {
	*metrics.StandardRegistry
}

func NewGCMetricaDataSource() GCMetricaDataSource {
	r := metrics.NewRegistry()

	metrics.RegisterDebugGCStats(r)
	go metrics.CaptureDebugGCStats(r, time.Duration(GC_POLL_INTERVAL_IN_SECONDS)*time.Second)
	return GCMetricaDataSource{r}
}

func (ds GCMetricaDataSource) GetValue(key string, statFunction string) (float64, error) {
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

type GCMetrica struct {
	dataSource    GCMetricaDataSource
	basePath      string
	name          string
	units         string
	dataSourceKey string
	previousValue float64
	statFunction  string
}

func (metrica *GCMetrica) GetName() string {
	name := metrica.basePath + metrica.name
	if metrica.statFunction != "" {
		name += "/" + metrica.statFunction
	}
	return name
}
func (metrica *GCMetrica) GetUnits() string {
	return metrica.units
}
func (metrica *GCMetrica) GetValue() (float64, error) {
	if currentValue, err := metrica.dataSource.GetValue(metrica.dataSourceKey, metrica.statFunction); err != nil {
		return 0, err
	} else {
		value := currentValue - metrica.previousValue
		metrica.previousValue = currentValue
		return value, nil
	}
}

func addGCMericsToComponent(component newrelic_platform_go.IComponent) {
	metrics := []*GCMetrica{
		&GCMetrica{
			name:          "NumberOfGCCalls",
			units:         "calls",
			dataSourceKey: "debug.GCStats.NumGC",
		},
		&GCMetrica{
			name:          "PauseTotalTime",
			units:         "nanoseconds",
			dataSourceKey: "debug.GCStats.PauseTotal",
		},
		&GCMetrica{
			name:          "GCTime",
			units:         "nanoseconds",
			dataSourceKey: "debug.GCStats.Pause",
			statFunction:  "Max",
		},
		&GCMetrica{
			name:          "GCTime",
			units:         "nanoseconds",
			dataSourceKey: "debug.GCStats.Pause",
			statFunction:  "Min",
		},
		&GCMetrica{
			name:          "GCTime",
			units:         "nanoseconds",
			dataSourceKey: "debug.GCStats.Pause",
			statFunction:  "Mean",
		},
		&GCMetrica{
			name:          "GCTime",
			units:         "nanoseconds",
			dataSourceKey: "debug.GCStats.Pause",
			statFunction:  "Percentile95",
		},
	}

	ds := NewGCMetricaDataSource()
	for _, m := range metrics {
		m.basePath = "Runtime/GC/"
		m.dataSource = ds
		component.AddMetrica(m)
	}
}
