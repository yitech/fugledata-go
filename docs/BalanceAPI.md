# \BalanceAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BalanceDelete**](BalanceAPI.md#BalanceDelete) | **Delete** /balance | 
[**BalanceGet**](BalanceAPI.md#BalanceGet) | **Get** /balance | 
[**BalancePatch**](BalanceAPI.md#BalancePatch) | **Patch** /balance | 
[**BalancePost**](BalanceAPI.md#BalancePost) | **Post** /balance | 



## BalanceDelete

> BalanceDelete(ctx).Dt(dt).Available(available).PresaveAmount(presaveAmount).Id(id).Prefer(prefer).Execute()



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
	available := "available_example" // string |  (optional)
	presaveAmount := "presaveAmount_example" // string |  (optional)
	id := "id_example" // string |  (optional)
	prefer := "prefer_example" // string | Preference (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BalanceAPI.BalanceDelete(context.Background()).Dt(dt).Available(available).PresaveAmount(presaveAmount).Id(id).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BalanceAPI.BalanceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBalanceDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dt** | **string** |  | 
 **available** | **string** |  | 
 **presaveAmount** | **string** |  | 
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


## BalanceGet

> []Balance BalanceGet(ctx).Dt(dt).Available(available).PresaveAmount(presaveAmount).Id(id).Select_(select_).Order(order).Range_(range_).RangeUnit(rangeUnit).Offset(offset).Limit(limit).Prefer(prefer).Execute()



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
	available := "available_example" // string |  (optional)
	presaveAmount := "presaveAmount_example" // string |  (optional)
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
	resp, r, err := apiClient.BalanceAPI.BalanceGet(context.Background()).Dt(dt).Available(available).PresaveAmount(presaveAmount).Id(id).Select_(select_).Order(order).Range_(range_).RangeUnit(rangeUnit).Offset(offset).Limit(limit).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BalanceAPI.BalanceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BalanceGet`: []Balance
	fmt.Fprintf(os.Stdout, "Response from `BalanceAPI.BalanceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBalanceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dt** | **string** |  | 
 **available** | **string** |  | 
 **presaveAmount** | **string** |  | 
 **id** | **string** |  | 
 **select_** | **string** | Filtering Columns | 
 **order** | **string** | Ordering | 
 **range_** | **string** | Limiting and Pagination | 
 **rangeUnit** | **string** | Limiting and Pagination | [default to &quot;items&quot;]
 **offset** | **string** | Limiting and Pagination | 
 **limit** | **string** | Limiting and Pagination | 
 **prefer** | **string** | Preference | 

### Return type

[**[]Balance**](Balance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.pgrst.object+json;nulls=stripped, application/vnd.pgrst.object+json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BalancePatch

> BalancePatch(ctx).Dt(dt).Available(available).PresaveAmount(presaveAmount).Id(id).Prefer(prefer).Balance(balance).Execute()



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
	available := "available_example" // string |  (optional)
	presaveAmount := "presaveAmount_example" // string |  (optional)
	id := "id_example" // string |  (optional)
	prefer := "prefer_example" // string | Preference (optional)
	balance := *openapiclient.NewBalance(time.Now(), int32(123)) // Balance | balance (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BalanceAPI.BalancePatch(context.Background()).Dt(dt).Available(available).PresaveAmount(presaveAmount).Id(id).Prefer(prefer).Balance(balance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BalanceAPI.BalancePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBalancePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dt** | **string** |  | 
 **available** | **string** |  | 
 **presaveAmount** | **string** |  | 
 **id** | **string** |  | 
 **prefer** | **string** | Preference | 
 **balance** | [**Balance**](Balance.md) | balance | 

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


## BalancePost

> BalancePost(ctx).Select_(select_).Prefer(prefer).Balance(balance).Execute()



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
	balance := *openapiclient.NewBalance(time.Now(), int32(123)) // Balance | balance (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BalanceAPI.BalancePost(context.Background()).Select_(select_).Prefer(prefer).Balance(balance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BalanceAPI.BalancePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBalancePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **string** | Filtering Columns | 
 **prefer** | **string** | Preference | 
 **balance** | [**Balance**](Balance.md) | balance | 

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

