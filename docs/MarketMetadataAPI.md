# \MarketMetadataAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarketMetadataDelete**](MarketMetadataAPI.md#MarketMetadataDelete) | **Delete** /market_metadata | 
[**MarketMetadataGet**](MarketMetadataAPI.md#MarketMetadataGet) | **Get** /market_metadata | 
[**MarketMetadataPatch**](MarketMetadataAPI.md#MarketMetadataPatch) | **Patch** /market_metadata | 
[**MarketMetadataPost**](MarketMetadataAPI.md#MarketMetadataPost) | **Post** /market_metadata | 



## MarketMetadataDelete

> MarketMetadataDelete(ctx).Id(id).Symbol(symbol).Category(category).Prefer(prefer).Execute()



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
	category := "category_example" // string |  (optional)
	prefer := "prefer_example" // string | Preference (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MarketMetadataAPI.MarketMetadataDelete(context.Background()).Id(id).Symbol(symbol).Category(category).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketMetadataAPI.MarketMetadataDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketMetadataDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **symbol** | **string** |  | 
 **category** | **string** |  | 
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


## MarketMetadataGet

> []MarketMetadata MarketMetadataGet(ctx).Id(id).Symbol(symbol).Category(category).Select_(select_).Order(order).Range_(range_).RangeUnit(rangeUnit).Offset(offset).Limit(limit).Prefer(prefer).Execute()



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
	category := "category_example" // string |  (optional)
	select_ := "select__example" // string | Filtering Columns (optional)
	order := "order_example" // string | Ordering (optional)
	range_ := "range__example" // string | Limiting and Pagination (optional)
	rangeUnit := "rangeUnit_example" // string | Limiting and Pagination (optional) (default to "items")
	offset := "offset_example" // string | Limiting and Pagination (optional)
	limit := "limit_example" // string | Limiting and Pagination (optional)
	prefer := "prefer_example" // string | Preference (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketMetadataAPI.MarketMetadataGet(context.Background()).Id(id).Symbol(symbol).Category(category).Select_(select_).Order(order).Range_(range_).RangeUnit(rangeUnit).Offset(offset).Limit(limit).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketMetadataAPI.MarketMetadataGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketMetadataGet`: []MarketMetadata
	fmt.Fprintf(os.Stdout, "Response from `MarketMetadataAPI.MarketMetadataGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketMetadataGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **symbol** | **string** |  | 
 **category** | **string** |  | 
 **select_** | **string** | Filtering Columns | 
 **order** | **string** | Ordering | 
 **range_** | **string** | Limiting and Pagination | 
 **rangeUnit** | **string** | Limiting and Pagination | [default to &quot;items&quot;]
 **offset** | **string** | Limiting and Pagination | 
 **limit** | **string** | Limiting and Pagination | 
 **prefer** | **string** | Preference | 

### Return type

[**[]MarketMetadata**](MarketMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.pgrst.object+json;nulls=stripped, application/vnd.pgrst.object+json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketMetadataPatch

> MarketMetadataPatch(ctx).Id(id).Symbol(symbol).Category(category).Prefer(prefer).MarketMetadata(marketMetadata).Execute()



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
	category := "category_example" // string |  (optional)
	prefer := "prefer_example" // string | Preference (optional)
	marketMetadata := *openapiclient.NewMarketMetadata(int32(123), "Symbol_example") // MarketMetadata | market_metadata (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MarketMetadataAPI.MarketMetadataPatch(context.Background()).Id(id).Symbol(symbol).Category(category).Prefer(prefer).MarketMetadata(marketMetadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketMetadataAPI.MarketMetadataPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketMetadataPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **symbol** | **string** |  | 
 **category** | **string** |  | 
 **prefer** | **string** | Preference | 
 **marketMetadata** | [**MarketMetadata**](MarketMetadata.md) | market_metadata | 

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


## MarketMetadataPost

> MarketMetadataPost(ctx).Select_(select_).Prefer(prefer).MarketMetadata(marketMetadata).Execute()



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
	select_ := "select__example" // string | Filtering Columns (optional)
	prefer := "prefer_example" // string | Preference (optional)
	marketMetadata := *openapiclient.NewMarketMetadata(int32(123), "Symbol_example") // MarketMetadata | market_metadata (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MarketMetadataAPI.MarketMetadataPost(context.Background()).Select_(select_).Prefer(prefer).MarketMetadata(marketMetadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketMetadataAPI.MarketMetadataPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketMetadataPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **string** | Filtering Columns | 
 **prefer** | **string** | Preference | 
 **marketMetadata** | [**MarketMetadata**](MarketMetadata.md) | market_metadata | 

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

