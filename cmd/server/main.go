package main

import (
	"apihut-layout/internal/conf"
	logger2 "apihut-layout/internal/logger"
	"flag"
	"os"

	"go.uber.org/zap"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// confPath is the config flag.
	confPath string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&confPath, "conf", "../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
		),
	)
}

func main() {
	// 解析配置文件路径
	flag.Parse()
	// 初始化配置
	c := config.New(
		config.WithSource(
			file.NewSource(confPath),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	// 赋值全局变量
	Name = bc.GetName()
	Version = bc.GetVersion()

	// 初始化日志
	zapLogger := logger2.NewZapLogger(
		logger2.NewEncoder(),
		logger2.NewLumberWriter(bc.GetLog().GetPath(), conf.IsDevMode(bc.Mode)),
		zap.NewAtomicLevelAt(zap.DebugLevel),
		zap.AddCaller(),
	)
	defer func() { _ = zapLogger.Sync }()
	logger := log.With(zapLogger)

	// 初始化App
	app, cleanup, err := initApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// 运行
	if err = app.Run(); err != nil {
		panic(err)
	}
}
