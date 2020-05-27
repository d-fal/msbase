package conf

// Errors errors
type Errors struct {
	Errors ErrorTypes `yaml:"Errors"`
}

// ErrorBlock get input block
type ErrorBlock struct {
	Code         interface{} `yaml:"code"`
	HTTPCode     interface{} `yaml:"HTTPCode"`
	Message      string      `yaml:"message"`
	MessageFarsi string      `yaml:"messageFarsi"`
	Status       string      `yaml:"Status"`
	TraceID      string      `yaml:"trace-id"`
}

// ErrorTypes types of possible errors one could see in the ecosystem
type ErrorTypes struct {
	NoError                     ErrorBlock `yaml:"NoError"`
	ServerError                 ErrorBlock `yaml:"ServerError"`
	InputError                  ErrorBlock `yaml:"InputError"`
	NoRouteToDestination        ErrorBlock `yaml:"NoRouteToDestination"`
	FailedToParseServerResponse ErrorBlock `yaml:"FailedToParseServerResponse"`
	NonStandardMobileNumber     ErrorBlock `yaml:"NonStandardMobileNumber"`
	ExcessiveLengthMobileNumber ErrorBlock `yaml:"ExcessiveLengthMobileNumber"`
	// ErrorPatternNotFound raises this error when the expected error is not found
	ErrorPatternNotFound ErrorBlock `yaml:"ErrorPatternNotFound"`
	// ErrorStatusCodeNotFound raises when the status code is not found
	ErrorStatusCodeNotFound ErrorBlock `yaml:"ErrorStatusCodeNotFound"`
	// ErrorURLNotFound Cannot find url
	ErrorURLNotFound ErrorBlock `yaml:"ErrorURLNotFound"`
	// ErrorCannotReachOutEventHost cannot dial logger server
	ErrorCannotReachOutEventHost ErrorBlock `yaml:"ErrorCannotReachOutEventHost"`
	// ErrorCannotMakeReportRequest cannot prepare http request
	ErrorCannotMakeReportRequest ErrorBlock `yaml:"ErrorCannotMakeReportRequest"`
	// ErrorCannotGetToken error fetch token
	ErrorCannotGetToken ErrorBlock `yaml:"ErrorCannotGetToken"`
	// ErrorCannotInspectToken cannot inspect token
	ErrorCannotInspectToken ErrorBlock `yaml:"ErrorCannotInspectToken"`
	// ActivationFailed cannot update user
	ActivationFailed ErrorBlock `yaml:"ActivationFailed"`
	// ErrorCannotFetchUserData user data not accessible
	ErrorCannotFetchUserData                 ErrorBlock `yaml:"ErrorCannotFetchUserData"`
	ErrorDatabaseCredentials                 ErrorBlock `yaml:"ErrorDatabaseCredentials"`
	ErrorBillServiceIsNotSupported           ErrorBlock `yaml:"ErrorBillServiceIsNotSupported"`
	ErrorBillServicesNotLoaded               ErrorBlock `yaml:"ErrorBillServicesNotLoaded"`
	ErrorCreatingHTTPRequest                 ErrorBlock `yaml:"ErrorCreatingHTTPRequest"`
	ErrorCannotFetchHTTPData                 ErrorBlock `yaml:"ErrorCannotFetchHTTPData"`
	ErrorCannotEmulateData                   ErrorBlock `yaml:"ErrorCannotEmulateData"`
	InputErrorInsufficientData               ErrorBlock `yaml:"InputErrorInsufficientData"`
	FlowNotFound                             ErrorBlock `yaml:"FlowNotFound"`
	ErrorTokenExpired                        ErrorBlock `yaml:"ErrorTokenExpired"`
	ErrorConfigPayloadIsMalformed            ErrorBlock `yaml:"ErrorConfigPayloadIsMalformed"`
	ErrorConnectionTimeout                   ErrorBlock `yaml:"ErrorConnectionTimeout"`
	ErrorCannotLoadImage                     ErrorBlock `yaml:"ErrorCannotLoadImage"`
	ErrorDecodeImage                         ErrorBlock `yaml:"ErrorDecodeImage"`
	ErrorCannotQueryCacheServer              ErrorBlock `yaml:"ErrorCannotQueryCacheServer"`
	ErrorFilePathIsEmpty                     ErrorBlock `yaml:"ErrorFilePathIsEmpty"`
	WarningKeyDoesNotExistInCache            ErrorBlock `yaml:"WarningKeyDoesNotExistInCache"`
	ErrorCannotUnmarshallCacheServer         ErrorBlock `yaml:"ErrorCannotUnmarshallCacheServer"`
	ErrorCannotRetrieveThirdPartyInformation ErrorBlock `yaml:"ErrorCannotRetrieveThirdPartyInformation"`
	ErrorCustom                              ErrorBlock `yaml:"ErrorCustom"`
	ErrorTokenGenerationFailed               ErrorBlock `yaml:"ErrorTokenGenerationFailed:"`
	ErrorAuthorizationFailed                 ErrorBlock `yaml:"ErrorAuthorizationFailed"`
	ErrorNoTokensPresentedInHeader           ErrorBlock `yaml:"ErrorNoTokensPresentedInHeader"`
	UnknownError                             ErrorBlock `yaml:"UnknownError"`
}

// MessageMicroBlock description the provider offers
type MessageMicroBlock struct {
	Info        string `yaml:"Info"`
	InfoFa      string `yaml:"InfoFa"`
	Type        string `yaml:"Type"`
	IsEncrypted bool   `yaml:"IsEncrypted"`
	Action      string `yaml:"Action"`
}

func (e *ErrorBlock) IsOk() bool {
	if (*e == ErrorBlock{} || *e == (GetConfigObject().GetErrorList().NoError)) {
		return true
	}
	return false
}

func (e *ErrorBlock) IsCached() bool {

	if *e == confRcv.GetErrorList().WarningKeyDoesNotExistInCache ||
		*e == confRcv.GetErrorList().ErrorCannotUnmarshallCacheServer ||
		*e == confRcv.GetErrorList().ErrorCannotQueryCacheServer {
		return false
	}
	return true
}
