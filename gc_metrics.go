package main

import (
	"fmt"
	metrics "github.com/yvasiyarov/go-metrics"
	"github.com/yvasiyarov/newrelic_platform_go"
	"time"
)

//metrics:
//number of GC calls/second
//Total durations of all GC calls
//Quantiles of GC call durations 

type GCMetricaDataSource metrics.Registry

func NewGCMetricaDataSource() GCMetricaDataSource {
	r := metrics.NewRegistry()

	metrics.RegisterDebugGCStats(r)
	go metrics.CaptureDebugGCStats(r, time.Duration(GC_STAT_CAPTURING_INTERVAL_IN_SECONDS)*time.Second)
	return r
}

type GCMetrica struct {
	dataSource    GCMetricaDataSource
	basePath      string
	name          string
	units         string
	dataSourceKey string
}

func (metrica *GCMetrica) GetName() string {
	return metrica.basePath + metrica.name
}
func (metrica *GCMetrica) GetUnits() string {
	return metrica.units
}
func (metrica *GCMetrica) GetValue() (float64, error) {
	if valueContainer := metrica.dataSource.Get(metrica.dataSourceKey); valueContainer == nil {
		return 0, fmt.Errorf("Metrica with name %s is not registered\n", metrica.name)
	} else {
		switch valueContainer := valueContainer.(type) {
		default:
			return 0, fmt.Errorf("Metrica container has unexpected type: %T\n", valueContainer)
		case *metrics.StandardGauge:
			return float64(valueContainer.Value()), nil
		}
	}
}

func addGCMericsToComponent(component newrelic_platform_go.IComponent) {
	metrics := []*GCMetrica{
		&GCMetrica{
			name:          "NumberOfGCCalls",
			units:         "calls",
			dataSourceKey: "debug.GCStats.NumGC",
		},
	}

	ds := NewGCMetricaDataSource()
	for _, m := range metrics {
		m.basePath = "Runtime/GC/"
		m.dataSource = ds
		component.AddMetrica(m)
	}
}
