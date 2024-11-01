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


// KlineAPIService KlineAPI service
type KlineAPIService service

type ApiKlineDeleteRequest struct {
	ctx context.Context
	ApiService *KlineAPIService
	id *string
	symbol *string
	dt *string
	open *string
	high *string
	low *string
	close *string
	volume *string
	prefer *string
}

func (r ApiKlineDeleteRequest) Id(id string) ApiKlineDeleteRequest {
	r.id = &id
	return r
}

func (r ApiKlineDeleteRequest) Symbol(symbol string) ApiKlineDeleteRequest {
	r.symbol = &symbol
	return r
}

func (r ApiKlineDeleteRequest) Dt(dt string) ApiKlineDeleteRequest {
	r.dt = &dt
	return r
}

func (r ApiKlineDeleteRequest) Open(open string) ApiKlineDeleteRequest {
	r.open = &open
	return r
}

func (r ApiKlineDeleteRequest) High(high string) ApiKlineDeleteRequest {
	r.high = &high
	return r
}

func (r ApiKlineDeleteRequest) Low(low string) ApiKlineDeleteRequest {
	r.low = &low
	return r
}

func (r ApiKlineDeleteRequest) Close(close string) ApiKlineDeleteRequest {
	r.close = &close
	return r
}

func (r ApiKlineDeleteRequest) Volume(volume string) ApiKlineDeleteRequest {
	r.volume = &volume
	return r
}

// Preference
func (r ApiKlineDeleteRequest) Prefer(prefer string) ApiKlineDeleteRequest {
	r.prefer = &prefer
	return r
}

func (r ApiKlineDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.KlineDeleteExecute(r)
}

/*
KlineDelete Method for KlineDelete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiKlineDeleteRequest
*/
func (a *KlineAPIService) KlineDelete(ctx context.Context) ApiKlineDeleteRequest {
	return ApiKlineDeleteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *KlineAPIService) KlineDeleteExecute(r ApiKlineDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KlineAPIService.KlineDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/kline"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "", "")
	}
	if r.symbol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "", "")
	}
	if r.dt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dt", r.dt, "", "")
	}
	if r.open != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "open", r.open, "", "")
	}
	if r.high != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "high", r.high, "", "")
	}
	if r.low != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "low", r.low, "", "")
	}
	if r.close != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "close", r.close, "", "")
	}
	if r.volume != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "volume", r.volume, "", "")
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

type ApiKlineGetRequest struct {
	ctx context.Context
	ApiService *KlineAPIService
	id *string
	symbol *string
	dt *string
	open *string
	high *string
	low *string
	close *string
	volume *string
	select_ *string
	order *string
	range_ *string
	rangeUnit *string
	offset *string
	limit *string
	prefer *string
}

func (r ApiKlineGetRequest) Id(id string) ApiKlineGetRequest {
	r.id = &id
	return r
}

func (r ApiKlineGetRequest) Symbol(symbol string) ApiKlineGetRequest {
	r.symbol = &symbol
	return r
}

func (r ApiKlineGetRequest) Dt(dt string) ApiKlineGetRequest {
	r.dt = &dt
	return r
}

func (r ApiKlineGetRequest) Open(open string) ApiKlineGetRequest {
	r.open = &open
	return r
}

func (r ApiKlineGetRequest) High(high string) ApiKlineGetRequest {
	r.high = &high
	return r
}

func (r ApiKlineGetRequest) Low(low string) ApiKlineGetRequest {
	r.low = &low
	return r
}

func (r ApiKlineGetRequest) Close(close string) ApiKlineGetRequest {
	r.close = &close
	return r
}

func (r ApiKlineGetRequest) Volume(volume string) ApiKlineGetRequest {
	r.volume = &volume
	return r
}

// Filtering Columns
func (r ApiKlineGetRequest) Select_(select_ string) ApiKlineGetRequest {
	r.select_ = &select_
	return r
}

// Ordering
func (r ApiKlineGetRequest) Order(order string) ApiKlineGetRequest {
	r.order = &order
	return r
}

// Limiting and Pagination
func (r ApiKlineGetRequest) Range_(range_ string) ApiKlineGetRequest {
	r.range_ = &range_
	return r
}

// Limiting and Pagination
func (r ApiKlineGetRequest) RangeUnit(rangeUnit string) ApiKlineGetRequest {
	r.rangeUnit = &rangeUnit
	return r
}

// Limiting and Pagination
func (r ApiKlineGetRequest) Offset(offset string) ApiKlineGetRequest {
	r.offset = &offset
	return r
}

// Limiting and Pagination
func (r ApiKlineGetRequest) Limit(limit string) ApiKlineGetRequest {
	r.limit = &limit
	return r
}

