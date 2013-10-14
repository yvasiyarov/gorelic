package gorelic

import (
	"runtime"
)

type NOGoroutinesMetrica struct{}

func (metrica *NOGoroutinesMetrica) GetName() string {
	return "Runtime/General/NOGoroutines"
}
func (metrica *NOGoroutinesMetrica) GetUnits() string {
	return "goroutines"
}
func (metrica *NOGoroutinesMetrica) GetValue() (float64, error) {
	return float64(runtime.NumGoroutine()), nil
}

type NOCgoCallsMetrica struct {
	lastValue int64
}

func (metrica *NOCgoCallsMetrica) GetName() string {
	return "Runtime/General/NOCgoCalls"
}
func (metrica *NOCgoCallsMetrica) GetUnits() string {
	return "calls"
}
func (metrica *NOCgoCallsMetrica) GetValue() (float64, error) {
	currentValue := runtime.NumCgoCall()
	value := float64(currentValue - metrica.lastValue)
	metrica.lastValue = currentValue

	return value, nil
}
