package model

import (
	"msbase/pkg/libs"
)

// SetLoggingMode which mode steve is going to generate logs
func SetLoggingMode(mode int) {
	switch mode {
	case libs.LoggingModeVerbose:
		loggingMode = libs.LoggingModeVerbose
	case libs.LoggingModeDebug:
		loggingMode = libs.LoggingModeDebug
	case libs.LoggingModeRelease:
		loggingMode = libs.LoggingModeRelease
	}
}

// GetLoggingMode which mode STEVE is working with
func GetLoggingMode() int {
	return loggingMode
}
