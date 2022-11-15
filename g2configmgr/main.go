/*
Package G2configmgr is a Go wrapper over Senzing's G2Configmgr C binding.

To use G2configmgr, the LD_LIBRARY_PATH environment variable must include
a path to Senzing's libraries.  Example:

	export LD_LIBRARY_PATH=/opt/senzing/g2/lib
*/
package g2configmgr

import (
	"context"

	"github.com/senzing/go-logging/logger"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type G2configmgr interface {
	AddConfig(ctx context.Context, configStr string, configComments string) (int64, error)
	ClearLastException(ctx context.Context) error
	Destroy(ctx context.Context) error
	GetConfig(ctx context.Context, configID int64) (string, error)
	GetConfigList(ctx context.Context) (string, error)
	GetDefaultConfigID(ctx context.Context) (int64, error)
	GetLastException(ctx context.Context) (string, error)
	GetLastExceptionCode(ctx context.Context) (int, error)
	Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error
	ReplaceDefaultConfigID(ctx context.Context, oldConfigID int64, newConfigID int64) error
	SetDefaultConfigID(ctx context.Context, configID int64) error
	SetLogLevel(ctx context.Context, logLevel logger.Level) error
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const ProductId = 6002

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var IdMessages = map[int]string{
	1:    "Enter AddConfig(%s, %s).",
	2:    "Exit  AddConfig(%s, %s) returned (%d, %v).",
	3:    "Enter ClearLastException().",
	4:    "Exit  ClearLastException() returned (%v).",
	5:    "Enter Destroy().",
	6:    "Exit  Destroy() returned (%v).",
	7:    "Enter GetConfig(%d).",
	8:    "Exit  GetConfig(%d) returned (%s, %v).",
	9:    "Enter GetConfigList().",
	10:   "Exit  GetConfigList() returned (%s, %v).",
	11:   "Enter GetDefaultConfigID().",
	12:   "Exit  GetDefaultConfigID() returned (%d, %v).",
	13:   "Enter GetLastException().",
	14:   "Exit  GetLastException() returned (%s, %v).",
	15:   "Enter GetLastExceptionCode().",
	16:   "Exit  GetLastExceptionCode() returned (%d, %v).",
	17:   "Enter Init(%s, %s, %d).",
	18:   "Exit  Init(%s, %s, %d) returned (%v).",
	19:   "Enter ReplaceDefaultConfigID(%d, %d).",
	20:   "Exit  ReplaceDefaultConfigID(%d, %d) returned (%v).",
	21:   "Enter SetDefaultConfigID(%d).",
	22:   "Exit  SetDefaultConfigID(%d) returned (%v).",
	23:   "Enter SetLogLevel(%v).",
	24:   "Exit  SetLogLevel(%v) returned (%v).",
	4001: "Call to G2ConfigMgr_addConfig(%s, %s) failed. Return code: %d",
	4002: "Call to G2ConfigMgr_destroy() failed. Return code: %d",
	4003: "Call to G2ConfigMgr_getConfig(%d) failed. Return code: %d",
	4004: "Call to G2ConfigMgr_getConfigList() failed. Return code: %d",
	4005: "Call to G2ConfigMgr_getDefaultConfigID() failed. Return code: %d",
	4006: "Call to G2ConfigMgr_getLastException() failed. Return code: %d",
	4007: "Call to G2ConfigMgr_init(%s, %s, %d) failed. Return code: %d",
	4008: "Call to G2ConfigMgr_replaceDefaultConfigID(%d, %d) failed. Return code: %d",
	4009: "Call to G2ConfigMgr_setDefaultConfigID(%d) failed. Return code: %d",
}

var IdRanges = map[int]string{
	0000: logger.LevelTraceName,
	1000: logger.LevelDebugName,
	2000: logger.LevelInfoName,
	3000: logger.LevelWarnName,
	4000: logger.LevelErrorName,
	5000: logger.LevelFatalName,
	6000: logger.LevelPanicName,
}

var IdRangesLogLevel = map[int]logger.Level{
	0000: logger.LevelTrace,
	1000: logger.LevelDebug,
	2000: logger.LevelInfo,
	3000: logger.LevelWarn,
	4000: logger.LevelError,
	5000: logger.LevelFatal,
	6000: logger.LevelPanic,
}

var IdStatuses = map[int]string{
	4001: logger.LevelErrorName,
	4002: logger.LevelErrorName,
	4003: logger.LevelErrorName,
	4004: logger.LevelErrorName,
	4005: logger.LevelErrorName,
	4006: logger.LevelErrorName,
	4007: logger.LevelErrorName,
	4008: logger.LevelErrorName,
	4009: logger.LevelErrorName,
}
