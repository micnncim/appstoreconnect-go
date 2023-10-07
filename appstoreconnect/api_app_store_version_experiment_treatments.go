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

// AppStoreVersionExperimentTreatmentsApiService AppStoreVersionExperimentTreatmentsApi service
type AppStoreVersionExperimentTreatmentsApiService service

type ApiAppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelatedRequest struct {
	ctx                                                   context.Context
	ApiService                                            *AppStoreVersionExperimentTreatmentsApiService
	id                                                    string
	filterLocale                                          *[]string
	fieldsAppScreenshotSets                               *[]string
	fieldsAppStoreVersionExperimentTreatments             *[]string
	fieldsAppStoreVersionExperimentTreatmentLocalizations *[]string
	fieldsAppPreviewSets                                  *[]string
	limit                                                 *int32
	limitAppScreenshotSets                                *int32
	limitAppPreviewSets                                   *int32
	include                                               *[]string
}

// filter by attribute &#39;locale&#39;
func (r ApiAppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelatedRequest) FilterLocale(filterLocale []string) ApiAppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelatedRequest {
	r.filterLocale = &filterLocale
	return r
}

// the fields to include for returned resources of type appScreenshotSets
func (r ApiAppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelatedRequest) FieldsAppScreenshotSets(fieldsAppScreenshotSets []string) ApiAppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelatedRequest {
	r.fieldsAppScreenshotSets = &fieldsAppScreenshotSets
	return r
}

// the fields to include for returned resources of type appStoreVersionExperimentTreatments
func (r ApiAppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelatedRequest) FieldsAppStoreVersionExperimentTreatments(fieldsAppStoreVersionExperimentTreatments []string) ApiAppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelatedRequest {
	r.fieldsAppStoreVersionExperimentTreatments = &fieldsAppStoreVersionExperimentTreatments
	return r
}

// the fields to include for returned resources of type appStoreVersionExperimentTreatmentLocalizations
func (r ApiAppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelatedRequest) FieldsAppStoreVersionExperimentTreatmentLocalizations(fieldsAppStoreVersionExperimentTreatmentLocalizations []string) ApiAppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelatedRequest {
	r.fieldsAppStoreVersionExperimentTreatmentLocalizations = &fieldsAppStoreVersionExperimentTreatmentLocalizations
	return r
}

// the fields to include for returned resources of type appPreviewSets
func (r ApiAppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelatedRequest) FieldsAppPreviewSets(fieldsAppPreviewSets []string) ApiAppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelatedRequest {
	r.fieldsAppPreviewSets = &fieldsAppPreviewSets
	return r
}

// maximum resources per page
func (r ApiAppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelatedRequest) Limit(limit int32) ApiAppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelatedRequest {
	r.limit = &limit
	return r
}

// maximum number of related appScreenshotSets returned (when they are included)
func (r ApiAppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelatedRequest) LimitAppScreenshotSets(limitAppScreenshotSets int32) ApiAppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelatedRequest {
	r.limitAppScreenshotSets = &limitAppScreenshotSets
	return r
}

// maximum number of related appPreviewSets returned (when they are included)
func (r ApiAppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelatedRequest) LimitAppPreviewSets(limitAppPreviewSets int32) ApiAppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelatedRequest {
	r.limitAppPreviewSets = &limitAppPreviewSets
	return r
}

// comma-separated list of relationships to include
func (r ApiAppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelatedRequest) Include(include []string) ApiAppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelatedRequest {
	r.include = &include
	return r
}

func (r ApiAppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelatedRequest) Execute() (*AppStoreVersionExperimentTreatmentLocalizationsResponse, *http.Response, error) {
	return r.ApiService.AppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelatedExecute(r)
}

