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
	"os"
)

// SalesReportsApiService SalesReportsApi service
type SalesReportsApiService service

type ApiSalesReportsGetCollectionRequest struct {
	ctx                 context.Context
	ApiService          *SalesReportsApiService
	filterFrequency     *[]string
	filterReportSubType *[]string
	filterReportType    *[]string
	filterVendorNumber  *[]string
	filterReportDate    *[]string
	filterVersion       *[]string
}

// filter by attribute &#39;frequency&#39;
func (r ApiSalesReportsGetCollectionRequest) FilterFrequency(filterFrequency []string) ApiSalesReportsGetCollectionRequest {
	r.filterFrequency = &filterFrequency
	return r
}

// filter by attribute &#39;reportSubType&#39;
func (r ApiSalesReportsGetCollectionRequest) FilterReportSubType(filterReportSubType []string) ApiSalesReportsGetCollectionRequest {
	r.filterReportSubType = &filterReportSubType
	return r
}

// filter by attribute &#39;reportType&#39;
func (r ApiSalesReportsGetCollectionRequest) FilterReportType(filterReportType []string) ApiSalesReportsGetCollectionRequest {
	r.filterReportType = &filterReportType
	return r
}

// filter by attribute &#39;vendorNumber&#39;
func (r ApiSalesReportsGetCollectionRequest) FilterVendorNumber(filterVendorNumber []string) ApiSalesReportsGetCollectionRequest {
	r.filterVendorNumber = &filterVendorNumber
	return r
}

// filter by attribute &#39;reportDate&#39;
func (r ApiSalesReportsGetCollectionRequest) FilterReportDate(filterReportDate []string) ApiSalesReportsGetCollectionRequest {
	r.filterReportDate = &filterReportDate
	return r
}

// filter by attribute &#39;version&#39;
func (r ApiSalesReportsGetCollectionRequest) FilterVersion(filterVersion []string) ApiSalesReportsGetCollectionRequest {
	r.filterVersion = &filterVersion
	return r
}

func (r ApiSalesReportsGetCollectionRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.SalesReportsGetCollectionExecute(r)
}

/*
SalesReportsGetCollection Method for SalesReportsGetCollection

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSalesReportsGetCollectionRequest
*/
func (a *SalesReportsApiService) SalesReportsGetCollection(ctx context.Context) ApiSalesReportsGetCollectionRequest {
	return ApiSalesReportsGetCollectionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return *os.File
func (a *SalesReportsApiService) SalesReportsGetCollectionExecute(r ApiSalesReportsGetCollectionRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SalesReportsApiService.SalesReportsGetCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/salesReports"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.filterFrequency == nil {
		return localVarReturnValue, nil, reportError("filterFrequency is required and must be specified")
	}
	if r.filterReportSubType == nil {
		return localVarReturnValue, nil, reportError("filterReportSubType is required and must be specified")
	}
	if r.filterReportType == nil {
		return localVarReturnValue, nil, reportError("filterReportType is required and must be specified")
	}
	if r.filterVendorNumber == nil {
		return localVarReturnValue, nil, reportError("filterVendorNumber is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "filter[frequency]", r.filterFrequency, "csv")
	if r.filterReportDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[reportDate]", r.filterReportDate, "csv")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "filter[reportSubType]", r.filterReportSubType, "csv")
	parameterAddToHeaderOrQuery(localVarQueryParams, "filter[reportType]", r.filterReportType, "csv")
	parameterAddToHeaderOrQuery(localVarQueryParams, "filter[vendorNumber]", r.filterVendorNumber, "csv")
	if r.filterVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[version]", r.filterVersion, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/a-gzip"}

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
