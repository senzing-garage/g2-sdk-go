package senzing

import (
	"context"
	"encoding/json"
)

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

// --- Config -----------------------------------------------------------------

func UnmarshalConfigAddDataSourceResponse(ctx context.Context, jsonString string) (*ConfigAddDataSourceResponse, error) {
	result := &ConfigAddDataSourceResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalConfigListDataSourcesResponse(ctx context.Context, jsonString string) (*ConfigListDataSourcesResponse, error) {
	result := &ConfigListDataSourcesResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalConfigSaveResponse(ctx context.Context, jsonString string) (*ConfigSaveResponse, error) {
	result := &ConfigSaveResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// --- Configmgr --------------------------------------------------------------

func UnmarshalConfigmgrGetConfigResponse(ctx context.Context, jsonString string) (*ConfigmgrGetConfigResponse, error) {
	result := &ConfigmgrGetConfigResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalConfigmgrGetConfigListResponse(ctx context.Context, jsonString string) (*ConfigmgrGetConfigListResponse, error) {
	result := &ConfigmgrGetConfigListResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// --- Diagnostic -------------------------------------------------------------

func UnmarshalDiagnosticCheckDBPerfResponse(ctx context.Context, jsonString string) (*DiagnosticCheckDBPerfResponse, error) {
	result := &DiagnosticCheckDBPerfResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalDiagnosticFetchNextEntityBySizeResponse(ctx context.Context, jsonString string) (*DiagnosticFetchNextEntityBySizeResponse, error) {
	result := &DiagnosticFetchNextEntityBySizeResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalDiagnosticFindEntitiesByFeatureIDsResponse(ctx context.Context, jsonString string) (*DiagnosticFindEntitiesByFeatureIDsResponse, error) {
	result := &DiagnosticFindEntitiesByFeatureIDsResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalDiagnosticGetDataSourceCountsResponse(ctx context.Context, jsonString string) (*DiagnosticGetDataSourceCountsResponse, error) {
	result := &DiagnosticGetDataSourceCountsResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalDiagnosticGetDBInfoResponse(ctx context.Context, jsonString string) (*DiagnosticGetDBInfoResponse, error) {
	result := &DiagnosticGetDBInfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalDiagnosticGetEntityDetailsResponse(ctx context.Context, jsonString string) (*DiagnosticGetEntityDetailsResponse, error) {
	result := &DiagnosticGetEntityDetailsResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalDiagnosticGetEntityListBySizeResponse(ctx context.Context, jsonString string) (*DiagnosticGetEntityListBySizeResponse, error) {
	result := &DiagnosticGetEntityListBySizeResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalDiagnosticGetEntityResumeResponse(ctx context.Context, jsonString string) (*DiagnosticGetEntityResumeResponse, error) {
	result := &DiagnosticGetEntityResumeResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalDiagnosticGetEntitySizeBreakdownResponse(ctx context.Context, jsonString string) (*DiagnosticGetEntitySizeBreakdownResponse, error) {
	result := &DiagnosticGetEntitySizeBreakdownResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalDiagnosticGetFeatureResponse(ctx context.Context, jsonString string) (*DiagnosticGetFeatureResponse, error) {
	result := &DiagnosticGetFeatureResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalDiagnosticGetGenericFeaturesResponse(ctx context.Context, jsonString string) (*DiagnosticGetGenericFeaturesResponse, error) {
	result := &DiagnosticGetGenericFeaturesResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalDiagnosticGetMappingStatisticsResponse(ctx context.Context, jsonString string) (*DiagnosticGetMappingStatisticsResponse, error) {
	result := &DiagnosticGetMappingStatisticsResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalDiagnosticGetRelationshipDetailsResponse(ctx context.Context, jsonString string) (*DiagnosticGetRelationshipDetailsResponse, error) {
	result := &DiagnosticGetRelationshipDetailsResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalDiagnosticGetResolutionStatisticsResponse(ctx context.Context, jsonString string) (*DiagnosticGetResolutionStatisticsResponse, error) {
	result := &DiagnosticGetResolutionStatisticsResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalDiagnosticStreamEntityListBySizeResponse(ctx context.Context, jsonString string) (*DiagnosticStreamEntityListBySizeResponse, error) {
	result := &DiagnosticStreamEntityListBySizeResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// --- Engine -----------------------------------------------------------------

func UnmarshalEngineAddRecordWithInfoResponse(ctx context.Context, jsonString string) (*EngineAddRecordWithInfoResponse, error) {
	result := &EngineAddRecordWithInfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineAddRecordWithInfoWithReturnedRecordIDResponse(ctx context.Context, jsonString string) (*EngineAddRecordWithInfoWithReturnedRecordIDResponse, error) {
	result := &EngineAddRecordWithInfoWithReturnedRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineAddRecordWithReturnedRecordIDResponse(ctx context.Context, jsonString string) (*EngineAddRecordWithReturnedRecordIDResponse, error) {
	result := &EngineAddRecordWithReturnedRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineCheckRecordResponse(ctx context.Context, jsonString string) (*EngineCheckRecordResponse, error) {
	result := &EngineCheckRecordResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineDeleteRecordWithInfoResponse(ctx context.Context, jsonString string) (*EngineDeleteRecordWithInfoResponse, error) {
	result := &EngineDeleteRecordWithInfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineExportConfigAndConfigIDResponse(ctx context.Context, jsonString string) (*EngineExportConfigAndConfigIDResponse, error) {
	result := &EngineExportConfigAndConfigIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineExportConfigResponse(ctx context.Context, jsonString string) (*EngineExportConfigResponse, error) {
	result := &EngineExportConfigResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFetchNextResponse(ctx context.Context, jsonString string) (*EngineFetchNextResponse, error) {
	result := &EngineFetchNextResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindInterestingEntitiesByEntityIDResponse(ctx context.Context, jsonString string) (*EngineFindInterestingEntitiesByEntityIDResponse, error) {
	result := &EngineFindInterestingEntitiesByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindInterestingEntitiesByRecordIDResponse(ctx context.Context, jsonString string) (*EngineFindInterestingEntitiesByRecordIDResponse, error) {
	result := &EngineFindInterestingEntitiesByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindNetworkByEntityID_V2Response(ctx context.Context, jsonString string) (*EngineFindNetworkByEntityID_V2Response, error) {
	result := &EngineFindNetworkByEntityID_V2Response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindNetworkByEntityIDResponse(ctx context.Context, jsonString string) (*EngineFindNetworkByEntityIDResponse, error) {
	result := &EngineFindNetworkByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindNetworkByRecordID_V2Response(ctx context.Context, jsonString string) (*EngineFindNetworkByRecordID_V2Response, error) {
	result := &EngineFindNetworkByRecordID_V2Response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindNetworkByRecordIDResponse(ctx context.Context, jsonString string) (*EngineFindNetworkByRecordIDResponse, error) {
	result := &EngineFindNetworkByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindPathByEntityID_V2Response(ctx context.Context, jsonString string) (*EngineFindPathByEntityID_V2Response, error) {
	result := &EngineFindPathByEntityID_V2Response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindPathByEntityIDResponse(ctx context.Context, jsonString string) (*EngineFindPathByEntityIDResponse, error) {
	result := &EngineFindPathByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindPathByRecordID_V2Response(ctx context.Context, jsonString string) (*EngineFindPathByRecordID_V2Response, error) {
	result := &EngineFindPathByRecordID_V2Response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindPathByRecordIDResponse(ctx context.Context, jsonString string) (*EngineFindPathByRecordIDResponse, error) {
	result := &EngineFindPathByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindPathExcludingByEntityID_V2Response(ctx context.Context, jsonString string) (*EngineFindPathExcludingByEntityID_V2Response, error) {
	result := &EngineFindPathExcludingByEntityID_V2Response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindPathExcludingByEntityIDResponse(ctx context.Context, jsonString string) (*EngineFindPathExcludingByEntityIDResponse, error) {
	result := &EngineFindPathExcludingByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindPathExcludingByRecordID_V2Response(ctx context.Context, jsonString string) (*EngineFindPathExcludingByRecordID_V2Response, error) {
	result := &EngineFindPathExcludingByRecordID_V2Response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindPathExcludingByRecordIDResponse(ctx context.Context, jsonString string) (*EngineFindPathExcludingByRecordIDResponse, error) {
	result := &EngineFindPathExcludingByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindPathIncludingSourceByEntityID_V2Response(ctx context.Context, jsonString string) (*EngineFindPathIncludingSourceByEntityID_V2Response, error) {
	result := &EngineFindPathIncludingSourceByEntityID_V2Response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindPathIncludingSourceByEntityIDResponse(ctx context.Context, jsonString string) (*EngineFindPathIncludingSourceByEntityIDResponse, error) {
	result := &EngineFindPathIncludingSourceByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindPathIncludingSourceByRecordID_V2Response(ctx context.Context, jsonString string) (*EngineFindPathIncludingSourceByRecordID_V2Response, error) {
	result := &EngineFindPathIncludingSourceByRecordID_V2Response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineFindPathIncludingSourceByRecordIDResponse(ctx context.Context, jsonString string) (*EngineFindPathIncludingSourceByRecordIDResponse, error) {
	result := &EngineFindPathIncludingSourceByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineGetEntityByEntityID_V2Response(ctx context.Context, jsonString string) (*EngineGetEntityByEntityID_V2Response, error) {
	result := &EngineGetEntityByEntityID_V2Response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineGetEntityByEntityIDResponse(ctx context.Context, jsonString string) (*EngineGetEntityByEntityIDResponse, error) {
	result := &EngineGetEntityByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineGetEntityByRecordID_V2Response(ctx context.Context, jsonString string) (*EngineGetEntityByRecordID_V2Response, error) {
	result := &EngineGetEntityByRecordID_V2Response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineGetEntityByRecordIDResponse(ctx context.Context, jsonString string) (*EngineGetEntityByRecordIDResponse, error) {
	result := &EngineGetEntityByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineGetRecord_V2Response(ctx context.Context, jsonString string) (*EngineGetRecord_V2Response, error) {
	result := &EngineGetRecord_V2Response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineGetRecordResponse(ctx context.Context, jsonString string) (*EngineGetRecordResponse, error) {
	result := &EngineGetRecordResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineGetRedoRecordResponse(ctx context.Context, jsonString string) (*EngineGetRedoRecordResponse, error) {
	result := &EngineGetRedoRecordResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineGetVirtualEntityByRecordID_V2Response(ctx context.Context, jsonString string) (*EngineGetVirtualEntityByRecordID_V2Response, error) {
	result := &EngineGetVirtualEntityByRecordID_V2Response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineGetVirtualEntityByRecordIDResponse(ctx context.Context, jsonString string) (*EngineGetVirtualEntityByRecordIDResponse, error) {
	result := &EngineGetVirtualEntityByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineHowEntityByEntityID_V2Response(ctx context.Context, jsonString string) (*EngineHowEntityByEntityID_V2Response, error) {
	result := &EngineHowEntityByEntityID_V2Response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineHowEntityByEntityIDResponse(ctx context.Context, jsonString string) (*EngineHowEntityByEntityIDResponse, error) {
	result := &EngineHowEntityByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineProcessRedoRecordResponse(ctx context.Context, jsonString string) (*EngineProcessRedoRecordResponse, error) {
	result := &EngineProcessRedoRecordResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineProcessRedoRecordWithInfoResponse(ctx context.Context, jsonString string) (*EngineProcessRedoRecordWithInfoResponse, error) {
	result := &EngineProcessRedoRecordWithInfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineProcessWithInfoResponse(ctx context.Context, jsonString string) (*EngineProcessWithInfoResponse, error) {
	result := &EngineProcessWithInfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineProcessWithResponseResizeResponse(ctx context.Context, jsonString string) (*EngineProcessWithResponseResizeResponse, error) {
	result := &EngineProcessWithResponseResizeResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineProcessWithResponseResponse(ctx context.Context, jsonString string) (*EngineProcessWithResponseResponse, error) {
	result := &EngineProcessWithResponseResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineReevaluateEntityWithInfoResponse(ctx context.Context, jsonString string) (*EngineReevaluateEntityWithInfoResponse, error) {
	result := &EngineReevaluateEntityWithInfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineReevaluateRecordWithInfoResponse(ctx context.Context, jsonString string) (*EngineReevaluateRecordWithInfoResponse, error) {
	result := &EngineReevaluateRecordWithInfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineReplaceRecordWithInfoResponse(ctx context.Context, jsonString string) (*EngineReplaceRecordWithInfoResponse, error) {
	result := &EngineReplaceRecordWithInfoResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineSearchByAttributes_V2Response(ctx context.Context, jsonString string) (*EngineSearchByAttributes_V2Response, error) {
	result := &EngineSearchByAttributes_V2Response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineSearchByAttributesResponse(ctx context.Context, jsonString string) (*EngineSearchByAttributesResponse, error) {
	result := &EngineSearchByAttributesResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineStatsResponse(ctx context.Context, jsonString string) (*EngineStatsResponse, error) {
	result := &EngineStatsResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineStreamExportJSONEntityReportResponse(ctx context.Context, jsonString string) (*EngineStreamExportJSONEntityReportResponse, error) {
	result := &EngineStreamExportJSONEntityReportResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineWhyEntities_V2Response(ctx context.Context, jsonString string) (*EngineWhyEntities_V2Response, error) {
	result := &EngineWhyEntities_V2Response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineWhyEntitiesResponse(ctx context.Context, jsonString string) (*EngineWhyEntitiesResponse, error) {
	result := &EngineWhyEntitiesResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineWhyEntityByEntityID_V2Response(ctx context.Context, jsonString string) (*EngineWhyEntityByEntityID_V2Response, error) {
	result := &EngineWhyEntityByEntityID_V2Response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineWhyEntityByEntityIDResponse(ctx context.Context, jsonString string) (*EngineWhyEntityByEntityIDResponse, error) {
	result := &EngineWhyEntityByEntityIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineWhyEntityByRecordID_V2Response(ctx context.Context, jsonString string) (*EngineWhyEntityByRecordID_V2Response, error) {
	result := &EngineWhyEntityByRecordID_V2Response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineWhyEntityByRecordIDResponse(ctx context.Context, jsonString string) (*EngineWhyEntityByRecordIDResponse, error) {
	result := &EngineWhyEntityByRecordIDResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineWhyRecords_V2Response(ctx context.Context, jsonString string) (*EngineWhyRecords_V2Response, error) {
	result := &EngineWhyRecords_V2Response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalEngineWhyRecordsResponse(ctx context.Context, jsonString string) (*EngineWhyRecordsResponse, error) {
	result := &EngineWhyRecordsResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

// --- Product ----------------------------------------------------------------

/*
UnmarshalProductVersionResponse...
*/
func UnmarshalProductLicenseResponse(ctx context.Context, jsonString string) (*ProductLicenseResponse, error) {
	result := &ProductLicenseResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalProductValidateLicenseFileResponse(ctx context.Context, jsonString string) (*ProductValidateLicenseFileResponse, error) {
	result := &ProductValidateLicenseFileResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

func UnmarshalProductValidateLicenseStringBase64Response(ctx context.Context, jsonString string) (*ProductValidateLicenseStringBase64Response, error) {
	result := &ProductValidateLicenseStringBase64Response{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}

/*
UnmarshalProductVersionResponse...
*/
func UnmarshalProductVersionResponse(ctx context.Context, jsonString string) (*ProductVersionResponse, error) {
	result := &ProductVersionResponse{}
	err := json.Unmarshal([]byte(jsonString), result)
	return result, err
}