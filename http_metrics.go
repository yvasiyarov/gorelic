package gorelic

import (
	metrics "github.com/yvasiyarov/go-metrics"
	"github.com/yvasiyarov/newrelic_platform_go"
	"net/http"
	"time"
)

type THttpHandlerFunc func(http.ResponseWriter, *http.Request)
type THttpHandler struct {
	originalHandler     http.Handler
	originalHandlerFunc THttpHandlerFunc
	isFunc              bool
	timer               metrics.Timer
}

var httpTimer metrics.Timer

func NewHttpHandlerFunc(h THttpHandlerFunc) *THttpHandler {
	return &THttpHandler{
		isFunc:              true,
		originalHandlerFunc: h,
	}
}
func NewHttpHandler(h http.Handler) *THttpHandler {
	return &THttpHandler{
		isFunc:          false,
		originalHandler: h,
	}
}

func (handler *THttpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	startTime := time.Now()
	defer handler.timer.UpdateSince(startTime)

	if handler.isFunc {
		handler.originalHandlerFunc(w, req)
	} else {
		handler.originalHandler.ServeHTTP(w, req)
	}
}

func addHttpMericsToComponent(component newrelic_platform_go.IComponent, timer metrics.Timer) {
	rate1 := &TimerRate1Metrica{
		BaseTimerMetrica: &BaseTimerMetrica{
			name:       "http/throughput/1minute",
			units:      "rps",
			dataSource: timer,
		},
	}
	component.AddMetrica(rate1)

	rateMean := &TimerRateMeanMetrica{
		BaseTimerMetrica: &BaseTimerMetrica{
			name:       "http/throughput/rateMean",
			units:      "rps",
			dataSource: timer,
		},
	}
	component.AddMetrica(rateMean)

	responseTimeMean := &TimerMeanMetrica{
		BaseTimerMetrica: &BaseTimerMetrica{
			name:       "http/responseTime/mean",
			units:      "nanoseconds",
			dataSource: timer,
		},
	}
	component.AddMetrica(responseTimeMean)

	responseTimeMax := &TimerMaxMetrica{
		BaseTimerMetrica: &BaseTimerMetrica{
			name:       "http/responseTime/max",
			units:      "nanoseconds",
			dataSource: timer,
		},
	}
	component.AddMetrica(responseTimeMax)

	responseTimeMin := &TimerMinMetrica{
		BaseTimerMetrica: &BaseTimerMetrica{
			name:       "http/responseTime/min",
			units:      "nanoseconds",
			dataSource: timer,
		},
	}
	component.AddMetrica(responseTimeMin)

	responseTimePercentile75 := &TimerPercentile75Metrica{
		BaseTimerMetrica: &BaseTimerMetrica{
			name:       "http/responseTime/percentile75",
			units:      "nanoseconds",
			dataSource: timer,
		},
	}
	component.AddMetrica(responseTimePercentile75)

	responseTimePercentile90 := &TimerPercentile90Metrica{
		BaseTimerMetrica: &BaseTimerMetrica{
			name:       "http/responseTime/percentile90",
			units:      "nanoseconds",
			dataSource: timer,
		},
	}
	component.AddMetrica(responseTimePercentile90)

	responseTimePercentile95 := &TimerPercentile95Metrica{
		BaseTimerMetrica: &BaseTimerMetrica{
			name:       "http/responseTime/percentile95",
			units:      "nanoseconds",
			dataSource: timer,
		},
	}
	component.AddMetrica(responseTimePercentile95)
}
