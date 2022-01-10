package logger

import (
	"fmt"
	"os"

	"github.com/natefinch/lumberjack"

	"github.com/go-kratos/kratos/v2/log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ log.Logger = (*ZapLogger)(nil)

type ZapLogger struct {
	log  *zap.Logger
	Sync func() error
}

func NewZapLogger(encoder zapcore.Encoder, writer zapcore.WriteSyncer, level zap.AtomicLevel, opts ...zap.Option) *ZapLogger {
	core := zapcore.NewCore(encoder, writer, level)
	zapLogger := zap.New(core, opts...)
	return &ZapLogger{log: zapLogger, Sync: zapLogger.Sync}
}

// Log 实现 Kratos 的 Logger 接口
func (z *ZapLogger) Log(level log.Level, kvs ...interface{}) error {
	if len(kvs) == 0 || len(kvs)%2 != 0 {
		z.log.Warn(fmt.Sprint("key and values must appear in pairs: ", kvs))
		return nil
	}

	var data []zap.Field
	for i := 0; i < len(kvs); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(kvs[i]), fmt.Sprint(kvs[i+1])))
	}

	switch level {
	case log.LevelDebug:
		z.log.Debug("", data...)
	case log.LevelInfo:
		z.log.Info("", data...)
	case log.LevelWarn:
		z.log.Warn("", data...)
	case log.LevelError:
		z.log.Error("", data...)
	}
	return nil
}

func (z *ZapLogger) L() *zap.Logger {
	return z.log
}

func (z *ZapLogger) Debug(msg string, fields ...zap.Field) {
	z.log.Debug(msg, fields...)
}

func (z *ZapLogger) Info(msg string, fields ...zap.Field) {
	z.log.Info(msg, fields...)
}

func (z *ZapLogger) Warn(msg string, fields ...zap.Field) {
	z.log.Warn(msg, fields...)
}

func (z *ZapLogger) Error(msg string, fields ...zap.Field) {
	z.log.Error(msg, fields...)
}

func (z *ZapLogger) Fatal(msg string, fields ...zap.Field) {
	z.log.Fatal(msg, fields...)
}

func NewEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:   "ts",
		LevelKey:  "level",
		NameKey:   "logger",
		CallerKey: "caller",
		//MessageKey:     "msg", // 屏蔽msg以兼容kratos的logger，否则日志将出现两个msg
		StacktraceKey:  "stack",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

func NewWriter(name string) (zapcore.WriteSyncer, error) {
	var (
		file *os.File
		err  error
	)
	file, err = os.Create(name)
	if err != nil {
		return nil, err
	}
	return zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(file),
	), nil
}

func NewLumberWriter(name string, devMode bool) zapcore.WriteSyncer {
	logger := &lumberjack.Logger{
		Filename:   name,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}

	var writer []zapcore.WriteSyncer
	writer = append(writer, zapcore.AddSync(logger))
	if devMode {
		writer = append(writer, zapcore.AddSync(os.Stdout))
	}

	return zapcore.NewMultiWriteSyncer(writer...)
}
