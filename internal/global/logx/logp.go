package logx

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

var (
	SystemLogger  *LogP
	ServiceLogger *LogP
	topicCache    sync.Map
	sourceCache   sync.Map
)

type Option func(p *LogP)

type LogP struct {
	logger  *zap.Logger
	zapConf *zap.Config
	sinks   []Sink //The sinks field represents different output destinations for the logs.
}

func WithZapConfig(config *zap.Config) Option {
	return func(cfg *LogP) {
		cfg.zapConf = config
	}
}

func WithSink(s ...Sink) Option {
	return func(cfg *LogP) {
		cfg.sinks = append(cfg.sinks, s...)
	}
}

func Setup(opts ...Option) *LogP {
	l := &LogP{}
	for _, opt := range opts {
		opt(l)
	}
	if l.zapConf == nil {
		l.zapConf = defaultConfig()
	}
	log, err := l.zapConf.Build()
	if err != nil {
		panic(err)
	}
	l.logger = log
	if len(l.sinks) > 0 {
		for _, sink := range l.sinks {
			sink.Open()
		}
		core := NewCoreX(l.logger, l.sinks)
		l.logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.DPanicLevel), zap.AddCallerSkip(1))
	}
	return l
}

func (l *LogP) Stop() {
	for _, sink := range l.sinks {
		sink.Close()
	}
	l.logger.Sync()
}

func defaultConfig() *zap.Config {
	return &zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Development:      false,
		Sampling:         nil,
		Encoding:         "console",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}
}

func (l *LogP) WithTopic(topic string) *LogP {
	value, ok := topicCache.Load(topic)
	if ok {
		return value.(*LogP)
	}
	newSinkList := make([]Sink, 0, len(l.sinks))
	for _, sink := range l.sinks {
		newSink := sink.WithTopic(topic)
		newSinkList = append(newSinkList, newSink)
	}
	loggerCreateByTopic := Setup(WithSink(newSinkList...))
	loggerCreateByTopic.SetLevel(l.zapConf.Level.Level())
	topicCache.Store(topic, loggerCreateByTopic)
	return loggerCreateByTopic
}

func (l *LogP) WithSource(source string) *LogP {
	value, ok := sourceCache.Load(source)
	if ok {
		return value.(*LogP)
	}
	newSinkList := make([]Sink, 0, len(l.sinks))
	for _, sink := range l.sinks {
		newSink := sink.WithSource(source)
		newSinkList = append(newSinkList, newSink)
	}
	loggerCreateBySource := Setup(WithSink(newSinkList...))
	loggerCreateBySource.SetLevel(l.zapConf.Level.Level())
	sourceCache.Store(source, loggerCreateBySource)
	return loggerCreateBySource
}

func (l *LogP) SetLevel(level zapcore.Level) {
	l.zapConf.Level.SetLevel(level)
}

func (l *LogP) Debugf(s string, a ...any) {
	l.logger.Sugar().Debugf(s, a...)
}

func (l *LogP) Debugv(a any) {
	l.logger.Sugar().Debug(a)
}

func (l *LogP) Debugw(s string, field ...zapcore.Field) {
	l.logger.Debug(s, field...)
}

func (l *LogP) Error(a ...any) {
	l.logger.Sugar().Error(a...)
}

func (l *LogP) Errorf(s string, a ...any) {
	l.logger.Sugar().Errorf(s, a...)
}

func (l *LogP) Errorv(a any) {
	l.logger.Sugar().Error(a)
}

func (l *LogP) Errorw(s string, field ...zapcore.Field) {
	l.logger.Error(s, field...)
}

func (l *LogP) Warnf(s string, a ...any) {
	l.logger.Sugar().Warnf(s, a...)
}

func (l *LogP) Warnv(a any) {
	l.logger.Sugar().Warn(a)
}

func (l *LogP) Warnw(s string, field ...zapcore.Field) {
	l.logger.Warn(s, field...)
}

func (l *LogP) Infof(s string, a ...any) {
	l.logger.Sugar().Infof(s, a...)
}

func (l *LogP) Infov(a any) {
	l.logger.Sugar().Info(a)
}

func (l *LogP) Infow(s string, field ...zapcore.Field) {
	l.logger.Info(s, field...)
}

func (l *LogP) Debug(v ...any) {
	l.logger.Sugar().Debug(v...)
}

func (l *LogP) Info(v ...any) {
	l.logger.Sugar().Info(v...)
}

func (l *LogP) Warn(v ...any) {
	l.logger.Sugar().Warn(v...)
}
