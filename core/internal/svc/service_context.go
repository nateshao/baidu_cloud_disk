package svc

import (
	"cloud_disk_demo/core/internal/config"
	"cloud_disk_demo/core/internal/middleware"
	"cloud_disk_demo/core/models"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/rest"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	RDB    *redis.Client
	Auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: models.Init(c.Mysql.DataSource),
		RDB:    models.InitRedis(c),
		Auth:   middleware.NewAuthMiddleware().Handle,
	}
}
