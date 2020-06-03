package middleware

func init() {
	middlewaresSet = MiddlewaresSet{"cors", "token_validator", "authenticator", "logger"}
	// authServer = cconf.GetAuthServer()

}
