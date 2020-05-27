package utils

import (
	"time"

	"msbase/pkg/conf"
	"msbase/pkg/libs"
)

func GetTimeout() time.Duration {
	timeout := libs.RequestTimeout
	if conf.GetConfigObject().GetServerConfig().RequestTimeout != 0 {
		timeout = conf.GetConfigObject().GetServerConfig().RequestTimeout
	}

	return time.Duration(timeout) * time.Second
}
