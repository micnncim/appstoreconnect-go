/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// AppClipsApiService AppClipsApi service
type AppClipsApiService service

type ApiAppClipsAppClipAdvancedExperiencesGetToManyRelatedRequest struct {
	ctx                                          context.Context
	ApiService                                   *AppClipsApiService
	id                                           string
	filterAction                                 *[]string
	filterPlaceStatus                            *[]string
	filterStatus                                 *[]string
	fieldsAppClipAdvancedExperiences             *[]string
	fieldsAppClips                               *[]string
	fieldsAppClipAdvancedExperienceImages        *[]string
	fieldsAppClipAdvancedExperienceLocalizations *[]string
	limit                                        *int32
	limitLocalizations                           *int32
	include                                      *[]string
}

// filter by attribute &#39;action&#39;
func (r ApiAppClipsAppClipAdvancedExperiencesGetToManyRelatedRequest) FilterAction(filterAction []string) ApiAppClipsAppClipAdvancedExperiencesGetToManyRelatedRequest {
	r.filterAction = &filterAction
	return r
}

// filter by attribute &#39;placeStatus&#39;
func (r ApiAppClipsAppClipAdvancedExperiencesGetToManyRelatedRequest) FilterPlaceStatus(filterPlaceStatus []string) ApiAppClipsAppClipAdvancedExperiencesGetToManyRelatedRequest {
	r.filterPlaceStatus = &filterPlaceStatus
	return r
}

// filter by attribute &#39;status&#39;
func (r ApiAppClipsAppClipAdvancedExperiencesGetToManyRelatedRequest) FilterStatus(filterStatus []string) ApiAppClipsAppClipAdvancedExperiencesGetToManyRelatedRequest {
	r.filterStatus = &filterStatus
	return r
}

// the fields to include for returned resources of type appClipAdvancedExperiences
func (r ApiAppClipsAppClipAdvancedExperiencesGetToManyRelatedRequest) FieldsAppClipAdvancedExperiences(fieldsAppClipAdvancedExperiences []string) ApiAppClipsAppClipAdvancedExperiencesGetToManyRelatedRequest {
	r.fieldsAppClipAdvancedExperiences = &fieldsAppClipAdvancedExperiences
	return r
}

// the fields to include for returned resources of type appClips
func (r ApiAppClipsAppClipAdvancedExperiencesGetToManyRelatedRequest) FieldsAppClips(fieldsAppClips []string) ApiAppClipsAppClipAdvancedExperiencesGetToManyRelatedRequest {
	r.fieldsAppClips = &fieldsAppClips
	return r
}

// the fields to include for returned resources of type appClipAdvancedExperienceImages
func (r ApiAppClipsAppClipAdvancedExperiencesGetToManyRelatedRequest) FieldsAppClipAdvancedExperienceImages(fieldsAppClipAdvancedExperienceImages []string) ApiAppClipsAppClipAdvancedExperiencesGetToManyRelatedRequest {
	r.fieldsAppClipAdvancedExperienceImages = &fieldsAppClipAdvancedExperienceImages
	return r
}

// the fields to include for returned resources of type appClipAdvancedExperienceLocalizations
func (r ApiAppClipsAppClipAdvancedExperiencesGetToManyRelatedRequest) FieldsAppClipAdvancedExperienceLocalizations(fieldsAppClipAdvancedExperienceLocalizations []string) ApiAppClipsAppClipAdvancedExperiencesGetToManyRelatedRequest {
	r.fieldsAppClipAdvancedExperienceLocalizations = &fieldsAppClipAdvancedExperienceLocalizations
	return r
}

// maximum resources per page
func (r ApiAppClipsAppClipAdvancedExperiencesGetToManyRelatedRequest) Limit(limit int32) ApiAppClipsAppClipAdvancedExperiencesGetToManyRelatedRequest {
	r.limit = &limit
	return r
}

// maximum number of related localizations returned (when they are included)
func (r ApiAppClipsAppClipAdvancedExperiencesGetToManyRelatedRequest) LimitLocalizations(limitLocalizations int32) ApiAppClipsAppClipAdvancedExperiencesGetToManyRelatedRequest {
	r.limitLocalizations = &limitLocalizations
	return r
}

