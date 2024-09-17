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

// checks if the RunFaucetStx4XXResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunFaucetStx4XXResponse{}

// RunFaucetStx4XXResponse struct for RunFaucetStx4XXResponse
type RunFaucetStx4XXResponse struct {
	// Indicates if the faucet call was successful
	Success bool `json:"success"`
	// Error message
	Error string `json:"error"`
	Help *string `json:"help,omitempty"`
}

type _RunFaucetStx4XXResponse RunFaucetStx4XXResponse

// NewRunFaucetStx4XXResponse instantiates a new RunFaucetStx4XXResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunFaucetStx4XXResponse(success bool, error_ string) *RunFaucetStx4XXResponse {
	this := RunFaucetStx4XXResponse{}
	this.Success = success
	this.Error = error_
	return &this
}

// NewRunFaucetStx4XXResponseWithDefaults instantiates a new RunFaucetStx4XXResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunFaucetStx4XXResponseWithDefaults() *RunFaucetStx4XXResponse {
	this := RunFaucetStx4XXResponse{}
	return &this
}

// GetSuccess returns the Success field value
func (o *RunFaucetStx4XXResponse) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *RunFaucetStx4XXResponse) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *RunFaucetStx4XXResponse) SetSuccess(v bool) {
	o.Success = v
}

// GetError returns the Error field value
func (o *RunFaucetStx4XXResponse) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *RunFaucetStx4XXResponse) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *RunFaucetStx4XXResponse) SetError(v string) {
	o.Error = v
}

// GetHelp returns the Help field value if set, zero value otherwise.
func (o *RunFaucetStx4XXResponse) GetHelp() string {
	if o == nil || IsNil(o.Help) {
		var ret string
		return ret
	}
	return *o.Help
}

// GetHelpOk returns a tuple with the Help field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunFaucetStx4XXResponse) GetHelpOk() (*string, bool) {
	if o == nil || IsNil(o.Help) {
		return nil, false
	}
	return o.Help, true
}

// HasHelp returns a boolean if a field has been set.
func (o *RunFaucetStx4XXResponse) HasHelp() bool {
	if o != nil && !IsNil(o.Help) {
		return true
	}

	return false
}

// SetHelp gets a reference to the given string and assigns it to the Help field.
func (o *RunFaucetStx4XXResponse) SetHelp(v string) {
	o.Help = &v
}

func (o RunFaucetStx4XXResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunFaucetStx4XXResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	toSerialize["error"] = o.Error
	if !IsNil(o.Help) {
		toSerialize["help"] = o.Help
	}
	return toSerialize, nil
}

func (o *RunFaucetStx4XXResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"success",
		"error",
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

	varRunFaucetStx4XXResponse := _RunFaucetStx4XXResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRunFaucetStx4XXResponse)

	if err != nil {
		return err
	}

	*o = RunFaucetStx4XXResponse(varRunFaucetStx4XXResponse)

	return err
}

type NullableRunFaucetStx4XXResponse struct {
	value *RunFaucetStx4XXResponse
	isSet bool
}

func (v NullableRunFaucetStx4XXResponse) Get() *RunFaucetStx4XXResponse {
	return v.value
}

func (v *NullableRunFaucetStx4XXResponse) Set(val *RunFaucetStx4XXResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRunFaucetStx4XXResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRunFaucetStx4XXResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunFaucetStx4XXResponse(val *RunFaucetStx4XXResponse) *NullableRunFaucetStx4XXResponse {
	return &NullableRunFaucetStx4XXResponse{value: val, isSet: true}
}

func (v NullableRunFaucetStx4XXResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunFaucetStx4XXResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


