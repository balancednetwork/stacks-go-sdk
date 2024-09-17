/*
Stacks Blockchain API

Welcome to the API reference overview for the [Stacks Blockchain API](https://docs.hiro.so/stacks-blockchain-api).        [Download Postman collection](https://hirosystems.github.io/stacks-blockchain-api/collection.json)

API version: v7.14.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stacks_blockchain_api_client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GetNftHistory200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNftHistory200Response{}

// GetNftHistory200Response List of Non-Fungible Token history events
type GetNftHistory200Response struct {
	Limit int32 `json:"limit"`
	Offset int32 `json:"offset"`
	Total int32 `json:"total"`
	Results []NonFungibleTokenHistoryEvent `json:"results"`
}

type _GetNftHistory200Response GetNftHistory200Response

// NewGetNftHistory200Response instantiates a new GetNftHistory200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNftHistory200Response(limit int32, offset int32, total int32, results []NonFungibleTokenHistoryEvent) *GetNftHistory200Response {
	this := GetNftHistory200Response{}
	this.Limit = limit
	this.Offset = offset
	this.Total = total
	this.Results = results
	return &this
}

// NewGetNftHistory200ResponseWithDefaults instantiates a new GetNftHistory200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNftHistory200ResponseWithDefaults() *GetNftHistory200Response {
	this := GetNftHistory200Response{}
	return &this
}

// GetLimit returns the Limit field value
func (o *GetNftHistory200Response) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *GetNftHistory200Response) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *GetNftHistory200Response) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *GetNftHistory200Response) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *GetNftHistory200Response) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *GetNftHistory200Response) SetOffset(v int32) {
	o.Offset = v
}

// GetTotal returns the Total field value
func (o *GetNftHistory200Response) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *GetNftHistory200Response) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *GetNftHistory200Response) SetTotal(v int32) {
	o.Total = v
}

// GetResults returns the Results field value
func (o *GetNftHistory200Response) GetResults() []NonFungibleTokenHistoryEvent {
	if o == nil {
		var ret []NonFungibleTokenHistoryEvent
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *GetNftHistory200Response) GetResultsOk() ([]NonFungibleTokenHistoryEvent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *GetNftHistory200Response) SetResults(v []NonFungibleTokenHistoryEvent) {
	o.Results = v
}

func (o GetNftHistory200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNftHistory200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["total"] = o.Total
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *GetNftHistory200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"limit",
		"offset",
		"total",
		"results",
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

	varGetNftHistory200Response := _GetNftHistory200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetNftHistory200Response)

	if err != nil {
		return err
	}

	*o = GetNftHistory200Response(varGetNftHistory200Response)

	return err
}

type NullableGetNftHistory200Response struct {
	value *GetNftHistory200Response
	isSet bool
}

func (v NullableGetNftHistory200Response) Get() *GetNftHistory200Response {
	return v.value
}

func (v *NullableGetNftHistory200Response) Set(val *GetNftHistory200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNftHistory200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNftHistory200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNftHistory200Response(val *GetNftHistory200Response) *NullableGetNftHistory200Response {
	return &NullableGetNftHistory200Response{value: val, isSet: true}
}

func (v NullableGetNftHistory200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNftHistory200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


