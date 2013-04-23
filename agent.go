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
	AppName      []string          `json:"app_name"`
	Language     string            `json:"language"`
	Settings     *AgentSettings    `json:"settings"`
	Pid          int               `json:"pid"`
	Environment  *AgentEnvironment `json:"environment"`
	Host         string            `json:"host"`
	Identifier   string            `json:"identifier"`
	AgentVersion string            `json:"agent_version"`
}

func NewAgent() *Agent {
	a := &Agent{
		AppName:      []string{"Python Agent Test"},
		Language:     "python",
		Identifier:   "Python Agent Test",
		AgentVersion: "1.10.2.38",
		Environment:  NewAgentEnvironment(),
	}
	return a
}

type AgentSettings struct {
	StartupTimeout                        float32     `json:"startup_timeout"`
	DebugLogDataCollectorCalls            bool        `json:"debug.log_data_collector_calls"`
	EncodingKey                           string      `json:"encoding_key"`
	ApplicationId                         string      `json:"application_id"`
	ThreadProfilerEnabled                 bool        `json:"thread_profiler.enabled"`
	ErrorCollectorCaptureSource           bool        `json:"error_collector.capture_source"`
	CaptureParams                         bool        `json:"capture_params"`
	AgentLimitsSqlQueryLengthMaximum      int         `json:"agent_limits.sql_query_length_maximum"`
	ProxyPort                             int         `json:"proxy_port"`
	IncludeEnviron                        []string    `json:"include_environ"`
	TransactionNameLimit                  int         `json:"transaction_name.limit"`
	BrowserKey                            string      `json:"browser_key"`
	DebugLogTransactionTracePayload       bool        `json:"debug.log_transaction_trace_payload"`
	ShutdownTimeout                       float32     `json:"shutdown_timeout"`
	TrustedAccountIds                     []int       `json:"trusted_account_ids"`
	WebTransactionsApdex                  interface{} `json:"web_transactions_apdex"`
	Port                                  int         `json:"port"`
	AppName                               string      `json:"app_name"`
	TransactionNameRules                  []string    `json:"transaction_name_rules"`
	AgentLimitsTransactionTracesNodes     int         `json:"agent_limits.transaction_traces_nodes"`
	TransactionTracerEnabled              bool        `json:"transaction_tracer.enabled"`
	LogLevel                              int         `json:"log_level"`
	ProxyHost                             string      `json:"proxy_host"`
	IgnoredParams                         []string    `json:"ignored_params"`
	AgentLimitsSqlExplainPlans            int         `json:"agent_limits.sql_explain_plans"`
	ErrorCollectorEnabled                 bool        `json:"error_collector.enabled"`
	TransactionTracerFunctionTrace        []int       `json:"transaction_tracer.function_trace"`
	RumLoadEpisodesFile                   bool        `json:"rum.load_episodes_file"`
	AgentLimitsErrorsPerHarvest           int         `json:"agent_limits.errors_per_harvest"`
	TransactionTracerStackTraceThreshold  int         `json:"transaction_tracer.stack_trace_threshold"`
	AgentLimitsSlowTransactionDryHarvests int         `json:"agent_limits.slow_transaction_dry_harvests"`
	TransactionNameNamingScheme           interface{} `json:"transaction_name.naming_scheme"`
	UrlRules                              []string    `json:"url_rules"`
	RumJsonp                              bool        `json:"rum.jsonp"`
	ErrorCollectorIgnoreErrors            []string    `json:"error_collector.ignore_errors"`
	RumEnabled                            bool        `json:"rum.enabled"`
	EpisodesUrl                           interface{} `json:"episodes_url"`
	DebugLogNormalizedMetricData          bool        `json:"debug.log_normalized_metric_data"`
	TransactionTracerExplainEnabled       bool        `json:"transaction_tracer.explain_enabled"`
	TransactionTracerTopN                 int         `json:"transaction_tracer.top_n"`
	ConsoleListenerSocket                 interface{} `json:"console.listener_socket"`
	AgentLimitsSlowSqlData                int         `json:"agent_limits.slow_sql_data"`
	Enabled                               bool        `json:enabled`
	DebugLocalSettingsOverrides           []string    `json:"debug.local_settings_overrides"`
	DebugLogDataCollectorPayloads         bool        `json:"debug.log_data_collector_payloads"`
	ApdexT                                float32     `json:"apdex_t"`
	AgentLimitsThreadProfilerNodes        int         `json:"agent_limits.thread_profiler_nodes"`
	SSL                                   bool        `json:"ssl"`
	Host                                  string      `json:"host"`
	MetricNameRules                       []string    `json:"metric_name_rules"`
	TransactionTracerRecordSql            string      `json:"transaction_tracer.record_sql"`
	TransactionTracerTransactionThreshold int         `json:"transaction_tracer.transaction_threshold"`
	SamplingRate                          int         `json:"sampling_rate"`
	CollectErrors                         bool        `json:"collect_errors"`
	AgentLimitsMergeStatsMaximum          int         `json:"agent_limits.merge_stats_maximum"`
	DebugLogMalformedJsonData             bool        `json:"debug.log_malformed_json_data"`
	TransactionTracerExplainThreshold     float32     `json:"transaction_tracer.explain_threshold"`
	ConsoleAllowInterpreterCmd            bool        `json:"console.allow_interpreter_cmd"`
	DebugIgnoreAllServerSettings          bool        `json:"debug.ignore_all_server_settings"`
	AgentLimitsSavedTransactions          int         `json:"agent_limits.saved_transactions"`
	CollectTraces                         bool        `json:"collect_traces"`
	CrossProcessEnabled                   bool        `json:"cross_process.enabled"`
	SlowSqlEnabled                        bool        `json:"slow_sql.enabled"`
	AgentLimitsSlowSqlStackTrace          int         `json:"agent_limits.slow_sql_stack_trace"`
	DebugLogNormalizationRules            bool        `json:"debug.log_normalization_rules"`
	AgentLimitsErrorsPerTransaction       int         `json:"agent_limits.errors_per_transaction"`
	CaptureEnviron                        bool        `json:"capture_environ"`
	DebugLogRawMetricData                 bool        `json:"debug.log_raw_metric_data"`
	CrossProcessId                        int         `json:"cross_process_id"`
	DebugLogAgentInitialization           bool        `json:"debug.log_agent_initialization"`
	LogFile                               string      `json:"log_file"`
	ConfigFile                            string      `json:"config_file"`
	BrowserMonitoringAutoInstrument       bool        `json:"browser_monitoring.auto_instrument"`
	MonitorMode                           bool        `json:"monitor_mode"`
}

