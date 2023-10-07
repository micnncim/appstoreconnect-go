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

// AppCustomProductPagesApiService AppCustomProductPagesApi service
type AppCustomProductPagesApiService service

type ApiAppCustomProductPagesAppCustomProductPageVersionsGetToManyRelatedRequest struct {
	ctx                                     context.Context
	ApiService                              *AppCustomProductPagesApiService
	id                                      string
	filterState                             *[]string
	fieldsAppCustomProductPageLocalizations *[]string
	fieldsAppCustomProductPageVersions      *[]string
	fieldsAppCustomProductPages             *[]string
	limit                                   *int32
	limitAppCustomProductPageLocalizations  *int32
	include                                 *[]string
}

// filter by attribute &#39;state&#39;
func (r ApiAppCustomProductPagesAppCustomProductPageVersionsGetToManyRelatedRequest) FilterState(filterState []string) ApiAppCustomProductPagesAppCustomProductPageVersionsGetToManyRelatedRequest {
	r.filterState = &filterState
	return r
}

// the fields to include for returned resources of type appCustomProductPageLocalizations
func (r ApiAppCustomProductPagesAppCustomProductPageVersionsGetToManyRelatedRequest) FieldsAppCustomProductPageLocalizations(fieldsAppCustomProductPageLocalizations []string) ApiAppCustomProductPagesAppCustomProductPageVersionsGetToManyRelatedRequest {
	r.fieldsAppCustomProductPageLocalizations = &fieldsAppCustomProductPageLocalizations
	return r
}

// the fields to include for returned resources of type appCustomProductPageVersions
func (r ApiAppCustomProductPagesAppCustomProductPageVersionsGetToManyRelatedRequest) FieldsAppCustomProductPageVersions(fieldsAppCustomProductPageVersions []string) ApiAppCustomProductPagesAppCustomProductPageVersionsGetToManyRelatedRequest {
	r.fieldsAppCustomProductPageVersions = &fieldsAppCustomProductPageVersions
	return r
}

// the fields to include for returned resources of type appCustomProductPages
func (r ApiAppCustomProductPagesAppCustomProductPageVersionsGetToManyRelatedRequest) FieldsAppCustomProductPages(fieldsAppCustomProductPages []string) ApiAppCustomProductPagesAppCustomProductPageVersionsGetToManyRelatedRequest {
	r.fieldsAppCustomProductPages = &fieldsAppCustomProductPages
	return r
}

// maximum resources per page
func (r ApiAppCustomProductPagesAppCustomProductPageVersionsGetToManyRelatedRequest) Limit(limit int32) ApiAppCustomProductPagesAppCustomProductPageVersionsGetToManyRelatedRequest {
	r.limit = &limit
	return r
}

// maximum number of related appCustomProductPageLocalizations returned (when they are included)
func (r ApiAppCustomProductPagesAppCustomProductPageVersionsGetToManyRelatedRequest) LimitAppCustomProductPageLocalizations(limitAppCustomProductPageLocalizations int32) ApiAppCustomProductPagesAppCustomProductPageVersionsGetToManyRelatedRequest {
	r.limitAppCustomProductPageLocalizations = &limitAppCustomProductPageLocalizations
	return r
}

// comma-separated list of relationships to include
func (r ApiAppCustomProductPagesAppCustomProductPageVersionsGetToManyRelatedRequest) Include(include []string) ApiAppCustomProductPagesAppCustomProductPageVersionsGetToManyRelatedRequest {
	r.include = &include
	return r
}

func (r ApiAppCustomProductPagesAppCustomProductPageVersionsGetToManyRelatedRequest) Execute() (*AppCustomProductPageVersionsResponse, *http.Response, error) {
	return r.ApiService.AppCustomProductPagesAppCustomProductPageVersionsGetToManyRelatedExecute(r)
}

