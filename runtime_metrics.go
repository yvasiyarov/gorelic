package gorelic

import (
	"fmt"
	"github.com/yvasiyarov/newrelic_platform_go"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const LINUX_SYSTEM_QUERY_INTERVAL = 60

// Number of goroutines metrica
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

// Number of CGO calls metrica
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

//OS specific metrics data source interface
type ISystemMetricaDataSource interface {
	GetValue(key string) (float64, error)
}

// ISystemMetricaDataSource fabrica
func NewSystemMetricaDataSource() ISystemMetricaDataSource {
	var ds ISystemMetricaDataSource
	switch runtime.GOOS {
	default:
		ds = &SystemMetricaDataSource{}
	case "linux":
		ds = &LinuxSystemMetricaDataSource{
			systemData: make(map[string]string),
		}
	}
	return ds
}

//Default implementation of ISystemMetricaDataSource. Just return an error
type SystemMetricaDataSource struct{}

func (ds *SystemMetricaDataSource) GetValue(key string) (float64, error) {
	return 0, fmt.Errorf("This metrica was not implemented yet for %s", runtime.GOOS)
}

// Linux OS implementation of ISystemMetricaDataSource
type LinuxSystemMetricaDataSource struct {
	lastUpdate time.Time
	systemData map[string]string
}

func (ds *LinuxSystemMetricaDataSource) GetValue(key string) (float64, error) {
	if err := ds.checkAndUpdateData(); err != nil {
		return 0, err
	} else if val, ok := ds.systemData[key]; !ok {
		return 0, fmt.Errorf("System data with key %s was not found.", key)
	} else if key == "VmSize" || key == "VmPeak" || key == "VmHWM" || key == "VmRSS" {
		valueParts := strings.Split(val, " ")
		if len(valueParts) != 2 {
			return 0, fmt.Errorf("Invalid format for value %s", key)
		}
		valConverted, err := strconv.ParseFloat(valueParts[0], 64)
		if err != nil {
			return 0, err
		}
		switch valueParts[1] {
		case "kB":
			valConverted *= 1 << 10
		case "mB":
			valConverted *= 1 << 20
		case "gB":
			valConverted *= 1 << 30
		}
		return valConverted, nil
	} else if valConverted, err := strconv.ParseFloat(val, 64); err != nil {
		return valConverted, nil
	} else {
		return valConverted, nil
	}
}
func (ds *LinuxSystemMetricaDataSource) checkAndUpdateData() error {
	startTime := time.Now()
	if startTime.Sub(ds.lastUpdate) > time.Second*LINUX_SYSTEM_QUERY_INTERVAL {
		path := fmt.Sprintf("/proc/%d/status", os.Getpid())
		rawStats, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		lines := strings.Split(string(rawStats), "\n")
		for _, line := range lines {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				k := strings.TrimSpace(parts[0])
				v := strings.TrimSpace(parts[1])

				ds.systemData[k] = v
			}
		}
		ds.lastUpdate = startTime
	}
	return nil
}

// OS specific metrica 
type SystemMetrica struct {
	sourceKey    string
	newrelicName string
	units        string
	dataSource   ISystemMetricaDataSource
}

func (metrica *SystemMetrica) GetName() string {
	return metrica.newrelicName
}
func (metrica *SystemMetrica) GetUnits() string {
	return metrica.units
}
func (metrica *SystemMetrica) GetValue() (float64, error) {
	return metrica.dataSource.GetValue(metrica.sourceKey)
}

func addRuntimeMericsToComponent(component newrelic_platform_go.IComponent) {
	component.AddMetrica(&NOGoroutinesMetrica{})
	component.AddMetrica(&NOCgoCallsMetrica{})

	ds := NewSystemMetricaDataSource()
	metrics := []*SystemMetrica{
		&SystemMetrica{
			sourceKey:    "Threads",
			units:        "Threads",
			newrelicName: "Runtime/System/Threads",
		},
		&SystemMetrica{
			sourceKey:    "FDSize",
			units:        "fd",
			newrelicName: "Runtime/System/FDSize",
		},
		// Peak virtual memory size
		&SystemMetrica{
			sourceKey:    "VmPeak",
			units:        "bytes",
			newrelicName: "Runtime/System/Memory/VmPeakSize",
		},
		//Virtual memory size
		&SystemMetrica{
			sourceKey:    "VmSize",
			units:        "bytes",
			newrelicName: "Runtime/System/Memory/VmCurrent",
		},
		//Peak resident set size
		&SystemMetrica{
			sourceKey:    "VmHWM",
			units:        "bytes",
			newrelicName: "Runtime/System/Memory/RssPeak",
		},
		//Resident set size
		&SystemMetrica{
			sourceKey:    "VmRSS",
			units:        "bytes",
			newrelicName: "Runtime/System/Memory/RssCurrent",
		},
	}
	for _, m := range metrics {
		m.dataSource = ds
		component.AddMetrica(m)
	}
}
