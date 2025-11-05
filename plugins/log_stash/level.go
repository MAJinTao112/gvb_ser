package log_stash

import "encoding/json"

type Level int

const (
	DebugLevel Level = iota + 1
	InfoLevel
	WarnLevel
	ErrorLevel
)

func (s Level) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s Level) String() string {
	var str string
	switch s {
	case DebugLevel:
		str = "debug"
	case InfoLevel:
		str = "info"
	case WarnLevel:
		str = "warn"
	case ErrorLevel:
		str = "error"
	default:
		str = "其他"
	}
	return str

}
