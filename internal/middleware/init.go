package middleware

import (
	"msbase/internal/microservice/auth/cconf"
)

func init() {
	middlewaresSet = MiddlewaresSet{"cors", "token_validator", "authenticator", "logger"}
	authServer = cconf.GetAuthServer()

}
