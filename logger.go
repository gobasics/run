package run

import (
	"io"
	"os"

	"gobasics.dev/log"
)

type Logger struct {
	Close   func() error
	Fatal   log.Logger
	Error   log.Logger
	Warning log.Logger
	Info    log.Logger
	V       func(log.Level) log.Logger
}

func (l Logger) Init(config Config) error {
	return nil
}

var DefaultLogger = func() Logger {
	var l Logger
	l.Close = func() error { return nil }
	var w io.Writer
	path, ok := os.LookupEnv("GOBASICS_LOG_FILE")
	if !ok {
		w = os.Stdout
	} else {
		f, err := os.Create(path)
		if err != nil {
			w = os.Stdout
		} else {
			w = f
			l.Close = f.Close
		}
	}
	v := log.New(log.WithWriter(w))
	l.Fatal = v(0)
	l.Error = v(1)
	l.Warning = v(2)
	l.Info = v(3)
	l.V = v
	return l
}()
