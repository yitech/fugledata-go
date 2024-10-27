# Settlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dt** | **string** |  | 
**NextDay** | Pointer to **int32** |  | [optional] 
**AfterNext** | Pointer to **int32** |  | [optional] 
**Id** | **int32** | Note: This is a Primary Key.&lt;pk/&gt; | 

## Methods

### NewSettlement

`func NewSettlement(dt string, id int32, ) *Settlement`

NewSettlement instantiates a new Settlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettlementWithDefaults

`func NewSettlementWithDefaults() *Settlement`

NewSettlementWithDefaults instantiates a new Settlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDt

`func (o *Settlement) GetDt() string`

GetDt returns the Dt field if non-nil, zero value otherwise.

### GetDtOk

`func (o *Settlement) GetDtOk() (*string, bool)`

GetDtOk returns a tuple with the Dt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDt

`func (o *Settlement) SetDt(v string)`

SetDt sets Dt field to given value.


### GetNextDay

`func (o *Settlement) GetNextDay() int32`

GetNextDay returns the NextDay field if non-nil, zero value otherwise.

### GetNextDayOk

`func (o *Settlement) GetNextDayOk() (*int32, bool)`

GetNextDayOk returns a tuple with the NextDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextDay

`func (o *Settlement) SetNextDay(v int32)`

SetNextDay sets NextDay field to given value.

### HasNextDay

`func (o *Settlement) HasNextDay() bool`

HasNextDay returns a boolean if a field has been set.

### GetAfterNext

`func (o *Settlement) GetAfterNext() int32`

GetAfterNext returns the AfterNext field if non-nil, zero value otherwise.

### GetAfterNextOk

`func (o *Settlement) GetAfterNextOk() (*int32, bool)`

GetAfterNextOk returns a tuple with the AfterNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterNext

`func (o *Settlement) SetAfterNext(v int32)`

SetAfterNext sets AfterNext field to given value.

### HasAfterNext

`func (o *Settlement) HasAfterNext() bool`

HasAfterNext returns a boolean if a field has been set.

### GetId

`func (o *Settlement) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Settlement) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Settlement) SetId(v int32)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


