package server

import (
	"github.com/apihutco/apihut-layout/internal/conf"
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewHTTPServer, NewGRPCServer, NewRegister)

func NewRegister(conf *conf.Registry, logger log.Logger) registry.Registrar {
	log := log.NewHelper(logger)

	sc := []constant.ServerConfig{
		*constant.NewServerConfig(conf.Nacos.GetAddress(), conf.Nacos.GetPort()),
	}

	cc := constant.ClientConfig{
		NamespaceId:         conf.Nacos.GetNamespaceId(),
		TimeoutMs:           uint64(conf.Nacos.GetTimeout().GetNanos()),
		NotLoadCacheAtStart: true,
		LogDir:              conf.Nacos.GetLogDir(),
		CacheDir:            conf.Nacos.GetCacheDir(),
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}

	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	return nacos.New(client)
}
