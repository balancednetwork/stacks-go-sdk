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

// checks if the TokenTransferTransactionPostConditionsInnerAnyOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenTransferTransactionPostConditionsInnerAnyOf{}

// TokenTransferTransactionPostConditionsInnerAnyOf struct for TokenTransferTransactionPostConditionsInnerAnyOf
type TokenTransferTransactionPostConditionsInnerAnyOf struct {
	Principal TokenTransferTransactionPostConditionsInnerAnyOfPrincipal `json:"principal"`
	ConditionCode TokenTransferTransactionPostConditionsInnerAnyOfConditionCode `json:"condition_code"`
	Amount string `json:"amount"`
	Type string `json:"type"`
}

type _TokenTransferTransactionPostConditionsInnerAnyOf TokenTransferTransactionPostConditionsInnerAnyOf

// NewTokenTransferTransactionPostConditionsInnerAnyOf instantiates a new TokenTransferTransactionPostConditionsInnerAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenTransferTransactionPostConditionsInnerAnyOf(principal TokenTransferTransactionPostConditionsInnerAnyOfPrincipal, conditionCode TokenTransferTransactionPostConditionsInnerAnyOfConditionCode, amount string, type_ string) *TokenTransferTransactionPostConditionsInnerAnyOf {
	this := TokenTransferTransactionPostConditionsInnerAnyOf{}
	this.Principal = principal
	this.ConditionCode = conditionCode
	this.Amount = amount
	this.Type = type_
	return &this
}

// NewTokenTransferTransactionPostConditionsInnerAnyOfWithDefaults instantiates a new TokenTransferTransactionPostConditionsInnerAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenTransferTransactionPostConditionsInnerAnyOfWithDefaults() *TokenTransferTransactionPostConditionsInnerAnyOf {
	this := TokenTransferTransactionPostConditionsInnerAnyOf{}
	return &this
}

// GetPrincipal returns the Principal field value
func (o *TokenTransferTransactionPostConditionsInnerAnyOf) GetPrincipal() TokenTransferTransactionPostConditionsInnerAnyOfPrincipal {
	if o == nil {
		var ret TokenTransferTransactionPostConditionsInnerAnyOfPrincipal
		return ret
	}

	return o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransactionPostConditionsInnerAnyOf) GetPrincipalOk() (*TokenTransferTransactionPostConditionsInnerAnyOfPrincipal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Principal, true
}

// SetPrincipal sets field value
func (o *TokenTransferTransactionPostConditionsInnerAnyOf) SetPrincipal(v TokenTransferTransactionPostConditionsInnerAnyOfPrincipal) {
	o.Principal = v
}

// GetConditionCode returns the ConditionCode field value
func (o *TokenTransferTransactionPostConditionsInnerAnyOf) GetConditionCode() TokenTransferTransactionPostConditionsInnerAnyOfConditionCode {
	if o == nil {
		var ret TokenTransferTransactionPostConditionsInnerAnyOfConditionCode
		return ret
	}

	return o.ConditionCode
}

// GetConditionCodeOk returns a tuple with the ConditionCode field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransactionPostConditionsInnerAnyOf) GetConditionCodeOk() (*TokenTransferTransactionPostConditionsInnerAnyOfConditionCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConditionCode, true
}

// SetConditionCode sets field value
func (o *TokenTransferTransactionPostConditionsInnerAnyOf) SetConditionCode(v TokenTransferTransactionPostConditionsInnerAnyOfConditionCode) {
	o.ConditionCode = v
}

// GetAmount returns the Amount field value
func (o *TokenTransferTransactionPostConditionsInnerAnyOf) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransactionPostConditionsInnerAnyOf) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TokenTransferTransactionPostConditionsInnerAnyOf) SetAmount(v string) {
	o.Amount = v
}

// GetType returns the Type field value
func (o *TokenTransferTransactionPostConditionsInnerAnyOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransactionPostConditionsInnerAnyOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TokenTransferTransactionPostConditionsInnerAnyOf) SetType(v string) {
	o.Type = v
}

func (o TokenTransferTransactionPostConditionsInnerAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenTransferTransactionPostConditionsInnerAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["principal"] = o.Principal
	toSerialize["condition_code"] = o.ConditionCode
	toSerialize["amount"] = o.Amount
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *TokenTransferTransactionPostConditionsInnerAnyOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"principal",
		"condition_code",
		"amount",
		"type",
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

	varTokenTransferTransactionPostConditionsInnerAnyOf := _TokenTransferTransactionPostConditionsInnerAnyOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTokenTransferTransactionPostConditionsInnerAnyOf)

	if err != nil {
		return err
	}

	*o = TokenTransferTransactionPostConditionsInnerAnyOf(varTokenTransferTransactionPostConditionsInnerAnyOf)

	return err
}

type NullableTokenTransferTransactionPostConditionsInnerAnyOf struct {
	value *TokenTransferTransactionPostConditionsInnerAnyOf
	isSet bool
}

func (v NullableTokenTransferTransactionPostConditionsInnerAnyOf) Get() *TokenTransferTransactionPostConditionsInnerAnyOf {
	return v.value
}

func (v *NullableTokenTransferTransactionPostConditionsInnerAnyOf) Set(val *TokenTransferTransactionPostConditionsInnerAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenTransferTransactionPostConditionsInnerAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenTransferTransactionPostConditionsInnerAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenTransferTransactionPostConditionsInnerAnyOf(val *TokenTransferTransactionPostConditionsInnerAnyOf) *NullableTokenTransferTransactionPostConditionsInnerAnyOf {
	return &NullableTokenTransferTransactionPostConditionsInnerAnyOf{value: val, isSet: true}
}

func (v NullableTokenTransferTransactionPostConditionsInnerAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenTransferTransactionPostConditionsInnerAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


