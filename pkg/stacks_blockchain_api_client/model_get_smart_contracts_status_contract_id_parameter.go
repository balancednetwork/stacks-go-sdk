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


// GetSmartContractsStatusContractIdParameter struct for GetSmartContractsStatusContractIdParameter
type GetSmartContractsStatusContractIdParameter struct {
	ArrayOfString *[]string
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *GetSmartContractsStatusContractIdParameter) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ArrayOfString
	err = json.Unmarshal(data, &dst.ArrayOfString);
	if err == nil {
		jsonArrayOfString, _ := json.Marshal(dst.ArrayOfString)
		if string(jsonArrayOfString) == "{}" { // empty struct
			dst.ArrayOfString = nil
		} else {
			return nil // data stored in dst.ArrayOfString, return on the first match
		}
	} else {
		dst.ArrayOfString = nil
	}

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

	return fmt.Errorf("data failed to match schemas in anyOf(GetSmartContractsStatusContractIdParameter)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *GetSmartContractsStatusContractIdParameter) MarshalJSON() ([]byte, error) {
	if src.ArrayOfString != nil {
		return json.Marshal(&src.ArrayOfString)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableGetSmartContractsStatusContractIdParameter struct {
	value *GetSmartContractsStatusContractIdParameter
	isSet bool
}

func (v NullableGetSmartContractsStatusContractIdParameter) Get() *GetSmartContractsStatusContractIdParameter {
	return v.value
}

func (v *NullableGetSmartContractsStatusContractIdParameter) Set(val *GetSmartContractsStatusContractIdParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSmartContractsStatusContractIdParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSmartContractsStatusContractIdParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSmartContractsStatusContractIdParameter(val *GetSmartContractsStatusContractIdParameter) *NullableGetSmartContractsStatusContractIdParameter {
	return &NullableGetSmartContractsStatusContractIdParameter{value: val, isSet: true}
}

func (v NullableGetSmartContractsStatusContractIdParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSmartContractsStatusContractIdParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


