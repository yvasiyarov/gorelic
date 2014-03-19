package main

import (
	"flag"
	"github.com/yvasiyarov/gorelic"
	"log"
	"math/rand"
	"runtime"
	"time"
    "expvar"
    "io"
    "net/http"
)

var newrelicLicense = flag.String("newrelic-license", "", "Newrelic license")

var numCalls = expvar.NewInt("num_calls")

func allocateAndSum(arraySize int) int {
	arr := make([]int, arraySize, arraySize)
	for i, _ := range arr {
		arr[i] = rand.Int()
	}
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)

	result := 0
	for _, v := range arr {
		result += v
	}
	//log.Printf("Array size is: %d, sum is: %d\n", arraySize, result)
	return result
}

func doSomeJob(numRoutines int) {
    for i := 0; i < numRoutines; i++ {
        go allocateAndSum(rand.Intn(1024) * 1024)
    }
    log.Printf("All %d routines started\n", numRoutines)
    time.Sleep(1000 * time.Millisecond)
    runtime.GC()
}

func HelloServer(w http.ResponseWriter, req *http.Request) {

	doSomeJob(5)
    io.WriteString(w, "Did some work")
}

type TWebHandlerFunc func(http.ResponseWriter, *http.Request)
type TWebHandler struct {
    originalHandler http.Handler
    originalHandlerFunc  TWebHandlerFunc
    isFunc bool
}

func NewWebHandlerFunc(h TWebHandlerFunc) *TWebHandler {
    return &TWebHandler{
        isFunc: true,
        originalHandlerFunc: h,
    }
}
func NewWebHandler(h http.Handler) *TWebHandler {
    return &TWebHandler{
        isFunc: false,
        originalHandler: h,
    }
}

func (handler *TWebHandler) BeforeStart() {
}
func (handler *TWebHandler) AfterEnd() {
    numCalls.Add(1)
}

func (handler *TWebHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    handler.BeforeStart()
    defer handler.AfterEnd()

    if handler.isFunc {
        handler.originalHandlerFunc(w, req)
    } else {
        handler.originalHandler.ServeHTTP(w, req)
    }
}

func InstrumentWebHandlerFunc(h TWebHandlerFunc) TWebHandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        proxy := NewWebHandlerFunc(h)
        proxy.ServeHTTP(w, req)
    }
}

func InstrumentWebHandler(h http.Handler) http.Handler {
    return NewWebHandler(h)
}

func main() {
	flag.Parse()
	if *newrelicLicense == "" {
		log.Fatalf("Please, pass a valid newrelic license key.\n Use --help to get more information about available options\n")
	}
	agent := gorelic.NewAgent()
	agent.Verbose = true
	agent.NewrelicLicense = *newrelicLicense
	agent.Run()

    http.HandleFunc("/", InstrumentWebHandlerFunc(HelloServer))
    http.ListenAndServe(":8080", nil)

}