/*
AppCustomProductPagesAppCustomProductPageVersionsGetToManyRelated Method for AppCustomProductPagesAppCustomProductPageVersionsGetToManyRelated

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id the id of the requested resource
	@return ApiAppCustomProductPagesAppCustomProductPageVersionsGetToManyRelatedRequest
*/
func (a *AppCustomProductPagesApiService) AppCustomProductPagesAppCustomProductPageVersionsGetToManyRelated(ctx context.Context, id string) ApiAppCustomProductPagesAppCustomProductPageVersionsGetToManyRelatedRequest {
	return ApiAppCustomProductPagesAppCustomProductPageVersionsGetToManyRelatedRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return AppCustomProductPageVersionsResponse
func (a *AppCustomProductPagesApiService) AppCustomProductPagesAppCustomProductPageVersionsGetToManyRelatedExecute(r ApiAppCustomProductPagesAppCustomProductPageVersionsGetToManyRelatedRequest) (*AppCustomProductPageVersionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AppCustomProductPageVersionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppCustomProductPagesApiService.AppCustomProductPagesAppCustomProductPageVersionsGetToManyRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appCustomProductPages/{id}/appCustomProductPageVersions"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filterState != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[state]", r.filterState, "csv")
	}
	if r.fieldsAppCustomProductPageLocalizations != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appCustomProductPageLocalizations]", r.fieldsAppCustomProductPageLocalizations, "csv")
	}
	if r.fieldsAppCustomProductPageVersions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appCustomProductPageVersions]", r.fieldsAppCustomProductPageVersions, "csv")
	}
	if r.fieldsAppCustomProductPages != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appCustomProductPages]", r.fieldsAppCustomProductPages, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.limitAppCustomProductPageLocalizations != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit[appCustomProductPageLocalizations]", r.limitAppCustomProductPageLocalizations, "")
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

type ApiAppCustomProductPagesCreateInstanceRequest struct {
	ctx                               context.Context
	ApiService                        *AppCustomProductPagesApiService
	appCustomProductPageCreateRequest *AppCustomProductPageCreateRequest
}

// AppCustomProductPage representation
func (r ApiAppCustomProductPagesCreateInstanceRequest) AppCustomProductPageCreateRequest(appCustomProductPageCreateRequest AppCustomProductPageCreateRequest) ApiAppCustomProductPagesCreateInstanceRequest {
	r.appCustomProductPageCreateRequest = &appCustomProductPageCreateRequest
	return r
}

func (r ApiAppCustomProductPagesCreateInstanceRequest) Execute() (*AppCustomProductPageResponse, *http.Response, error) {
	return r.ApiService.AppCustomProductPagesCreateInstanceExecute(r)
}

