package main

import (
	"flag"
	"os"

	"github.com/go-kratos/kratos/v2/registry"

	zapLogger "github.com/apihutco/apihut-layout/pkg/zap_logger"

	"github.com/apihutco/apihut-layout/internal/conf"
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

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server, rr registry.Registrar) *kratos.App {
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
		kratos.Registrar(rr),
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

	var rc conf.Registry
	if err := c.Scan(&rc); err != nil {
		panic(err)
	}

	// 赋值全局变量
	Name = bc.GetName()
	Version = bc.GetVersion()

	// 初始化日志
	logger := zapLogger.NewZapLogger(
		zapLogger.NewEncoder(),
		zapLogger.NewLumberWriter(bc.GetLog().GetPath(), conf.IsDevMode(bc.Mode)),
		zap.NewAtomicLevelAt(zap.DebugLevel),
		zap.AddCaller(),
	)
	defer func() { _ = logger.Sync }()

	// 初始化App
	app, cleanup, err := initApp(bc.Server, bc.Data, &rc, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// 运行
	if err = app.Run(); err != nil {
		panic(err)
	}
}
