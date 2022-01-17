package zapLogger

import (
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
)

func TestLog(t *testing.T) {
	writer, err := NewWriter("./test.log")
	if err != nil {
		t.Error(err)
		return
	}
	logger := NewZapLogger(NewEncoder(), writer, zap.NewAtomicLevelAt(zap.DebugLevel))

	log.NewHelper(logger).Error("sss")
}

func TestNewLumberWriter(t *testing.T) {
	logger := NewZapLogger(
		NewEncoder(),
		NewLumberWriter("./test.log", true),
		zap.NewAtomicLevelAt(zap.DebugLevel),
		//zap.AddCaller(),
		zap.AddStacktrace(zap.FatalLevel),
	)

	logger.Debug("debug")
	logger.L().Error("named.error")
	//logger.Info("info")
	//logger.Warn("warn")
	//logger.Error("error")
	//logger.Fatal("fatal")

	zlog := log.NewHelper(logger)
	zlog.Debug("zlog.debug")
	zlog.Info("zlog.info")
	zlog.Warn("zlog.warn")
	zlog.Error("zlog.error")
	//zlog.Fatal("zlog.fatal")
}
