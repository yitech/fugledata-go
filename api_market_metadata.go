/*
standard public schema

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 12.2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fugledata

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// MarketMetadataAPIService MarketMetadataAPI service
type MarketMetadataAPIService service

type ApiMarketMetadataDeleteRequest struct {
	ctx context.Context
	ApiService *MarketMetadataAPIService
	id *string
	symbol *string
	category *string
	prefer *string
}

func (r ApiMarketMetadataDeleteRequest) Id(id string) ApiMarketMetadataDeleteRequest {
	r.id = &id
	return r
}

func (r ApiMarketMetadataDeleteRequest) Symbol(symbol string) ApiMarketMetadataDeleteRequest {
	r.symbol = &symbol
	return r
}

func (r ApiMarketMetadataDeleteRequest) Category(category string) ApiMarketMetadataDeleteRequest {
	r.category = &category
	return r
}

// Preference
func (r ApiMarketMetadataDeleteRequest) Prefer(prefer string) ApiMarketMetadataDeleteRequest {
	r.prefer = &prefer
	return r
}

func (r ApiMarketMetadataDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.MarketMetadataDeleteExecute(r)
}

/*
MarketMetadataDelete Method for MarketMetadataDelete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMarketMetadataDeleteRequest
*/
func (a *MarketMetadataAPIService) MarketMetadataDelete(ctx context.Context) ApiMarketMetadataDeleteRequest {
	return ApiMarketMetadataDeleteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *MarketMetadataAPIService) MarketMetadataDeleteExecute(r ApiMarketMetadataDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketMetadataAPIService.MarketMetadataDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/market_metadata"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "", "")
	}
	if r.symbol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "", "")
	}
	if r.category != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "category", r.category, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.prefer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Prefer", r.prefer, "", "")
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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiMarketMetadataGetRequest struct {
	ctx context.Context
	ApiService *MarketMetadataAPIService
	id *string
	symbol *string
	category *string
	select_ *string
	order *string
	range_ *string
	rangeUnit *string
	offset *string
	limit *string
	prefer *string
}

func (r ApiMarketMetadataGetRequest) Id(id string) ApiMarketMetadataGetRequest {
	r.id = &id
	return r
}

func (r ApiMarketMetadataGetRequest) Symbol(symbol string) ApiMarketMetadataGetRequest {
	r.symbol = &symbol
	return r
}

func (r ApiMarketMetadataGetRequest) Category(category string) ApiMarketMetadataGetRequest {
	r.category = &category
	return r
}

// Filtering Columns
func (r ApiMarketMetadataGetRequest) Select_(select_ string) ApiMarketMetadataGetRequest {
	r.select_ = &select_
	return r
}

// Ordering
func (r ApiMarketMetadataGetRequest) Order(order string) ApiMarketMetadataGetRequest {
	r.order = &order
	return r
}

// Limiting and Pagination
func (r ApiMarketMetadataGetRequest) Range_(range_ string) ApiMarketMetadataGetRequest {
	r.range_ = &range_
	return r
}

// Limiting and Pagination
func (r ApiMarketMetadataGetRequest) RangeUnit(rangeUnit string) ApiMarketMetadataGetRequest {
	r.rangeUnit = &rangeUnit
	return r
}

// Limiting and Pagination
func (r ApiMarketMetadataGetRequest) Offset(offset string) ApiMarketMetadataGetRequest {
	r.offset = &offset
	return r
}

// Limiting and Pagination
func (r ApiMarketMetadataGetRequest) Limit(limit string) ApiMarketMetadataGetRequest {
	r.limit = &limit
	return r
}

// Preference
func (r ApiMarketMetadataGetRequest) Prefer(prefer string) ApiMarketMetadataGetRequest {
	r.prefer = &prefer
	return r
}

func (r ApiMarketMetadataGetRequest) Execute() ([]MarketMetadata, *http.Response, error) {
	return r.ApiService.MarketMetadataGetExecute(r)
}

