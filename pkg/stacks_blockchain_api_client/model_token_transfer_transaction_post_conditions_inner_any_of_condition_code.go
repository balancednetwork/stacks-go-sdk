/*
Stacks Blockchain API

Welcome to the API reference overview for the [Stacks Blockchain API](https://docs.hiro.so/stacks-blockchain-api).        [Download Postman collection](https://hirosystems.github.io/stacks-blockchain-api/collection.json)

API version: v7.14.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stacks_blockchain_api_client

import (
	"encoding/json"
	"fmt"
)


// TokenTransferTransactionPostConditionsInnerAnyOfConditionCode struct for TokenTransferTransactionPostConditionsInnerAnyOfConditionCode
type TokenTransferTransactionPostConditionsInnerAnyOfConditionCode struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TokenTransferTransactionPostConditionsInnerAnyOfConditionCode) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into String
	err = json.Unmarshal(data, &dst.String);
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(TokenTransferTransactionPostConditionsInnerAnyOfConditionCode)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TokenTransferTransactionPostConditionsInnerAnyOfConditionCode) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableTokenTransferTransactionPostConditionsInnerAnyOfConditionCode struct {
	value *TokenTransferTransactionPostConditionsInnerAnyOfConditionCode
	isSet bool
}

func (v NullableTokenTransferTransactionPostConditionsInnerAnyOfConditionCode) Get() *TokenTransferTransactionPostConditionsInnerAnyOfConditionCode {
	return v.value
}

func (v *NullableTokenTransferTransactionPostConditionsInnerAnyOfConditionCode) Set(val *TokenTransferTransactionPostConditionsInnerAnyOfConditionCode) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenTransferTransactionPostConditionsInnerAnyOfConditionCode) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenTransferTransactionPostConditionsInnerAnyOfConditionCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenTransferTransactionPostConditionsInnerAnyOfConditionCode(val *TokenTransferTransactionPostConditionsInnerAnyOfConditionCode) *NullableTokenTransferTransactionPostConditionsInnerAnyOfConditionCode {
	return &NullableTokenTransferTransactionPostConditionsInnerAnyOfConditionCode{value: val, isSet: true}
}

func (v NullableTokenTransferTransactionPostConditionsInnerAnyOfConditionCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenTransferTransactionPostConditionsInnerAnyOfConditionCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


