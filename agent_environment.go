package gorelic

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

