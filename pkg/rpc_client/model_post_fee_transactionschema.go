/*
Stacks 2.0+ RPC API

This is the documentation for the `stacks-node` RPC interface. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rpc_client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PostFeeTransactionschema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostFeeTransactionschema{}

// PostFeeTransactionschema POST request for estimated fee
type PostFeeTransactionschema struct {
	TransactionPayload string `json:"transaction_payload"`
	EstimatedLen *int32 `json:"estimated_len,omitempty"`
}

type _PostFeeTransactionschema PostFeeTransactionschema

// NewPostFeeTransactionschema instantiates a new PostFeeTransactionschema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostFeeTransactionschema(transactionPayload string) *PostFeeTransactionschema {
	this := PostFeeTransactionschema{}
	this.TransactionPayload = transactionPayload
	return &this
}

// NewPostFeeTransactionschemaWithDefaults instantiates a new PostFeeTransactionschema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostFeeTransactionschemaWithDefaults() *PostFeeTransactionschema {
	this := PostFeeTransactionschema{}
	return &this
}

// GetTransactionPayload returns the TransactionPayload field value
func (o *PostFeeTransactionschema) GetTransactionPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionPayload
}

// GetTransactionPayloadOk returns a tuple with the TransactionPayload field value
// and a boolean to check if the value has been set.
func (o *PostFeeTransactionschema) GetTransactionPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionPayload, true
}

// SetTransactionPayload sets field value
func (o *PostFeeTransactionschema) SetTransactionPayload(v string) {
	o.TransactionPayload = v
}

// GetEstimatedLen returns the EstimatedLen field value if set, zero value otherwise.
func (o *PostFeeTransactionschema) GetEstimatedLen() int32 {
	if o == nil || IsNil(o.EstimatedLen) {
		var ret int32
		return ret
	}
	return *o.EstimatedLen
}

// GetEstimatedLenOk returns a tuple with the EstimatedLen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostFeeTransactionschema) GetEstimatedLenOk() (*int32, bool) {
	if o == nil || IsNil(o.EstimatedLen) {
		return nil, false
	}
	return o.EstimatedLen, true
}

// HasEstimatedLen returns a boolean if a field has been set.
func (o *PostFeeTransactionschema) HasEstimatedLen() bool {
	if o != nil && !IsNil(o.EstimatedLen) {
		return true
	}

	return false
}

// SetEstimatedLen gets a reference to the given int32 and assigns it to the EstimatedLen field.
func (o *PostFeeTransactionschema) SetEstimatedLen(v int32) {
	o.EstimatedLen = &v
}

func (o PostFeeTransactionschema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostFeeTransactionschema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transaction_payload"] = o.TransactionPayload
	if !IsNil(o.EstimatedLen) {
		toSerialize["estimated_len"] = o.EstimatedLen
	}
	return toSerialize, nil
}

func (o *PostFeeTransactionschema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transaction_payload",
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

	varPostFeeTransactionschema := _PostFeeTransactionschema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostFeeTransactionschema)

	if err != nil {
		return err
	}

	*o = PostFeeTransactionschema(varPostFeeTransactionschema)

	return err
}

type NullablePostFeeTransactionschema struct {
	value *PostFeeTransactionschema
	isSet bool
}

func (v NullablePostFeeTransactionschema) Get() *PostFeeTransactionschema {
	return v.value
}

func (v *NullablePostFeeTransactionschema) Set(val *PostFeeTransactionschema) {
	v.value = val
	v.isSet = true
}

func (v NullablePostFeeTransactionschema) IsSet() bool {
	return v.isSet
}

func (v *NullablePostFeeTransactionschema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostFeeTransactionschema(val *PostFeeTransactionschema) *NullablePostFeeTransactionschema {
	return &NullablePostFeeTransactionschema{value: val, isSet: true}
}

func (v NullablePostFeeTransactionschema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostFeeTransactionschema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


