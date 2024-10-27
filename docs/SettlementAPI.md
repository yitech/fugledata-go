# \SettlementAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SettlementDelete**](SettlementAPI.md#SettlementDelete) | **Delete** /settlement | 
[**SettlementGet**](SettlementAPI.md#SettlementGet) | **Get** /settlement | 
[**SettlementPatch**](SettlementAPI.md#SettlementPatch) | **Patch** /settlement | 
[**SettlementPost**](SettlementAPI.md#SettlementPost) | **Post** /settlement | 



## SettlementDelete

> SettlementDelete(ctx).Dt(dt).NextDay(nextDay).AfterNext(afterNext).Id(id).Prefer(prefer).Execute()



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
	dt := "dt_example" // string |  (optional)
	nextDay := "nextDay_example" // string |  (optional)
	afterNext := "afterNext_example" // string |  (optional)
	id := "id_example" // string |  (optional)
	prefer := "prefer_example" // string | Preference (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettlementAPI.SettlementDelete(context.Background()).Dt(dt).NextDay(nextDay).AfterNext(afterNext).Id(id).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettlementAPI.SettlementDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSettlementDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dt** | **string** |  | 
 **nextDay** | **string** |  | 
 **afterNext** | **string** |  | 
 **id** | **string** |  | 
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


## SettlementGet

> []Settlement SettlementGet(ctx).Dt(dt).NextDay(nextDay).AfterNext(afterNext).Id(id).Select_(select_).Order(order).Range_(range_).RangeUnit(rangeUnit).Offset(offset).Limit(limit).Prefer(prefer).Execute()



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
	dt := "dt_example" // string |  (optional)
	nextDay := "nextDay_example" // string |  (optional)
	afterNext := "afterNext_example" // string |  (optional)
	id := "id_example" // string |  (optional)
	select_ := "select__example" // string | Filtering Columns (optional)
	order := "order_example" // string | Ordering (optional)
	range_ := "range__example" // string | Limiting and Pagination (optional)
	rangeUnit := "rangeUnit_example" // string | Limiting and Pagination (optional) (default to "items")
	offset := "offset_example" // string | Limiting and Pagination (optional)
	limit := "limit_example" // string | Limiting and Pagination (optional)
	prefer := "prefer_example" // string | Preference (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettlementAPI.SettlementGet(context.Background()).Dt(dt).NextDay(nextDay).AfterNext(afterNext).Id(id).Select_(select_).Order(order).Range_(range_).RangeUnit(rangeUnit).Offset(offset).Limit(limit).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettlementAPI.SettlementGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettlementGet`: []Settlement
	fmt.Fprintf(os.Stdout, "Response from `SettlementAPI.SettlementGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSettlementGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dt** | **string** |  | 
 **nextDay** | **string** |  | 
 **afterNext** | **string** |  | 
 **id** | **string** |  | 
 **select_** | **string** | Filtering Columns | 
 **order** | **string** | Ordering | 
 **range_** | **string** | Limiting and Pagination | 
 **rangeUnit** | **string** | Limiting and Pagination | [default to &quot;items&quot;]
 **offset** | **string** | Limiting and Pagination | 
 **limit** | **string** | Limiting and Pagination | 
 **prefer** | **string** | Preference | 

### Return type

[**[]Settlement**](Settlement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.pgrst.object+json;nulls=stripped, application/vnd.pgrst.object+json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SettlementPatch

> SettlementPatch(ctx).Dt(dt).NextDay(nextDay).AfterNext(afterNext).Id(id).Prefer(prefer).Settlement(settlement).Execute()



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
	dt := "dt_example" // string |  (optional)
	nextDay := "nextDay_example" // string |  (optional)
	afterNext := "afterNext_example" // string |  (optional)
	id := "id_example" // string |  (optional)
	prefer := "prefer_example" // string | Preference (optional)
	settlement := *openapiclient.NewSettlement(time.Now(), int32(123)) // Settlement | settlement (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettlementAPI.SettlementPatch(context.Background()).Dt(dt).NextDay(nextDay).AfterNext(afterNext).Id(id).Prefer(prefer).Settlement(settlement).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettlementAPI.SettlementPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSettlementPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dt** | **string** |  | 
 **nextDay** | **string** |  | 
 **afterNext** | **string** |  | 
 **id** | **string** |  | 
 **prefer** | **string** | Preference | 
 **settlement** | [**Settlement**](Settlement.md) | settlement | 

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


## SettlementPost

> SettlementPost(ctx).Select_(select_).Prefer(prefer).Settlement(settlement).Execute()



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
	settlement := *openapiclient.NewSettlement(time.Now(), int32(123)) // Settlement | settlement (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettlementAPI.SettlementPost(context.Background()).Select_(select_).Prefer(prefer).Settlement(settlement).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettlementAPI.SettlementPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSettlementPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **string** | Filtering Columns | 
 **prefer** | **string** | Preference | 
 **settlement** | [**Settlement**](Settlement.md) | settlement | 

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

