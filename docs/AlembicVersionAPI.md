# \AlembicVersionAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlembicVersionDelete**](AlembicVersionAPI.md#AlembicVersionDelete) | **Delete** /alembic_version | 
[**AlembicVersionGet**](AlembicVersionAPI.md#AlembicVersionGet) | **Get** /alembic_version | 
[**AlembicVersionPatch**](AlembicVersionAPI.md#AlembicVersionPatch) | **Patch** /alembic_version | 
[**AlembicVersionPost**](AlembicVersionAPI.md#AlembicVersionPost) | **Post** /alembic_version | 



## AlembicVersionDelete

> AlembicVersionDelete(ctx).VersionNum(versionNum).Prefer(prefer).Execute()



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
	versionNum := "versionNum_example" // string |  (optional)
	prefer := "prefer_example" // string | Preference (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlembicVersionAPI.AlembicVersionDelete(context.Background()).VersionNum(versionNum).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlembicVersionAPI.AlembicVersionDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlembicVersionDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **versionNum** | **string** |  | 
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


## AlembicVersionGet

> []AlembicVersion AlembicVersionGet(ctx).VersionNum(versionNum).Select_(select_).Order(order).Range_(range_).RangeUnit(rangeUnit).Offset(offset).Limit(limit).Prefer(prefer).Execute()



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
	versionNum := "versionNum_example" // string |  (optional)
	select_ := "select__example" // string | Filtering Columns (optional)
	order := "order_example" // string | Ordering (optional)
	range_ := "range__example" // string | Limiting and Pagination (optional)
	rangeUnit := "rangeUnit_example" // string | Limiting and Pagination (optional) (default to "items")
	offset := "offset_example" // string | Limiting and Pagination (optional)
	limit := "limit_example" // string | Limiting and Pagination (optional)
	prefer := "prefer_example" // string | Preference (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlembicVersionAPI.AlembicVersionGet(context.Background()).VersionNum(versionNum).Select_(select_).Order(order).Range_(range_).RangeUnit(rangeUnit).Offset(offset).Limit(limit).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlembicVersionAPI.AlembicVersionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlembicVersionGet`: []AlembicVersion
	fmt.Fprintf(os.Stdout, "Response from `AlembicVersionAPI.AlembicVersionGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlembicVersionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **versionNum** | **string** |  | 
 **select_** | **string** | Filtering Columns | 
 **order** | **string** | Ordering | 
 **range_** | **string** | Limiting and Pagination | 
 **rangeUnit** | **string** | Limiting and Pagination | [default to &quot;items&quot;]
 **offset** | **string** | Limiting and Pagination | 
 **limit** | **string** | Limiting and Pagination | 
 **prefer** | **string** | Preference | 

### Return type

[**[]AlembicVersion**](AlembicVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.pgrst.object+json;nulls=stripped, application/vnd.pgrst.object+json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlembicVersionPatch

> AlembicVersionPatch(ctx).VersionNum(versionNum).Prefer(prefer).AlembicVersion(alembicVersion).Execute()



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
	versionNum := "versionNum_example" // string |  (optional)
	prefer := "prefer_example" // string | Preference (optional)
	alembicVersion := *openapiclient.NewAlembicVersion("VersionNum_example") // AlembicVersion | alembic_version (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlembicVersionAPI.AlembicVersionPatch(context.Background()).VersionNum(versionNum).Prefer(prefer).AlembicVersion(alembicVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlembicVersionAPI.AlembicVersionPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlembicVersionPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **versionNum** | **string** |  | 
 **prefer** | **string** | Preference | 
 **alembicVersion** | [**AlembicVersion**](AlembicVersion.md) | alembic_version | 

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


## AlembicVersionPost

> AlembicVersionPost(ctx).Select_(select_).Prefer(prefer).AlembicVersion(alembicVersion).Execute()



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
	alembicVersion := *openapiclient.NewAlembicVersion("VersionNum_example") // AlembicVersion | alembic_version (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlembicVersionAPI.AlembicVersionPost(context.Background()).Select_(select_).Prefer(prefer).AlembicVersion(alembicVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlembicVersionAPI.AlembicVersionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlembicVersionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **string** | Filtering Columns | 
 **prefer** | **string** | Preference | 
 **alembicVersion** | [**AlembicVersion**](AlembicVersion.md) | alembic_version | 

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

