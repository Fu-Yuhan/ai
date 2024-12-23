package svc

import (
	"gin-skeleton/config"
	"gin-skeleton/svc/logx"
	"gin-skeleton/svc/redisx"
	"gin-skeleton/svc/sqlx"

	_ "github.com/golang-module/carbon/v2"
)

type ServiceContext struct {
	Config config.Config
	Redis  *redisx.Client
	Postgres  sqlx.SqlConn
	CacheX *CacheX
}

func NewServiceContext(c config.Config) *ServiceContext {
	logx.New(c.Log.Level, 2)
	return &ServiceContext{
		Config: c,
		Redis: redisx.New(&redisx.Options{
			Network:  c.Redis.Network,
			Addr:     c.Redis.Addr,
			Username: c.Redis.Username,
			Password: c.Redis.Password,
			DB:       c.Redis.DB,
		}),
		Postgres:  sqlx.NewPostgres(c.Postgres.DataSource),
		CacheX: NewCacheX(c),
	}
}
