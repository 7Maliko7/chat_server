package log

type Logger interface {
	Debug() Event
	Info() Event
	Warn() Event
	Error() Event
	Panic() Event
}

type Event interface {
	Str(key, value string) Event
	Int(key string, value int) Event
	Bool(key string, value bool) Event
	Send()
	Msg(value string)
}
