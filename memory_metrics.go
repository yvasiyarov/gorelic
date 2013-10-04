package main

import (
	metrics "github.com/yvasiyarov/go-metrics"
	"github.com/yvasiyarov/newrelic_platform_go"
	"time"
)


func NewMemoryMetricaDataSource() GoMetricaDataSource {
	r := metrics.NewRegistry()

	metrics.RegisterRuntimeMemStats(r)
	metrics.CaptureRuntimeMemStatsOnce(r)
	go metrics.CaptureRuntimeMemStats(r, time.Duration(MEMORY_ALLOCATOR_POLL_INTERVAL_IN_SECONDS)*time.Second)
	return GoMetricaDataSource{r}
}

func addMemoryMericsToComponent(component newrelic_platform_go.IComponent) {
	metrics := []*BaseGoMetrica{
		&BaseGoMetrica{
			name:          "MemoryAllocated",
			units:         "bytes",
			dataSourceKey: "runtime.MemStats.Alloc",
		},
	}

	ds := NewMemoryMetricaDataSource()
	for _, m := range metrics {
		m.basePath = "Runtime/Memory/"
		m.dataSource = ds
		component.AddMetrica(&GaugeMetrica{m})
	}
}
