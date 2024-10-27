# MarketMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Note: This is a Primary Key.&lt;pk/&gt; | 
**Symbol** | **string** |  | 
**Category** | Pointer to **string** |  | [optional] 

## Methods

### NewMarketMetadata

`func NewMarketMetadata(id int32, symbol string, ) *MarketMetadata`

NewMarketMetadata instantiates a new MarketMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketMetadataWithDefaults

`func NewMarketMetadataWithDefaults() *MarketMetadata`

NewMarketMetadataWithDefaults instantiates a new MarketMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MarketMetadata) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketMetadata) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketMetadata) SetId(v int32)`

SetId sets Id field to given value.


### GetSymbol

`func (o *MarketMetadata) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MarketMetadata) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MarketMetadata) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetCategory

`func (o *MarketMetadata) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *MarketMetadata) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *MarketMetadata) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *MarketMetadata) HasCategory() bool`

HasCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


