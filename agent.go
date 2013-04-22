package gorelic

import (
//	"encoding/json"
//	"errors"
//	"fmt"
//	"io/ioutil"
//	"net/http"
//	"net/url"
//	"strings"
)

type AgentSettings struct {
}

type environmentAttribute []interface{}
type AgentEnvironment []environmentAttribute

func NewAgentEnvironment() *AgentEnvironment {
    //TODO:  ["Plugin List", []]

    env := &AgentEnvironment{
        environmentAttribute{"Agent Version", "1.10.2.38"},
        environmentAttribute{"Arch", "x86_64"},
        environmentAttribute{"OS", "Linux"},
        environmentAttribute{"OS version", "3.2.0-24-generic"},
        environmentAttribute{"CPU Count", "1"},
        environmentAttribute{"System Memory", "2003.6328125"},
        environmentAttribute{"Python Program Name", "/usr/local/bin/newrelic-admin"},
        environmentAttribute{"Python Executable", "/usr/bin/python"},
        environmentAttribute{"Python Home", ""},
        environmentAttribute{"Python Path", ""},
        environmentAttribute{"Python Prefix", "/usr"},
        environmentAttribute{"Python Exec Prefix", "/usr"},
        environmentAttribute{"Python Version", "2.7.3 (default, Apr 20 2012, 22:39:59) \n[GCC 4.6.3]"},
        environmentAttribute{"Python Platform", "linux2"},
        environmentAttribute{"Python Max Unicode", "1114111"},
        environmentAttribute{"Compiled Extensions", ""},
    }
    return env
}


type Agent struct {
    AppName []string `json:"app_name"`
    Language string `json:"language"`
    Settings *AgentSettings `json:"settings"`
    Pid      int `json:"pid"`
    Environment *AgentEnvironment `json:"environment"`
    Host  string `json:"host"`
    Identifier  string `json:"identifier"`
    AgentVersion string `json:"agent_version"`
}

func NewAgent() *Agent {
    a := &Agent{
        AppName: []string{"Python Agent Test"},
        Language: "python",
        Identifier: "Python Agent Test",
        AgentVersion: "1.10.2.38",
        Environment: NewAgentEnvironment(),
    }
    return a
}

type AgentSettings struct {
    StartupTimeout float `json:"startup_timeout"`
    DebugLogDataCollectorCalls bool `json:"debug.log_data_collector_calls"`
    EncodingKey string `json:"encoding_key"`
    ApplicationId string `json:"application_id"`
    ThreadProfilerEnabled bool `json:"thread_profiler.enabled"` 
    ErrorCollectorCaptureSource bool `json:"error_collector.capture_source"` 
    CaptureParams bool `json:"capture_params"`
    AgentLimitsSqlQueryLengthMaximum int `json:"agent_limits.sql_query_length_maximum"`
    ProxyPort int `json:"proxy_port"`
    IncludeEnviron []string `json:"include_environ"`  
    TransactionNameLimit int `json:"transaction_name.limit"`
    BrowserKey string `json:"browser_key"`
    DebugLogTransactionTracePayload bool `json:"debug.log_transaction_trace_payload"`
    ShutdownTimeout float `json:"shutdown_timeout"`
    TrustedAccountIds []int `json:"trusted_account_ids"`
    WebTransactionsApdex interface{} `json:"web_transactions_apdex"`
    Port int `json:"port"`
    AppName string `json:"app_name"`
    TransactionNameRules []string `json:"transaction_name_rules"`
    AgentLimitsTransactionTracesNodes int `json:"agent_limits.transaction_traces_nodes"`
    TransactionTracerEnabled bool `json:"transaction_tracer.enabled"`
    LogLevel int `json:"log_level"`
    ProxyHost string `json:"proxy_host"`
    IgnoredParams []string `json:"ignored_params"
    AgentLimitsSqlExplainPlans int `json:"agent_limits.sql_explain_plans"`
}

func NewAgentSettings() *AgentSettings {
    s := &AgentSettings{
        StartupTimeout: 0.0,
        DebugLogDataCollectorCalls: true,
        ThreadProfilerEnabled: true,
        ErrorCollectorCaptureSource: false,
        CaptureParams: true,
        AgentLimitsSqlQueryLengthMaximum: 16384,
        ProxyPort: 0,
        IncludeEnviron: []string{"REQUEST_METHOD", "HTTP_USER_AGENT", "HTTP_REFERER", "CONTENT_TYPE", "CONTENT_LENGTH"},
        TransactionNameLimit: 0,
        DebugLogTransactionTracePayload: false,
        ShutdownTimeout: 30.0,
        TrustedAccountIds: []int{},
        WebTransactionsApdex: map[string]string{},
        Port: 0, 
        AppName: "Python Agent Test",
        TransactionNameRules: []string{},
        AgentLimitsTransactionTracesNodes: 2000,
        TransactionTracerEnabled: true,
        LogLevel: 10,
        IgnoredParams: []string{},
        AgentLimitsSqlExplainPlans: 30,
    }
    return s
}


