package logrusfluentd

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/sirupsen/logrus"
)

type Formatter struct {
	service string
}

func prefixMessage(service string, message string) string {
	bytes := bytes.Buffer{}
	bytes.WriteString(strings.ToUpper(service))
	bytes.WriteString(": ")
	bytes.WriteString(message)
	return bytes.String()
}

func severity(level logrus.Level) string {
	switch level {
	case logrus.PanicLevel:
		return "CRITICAL"
	case logrus.FatalLevel:
		return "ALERT"
	default:
		return strings.ToUpper(level.String())
	}
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields, len(entry.Data)+4)

	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			// Otherwise errors are ignored by `encoding/json`
			// https://github.com/Sirupsen/logrus/issues/137
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}

	prefixFieldClashes(data)

	data["message"] = prefixMessage(f.service, entry.Message)
	data["severity"] = severity(entry.Level)
	data["timestamp"] = entry.Time.Unix()
	data["serviceContext"] = map[string]string{
		"service": f.service,
	}

	ser, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return append(ser, '\n'), nil
}

func NewFormatter(service string) *Formatter {
	return &Formatter{service}
}

func prefixFieldClashes(data logrus.Fields) {
	clashes := []string{"time", "msg", "level", "message", "severity", "timestamp", "serviceContext"}
	for _, v := range clashes {
		if iv, ok := data[v]; ok {
			data["fields."+v] = iv
		}
	}
}

// To ensure our Formatter implement the logrus interface
var _ logrus.Formatter = (*Formatter)(nil)
