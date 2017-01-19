package gorelic

import (
	metrics "github.com/yvasiyarov/go-metrics"
	"github.com/yvasiyarov/newrelic_platform_go"
	"time"
)

type Tracer struct {
	metrics   map[string]*TraceTransaction
	component newrelic_platform_go.IComponent
}

func newTracer(component newrelic_platform_go.IComponent) *Tracer {
	return &Tracer{make(map[string]*TraceTransaction), component}
}

func (t *Tracer) Trace(name string, traceFunc func()) {
	trace := t.BeginTrace(name)
	defer trace.EndTrace()
	traceFunc()
}

func (t *Tracer) BeginTrace(name string) *Trace {
	tracerName := "Trace/" + name
	m := t.metrics[tracerName]
	if m == nil {
		t.metrics[tracerName] = &TraceTransaction{tracerName, metrics.NewTimer()}
		m = t.metrics[tracerName]
		m.addMetricsToComponent(t.component)
	}
	return &Trace{m, time.Now()}
}

type Trace struct {
	transaction *TraceTransaction
	startTime   time.Time
}

func (t *Trace) EndTrace() {
	t.transaction.timer.UpdateSince(t.startTime)
}

type TraceTransaction struct {
	name  string
	timer metrics.Timer
}

func (transaction *TraceTransaction) addMetricsToComponent(component newrelic_platform_go.IComponent) {

	rate1 := &timerRate1Metrica{
		baseTimerMetrica: &baseTimerMetrica{
			name:       transaction.name + "/throughput/1minute",
			units:      "rps",
			dataSource: transaction.timer,
		},
	}
	component.AddMetrica(rate1)

	rateMean := &timerRateMeanMetrica{
		baseTimerMetrica: &baseTimerMetrica{
			name:       transaction.name + "/throughput/rateMean",
			units:      "rps",
			dataSource: transaction.timer,
		},
	}
	component.AddMetrica(rateMean)

	tracerMean := &timerMeanMetrica{
		baseTimerMetrica: &baseTimerMetrica{
			name:       transaction.name + "/responseTime/mean",
			units:      "ms",
			dataSource: transaction.timer,
		},
	}
	component.AddMetrica(tracerMean)

	tracerMax := &timerMaxMetrica{
		baseTimerMetrica: &baseTimerMetrica{
			name:       transaction.name + "/responseTime/max",
			units:      "ms",
			dataSource: transaction.timer,
		},
	}
	component.AddMetrica(tracerMax)

	tracerMin := &timerMinMetrica{
		baseTimerMetrica: &baseTimerMetrica{
			name:       transaction.name + "/responseTime/min",
			units:      "ms",
			dataSource: transaction.timer,
		},
	}
	component.AddMetrica(tracerMin)

	tracer75 := &timerPercentile75Metrica{
		baseTimerMetrica: &baseTimerMetrica{
			name:       transaction.name + "/responseTime/percentile75",
			units:      "ms",
			dataSource: transaction.timer,
		},
	}
	component.AddMetrica(tracer75)

	tracer90 := &timerPercentile90Metrica{
		baseTimerMetrica: &baseTimerMetrica{
			name:       transaction.name + "/responseTime/percentile90",
			units:      "ms",
			dataSource: transaction.timer,
		},
	}
	component.AddMetrica(tracer90)

	tracer95 := &timerPercentile95Metrica{
		baseTimerMetrica: &baseTimerMetrica{
			name:       transaction.name + "/responseTime/percentile95",
			units:      "ms",
			dataSource: transaction.timer,
		},
	}
	component.AddMetrica(tracer95)

	tracer99 := &timerPercentile99Metrica{
		baseTimerMetrica: &baseTimerMetrica{
			name:       transaction.name + "/responseTime/percentile99",
			units:      "ms",
			dataSource: transaction.timer,
		},
	}
	component.AddMetrica(tracer99)
}
