# Inventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dt** | **string** |  | 
**Symbol** | **string** |  | 
**ShareQty** | **int32** |  | 
**Id** | **int32** | Note: This is a Primary Key.&lt;pk/&gt; | 

## Methods

### NewInventory

`func NewInventory(dt string, symbol string, shareQty int32, id int32, ) *Inventory`

NewInventory instantiates a new Inventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryWithDefaults

`func NewInventoryWithDefaults() *Inventory`

NewInventoryWithDefaults instantiates a new Inventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDt

`func (o *Inventory) GetDt() string`

GetDt returns the Dt field if non-nil, zero value otherwise.

### GetDtOk

`func (o *Inventory) GetDtOk() (*string, bool)`

GetDtOk returns a tuple with the Dt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDt

`func (o *Inventory) SetDt(v string)`

SetDt sets Dt field to given value.


### GetSymbol

`func (o *Inventory) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Inventory) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Inventory) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetShareQty

`func (o *Inventory) GetShareQty() int32`

GetShareQty returns the ShareQty field if non-nil, zero value otherwise.

### GetShareQtyOk

`func (o *Inventory) GetShareQtyOk() (*int32, bool)`

GetShareQtyOk returns a tuple with the ShareQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareQty

`func (o *Inventory) SetShareQty(v int32)`

SetShareQty sets ShareQty field to given value.


### GetId

`func (o *Inventory) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Inventory) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Inventory) SetId(v int32)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


