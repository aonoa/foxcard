package data

import (
	"foxcard/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewCardRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	rdb *redis.Client
}

// NewData .
func NewData(conf *conf.Data, logger log.Logger) (*Data, func(), error) {
	// return &Data{}, cleanup, nil

	d := &Data{
		rdb: redis.NewClient(&redis.Options{
			Addr:         conf.Redis.Addr,
			Password:     conf.Redis.Password,
			DB:           int(conf.Redis.Db),
			DialTimeout:  conf.Redis.DialTimeout.AsDuration(),
			WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
			ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
		}),
	}

	cleanup := func() {
		log.Info("message", "closing the data resources")
		if err := d.rdb.Close(); err != nil {
			log.Error(err)
		}
	}

	return d, cleanup, nil
}