// comma-separated list of relationships to include
func (r ApiAppClipsAppClipAdvancedExperiencesGetToManyRelatedRequest) Include(include []string) ApiAppClipsAppClipAdvancedExperiencesGetToManyRelatedRequest {
	r.include = &include
	return r
}

func (r ApiAppClipsAppClipAdvancedExperiencesGetToManyRelatedRequest) Execute() (*AppClipAdvancedExperiencesResponse, *http.Response, error) {
	return r.ApiService.AppClipsAppClipAdvancedExperiencesGetToManyRelatedExecute(r)
}

/*
AppClipsAppClipAdvancedExperiencesGetToManyRelated Method for AppClipsAppClipAdvancedExperiencesGetToManyRelated

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id the id of the requested resource
	@return ApiAppClipsAppClipAdvancedExperiencesGetToManyRelatedRequest
*/
func (a *AppClipsApiService) AppClipsAppClipAdvancedExperiencesGetToManyRelated(ctx context.Context, id string) ApiAppClipsAppClipAdvancedExperiencesGetToManyRelatedRequest {
	return ApiAppClipsAppClipAdvancedExperiencesGetToManyRelatedRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return AppClipAdvancedExperiencesResponse
func (a *AppClipsApiService) AppClipsAppClipAdvancedExperiencesGetToManyRelatedExecute(r ApiAppClipsAppClipAdvancedExperiencesGetToManyRelatedRequest) (*AppClipAdvancedExperiencesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AppClipAdvancedExperiencesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppClipsApiService.AppClipsAppClipAdvancedExperiencesGetToManyRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appClips/{id}/appClipAdvancedExperiences"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filterAction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[action]", r.filterAction, "csv")
	}
	if r.filterPlaceStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[placeStatus]", r.filterPlaceStatus, "csv")
	}
	if r.filterStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[status]", r.filterStatus, "csv")
	}
	if r.fieldsAppClipAdvancedExperiences != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appClipAdvancedExperiences]", r.fieldsAppClipAdvancedExperiences, "csv")
	}
	if r.fieldsAppClips != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appClips]", r.fieldsAppClips, "csv")
	}
	if r.fieldsAppClipAdvancedExperienceImages != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appClipAdvancedExperienceImages]", r.fieldsAppClipAdvancedExperienceImages, "csv")
	}
	if r.fieldsAppClipAdvancedExperienceLocalizations != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appClipAdvancedExperienceLocalizations]", r.fieldsAppClipAdvancedExperienceLocalizations, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.limitLocalizations != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit[localizations]", r.limitLocalizations, "")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiAppClipsAppClipDefaultExperiencesGetToManyRelatedRequest struct {
	ctx                                         context.Context
	ApiService                                  *AppClipsApiService
	id                                          string
	existsReleaseWithAppStoreVersion            *bool
	fieldsAppClips                              *[]string
	fieldsAppClipAppStoreReviewDetails          *[]string
	fieldsAppStoreVersions                      *[]string
	fieldsAppClipDefaultExperiences             *[]string
	fieldsAppClipDefaultExperienceLocalizations *[]string
	limit                                       *int32
	limitAppClipDefaultExperienceLocalizations  *int32
	include                                     *[]string
}

// filter by existence or non-existence of related &#39;releaseWithAppStoreVersion&#39;
func (r ApiAppClipsAppClipDefaultExperiencesGetToManyRelatedRequest) ExistsReleaseWithAppStoreVersion(existsReleaseWithAppStoreVersion bool) ApiAppClipsAppClipDefaultExperiencesGetToManyRelatedRequest {
	r.existsReleaseWithAppStoreVersion = &existsReleaseWithAppStoreVersion
	return r
}

// the fields to include for returned resources of type appClips
func (r ApiAppClipsAppClipDefaultExperiencesGetToManyRelatedRequest) FieldsAppClips(fieldsAppClips []string) ApiAppClipsAppClipDefaultExperiencesGetToManyRelatedRequest {
	r.fieldsAppClips = &fieldsAppClips
	return r
}

