package main

import (
	metrics "github.com/yvasiyarov/go-metrics"
	"runtime/debug"
    "time"
    "github.com/yvasiyarov/newrelic_platform_go"
)

//metrics:
//number of GC calls/second
//Total durations of all GC calls
//Quantiles of GC call durations 

type GCMetricaDataSource metrics.Registry 

func NewGCMetricaDataSource() GCMetricaDataSource {
    r := metrics.NewRegistry()

    metrics.RegisterDebugGCStats(r)
    go metrics.CaptureDebugGCStats(r, time.Duration(GC_STAT_CAPTURING_INTERVAL_IN_SECONDS) * time.Second)
    return r;
}

type GCMetrica struct{
    dataSource GCMetricaDataSource
    basePath string 
    name string 
    units string
}

func (metrica *GCMetrica) GetName() string {
	return metrica.basePath + metrica.name
}
func (metrica *GCMetrica) GetUnits() string {
	return metrica.units
}
func (metrica *GCMetrica) GetValue() (float64, error) {
    if val := metrica.dataSource.Get(metrica.name); val == nil {
        return 0, fmt.Errorf("Metrica with name %s is not registered\n")
    } else {
        return float64(val), nil
    }
}

func addGCMericsToComponent(component newrelic_platform_go.IComponent) {
    metrics := []*GCMetrica{
        &GCMetrica{
            name: "NumberOfGCCalls",
            units: "calls"
        }
    }

    ds := NewGCMetricaDataSource();
    for _, m := range metrics {
        m.basePath = "Runtime/GC/"
        m.dataSource = ds
        component.AddMetrica(m)
    }
}
