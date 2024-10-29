# \InventoryAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InventoryDelete**](InventoryAPI.md#InventoryDelete) | **Delete** /inventory | 
[**InventoryGet**](InventoryAPI.md#InventoryGet) | **Get** /inventory | 
[**InventoryPatch**](InventoryAPI.md#InventoryPatch) | **Patch** /inventory | 
[**InventoryPost**](InventoryAPI.md#InventoryPost) | **Post** /inventory | 



## InventoryDelete

> InventoryDelete(ctx).Dt(dt).Symbol(symbol).ShareQty(shareQty).Id(id).Prefer(prefer).Execute()



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
	symbol := "symbol_example" // string |  (optional)
	shareQty := "shareQty_example" // string |  (optional)
	id := "id_example" // string |  (optional)
	prefer := "prefer_example" // string | Preference (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryAPI.InventoryDelete(context.Background()).Dt(dt).Symbol(symbol).ShareQty(shareQty).Id(id).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.InventoryDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dt** | **string** |  | 
 **symbol** | **string** |  | 
 **shareQty** | **string** |  | 
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


## InventoryGet

> []Inventory InventoryGet(ctx).Dt(dt).Symbol(symbol).ShareQty(shareQty).Id(id).Select_(select_).Order(order).Range_(range_).RangeUnit(rangeUnit).Offset(offset).Limit(limit).Prefer(prefer).Execute()



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
	symbol := "symbol_example" // string |  (optional)
	shareQty := "shareQty_example" // string |  (optional)
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
	resp, r, err := apiClient.InventoryAPI.InventoryGet(context.Background()).Dt(dt).Symbol(symbol).ShareQty(shareQty).Id(id).Select_(select_).Order(order).Range_(range_).RangeUnit(rangeUnit).Offset(offset).Limit(limit).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.InventoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryGet`: []Inventory
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.InventoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dt** | **string** |  | 
 **symbol** | **string** |  | 
 **shareQty** | **string** |  | 
 **id** | **string** |  | 
 **select_** | **string** | Filtering Columns | 
 **order** | **string** | Ordering | 
 **range_** | **string** | Limiting and Pagination | 
 **rangeUnit** | **string** | Limiting and Pagination | [default to &quot;items&quot;]
 **offset** | **string** | Limiting and Pagination | 
 **limit** | **string** | Limiting and Pagination | 
 **prefer** | **string** | Preference | 

### Return type

[**[]Inventory**](Inventory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.pgrst.object+json;nulls=stripped, application/vnd.pgrst.object+json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryPatch

> InventoryPatch(ctx).Dt(dt).Symbol(symbol).ShareQty(shareQty).Id(id).Prefer(prefer).Inventory(inventory).Execute()



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
	symbol := "symbol_example" // string |  (optional)
	shareQty := "shareQty_example" // string |  (optional)
	id := "id_example" // string |  (optional)
	prefer := "prefer_example" // string | Preference (optional)
	inventory := *openapiclient.NewInventory(time.Now(), "Symbol_example", int32(123), int32(123)) // Inventory | inventory (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryAPI.InventoryPatch(context.Background()).Dt(dt).Symbol(symbol).ShareQty(shareQty).Id(id).Prefer(prefer).Inventory(inventory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.InventoryPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dt** | **string** |  | 
 **symbol** | **string** |  | 
 **shareQty** | **string** |  | 
 **id** | **string** |  | 
 **prefer** | **string** | Preference | 
 **inventory** | [**Inventory**](Inventory.md) | inventory | 

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


## InventoryPost

> InventoryPost(ctx).Select_(select_).Prefer(prefer).Inventory(inventory).Execute()



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
	inventory := *openapiclient.NewInventory(time.Now(), "Symbol_example", int32(123), int32(123)) // Inventory | inventory (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryAPI.InventoryPost(context.Background()).Select_(select_).Prefer(prefer).Inventory(inventory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.InventoryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **string** | Filtering Columns | 
 **prefer** | **string** | Preference | 
 **inventory** | [**Inventory**](Inventory.md) | inventory | 

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

