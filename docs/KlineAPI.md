# \KlineAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KlineDelete**](KlineAPI.md#KlineDelete) | **Delete** /kline | 
[**KlineGet**](KlineAPI.md#KlineGet) | **Get** /kline | 
[**KlinePatch**](KlineAPI.md#KlinePatch) | **Patch** /kline | 
[**KlinePost**](KlineAPI.md#KlinePost) | **Post** /kline | 



## KlineDelete

> KlineDelete(ctx).Id(id).Symbol(symbol).Dt(dt).Open(open).High(high).Low(low).Close(close).Volume(volume).Prefer(prefer).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yitech/fugledata-go"
)

func main() {
	id := "id_example" // string |  (optional)
	symbol := "symbol_example" // string |  (optional)
	dt := "dt_example" // string |  (optional)
	open := "open_example" // string |  (optional)
	high := "high_example" // string |  (optional)
	low := "low_example" // string |  (optional)
	close := "close_example" // string |  (optional)
	volume := "volume_example" // string |  (optional)
	prefer := "prefer_example" // string | Preference (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KlineAPI.KlineDelete(context.Background()).Id(id).Symbol(symbol).Dt(dt).Open(open).High(high).Low(low).Close(close).Volume(volume).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KlineAPI.KlineDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKlineDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **symbol** | **string** |  | 
 **dt** | **string** |  | 
 **open** | **string** |  | 
 **high** | **string** |  | 
 **low** | **string** |  | 
 **close** | **string** |  | 
 **volume** | **string** |  | 
 **prefer** | **string** | Preference | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KlineGet

> []Kline KlineGet(ctx).Id(id).Symbol(symbol).Dt(dt).Open(open).High(high).Low(low).Close(close).Volume(volume).Select_(select_).Order(order).Range_(range_).RangeUnit(rangeUnit).Offset(offset).Limit(limit).Prefer(prefer).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yitech/fugledata-go"
)

func main() {
	id := "id_example" // string |  (optional)
	symbol := "symbol_example" // string |  (optional)
	dt := "dt_example" // string |  (optional)
	open := "open_example" // string |  (optional)
	high := "high_example" // string |  (optional)
	low := "low_example" // string |  (optional)
	close := "close_example" // string |  (optional)
	volume := "volume_example" // string |  (optional)
	select_ := "select__example" // string | Filtering Columns (optional)
	order := "order_example" // string | Ordering (optional)
	range_ := "range__example" // string | Limiting and Pagination (optional)
	rangeUnit := "rangeUnit_example" // string | Limiting and Pagination (optional) (default to "items")
	offset := "offset_example" // string | Limiting and Pagination (optional)
	limit := "limit_example" // string | Limiting and Pagination (optional)
	prefer := "prefer_example" // string | Preference (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KlineAPI.KlineGet(context.Background()).Id(id).Symbol(symbol).Dt(dt).Open(open).High(high).Low(low).Close(close).Volume(volume).Select_(select_).Order(order).Range_(range_).RangeUnit(rangeUnit).Offset(offset).Limit(limit).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KlineAPI.KlineGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KlineGet`: []Kline
	fmt.Fprintf(os.Stdout, "Response from `KlineAPI.KlineGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKlineGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **symbol** | **string** |  | 
 **dt** | **string** |  | 
 **open** | **string** |  | 
 **high** | **string** |  | 
 **low** | **string** |  | 
 **close** | **string** |  | 
 **volume** | **string** |  | 
 **select_** | **string** | Filtering Columns | 
 **order** | **string** | Ordering | 
 **range_** | **string** | Limiting and Pagination | 
 **rangeUnit** | **string** | Limiting and Pagination | [default to &quot;items&quot;]
 **offset** | **string** | Limiting and Pagination | 
 **limit** | **string** | Limiting and Pagination | 
 **prefer** | **string** | Preference | 

### Return type

[**[]Kline**](Kline.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.pgrst.object+json;nulls=stripped, application/vnd.pgrst.object+json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KlinePatch

> KlinePatch(ctx).Id(id).Symbol(symbol).Dt(dt).Open(open).High(high).Low(low).Close(close).Volume(volume).Prefer(prefer).Kline(kline).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/yitech/fugledata-go"
)

func main() {
	id := "id_example" // string |  (optional)
	symbol := "symbol_example" // string |  (optional)
	dt := "dt_example" // string |  (optional)
	open := "open_example" // string |  (optional)
	high := "high_example" // string |  (optional)
	low := "low_example" // string |  (optional)
	close := "close_example" // string |  (optional)
	volume := "volume_example" // string |  (optional)
	prefer := "prefer_example" // string | Preference (optional)
	kline := *openapiclient.NewKline(int32(123), "Symbol_example", time.Now(), float32(123), float32(123), float32(123), float32(123), int32(123)) // Kline | kline (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KlineAPI.KlinePatch(context.Background()).Id(id).Symbol(symbol).Dt(dt).Open(open).High(high).Low(low).Close(close).Volume(volume).Prefer(prefer).Kline(kline).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KlineAPI.KlinePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKlinePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **symbol** | **string** |  | 
 **dt** | **string** |  | 
 **open** | **string** |  | 
 **high** | **string** |  | 
 **low** | **string** |  | 
 **close** | **string** |  | 
 **volume** | **string** |  | 
 **prefer** | **string** | Preference | 
 **kline** | [**Kline**](Kline.md) | kline | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/vnd.pgrst.object+json;nulls=stripped, application/vnd.pgrst.object+json, text/csv
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KlinePost

> KlinePost(ctx).Select_(select_).Prefer(prefer).Kline(kline).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/yitech/fugledata-go"
)

func main() {
	select_ := "select__example" // string | Filtering Columns (optional)
	prefer := "prefer_example" // string | Preference (optional)
	kline := *openapiclient.NewKline(int32(123), "Symbol_example", time.Now(), float32(123), float32(123), float32(123), float32(123), int32(123)) // Kline | kline (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KlineAPI.KlinePost(context.Background()).Select_(select_).Prefer(prefer).Kline(kline).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KlineAPI.KlinePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKlinePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **string** | Filtering Columns | 
 **prefer** | **string** | Preference | 
 **kline** | [**Kline**](Kline.md) | kline | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/vnd.pgrst.object+json;nulls=stripped, application/vnd.pgrst.object+json, text/csv
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

