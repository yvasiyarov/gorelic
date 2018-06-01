package gorelic

import (
	"reflect"
	"sort"
	"testing"

	"github.com/yvasiyarov/newrelic_platform_go"
)

var (
	dummyComponent = newrelic_platform_go.NewPluginComponent(DefaultAgentName, DefaultAgentGuid, false)
)

func TestBeginEndTrace(t *testing.T) {
	tracer := newTracer(dummyComponent)

	trace := tracer.BeginTrace("dummy_trace")

	expectedTraceName := "Trace/dummy_trace"
	if trace.transaction.name != expectedTraceName {
		t.Errorf("Expected the trace name to be %s but got %s instead", expectedTraceName, trace.transaction.name)
	}

	expectedMetricaModelCount := 6
	if len(dummyComponent.MetricaModels) != expectedMetricaModelCount {
		t.Errorf("Expected the number of metrica models to be %d but got %d instead", expectedMetricaModelCount, len(dummyComponent.MetricaModels))
	}

	var metricas []string
	for _, metrica := range dummyComponent.MetricaModels {
		metricas = append(metricas, metrica.GetName())
	}
	sort.Strings(metricas)

	expectedMetricas := []string{"Trace/dummy_trace/max", "Trace/dummy_trace/mean", "Trace/dummy_trace/min", "Trace/dummy_trace/percentile75", "Trace/dummy_trace/percentile90", "Trace/dummy_trace/percentile95"}
	if !reflect.DeepEqual(metricas, expectedMetricas) {
		t.Errorf("Expected metricas to be %v buty got %v instead", metricas, expectedMetricas)
	}

	startTime := trace.transaction.timer.Count()
	trace.EndTrace()
	if trace.transaction.timer.Count() == startTime {
		t.Error("Expected the transaction timer to be incremented")
	}
}

func TestTrace(t *testing.T) {
	tracer := newTracer(dummyComponent)

	traceFuncExecuted := false
	dummyTraceFunc := func() {
		traceFuncExecuted = true
	}
	tracer.Trace("dummy_trace", dummyTraceFunc)

	if !traceFuncExecuted {
		t.Fatal("Trace func was not executed")
	}
}
