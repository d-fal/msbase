package middleware

import (
	"go.uber.org/zap"
)

type MiddlewaresSet struct {
	// middleware sets
	CORSMiddleware   string
	TokenValidator   string
	Authenticator    string
	LoggerMiddleware string
}

var (
	middlewaresSet MiddlewaresSet
	// authServer     cconf.AuthServer
	zaplogger *zap.Logger
)
