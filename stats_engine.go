package gorelic

import (
    "time"
)

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

