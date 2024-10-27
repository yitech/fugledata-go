# Balance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dt** | **string** |  | 
**Available** | Pointer to **int32** |  | [optional] 
**PresaveAmount** | Pointer to **int32** |  | [optional] 
**Id** | **int32** | Note: This is a Primary Key.&lt;pk/&gt; | 

## Methods

### NewBalance

`func NewBalance(dt string, id int32, ) *Balance`

NewBalance instantiates a new Balance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceWithDefaults

`func NewBalanceWithDefaults() *Balance`

NewBalanceWithDefaults instantiates a new Balance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDt

`func (o *Balance) GetDt() string`

GetDt returns the Dt field if non-nil, zero value otherwise.

### GetDtOk

`func (o *Balance) GetDtOk() (*string, bool)`

GetDtOk returns a tuple with the Dt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDt

`func (o *Balance) SetDt(v string)`

SetDt sets Dt field to given value.


### GetAvailable

`func (o *Balance) GetAvailable() int32`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *Balance) GetAvailableOk() (*int32, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *Balance) SetAvailable(v int32)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *Balance) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetPresaveAmount

`func (o *Balance) GetPresaveAmount() int32`

GetPresaveAmount returns the PresaveAmount field if non-nil, zero value otherwise.

### GetPresaveAmountOk

`func (o *Balance) GetPresaveAmountOk() (*int32, bool)`

GetPresaveAmountOk returns a tuple with the PresaveAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresaveAmount

`func (o *Balance) SetPresaveAmount(v int32)`

SetPresaveAmount sets PresaveAmount field to given value.

### HasPresaveAmount

`func (o *Balance) HasPresaveAmount() bool`

HasPresaveAmount returns a boolean if a field has been set.

### GetId

`func (o *Balance) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Balance) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Balance) SetId(v int32)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