/*
AppCustomProductPagesCreateInstance Method for AppCustomProductPagesCreateInstance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAppCustomProductPagesCreateInstanceRequest
*/
func (a *AppCustomProductPagesApiService) AppCustomProductPagesCreateInstance(ctx context.Context) ApiAppCustomProductPagesCreateInstanceRequest {
	return ApiAppCustomProductPagesCreateInstanceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AppCustomProductPageResponse
func (a *AppCustomProductPagesApiService) AppCustomProductPagesCreateInstanceExecute(r ApiAppCustomProductPagesCreateInstanceRequest) (*AppCustomProductPageResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AppCustomProductPageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppCustomProductPagesApiService.AppCustomProductPagesCreateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appCustomProductPages"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.appCustomProductPageCreateRequest == nil {
		return localVarReturnValue, nil, reportError("appCustomProductPageCreateRequest is required and must be specified")
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
	localVarPostBody = r.appCustomProductPageCreateRequest
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

type ApiAppCustomProductPagesDeleteInstanceRequest struct {
	ctx        context.Context
	ApiService *AppCustomProductPagesApiService
	id         string
}

func (r ApiAppCustomProductPagesDeleteInstanceRequest) Execute() (*http.Response, error) {
	return r.ApiService.AppCustomProductPagesDeleteInstanceExecute(r)
}

/*
AppCustomProductPagesDeleteInstance Method for AppCustomProductPagesDeleteInstance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id the id of the requested resource
	@return ApiAppCustomProductPagesDeleteInstanceRequest
*/
func (a *AppCustomProductPagesApiService) AppCustomProductPagesDeleteInstance(ctx context.Context, id string) ApiAppCustomProductPagesDeleteInstanceRequest {
	return ApiAppCustomProductPagesDeleteInstanceRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *AppCustomProductPagesApiService) AppCustomProductPagesDeleteInstanceExecute(r ApiAppCustomProductPagesDeleteInstanceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppCustomProductPagesApiService.AppCustomProductPagesDeleteInstance")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appCustomProductPages/{id}"
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

type ApiAppCustomProductPagesGetInstanceRequest struct {
	ctx                                context.Context
	ApiService                         *AppCustomProductPagesApiService
	id                                 string
	fieldsAppCustomProductPages        *[]string
	include                            *[]string
	fieldsAppCustomProductPageVersions *[]string
	limitAppCustomProductPageVersions  *int32
}

// the fields to include for returned resources of type appCustomProductPages
func (r ApiAppCustomProductPagesGetInstanceRequest) FieldsAppCustomProductPages(fieldsAppCustomProductPages []string) ApiAppCustomProductPagesGetInstanceRequest {
	r.fieldsAppCustomProductPages = &fieldsAppCustomProductPages
	return r
}

// comma-separated list of relationships to include
func (r ApiAppCustomProductPagesGetInstanceRequest) Include(include []string) ApiAppCustomProductPagesGetInstanceRequest {
	r.include = &include
	return r
}

// the fields to include for returned resources of type appCustomProductPageVersions
func (r ApiAppCustomProductPagesGetInstanceRequest) FieldsAppCustomProductPageVersions(fieldsAppCustomProductPageVersions []string) ApiAppCustomProductPagesGetInstanceRequest {
	r.fieldsAppCustomProductPageVersions = &fieldsAppCustomProductPageVersions
	return r
}

// maximum number of related appCustomProductPageVersions returned (when they are included)
func (r ApiAppCustomProductPagesGetInstanceRequest) LimitAppCustomProductPageVersions(limitAppCustomProductPageVersions int32) ApiAppCustomProductPagesGetInstanceRequest {
	r.limitAppCustomProductPageVersions = &limitAppCustomProductPageVersions
	return r
}

func (r ApiAppCustomProductPagesGetInstanceRequest) Execute() (*AppCustomProductPageResponse, *http.Response, error) {
	return r.ApiService.AppCustomProductPagesGetInstanceExecute(r)
}

/*
AppCustomProductPagesGetInstance Method for AppCustomProductPagesGetInstance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id the id of the requested resource
	@return ApiAppCustomProductPagesGetInstanceRequest
*/
func (a *AppCustomProductPagesApiService) AppCustomProductPagesGetInstance(ctx context.Context, id string) ApiAppCustomProductPagesGetInstanceRequest {
	return ApiAppCustomProductPagesGetInstanceRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return AppCustomProductPageResponse
func (a *AppCustomProductPagesApiService) AppCustomProductPagesGetInstanceExecute(r ApiAppCustomProductPagesGetInstanceRequest) (*AppCustomProductPageResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AppCustomProductPageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppCustomProductPagesApiService.AppCustomProductPagesGetInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appCustomProductPages/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsAppCustomProductPages != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appCustomProductPages]", r.fieldsAppCustomProductPages, "csv")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "csv")
	}
	if r.fieldsAppCustomProductPageVersions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appCustomProductPageVersions]", r.fieldsAppCustomProductPageVersions, "csv")
	}
	if r.limitAppCustomProductPageVersions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit[appCustomProductPageVersions]", r.limitAppCustomProductPageVersions, "")
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

type ApiAppCustomProductPagesUpdateInstanceRequest struct {
	ctx                               context.Context
	ApiService                        *AppCustomProductPagesApiService
	id                                string
	appCustomProductPageUpdateRequest *AppCustomProductPageUpdateRequest
}

// AppCustomProductPage representation
func (r ApiAppCustomProductPagesUpdateInstanceRequest) AppCustomProductPageUpdateRequest(appCustomProductPageUpdateRequest AppCustomProductPageUpdateRequest) ApiAppCustomProductPagesUpdateInstanceRequest {
	r.appCustomProductPageUpdateRequest = &appCustomProductPageUpdateRequest
	return r
}

func (r ApiAppCustomProductPagesUpdateInstanceRequest) Execute() (*AppCustomProductPageResponse, *http.Response, error) {
	return r.ApiService.AppCustomProductPagesUpdateInstanceExecute(r)
}

/*
AppCustomProductPagesUpdateInstance Method for AppCustomProductPagesUpdateInstance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id the id of the requested resource
	@return ApiAppCustomProductPagesUpdateInstanceRequest
*/
func (a *AppCustomProductPagesApiService) AppCustomProductPagesUpdateInstance(ctx context.Context, id string) ApiAppCustomProductPagesUpdateInstanceRequest {
	return ApiAppCustomProductPagesUpdateInstanceRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return AppCustomProductPageResponse
func (a *AppCustomProductPagesApiService) AppCustomProductPagesUpdateInstanceExecute(r ApiAppCustomProductPagesUpdateInstanceRequest) (*AppCustomProductPageResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AppCustomProductPageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppCustomProductPagesApiService.AppCustomProductPagesUpdateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appCustomProductPages/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.appCustomProductPageUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("appCustomProductPageUpdateRequest is required and must be specified")
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
	localVarPostBody = r.appCustomProductPageUpdateRequest
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
