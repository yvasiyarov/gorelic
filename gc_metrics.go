package main

import (
	metrics "github.com/yvasiyarov/go-metrics"
	"github.com/yvasiyarov/newrelic_platform_go"
	"time"
)

func NewGCMetricaDataSource() GoMetricaDataSource {
	r := metrics.NewRegistry()

	metrics.RegisterDebugGCStats(r)
	go metrics.CaptureDebugGCStats(r, time.Duration(GC_POLL_INTERVAL_IN_SECONDS)*time.Second)
	return GoMetricaDataSource{r}
}


func addGCMericsToComponent(component newrelic_platform_go.IComponent) {
	metrics := []*BaseGoMetrica{
		&BaseGoMetrica{
			name:          "NumberOfGCCalls",
			units:         "calls",
			dataSourceKey: "debug.GCStats.NumGC",
		},
		&BaseGoMetrica{
			name:          "PauseTotalTime",
			units:         "nanoseconds",
			dataSourceKey: "debug.GCStats.PauseTotal",
		},
	}

	ds := NewGCMetricaDataSource()
	for _, m := range metrics {
		m.basePath = "Runtime/GC/"
		m.dataSource = ds
		component.AddMetrica(&GaugeMetrica{m})
	}

	histogramMetrics := []*HistogramMetrica{
		&HistogramMetrica{
			statFunction:  HISTOGRAM_MAX,
            BaseGoMetrica: &BaseGoMetrica{name: "Max"},
		},
		&HistogramMetrica{
			statFunction:  HISTOGRAM_MIN,
            BaseGoMetrica: &BaseGoMetrica{name: "Min"},
		},
		&HistogramMetrica{
			statFunction:  HISTOGRAM_MEAN,
            BaseGoMetrica: &BaseGoMetrica{name: "Mean"},
		},
		&HistogramMetrica{
			statFunction:  HISTOGRAM_PERCENTILE,
            percentileValue: 0.95,
            BaseGoMetrica: &BaseGoMetrica{name: "Percentile95"},
		},
	}
	for _, m := range histogramMetrics {
        m.BaseGoMetrica.units = "nanoseconds"
        m.BaseGoMetrica.dataSourceKey = "debug.GCStats.Pause"
        m.BaseGoMetrica.basePath = "Runtime/GC/GCTime/"
        m.BaseGoMetrica.dataSource =  ds

		component.AddMetrica(m)
	}
}