/*
MarketMetadataGet Method for MarketMetadataGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMarketMetadataGetRequest
*/
func (a *MarketMetadataAPIService) MarketMetadataGet(ctx context.Context) ApiMarketMetadataGetRequest {
	return ApiMarketMetadataGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []MarketMetadata
func (a *MarketMetadataAPIService) MarketMetadataGetExecute(r ApiMarketMetadataGetRequest) ([]MarketMetadata, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []MarketMetadata
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketMetadataAPIService.MarketMetadataGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/market_metadata"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "", "")
	}
	if r.symbol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "", "")
	}
	if r.category != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "category", r.category, "", "")
	}
	if r.select_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "select", r.select_, "", "")
	}
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "", "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.pgrst.object+json;nulls=stripped", "application/vnd.pgrst.object+json", "text/csv"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.range_ != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Range", r.range_, "", "")
	}
	if r.rangeUnit != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Range-Unit", r.rangeUnit, "", "")
	}
	if r.prefer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Prefer", r.prefer, "", "")
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

type ApiMarketMetadataPatchRequest struct {
	ctx context.Context
	ApiService *MarketMetadataAPIService
	id *string
	symbol *string
	category *string
	prefer *string
	marketMetadata *MarketMetadata
}

func (r ApiMarketMetadataPatchRequest) Id(id string) ApiMarketMetadataPatchRequest {
	r.id = &id
	return r
}

func (r ApiMarketMetadataPatchRequest) Symbol(symbol string) ApiMarketMetadataPatchRequest {
	r.symbol = &symbol
	return r
}

func (r ApiMarketMetadataPatchRequest) Category(category string) ApiMarketMetadataPatchRequest {
	r.category = &category
	return r
}

// Preference
func (r ApiMarketMetadataPatchRequest) Prefer(prefer string) ApiMarketMetadataPatchRequest {
	r.prefer = &prefer
	return r
}

// market_metadata
func (r ApiMarketMetadataPatchRequest) MarketMetadata(marketMetadata MarketMetadata) ApiMarketMetadataPatchRequest {
	r.marketMetadata = &marketMetadata
	return r
}

func (r ApiMarketMetadataPatchRequest) Execute() (*http.Response, error) {
	return r.ApiService.MarketMetadataPatchExecute(r)
}

/*
MarketMetadataPatch Method for MarketMetadataPatch

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMarketMetadataPatchRequest
*/
func (a *MarketMetadataAPIService) MarketMetadataPatch(ctx context.Context) ApiMarketMetadataPatchRequest {
	return ApiMarketMetadataPatchRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *MarketMetadataAPIService) MarketMetadataPatchExecute(r ApiMarketMetadataPatchRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketMetadataAPIService.MarketMetadataPatch")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/market_metadata"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "", "")
	}
	if r.symbol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "", "")
	}
	if r.category != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "category", r.category, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/vnd.pgrst.object+json;nulls=stripped", "application/vnd.pgrst.object+json", "text/csv"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.prefer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Prefer", r.prefer, "", "")
	}
	// body params
	localVarPostBody = r.marketMetadata
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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiMarketMetadataPostRequest struct {
	ctx context.Context
	ApiService *MarketMetadataAPIService
	select_ *string
	prefer *string
	marketMetadata *MarketMetadata
}

// Filtering Columns
func (r ApiMarketMetadataPostRequest) Select_(select_ string) ApiMarketMetadataPostRequest {
	r.select_ = &select_
	return r
}

// Preference
func (r ApiMarketMetadataPostRequest) Prefer(prefer string) ApiMarketMetadataPostRequest {
	r.prefer = &prefer
	return r
}

// market_metadata
func (r ApiMarketMetadataPostRequest) MarketMetadata(marketMetadata MarketMetadata) ApiMarketMetadataPostRequest {
	r.marketMetadata = &marketMetadata
	return r
}

func (r ApiMarketMetadataPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.MarketMetadataPostExecute(r)
}

/*
MarketMetadataPost Method for MarketMetadataPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMarketMetadataPostRequest
*/
func (a *MarketMetadataAPIService) MarketMetadataPost(ctx context.Context) ApiMarketMetadataPostRequest {
	return ApiMarketMetadataPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *MarketMetadataAPIService) MarketMetadataPostExecute(r ApiMarketMetadataPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketMetadataAPIService.MarketMetadataPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/market_metadata"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.select_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "select", r.select_, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/vnd.pgrst.object+json;nulls=stripped", "application/vnd.pgrst.object+json", "text/csv"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.prefer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Prefer", r.prefer, "", "")
	}
	// body params
	localVarPostBody = r.marketMetadata
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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
