package gorelic

import (
	//	"encoding/json"
	//	"errors"
	//	"fmt"
	//	"io/ioutil"
	//	"net/http"
	//	"net/url"
	"reflect"
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
		Host:         "web-v4.butik.ru", //replace with real host name
		Settings:     NewAgentSettings(),
	}
	return a
}

type AgentSettings struct {
	StartupTimeout       float32           `json:"startup_timeout"`
	AgentRunId           int               `json:"agent_run_id"`
	ProductLevel         int               `json:"product_level"`
	Beacon               string            `json:"beacon"`
	Messages             []SettingsMessage `json:"messages"`
	DataReportPeriod     int               `json:"data_report_period"`
	EncodingKey          string            `json:"encoding_key"`
	EpisodesFile         string            `json:"episodes_file"`
	ApplicationId        int               `json:"application_id"`
	CaptureParams        bool              `json:"capture_params"`
	ProxyPort            int               `json:"proxy_port"`
	IncludeEnviron       []string          `json:"include_environ"`
	BrowserKey           string            `json:"browser_key"`
	ShutdownTimeout      float32           `json:"shutdown_timeout"`
	TrustedAccountIds    []int             `json:"trusted_account_ids"`
	WebTransactionsApdex interface{}       `json:"web_transactions_apdex"`
	Port                 int               `json:"port"`
	AppName              string            `json:"app_name"`
	TransactionNameRules []string          `json:"transaction_name_rules"`
	LogLevel             int               `json:"log_level"`
	ProxyHost            string            `json:"proxy_host"`
	IgnoredParams        []string          `json:"ignored_params"`
	UrlRules             []SettingsUrlRule `json:"url_rules"`
	EpisodesUrl          interface{}       `json:"episodes_url"`
	Enabled              bool              `json:enabled`
	ApdexT               float32           `json:"apdex_t"`
	SSL                  bool              `json:"ssl"`
	Host                 string            `json:"host"`
	MetricNameRules      []string          `json:"metric_name_rules"`
	SamplingRate         int               `json:"sampling_rate"`
	CollectErrors        bool              `json:"collect_errors"`
	CollectTraces        bool              `json:"collect_traces"`
	CrossProcessEnabled  bool              `json:"cross_process.enabled"`
	CaptureEnviron       bool              `json:"capture_environ"`
	CrossProcessId       int               `json:"cross_process_id"`
	LogFile              string            `json:"log_file"`
	ConfigFile           string            `json:"config_file"`
	MonitorMode          bool              `json:"monitor_mode"`

	ErrorCollectorCaptureSource bool     `json:"error_collector.capture_source"`
	ErrorCollectorEnabled       bool     `json:"error_collector.enabled"`
	ErrorCollectorIgnoreErrors  []string `json:"error_collector.ignore_errors"`

	RumLoadEpisodesFile bool `json:"rum.load_episodes_file"`
	RumJsonp            bool `json:"rum.jsonp"`
	RumEnabled          bool `json:"rum.enabled"`

	ThreadProfilerEnabled bool `json:"thread_profiler.enabled"`

	SlowSqlEnabled bool `json:"slow_sql.enabled"`

	TransactionNameLimit        int         `json:"transaction_name.limit"`
	TransactionNameNamingScheme interface{} `json:"transaction_name.naming_scheme"`

	TransactionTracerEnabled              bool    `json:"transaction_tracer.enabled"`
	TransactionTracerFunctionTrace        []int   `json:"transaction_tracer.function_trace"`
	TransactionTracerStackTraceThreshold  int     `json:"transaction_tracer.stack_trace_threshold"`
	TransactionTracerExplainEnabled       bool    `json:"transaction_tracer.explain_enabled"`
	TransactionTracerTopN                 int     `json:"transaction_tracer.top_n"`
	TransactionTracerRecordSql            string  `json:"transaction_tracer.record_sql"`
	TransactionTracerTransactionThreshold int     `json:"transaction_tracer.transaction_threshold"`
	TransactionTracerExplainThreshold     float32 `json:"transaction_tracer.explain_threshold"`

	ConsoleListenerSocket      interface{} `json:"console.listener_socket"`
	ConsoleAllowInterpreterCmd bool        `json:"console.allow_interpreter_cmd"`

	AgentLimitsSqlQueryLengthMaximum      int `json:"agent_limits.sql_query_length_maximum"`
	AgentLimitsTransactionTracesNodes     int `json:"agent_limits.transaction_traces_nodes"`
	AgentLimitsSqlExplainPlans            int `json:"agent_limits.sql_explain_plans"`
	AgentLimitsErrorsPerHarvest           int `json:"agent_limits.errors_per_harvest"`
	AgentLimitsSlowTransactionDryHarvests int `json:"agent_limits.slow_transaction_dry_harvests"`
	AgentLimitsSlowSqlData                int `json:"agent_limits.slow_sql_data"`
	AgentLimitsThreadProfilerNodes        int `json:"agent_limits.thread_profiler_nodes"`
	AgentLimitsMergeStatsMaximum          int `json:"agent_limits.merge_stats_maximum"`
	AgentLimitsSlowSqlStackTrace          int `json:"agent_limits.slow_sql_stack_trace"`
	AgentLimitsSavedTransactions          int `json:"agent_limits.saved_transactions"`
	AgentLimitsErrorsPerTransaction       int `json:"agent_limits.errors_per_transaction"`

	BrowserMonitoringAutoInstrument bool `json:"browser_monitoring.auto_instrument"`

	DebugLogTransactionTracePayload bool     `json:"debug.log_transaction_trace_payload"`
	DebugLogNormalizedMetricData    bool     `json:"debug.log_normalized_metric_data"`
	DebugLocalSettingsOverrides     []string `json:"debug.local_settings_overrides"`
	DebugLogDataCollectorPayloads   bool     `json:"debug.log_data_collector_payloads"`
	DebugLogMalformedJsonData       bool     `json:"debug.log_malformed_json_data"`
	DebugLogNormalizationRules      bool     `json:"debug.log_normalization_rules"`
	DebugIgnoreAllServerSettings    bool     `json:"debug.ignore_all_server_settings"`
	DebugLogRawMetricData           bool     `json:"debug.log_raw_metric_data"`
	DebugLogAgentInitialization     bool     `json:"debug.log_agent_initialization"`
	DebugLogDataCollectorCalls      bool     `json:"debug.log_data_collector_calls"`
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
		UrlRules:                              []SettingsUrlRule{},
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

type SettingsUrlRule struct {
	Ignore          bool   `json:"ignore"`
	Replacement     string `json:"replacement"`
	ReplaceAll      bool   `json:"replace_all"`
	EachSegment     bool   `json:"each_segment"`
	TerminateChain  bool   `json:"terminate_chain"`
	EvalOrder       int    `json:"eval_order"`
	MatchExpression string `json:"match_expression"`
}

type SettingsMessage struct {
	Message string `json:"message"`
	Level   string `json:"level"`
}

func (agent *AgentSettings) ApplyConfigFromServer(serverConfig map[string]interface{}) {
	agentType := reflect.TypeOf(*agent)
	agentValue := reflect.ValueOf(agent)

	for i := 0; i < agentType.NumField(); i++ {
		field := agentType.Field(i)

		if field.Name == "UrlRules" || field.Name == "Messages" {
			continue //UrlRules and Messages don't implemented 
		}

		fieldName := field.Name
		if json := field.Tag.Get("json"); json != "" {
			fieldName = json
		}

		if v, ok := serverConfig[fieldName]; ok {
			fieldValue := agentValue.Field(i)
			newFieldValue := reflect.ValueOf(v)

			if fieldValue.CanSet() && newFieldValue.Type().AssignableTo(field.Type) {
				fieldValue.Set(newFieldValue)
			}
		}
	}
}
