package gomolgelf

import (
	"fmt"
	"time"

	"github.com/aphistic/golf"
	"github.com/aphistic/gomol"
)

type GelfLoggerConfig struct {
	Hostname string
	Port     int
}

func NewGelfLoggerConfig() *GelfLoggerConfig {
	return &GelfLoggerConfig{
		Hostname: "",
		Port:     -1,
	}
}

type gelfClient interface {
	Close() error
}

type gelfClientLogger interface {
	Dbgm(map[string]interface{}, string, ...interface{}) error
	Infom(map[string]interface{}, string, ...interface{}) error
	Warnm(map[string]interface{}, string, ...interface{}) error
	Errm(map[string]interface{}, string, ...interface{}) error
	Emergm(map[string]interface{}, string, ...interface{}) error
}

type GelfLogger struct {
	base          *gomol.Base
	client        gelfClient
	logger        gelfClientLogger
	isInitialized bool
	config        *GelfLoggerConfig
}

func NewGelfLogger(config *GelfLoggerConfig) (*GelfLogger, error) {
	l := &GelfLogger{
		config: config,
	}
	return l, nil
}

func (l *GelfLogger) getLogger() gelfClientLogger {
	return l.logger
}

func (l *GelfLogger) getAttrs(m map[string]interface{}) map[string]interface{} {
	attrs := make(map[string]interface{}, 0)
	for key, val := range l.base.BaseAttrs {
		attrs[key] = val
	}
	if m != nil {
		for key, val := range m {
			attrs[key] = fmt.Sprintf("%v", val)
		}
	}

	return attrs
}

func (l *GelfLogger) SetBase(base *gomol.Base) {
	l.base = base
}

func (l *GelfLogger) IsInitialized() bool {
	return l.isInitialized
}

func (l *GelfLogger) InitLogger() error {
	c, err := golf.NewClient()
	if err != nil {
		return err
	}
	err = c.Dial(fmt.Sprintf("udp://%v:%v", l.config.Hostname, l.config.Port))
	if err != nil {
		return err
	}
	nl, err := c.NewLogger()
	if err != nil {
		return err
	}

	l.client = c
	l.logger = nl
	l.isInitialized = true

	return nil
}

func (l *GelfLogger) ShutdownLogger() error {
	err := l.client.Close()
	if err != nil {
		return err
	}

	l.isInitialized = false
	return nil
}

func (l *GelfLogger) Logm(timestamp time.Time, level gomol.LogLevel, m map[string]interface{}, msg string) error {
	attrs := l.getAttrs(m)
	switch level {
	case gomol.LevelDebug:
		l.getLogger().Dbgm(attrs, msg)
	case gomol.LevelInfo:
		l.getLogger().Infom(attrs, msg)
	case gomol.LevelWarning:
		l.getLogger().Warnm(attrs, msg)
	case gomol.LevelError:
		l.getLogger().Errm(attrs, msg)
	case gomol.LevelFatal:
		l.getLogger().Emergm(attrs, msg)

	}
	return nil
}
