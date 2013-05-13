package gorelic

import (
    "time"
)

// {"scope": "WebTransaction/Uri/test", "name": "Function/newrelic.admin:_function3"}
// {"scope": "WebTransaction/Uri/test", "name": "Function/newrelic.admin:_wsgi_application"},
// {"scope": "WebTransaction/Uri/test", "name": "Function/newrelic.admin:_function1"},
// {"scope": "WebTransaction/Uri/test", "name": "Function/newrelic.admin:_function2"},

// {"scope": "", "name": "Supportability/StatsEngine/Calls/record_transaction"},
// {"scope": "", "name": "Supportability/TransactionNode/Calls/time_metrics"},
// {"scope": "", "name": "Supportability/TransactionNode/Calls/slow_sql_nodes"},
// {"scope": "", "name": "Supportability/Transaction/Counts/metric_data"},
//  {"scope": "", "name": "Supportability/TransactionNode/Calls/value_metrics"},
// {"scope": "", "name": "Supportability/TransactionNode/Calls/apdex_metrics"},
// {"scope": "", "name": "Supportability/TransactionNode/Calls/error_details"},

// {"scope": "", "name": "External/allWeb"}
// {"scope": "", "name": "External/localhost/all"},
// {"scope": "", "name": "External/localhost/test/GET"}, 
// {"scope": "WebTransaction/Uri/test", "name": "External/localhost/test/GET"},

// {"scope": "", "name": "Python/WSGI/Input/Time"},
// {"scope": "", "name": "Python/WSGI/Input/Bytes"},
// {"scope": "", "name": "Python/WSGI/Output/Calls/yield"},
// {"scope": "", "name": "Python/WSGI/Input/Calls/readlines"},
// {"scope": "", "name": "Python/WSGI/Input/Calls/readline"}
// {"scope": "", "name": "Python/WSGI/Input/Calls/read"},
// {"scope": "", "name": "Python/WSGI/Output/Calls/write"},
// {"scope": "", "name": "Python/WSGI/Output/Time"},
// {"scope": "", "name": "Python/WSGI/Output/Bytes"},
// {"scope": "WebTransaction/Uri/test", "name": "Python/WSGI/Application"},

// {"scope": "", "name": "Errors/all"},
// {"scope": "", "name": "Errors/allWeb"},
// {"scope": "", "name": "Errors/WebTransaction/Uri/test"},

// {"scope": "", "name": "WebTransaction"},
// {"scope": "", "name": "WebTransaction/Uri/test"},

// {"scope": "", "name": "CPU/User/Utilization"},
// {"scope": "", "name": "HttpDispatcher"},
// {"scope": "", "name": "Memory/Physical"},
// {"scope": "", "name": "Apdex/Uri/test"},
// {"scope": "", "name": "Instance/Reporting"},
// {"scope": "", "name": "Apdex"},
// {"scope": "", "name": "CPU/User Time"}
   

type StatsEngine struct{
    IsInitialized bool

    
}


func NewStatsEngine() *StatsEngine {
    e := &StatsEngine{}
    return e  
}

func (engine *StatsEngine) ResetStats() {
    //TODO: reset accamulated statistics
}
func (engine *StatsEngine) RecordTransaction() {
    if not engine.IsInitialized {
        return
    }
}

type Transaction struct {
    //priority int
    //name
    //errors
    //slow_sql
    //custom_metrics
    //dead bool
    //state int
    BackgroundTask bool
    QueueStart float32
    StartTime time.Time
    EndTime   time.Time
    Stopped bool
    ThreadId int

    Enabled bool
    IgnoreTransaction bool
    SuppressApdex bool
    ResponseCode int
}

func NewTransaction() *Transaction {
    t := &Transaction{}
    return t
}

func (tr *Transaction) Start() {
    if !tr.Enabled {
        return
    }

    tr.StartTime = time.Now()
    //TODO: record CPU time and UserTime

}

func (tr *Transaction) GetTransactionType() string {
    if tr.BackgroundTask {
        return "OtherTransaction"
    } else {
        return "WebTransaction"
    }
}
func (tr *Transaction) GetGroup() string {
    if tr.BackgroundTask {
        return "Go"
    } else {
        return "Uri"
    }
}
func (tr *Transaction) End() {
    if !tr.Enabled {
        return
    }
    //TODO:
    // Bytes read
    // Bytes write
    // calls_read
    // calls_readline
    // calls_readlines
    // calls_write
    // calls_yield
    // read_duration
    // sent_duration
    tr.EndTime = time.Now()
}

func (tr *Transaction) GetDuration() time.Duration {
    duration := time.Duration(0)
    if tr.StartTime && tr.EndTime {
        duration = tr.EndTime.Sub(tr.StartTime) 
    }
    return duration
}

