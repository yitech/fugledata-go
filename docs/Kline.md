# Kline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Note: This is a Primary Key.&lt;pk/&gt; | 
**Symbol** | **string** |  | 
**Dt** | **string** |  | 
**Open** | **float32** |  | 
**High** | **float32** |  | 
**Low** | **float32** |  | 
**Close** | **float32** |  | 
**Volume** | **int32** |  | 

## Methods

### NewKline

`func NewKline(id int32, symbol string, dt string, open float32, high float32, low float32, close float32, volume int32, ) *Kline`

NewKline instantiates a new Kline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKlineWithDefaults

`func NewKlineWithDefaults() *Kline`

NewKlineWithDefaults instantiates a new Kline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Kline) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Kline) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Kline) SetId(v int32)`

SetId sets Id field to given value.


### GetSymbol

`func (o *Kline) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Kline) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Kline) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetDt

`func (o *Kline) GetDt() string`

GetDt returns the Dt field if non-nil, zero value otherwise.

### GetDtOk

`func (o *Kline) GetDtOk() (*string, bool)`

GetDtOk returns a tuple with the Dt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDt

`func (o *Kline) SetDt(v string)`

SetDt sets Dt field to given value.


### GetOpen

`func (o *Kline) GetOpen() float32`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *Kline) GetOpenOk() (*float32, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *Kline) SetOpen(v float32)`

SetOpen sets Open field to given value.


### GetHigh

`func (o *Kline) GetHigh() float32`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *Kline) GetHighOk() (*float32, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *Kline) SetHigh(v float32)`

SetHigh sets High field to given value.


### GetLow

`func (o *Kline) GetLow() float32`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *Kline) GetLowOk() (*float32, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *Kline) SetLow(v float32)`

SetLow sets Low field to given value.


### GetClose

`func (o *Kline) GetClose() float32`

GetClose returns the Close field if non-nil, zero value otherwise.

### GetCloseOk

`func (o *Kline) GetCloseOk() (*float32, bool)`

GetCloseOk returns a tuple with the Close field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClose

`func (o *Kline) SetClose(v float32)`

SetClose sets Close field to given value.


### GetVolume

`func (o *Kline) GetVolume() int32`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *Kline) GetVolumeOk() (*int32, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *Kline) SetVolume(v int32)`

SetVolume sets Volume field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


