package libs

const (
	// URLSeparator separates url
	URLSeparator = "/"

	// DirectorySeparator separates directories
	DirectorySeparator = "/"
	// MaximumTolerableErrorCount maximum number of errors in which we would take action if reported errors exceed it
	MaximumTolerableErrorCount = 3
	// ReportReceiverDatabase send error to database
	ReportReceiverDatabase = "database"
	// ReportReceiverURL the url to which we report our errors
	ReportReceiverURL = "url"
	// ReportReceiverAll all othe methods to send report are used
	ReportReceiverAll = "all"
	// BlockDataLabel block data label , it can be changed in config
	BlockDataLabel = "value"
	// BlockTagLabel block tag label
	BlockTagLabel = "key"

	// RequestTimeout the timeout needed to perform an action
	RequestTimeout = 30

	// StatusSuccess success code by the error block
	StatusSuccess = "success"

	// StatusFailed failed status
	StatusFailed = "failed"

	// StatusError error status
	StatusError = "error"

	// AppVersion states the app version
	AppVersion = "1.0.0"

	// PingUpateInterval duration in seconds in which we want to update the db ping
	PingUpateInterval = 30

	//ServiceTypeSOAP soap service
	ServiceTypeSOAP = "soap"

	//ServiceTypeJSON json service
	ServiceTypeJSON = "json"

	//ServiceQueryString query string
	ServiceQueryString = "querystring"

	//StatusKey the key that will be represented by web services
	StatusKey = "Status"
	//InputParameterAlias alias for input
	InputParameterAlias = "input"

	// Keycloak auth server
	KeycloakServer = "Keycloak"
)

const (
	LoggingModeVerbose = iota
	LoggingModeDebug
	LoggingModeRelease
)
