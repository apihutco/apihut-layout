package data

import (
	"apihut-layout/internal/conf"
	"apihut-layout/internal/data/ent"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	_ "github.com/mattn/go-sqlite3"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewDB, NewData, NewGreeterRepo)

// Data .
type Data struct {
	db *ent.Client
}

func NewDB(conf *conf.Data, logger log.Logger) *ent.Client {

	log := log.NewHelper(logger)

	client, err := ent.Open(conf.Database.GetDriver(), conf.Database.GetSource())
	if err != nil {
		log.Fatal(err)
	}
	if err = client.Schema.Create(context.Background()); err != nil {
		log.Fatal(err)
	}

	return client
}

// NewData .
func NewData(client *ent.Client, logger log.Logger) (*Data, func(), error) {

	log := log.NewHelper(logger)

	d := &Data{db: client}

	cleanup := func() {
		log.Info("closing the data resources")
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}
	return d, cleanup, nil
}