/*
AppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelated Method for AppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelated

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id the id of the requested resource
	@return ApiAppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelatedRequest
*/
func (a *AppStoreVersionExperimentTreatmentsApiService) AppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelated(ctx context.Context, id string) ApiAppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelatedRequest {
	return ApiAppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelatedRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return AppStoreVersionExperimentTreatmentLocalizationsResponse
func (a *AppStoreVersionExperimentTreatmentsApiService) AppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelatedExecute(r ApiAppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelatedRequest) (*AppStoreVersionExperimentTreatmentLocalizationsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AppStoreVersionExperimentTreatmentLocalizationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppStoreVersionExperimentTreatmentsApiService.AppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appStoreVersionExperimentTreatments/{id}/appStoreVersionExperimentTreatmentLocalizations"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filterLocale != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[locale]", r.filterLocale, "csv")
	}
	if r.fieldsAppScreenshotSets != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appScreenshotSets]", r.fieldsAppScreenshotSets, "csv")
	}
	if r.fieldsAppStoreVersionExperimentTreatments != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appStoreVersionExperimentTreatments]", r.fieldsAppStoreVersionExperimentTreatments, "csv")
	}
	if r.fieldsAppStoreVersionExperimentTreatmentLocalizations != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appStoreVersionExperimentTreatmentLocalizations]", r.fieldsAppStoreVersionExperimentTreatmentLocalizations, "csv")
	}
	if r.fieldsAppPreviewSets != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appPreviewSets]", r.fieldsAppPreviewSets, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.limitAppScreenshotSets != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit[appScreenshotSets]", r.limitAppScreenshotSets, "")
	}
	if r.limitAppPreviewSets != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit[appPreviewSets]", r.limitAppPreviewSets, "")
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

type ApiAppStoreVersionExperimentTreatmentsCreateInstanceRequest struct {
	ctx                                             context.Context
	ApiService                                      *AppStoreVersionExperimentTreatmentsApiService
	appStoreVersionExperimentTreatmentCreateRequest *AppStoreVersionExperimentTreatmentCreateRequest
}

// AppStoreVersionExperimentTreatment representation
func (r ApiAppStoreVersionExperimentTreatmentsCreateInstanceRequest) AppStoreVersionExperimentTreatmentCreateRequest(appStoreVersionExperimentTreatmentCreateRequest AppStoreVersionExperimentTreatmentCreateRequest) ApiAppStoreVersionExperimentTreatmentsCreateInstanceRequest {
	r.appStoreVersionExperimentTreatmentCreateRequest = &appStoreVersionExperimentTreatmentCreateRequest
	return r
}

func (r ApiAppStoreVersionExperimentTreatmentsCreateInstanceRequest) Execute() (*AppStoreVersionExperimentTreatmentResponse, *http.Response, error) {
	return r.ApiService.AppStoreVersionExperimentTreatmentsCreateInstanceExecute(r)
}