func NewAgentSettings() *AgentSettings {
	s := &AgentSettings{
		StartupTimeout:                        0.0,
		DebugLogDataCollectorCalls:            true,
		ThreadProfilerEnabled:                 true,
		ErrorCollectorCaptureSource:           false,
		CaptureParams:                         true,
		AgentLimitsSqlQueryLengthMaximum:      16384,
		ProxyPort:                             0,
		IncludeEnviron:                        []string{"REQUEST_METHOD", "HTTP_USER_AGENT", "HTTP_REFERER", "CONTENT_TYPE", "CONTENT_LENGTH"},
		TransactionNameLimit:                  0,
		DebugLogTransactionTracePayload:       false,
		ShutdownTimeout:                       30.0,
		TrustedAccountIds:                     []int{},
		WebTransactionsApdex:                  map[string]string{},
		Port:                                  0,
		AppName:                               "Python Agent Test",
		TransactionNameRules:                  []string{},
		AgentLimitsTransactionTracesNodes:     2000,
		TransactionTracerEnabled:              true,
		LogLevel:                              10,
		IgnoredParams:                         []string{},
		AgentLimitsSqlExplainPlans:            30,
		ErrorCollectorEnabled:                 true,
		TransactionTracerFunctionTrace:        []int{},
		RumLoadEpisodesFile:                   true,
		AgentLimitsErrorsPerHarvest:           20,
		AgentLimitsSlowTransactionDryHarvests: 5,
		TransactionNameNamingScheme:           nil,
		UrlRules:                              []string{},
		RumJsonp:                              true,
		ErrorCollectorIgnoreErrors:            []string{},
		RumEnabled:                            true,
		DebugLogNormalizedMetricData:          false,
		TransactionTracerExplainEnabled:       true,
		TransactionTracerTopN:                 20,
		AgentLimitsSlowSqlData:                10,
		Enabled:                               true,
		DebugLocalSettingsOverrides:           []string{},
		DebugLogDataCollectorPayloads:         false,
		ApdexT:                                0.5,
		AgentLimitsThreadProfilerNodes:        20000,
		SSL:                                   false,
		Host:                                  START_COLLECTOR_URL,
		MetricNameRules:                       []string{},
		TransactionTracerRecordSql:            "obfuscated",
		TransactionTracerTransactionThreshold: 0,
		SamplingRate:                          0,
		CollectErrors:                         true,
		DebugLogNormalizationRules:            false,
		AgentLimitsErrorsPerTransaction:       5,
		DebugLogRawMetricData:                 false,
		LogFile:                               "/tmp/python-agent-test.log",
		DebugLogAgentInitialization:           false,
		ConfigFile:                            "newrelic.ini",
		BrowserMonitoringAutoInstrument:       true,
		MonitorMode:                           true,
	}
	return s
}