// the fields to include for returned resources of type appClipAppStoreReviewDetails
func (r ApiAppClipsAppClipDefaultExperiencesGetToManyRelatedRequest) FieldsAppClipAppStoreReviewDetails(fieldsAppClipAppStoreReviewDetails []string) ApiAppClipsAppClipDefaultExperiencesGetToManyRelatedRequest {
	r.fieldsAppClipAppStoreReviewDetails = &fieldsAppClipAppStoreReviewDetails
	return r
}

// the fields to include for returned resources of type appStoreVersions
func (r ApiAppClipsAppClipDefaultExperiencesGetToManyRelatedRequest) FieldsAppStoreVersions(fieldsAppStoreVersions []string) ApiAppClipsAppClipDefaultExperiencesGetToManyRelatedRequest {
	r.fieldsAppStoreVersions = &fieldsAppStoreVersions
	return r
}

// the fields to include for returned resources of type appClipDefaultExperiences
func (r ApiAppClipsAppClipDefaultExperiencesGetToManyRelatedRequest) FieldsAppClipDefaultExperiences(fieldsAppClipDefaultExperiences []string) ApiAppClipsAppClipDefaultExperiencesGetToManyRelatedRequest {
	r.fieldsAppClipDefaultExperiences = &fieldsAppClipDefaultExperiences
	return r
}

// the fields to include for returned resources of type appClipDefaultExperienceLocalizations
func (r ApiAppClipsAppClipDefaultExperiencesGetToManyRelatedRequest) FieldsAppClipDefaultExperienceLocalizations(fieldsAppClipDefaultExperienceLocalizations []string) ApiAppClipsAppClipDefaultExperiencesGetToManyRelatedRequest {
	r.fieldsAppClipDefaultExperienceLocalizations = &fieldsAppClipDefaultExperienceLocalizations
	return r
}

// maximum resources per page
func (r ApiAppClipsAppClipDefaultExperiencesGetToManyRelatedRequest) Limit(limit int32) ApiAppClipsAppClipDefaultExperiencesGetToManyRelatedRequest {
	r.limit = &limit
	return r
}

// maximum number of related appClipDefaultExperienceLocalizations returned (when they are included)
func (r ApiAppClipsAppClipDefaultExperiencesGetToManyRelatedRequest) LimitAppClipDefaultExperienceLocalizations(limitAppClipDefaultExperienceLocalizations int32) ApiAppClipsAppClipDefaultExperiencesGetToManyRelatedRequest {
	r.limitAppClipDefaultExperienceLocalizations = &limitAppClipDefaultExperienceLocalizations
	return r
}

// comma-separated list of relationships to include
func (r ApiAppClipsAppClipDefaultExperiencesGetToManyRelatedRequest) Include(include []string) ApiAppClipsAppClipDefaultExperiencesGetToManyRelatedRequest {
	r.include = &include
	return r
}

func (r ApiAppClipsAppClipDefaultExperiencesGetToManyRelatedRequest) Execute() (*AppClipDefaultExperiencesResponse, *http.Response, error) {
	return r.ApiService.AppClipsAppClipDefaultExperiencesGetToManyRelatedExecute(r)
}