/*
AppStoreVersionExperimentTreatmentsCreateInstance Method for AppStoreVersionExperimentTreatmentsCreateInstance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAppStoreVersionExperimentTreatmentsCreateInstanceRequest
*/
func (a *AppStoreVersionExperimentTreatmentsApiService) AppStoreVersionExperimentTreatmentsCreateInstance(ctx context.Context) ApiAppStoreVersionExperimentTreatmentsCreateInstanceRequest {
	return ApiAppStoreVersionExperimentTreatmentsCreateInstanceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AppStoreVersionExperimentTreatmentResponse
func (a *AppStoreVersionExperimentTreatmentsApiService) AppStoreVersionExperimentTreatmentsCreateInstanceExecute(r ApiAppStoreVersionExperimentTreatmentsCreateInstanceRequest) (*AppStoreVersionExperimentTreatmentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AppStoreVersionExperimentTreatmentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppStoreVersionExperimentTreatmentsApiService.AppStoreVersionExperimentTreatmentsCreateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appStoreVersionExperimentTreatments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.appStoreVersionExperimentTreatmentCreateRequest == nil {
		return localVarReturnValue, nil, reportError("appStoreVersionExperimentTreatmentCreateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.appStoreVersionExperimentTreatmentCreateRequest
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
		if localVarHTTPResponse.StatusCode == 409 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ApiAppStoreVersionExperimentTreatmentsDeleteInstanceRequest struct {
	ctx        context.Context
	ApiService *AppStoreVersionExperimentTreatmentsApiService
	id         string
}

func (r ApiAppStoreVersionExperimentTreatmentsDeleteInstanceRequest) Execute() (*http.Response, error) {
	return r.ApiService.AppStoreVersionExperimentTreatmentsDeleteInstanceExecute(r)
}

/*
AppStoreVersionExperimentTreatmentsDeleteInstance Method for AppStoreVersionExperimentTreatmentsDeleteInstance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id the id of the requested resource
	@return ApiAppStoreVersionExperimentTreatmentsDeleteInstanceRequest
*/
func (a *AppStoreVersionExperimentTreatmentsApiService) AppStoreVersionExperimentTreatmentsDeleteInstance(ctx context.Context, id string) ApiAppStoreVersionExperimentTreatmentsDeleteInstanceRequest {
	return ApiAppStoreVersionExperimentTreatmentsDeleteInstanceRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *AppStoreVersionExperimentTreatmentsApiService) AppStoreVersionExperimentTreatmentsDeleteInstanceExecute(r ApiAppStoreVersionExperimentTreatmentsDeleteInstanceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppStoreVersionExperimentTreatmentsApiService.AppStoreVersionExperimentTreatmentsDeleteInstance")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appStoreVersionExperimentTreatments/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiAppStoreVersionExperimentTreatmentsGetInstanceRequest struct {
	ctx                                                   context.Context
	ApiService                                            *AppStoreVersionExperimentTreatmentsApiService
	id                                                    string
	fieldsAppStoreVersionExperimentTreatments             *[]string
	include                                               *[]string
	fieldsAppStoreVersionExperimentTreatmentLocalizations *[]string
	limitAppStoreVersionExperimentTreatmentLocalizations  *int32
}

// the fields to include for returned resources of type appStoreVersionExperimentTreatments
func (r ApiAppStoreVersionExperimentTreatmentsGetInstanceRequest) FieldsAppStoreVersionExperimentTreatments(fieldsAppStoreVersionExperimentTreatments []string) ApiAppStoreVersionExperimentTreatmentsGetInstanceRequest {
	r.fieldsAppStoreVersionExperimentTreatments = &fieldsAppStoreVersionExperimentTreatments
	return r
}

// comma-separated list of relationships to include
func (r ApiAppStoreVersionExperimentTreatmentsGetInstanceRequest) Include(include []string) ApiAppStoreVersionExperimentTreatmentsGetInstanceRequest {
	r.include = &include
	return r
}

// the fields to include for returned resources of type appStoreVersionExperimentTreatmentLocalizations
func (r ApiAppStoreVersionExperimentTreatmentsGetInstanceRequest) FieldsAppStoreVersionExperimentTreatmentLocalizations(fieldsAppStoreVersionExperimentTreatmentLocalizations []string) ApiAppStoreVersionExperimentTreatmentsGetInstanceRequest {
	r.fieldsAppStoreVersionExperimentTreatmentLocalizations = &fieldsAppStoreVersionExperimentTreatmentLocalizations
	return r
}

// maximum number of related appStoreVersionExperimentTreatmentLocalizations returned (when they are included)
func (r ApiAppStoreVersionExperimentTreatmentsGetInstanceRequest) LimitAppStoreVersionExperimentTreatmentLocalizations(limitAppStoreVersionExperimentTreatmentLocalizations int32) ApiAppStoreVersionExperimentTreatmentsGetInstanceRequest {
	r.limitAppStoreVersionExperimentTreatmentLocalizations = &limitAppStoreVersionExperimentTreatmentLocalizations
	return r
}

func (r ApiAppStoreVersionExperimentTreatmentsGetInstanceRequest) Execute() (*AppStoreVersionExperimentTreatmentResponse, *http.Response, error) {
	return r.ApiService.AppStoreVersionExperimentTreatmentsGetInstanceExecute(r)
}

/*
AppStoreVersionExperimentTreatmentsGetInstance Method for AppStoreVersionExperimentTreatmentsGetInstance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id the id of the requested resource
	@return ApiAppStoreVersionExperimentTreatmentsGetInstanceRequest
*/
func (a *AppStoreVersionExperimentTreatmentsApiService) AppStoreVersionExperimentTreatmentsGetInstance(ctx context.Context, id string) ApiAppStoreVersionExperimentTreatmentsGetInstanceRequest {
	return ApiAppStoreVersionExperimentTreatmentsGetInstanceRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return AppStoreVersionExperimentTreatmentResponse
func (a *AppStoreVersionExperimentTreatmentsApiService) AppStoreVersionExperimentTreatmentsGetInstanceExecute(r ApiAppStoreVersionExperimentTreatmentsGetInstanceRequest) (*AppStoreVersionExperimentTreatmentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AppStoreVersionExperimentTreatmentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppStoreVersionExperimentTreatmentsApiService.AppStoreVersionExperimentTreatmentsGetInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appStoreVersionExperimentTreatments/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsAppStoreVersionExperimentTreatments != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appStoreVersionExperimentTreatments]", r.fieldsAppStoreVersionExperimentTreatments, "csv")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "csv")
	}
	if r.fieldsAppStoreVersionExperimentTreatmentLocalizations != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appStoreVersionExperimentTreatmentLocalizations]", r.fieldsAppStoreVersionExperimentTreatmentLocalizations, "csv")
	}
	if r.limitAppStoreVersionExperimentTreatmentLocalizations != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit[appStoreVersionExperimentTreatmentLocalizations]", r.limitAppStoreVersionExperimentTreatmentLocalizations, "")
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

type ApiAppStoreVersionExperimentTreatmentsUpdateInstanceRequest struct {
	ctx                                             context.Context
	ApiService                                      *AppStoreVersionExperimentTreatmentsApiService
	id                                              string
	appStoreVersionExperimentTreatmentUpdateRequest *AppStoreVersionExperimentTreatmentUpdateRequest
}

// AppStoreVersionExperimentTreatment representation
func (r ApiAppStoreVersionExperimentTreatmentsUpdateInstanceRequest) AppStoreVersionExperimentTreatmentUpdateRequest(appStoreVersionExperimentTreatmentUpdateRequest AppStoreVersionExperimentTreatmentUpdateRequest) ApiAppStoreVersionExperimentTreatmentsUpdateInstanceRequest {
	r.appStoreVersionExperimentTreatmentUpdateRequest = &appStoreVersionExperimentTreatmentUpdateRequest
	return r
}

func (r ApiAppStoreVersionExperimentTreatmentsUpdateInstanceRequest) Execute() (*AppStoreVersionExperimentTreatmentResponse, *http.Response, error) {
	return r.ApiService.AppStoreVersionExperimentTreatmentsUpdateInstanceExecute(r)
}

/*
AppStoreVersionExperimentTreatmentsUpdateInstance Method for AppStoreVersionExperimentTreatmentsUpdateInstance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id the id of the requested resource
	@return ApiAppStoreVersionExperimentTreatmentsUpdateInstanceRequest
*/
func (a *AppStoreVersionExperimentTreatmentsApiService) AppStoreVersionExperimentTreatmentsUpdateInstance(ctx context.Context, id string) ApiAppStoreVersionExperimentTreatmentsUpdateInstanceRequest {
	return ApiAppStoreVersionExperimentTreatmentsUpdateInstanceRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return AppStoreVersionExperimentTreatmentResponse
func (a *AppStoreVersionExperimentTreatmentsApiService) AppStoreVersionExperimentTreatmentsUpdateInstanceExecute(r ApiAppStoreVersionExperimentTreatmentsUpdateInstanceRequest) (*AppStoreVersionExperimentTreatmentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AppStoreVersionExperimentTreatmentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppStoreVersionExperimentTreatmentsApiService.AppStoreVersionExperimentTreatmentsUpdateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appStoreVersionExperimentTreatments/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.appStoreVersionExperimentTreatmentUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("appStoreVersionExperimentTreatmentUpdateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.appStoreVersionExperimentTreatmentUpdateRequest
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
		if localVarHTTPResponse.StatusCode == 409 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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