// Preference
func (r ApiKlineGetRequest) Prefer(prefer string) ApiKlineGetRequest {
	r.prefer = &prefer
	return r
}

func (r ApiKlineGetRequest) Execute() ([]Kline, *http.Response, error) {
	return r.ApiService.KlineGetExecute(r)
}

/*
KlineGet Method for KlineGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiKlineGetRequest
*/
func (a *KlineAPIService) KlineGet(ctx context.Context) ApiKlineGetRequest {
	return ApiKlineGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Kline
func (a *KlineAPIService) KlineGetExecute(r ApiKlineGetRequest) ([]Kline, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Kline
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KlineAPIService.KlineGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/kline"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "", "")
	}
	if r.symbol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "", "")
	}
	if r.dt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dt", r.dt, "", "")
	}
	if r.open != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "open", r.open, "", "")
	}
	if r.high != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "high", r.high, "", "")
	}
	if r.low != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "low", r.low, "", "")
	}
	if r.close != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "close", r.close, "", "")
	}
	if r.volume != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "volume", r.volume, "", "")
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

type ApiKlinePatchRequest struct {
	ctx context.Context
	ApiService *KlineAPIService
	id *string
	symbol *string
	dt *string
	open *string
	high *string
	low *string
	close *string
	volume *string
	prefer *string
	kline *Kline
}

func (r ApiKlinePatchRequest) Id(id string) ApiKlinePatchRequest {
	r.id = &id
	return r
}

func (r ApiKlinePatchRequest) Symbol(symbol string) ApiKlinePatchRequest {
	r.symbol = &symbol
	return r
}

func (r ApiKlinePatchRequest) Dt(dt string) ApiKlinePatchRequest {
	r.dt = &dt
	return r
}

func (r ApiKlinePatchRequest) Open(open string) ApiKlinePatchRequest {
	r.open = &open
	return r
}

func (r ApiKlinePatchRequest) High(high string) ApiKlinePatchRequest {
	r.high = &high
	return r
}

func (r ApiKlinePatchRequest) Low(low string) ApiKlinePatchRequest {
	r.low = &low
	return r
}

func (r ApiKlinePatchRequest) Close(close string) ApiKlinePatchRequest {
	r.close = &close
	return r
}

func (r ApiKlinePatchRequest) Volume(volume string) ApiKlinePatchRequest {
	r.volume = &volume
	return r
}

// Preference
func (r ApiKlinePatchRequest) Prefer(prefer string) ApiKlinePatchRequest {
	r.prefer = &prefer
	return r
}

// kline
func (r ApiKlinePatchRequest) Kline(kline Kline) ApiKlinePatchRequest {
	r.kline = &kline
	return r
}

func (r ApiKlinePatchRequest) Execute() (*http.Response, error) {
	return r.ApiService.KlinePatchExecute(r)
}

/*
KlinePatch Method for KlinePatch

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiKlinePatchRequest
*/
func (a *KlineAPIService) KlinePatch(ctx context.Context) ApiKlinePatchRequest {
	return ApiKlinePatchRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *KlineAPIService) KlinePatchExecute(r ApiKlinePatchRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KlineAPIService.KlinePatch")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/kline"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "", "")
	}
	if r.symbol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "", "")
	}
	if r.dt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dt", r.dt, "", "")
	}
	if r.open != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "open", r.open, "", "")
	}
	if r.high != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "high", r.high, "", "")
	}
	if r.low != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "low", r.low, "", "")
	}
	if r.close != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "close", r.close, "", "")
	}
	if r.volume != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "volume", r.volume, "", "")
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
	localVarPostBody = r.kline
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

type ApiKlinePostRequest struct {
	ctx context.Context
	ApiService *KlineAPIService
	select_ *string
	prefer *string
	kline *Kline
}

// Filtering Columns
func (r ApiKlinePostRequest) Select_(select_ string) ApiKlinePostRequest {
	r.select_ = &select_
	return r
}

// Preference
func (r ApiKlinePostRequest) Prefer(prefer string) ApiKlinePostRequest {
	r.prefer = &prefer
	return r
}

// kline
func (r ApiKlinePostRequest) Kline(kline Kline) ApiKlinePostRequest {
	r.kline = &kline
	return r
}

func (r ApiKlinePostRequest) Execute() (*http.Response, error) {
	return r.ApiService.KlinePostExecute(r)
}

/*
KlinePost Method for KlinePost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiKlinePostRequest
*/
func (a *KlineAPIService) KlinePost(ctx context.Context) ApiKlinePostRequest {
	return ApiKlinePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *KlineAPIService) KlinePostExecute(r ApiKlinePostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KlineAPIService.KlinePost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/kline"

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
	localVarPostBody = r.kline
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
