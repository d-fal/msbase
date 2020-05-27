package loader

import (
	"fmt"
	"log"
	"msbase/internal/microservice"
	"msbase/internal/routing"
	"msbase/pkg/cache"
	"msbase/pkg/conf"
	"msbase/pkg/identity"
	"msbase/pkg/logger"

	"go.uber.org/zap"
)

var (
	zapLogger *zap.Logger
)

func init() {
	logger.SetLogPath(conf.GetConfigObject().GetServerConfig().LogFile)

}
func LoadPackages() {

	/* REDIS CONNECTION STARTS HERE */
	cache.PrepareRedisPool()
	conn := cache.GetRedisConnection()
	defer (*conn).Close()

	if err := cache.Ping(); err != nil {
		zapLogger.Sugar().Fatal("Redis connection is not ready", err)
	}
	/* REDIS CONNECTION ends HERE */
	microservice.LoadActiveMicroservices()
	router, err := routing.Router()

	if err != nil {
		log.Fatal(err)
	}
	zapLogger = logger.GetZapLogger(identity.GetSignature().App.ID)
	zapLogger.Sugar().Info("App started")
	router.Run(fmt.Sprintf(":%s", conf.GetConfigObject().GetServerConfig().HTTPPort))
}
