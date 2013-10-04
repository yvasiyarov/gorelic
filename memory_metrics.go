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
	gaugeMetrics := []*BaseGoMetrica{
        //Memory in use metrics
		&BaseGoMetrica{
			name:          "InUse/Total",
			units:         "bytes",
			dataSourceKey: "runtime.MemStats.Alloc",
		},
		&BaseGoMetrica{
			name:          "InUse/Heap",
			units:         "bytes",
			dataSourceKey: "runtime.MemStats.HeapAlloc",
        },
		&BaseGoMetrica{
			name:          "InUse/Stack",
			units:         "bytes",
			dataSourceKey: "runtime.MemStats.StackInuse",
        },
		&BaseGoMetrica{
			name:          "InUse/MSpanInuse",
			units:         "bytes",
			dataSourceKey: "runtime.MemStats.MSpanInuse",
        },
		&BaseGoMetrica{
			name:          "InUse/MCacheInuse",
			units:         "bytes",
			dataSourceKey: "runtime.MemStats.MCacheInuse",
        },
	}
	ds := NewMemoryMetricaDataSource()
	for _, m := range gaugeMetrics {
		m.basePath = "Runtime/Memory/"
		m.dataSource = ds
		component.AddMetrica(&GaugeMetrica{m})
	}

	gaugeIncMetrics := []*BaseGoMetrica{
        //NO operations graph
		&BaseGoMetrica{
			name:          "Operations/NoPointerLookups",
			units:         "lookups",
			dataSourceKey: "runtime.MemStats.Lookups",
		},
		&BaseGoMetrica{
			name:          "Operations/NoMallocs",
			units:         "mallocs",
			dataSourceKey: "runtime.MemStats.Mallocs",
		},
		&BaseGoMetrica{
			name:          "Operations/NoFrees",
			units:         "frees",
			dataSourceKey: "runtime.MemStats.Frees",
        },

        // Sytem memory allocations
		&BaseGoMetrica{
			name:          "SysMem/Total",
			units:         "bytes",
			dataSourceKey: "runtime.MemStats.Sys",
		},
		&BaseGoMetrica{
			name:          "SysMem/Heap",
			units:         "bytes",
			dataSourceKey: "runtime.MemStats.HeapSys",
        },
		&BaseGoMetrica{
			name:          "SysMem/Stack",
			units:         "bytes",
			dataSourceKey: "runtime.MemStats.StackSys",
        },
		&BaseGoMetrica{
			name:          "SysMem/MSpan",
			units:         "bytes",
			dataSourceKey: "runtime.MemStats.MSpanSys",
        },
		&BaseGoMetrica{
			name:          "SysMem/MCache",
			units:         "bytes",
			dataSourceKey: "runtime.MemStats.MCacheSys",
        },
		&BaseGoMetrica{
			name:          "SysMem/BuckHash",
			units:         "bytes",
			dataSourceKey: "runtime.MemStats.BuckHashSys",
        },
    }

	for _, m := range gaugeIncMetrics {
		m.basePath = "Runtime/Memory/"
		m.dataSource = ds
        component.AddMetrica(&GaugeIncMetrica{BaseGoMetrica: m})
	}
}
