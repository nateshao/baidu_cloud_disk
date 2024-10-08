package svc

import (
	"baidu_cloud_disk/core/internal/config"
	"baidu_cloud_disk/core/internal/middleware"
	"baidu_cloud_disk/core/models"
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
