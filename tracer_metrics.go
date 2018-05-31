package gorelic

import (
	"sync"
	"time"

	metrics "github.com/yvasiyarov/go-metrics"
	"github.com/yvasiyarov/newrelic_platform_go"
)

type Tracer struct {
	sync.RWMutex
	metrics   map[string]*TraceTransaction
	component newrelic_platform_go.IComponent
}

func newTracer(component newrelic_platform_go.IComponent) *Tracer {
	return &Tracer{metrics: make(map[string]*TraceTransaction), component: component}
}

func (t *Tracer) Trace(name string, traceFunc func()) {
	trace := t.BeginTrace(name)
	defer trace.EndTrace()
	traceFunc()
}

func (t *Tracer) BeginTrace(name string) *Trace {
	tracerName := "Trace/" + name

	t.RLock()
	if m, ok := t.metrics[tracerName]; ok {
		return &Trace{m, time.Now()}
	}
	t.RUnlock()

	trans := TraceTransaction{name: tracerName, timer: metrics.NewTimer()}

	t.Lock()
	t.metrics[tracerName] = &trans
	t.Unlock()

	trans.addMetricsToComponent(t.component)

	return &Trace{transaction: &trans, startTime: time.Now()}
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
	tracerMean := &timerMeanMetrica{
		baseTimerMetrica: &baseTimerMetrica{
			name:       transaction.name + "/mean",
			units:      "ms",
			dataSource: transaction.timer,
		},
	}
	component.AddMetrica(tracerMean)

	tracerMax := &timerMaxMetrica{
		baseTimerMetrica: &baseTimerMetrica{
			name:       transaction.name + "/max",
			units:      "ms",
			dataSource: transaction.timer,
		},
	}
	component.AddMetrica(tracerMax)

	tracerMin := &timerMinMetrica{
		baseTimerMetrica: &baseTimerMetrica{
			name:       transaction.name + "/min",
			units:      "ms",
			dataSource: transaction.timer,
		},
	}
	component.AddMetrica(tracerMin)

	tracer75 := &timerPercentile75Metrica{
		baseTimerMetrica: &baseTimerMetrica{
			name:       transaction.name + "/percentile75",
			units:      "ms",
			dataSource: transaction.timer,
		},
	}
	component.AddMetrica(tracer75)

	tracer90 := &timerPercentile90Metrica{
		baseTimerMetrica: &baseTimerMetrica{
			name:       transaction.name + "/percentile90",
			units:      "ms",
			dataSource: transaction.timer,
		},
	}
	component.AddMetrica(tracer90)

	tracer95 := &timerPercentile95Metrica{
		baseTimerMetrica: &baseTimerMetrica{
			name:       transaction.name + "/percentile95",
			units:      "ms",
			dataSource: transaction.timer,
		},
	}
	component.AddMetrica(tracer95)
}