/*
AppClipsAppClipDefaultExperiencesGetToManyRelated Method for AppClipsAppClipDefaultExperiencesGetToManyRelated

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id the id of the requested resource
	@return ApiAppClipsAppClipDefaultExperiencesGetToManyRelatedRequest
*/
func (a *AppClipsApiService) AppClipsAppClipDefaultExperiencesGetToManyRelated(ctx context.Context, id string) ApiAppClipsAppClipDefaultExperiencesGetToManyRelatedRequest {
	return ApiAppClipsAppClipDefaultExperiencesGetToManyRelatedRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return AppClipDefaultExperiencesResponse
func (a *AppClipsApiService) AppClipsAppClipDefaultExperiencesGetToManyRelatedExecute(r ApiAppClipsAppClipDefaultExperiencesGetToManyRelatedRequest) (*AppClipDefaultExperiencesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AppClipDefaultExperiencesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppClipsApiService.AppClipsAppClipDefaultExperiencesGetToManyRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appClips/{id}/appClipDefaultExperiences"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.existsReleaseWithAppStoreVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exists[releaseWithAppStoreVersion]", r.existsReleaseWithAppStoreVersion, "")
	}
	if r.fieldsAppClips != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appClips]", r.fieldsAppClips, "csv")
	}
	if r.fieldsAppClipAppStoreReviewDetails != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appClipAppStoreReviewDetails]", r.fieldsAppClipAppStoreReviewDetails, "csv")
	}
	if r.fieldsAppStoreVersions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appStoreVersions]", r.fieldsAppStoreVersions, "csv")
	}
	if r.fieldsAppClipDefaultExperiences != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appClipDefaultExperiences]", r.fieldsAppClipDefaultExperiences, "csv")
	}
	if r.fieldsAppClipDefaultExperienceLocalizations != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appClipDefaultExperienceLocalizations]", r.fieldsAppClipDefaultExperienceLocalizations, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.limitAppClipDefaultExperienceLocalizations != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit[appClipDefaultExperienceLocalizations]", r.limitAppClipDefaultExperienceLocalizations, "")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiAppClipsGetInstanceRequest struct {
	ctx                              context.Context
	ApiService                       *AppClipsApiService
	id                               string
	fieldsAppClips                   *[]string
	include                          *[]string
	fieldsAppClipAdvancedExperiences *[]string
	fieldsAppClipDefaultExperiences  *[]string
	limitAppClipDefaultExperiences   *int32
}

// the fields to include for returned resources of type appClips
func (r ApiAppClipsGetInstanceRequest) FieldsAppClips(fieldsAppClips []string) ApiAppClipsGetInstanceRequest {
	r.fieldsAppClips = &fieldsAppClips
	return r
}

// comma-separated list of relationships to include
func (r ApiAppClipsGetInstanceRequest) Include(include []string) ApiAppClipsGetInstanceRequest {
	r.include = &include
	return r
}

// the fields to include for returned resources of type appClipAdvancedExperiences
func (r ApiAppClipsGetInstanceRequest) FieldsAppClipAdvancedExperiences(fieldsAppClipAdvancedExperiences []string) ApiAppClipsGetInstanceRequest {
	r.fieldsAppClipAdvancedExperiences = &fieldsAppClipAdvancedExperiences
	return r
}

// the fields to include for returned resources of type appClipDefaultExperiences
func (r ApiAppClipsGetInstanceRequest) FieldsAppClipDefaultExperiences(fieldsAppClipDefaultExperiences []string) ApiAppClipsGetInstanceRequest {
	r.fieldsAppClipDefaultExperiences = &fieldsAppClipDefaultExperiences
	return r
}

// maximum number of related appClipDefaultExperiences returned (when they are included)
func (r ApiAppClipsGetInstanceRequest) LimitAppClipDefaultExperiences(limitAppClipDefaultExperiences int32) ApiAppClipsGetInstanceRequest {
	r.limitAppClipDefaultExperiences = &limitAppClipDefaultExperiences
	return r
}

func (r ApiAppClipsGetInstanceRequest) Execute() (*AppClipResponse, *http.Response, error) {
	return r.ApiService.AppClipsGetInstanceExecute(r)
}

/*
AppClipsGetInstance Method for AppClipsGetInstance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id the id of the requested resource
	@return ApiAppClipsGetInstanceRequest
*/
func (a *AppClipsApiService) AppClipsGetInstance(ctx context.Context, id string) ApiAppClipsGetInstanceRequest {
	return ApiAppClipsGetInstanceRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return AppClipResponse
func (a *AppClipsApiService) AppClipsGetInstanceExecute(r ApiAppClipsGetInstanceRequest) (*AppClipResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AppClipResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppClipsApiService.AppClipsGetInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appClips/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsAppClips != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appClips]", r.fieldsAppClips, "csv")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "csv")
	}
	if r.fieldsAppClipAdvancedExperiences != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appClipAdvancedExperiences]", r.fieldsAppClipAdvancedExperiences, "csv")
	}
	if r.fieldsAppClipDefaultExperiences != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appClipDefaultExperiences]", r.fieldsAppClipDefaultExperiences, "csv")
	}
	if r.limitAppClipDefaultExperiences != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit[appClipDefaultExperiences]", r.limitAppClipDefaultExperiences, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
