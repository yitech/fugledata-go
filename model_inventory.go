/*
standard public schema

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 12.2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fugledata

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Inventory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Inventory{}

// Inventory struct for Inventory
type Inventory struct {
	Dt string `json:"dt"`
	Symbol string `json:"symbol"`
	ShareQty int32 `json:"share_qty"`
	// Note: This is a Primary Key.<pk/>
	Id int32 `json:"id"`
}

type _Inventory Inventory

// NewInventory instantiates a new Inventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventory(dt string, symbol string, shareQty int32, id int32) *Inventory {
	this := Inventory{}
	this.Dt = dt
	this.Symbol = symbol
	this.ShareQty = shareQty
	this.Id = id
	return &this
}

// NewInventoryWithDefaults instantiates a new Inventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryWithDefaults() *Inventory {
	this := Inventory{}
	return &this
}

// GetDt returns the Dt field value
func (o *Inventory) GetDt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dt
}

// GetDtOk returns a tuple with the Dt field value
// and a boolean to check if the value has been set.
func (o *Inventory) GetDtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dt, true
}

// SetDt sets field value
func (o *Inventory) SetDt(v string) {
	o.Dt = v
}

// GetSymbol returns the Symbol field value
func (o *Inventory) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *Inventory) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *Inventory) SetSymbol(v string) {
	o.Symbol = v
}

// GetShareQty returns the ShareQty field value
func (o *Inventory) GetShareQty() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ShareQty
}

// GetShareQtyOk returns a tuple with the ShareQty field value
// and a boolean to check if the value has been set.
func (o *Inventory) GetShareQtyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShareQty, true
}

// SetShareQty sets field value
func (o *Inventory) SetShareQty(v int32) {
	o.ShareQty = v
}

// GetId returns the Id field value
func (o *Inventory) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Inventory) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Inventory) SetId(v int32) {
	o.Id = v
}

func (o Inventory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Inventory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dt"] = o.Dt
	toSerialize["symbol"] = o.Symbol
	toSerialize["share_qty"] = o.ShareQty
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *Inventory) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dt",
		"symbol",
		"share_qty",
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varInventory := _Inventory{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInventory)

	if err != nil {
		return err
	}

	*o = Inventory(varInventory)

	return err
}

type NullableInventory struct {
	value *Inventory
	isSet bool
}

func (v NullableInventory) Get() *Inventory {
	return v.value
}

func (v *NullableInventory) Set(val *Inventory) {
	v.value = val
	v.isSet = true
}

func (v NullableInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventory(val *Inventory) *NullableInventory {
	return &NullableInventory{value: val, isSet: true}
}

func (v NullableInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


