package szproduct

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Log message prefix.
const Prefix = "szproduct."

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// Message templates for szproduct implementations.
var IdMessages = map[int]string{
	1:    "Enter " + Prefix + "ClearLastException().",
	2:    "Exit  " + Prefix + "ClearLastException() returned (%v).",
	3:    "Enter " + Prefix + "Destroy().",
	4:    "Exit  " + Prefix + "Destroy() returned (%v).",
	5:    "Enter " + Prefix + "GetLastException().",
	6:    "Exit  " + Prefix + "GetLastException() returned (%s, %v).",
	7:    "Enter " + Prefix + "GetLastExceptionCode().",
	8:    "Exit  " + Prefix + "GetLastExceptionCode() returned (%d, %v).",
	9:    "Enter " + Prefix + "Initialize(%s, %s, %d).",
	10:   "Exit  " + Prefix + "Initialize(%s, %s, %d) returned (%v).",
	11:   "Enter " + Prefix + "GetLicense().",
	12:   "Exit  " + Prefix + "GetLicense() returned (%s, %v).",
	13:   "Enter " + Prefix + "SetLogLevel(%v).",
	14:   "Exit  " + Prefix + "SetLogLevel(%v) returned (%v).",
	19:   "Enter " + Prefix + "GetVersion().",
	20:   "Exit  " + Prefix + "GetVersion() returned (%s, %v).",
	21:   "Enter " + Prefix + "RegisterObserver(%s).",
	22:   "Exit  " + Prefix + "RegisterObserver(%s) returned (%v).",
	23:   "Enter " + Prefix + "UnregisterObserver(%s).",
	24:   "Exit  " + Prefix + "UnregisterObserver(%s) returned (%v).",
	25:   "Enter " + Prefix + "GetSdkId().",
	26:   "Exit  " + Prefix + "GetSdkId() returned (%s).",
	4001: Prefix + "G2Product_destroy() failed. Return code: %d",
	4002: Prefix + "G2Product_getLastException() failed. Return code: %d",
	4003: Prefix + "G2Product_init(%s, %s, %d) failed. Return code: %d",
	5901: "During test setup, call to messagelogger.NewSenzingApiLogger() failed.",
	5902: "During test setup, call to szengineconfigurationjson.BuildSimpleSystemConfigurationJsonViaMap() failed.",
	5903: "During test setup, call to szengine.Initialize() failed.",
	5904: "During test setup, call to szdiagnostic.PurgeRepository() failed.",
	5905: "During test setup, call to szengine.Destroy() failed.",
	5906: "During test setup, call to szconfig.Initialize() failed.",
	5907: "During test setup, call to szconfig.Create() failed.",
	5908: "During test setup, call to szconfig.AddDataSource() failed.",
	5909: "During test setup, call to szconfig.GetJsonString() failed.",
	5910: "During test setup, call to szconfig.Close() failed.",
	5911: "During test setup, call to szconfig.Destroy() failed.",
	5912: "During test setup, call to szconfigmgr.Initialize() failed.",
	5913: "During test setup, call to szconfigmgr.AddConfig() failed.",
	5914: "During test setup, call to szconfigmgr.SetDefaultConfigId() failed.",
	5915: "During test setup, call to szconfigmgr.Destroy() failed.",
	5916: "During test setup, call to szengine.Initialize() failed.",
	5917: "During test setup, call to szengine.AddRecord() failed.",
	5918: "During test setup, call to szengine.Destroy() failed.",
	5920: "During test setup, call to setupSenzingConfig() failed.",
	5921: "During test setup, call to setupPurgeRepository() failed.",
	5922: "During test setup, call to setupAddRecords() failed.",
	5931: "During test setup, call to szengine.Initialize() failed.",
	5932: "During test setup, call to szdiagnostic.PurgeRepository() failed.",
	5933: "During test setup, call to szengine.Destroy() failed.",
	8001: Prefix + "Destroy",
	8002: Prefix + "Initialize",
	8003: Prefix + "GetLicense",
	8006: Prefix + "GetVersion",
	8007: Prefix + "GetSdkId",
	8008: Prefix + "RegisterObserver",
	8009: Prefix + "SetLogLevel",
	8010: Prefix + "UnregisterObserver",
}

// Status strings for specific szproduct messages.
var IdStatuses = map[int]string{}