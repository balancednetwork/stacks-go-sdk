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

// checks if the GetAccountDataschema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccountDataschema{}

// GetAccountDataschema GET request for account data
type GetAccountDataschema struct {
	Balance string `json:"balance"`
	Locked string `json:"locked"`
	UnlockHeight int32 `json:"unlock_height"`
	Nonce int32 `json:"nonce"`
	BalanceProof string `json:"balance_proof"`
	NonceProof string `json:"nonce_proof"`
}

type _GetAccountDataschema GetAccountDataschema

// NewGetAccountDataschema instantiates a new GetAccountDataschema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountDataschema(balance string, locked string, unlockHeight int32, nonce int32, balanceProof string, nonceProof string) *GetAccountDataschema {
	this := GetAccountDataschema{}
	this.Balance = balance
	this.Locked = locked
	this.UnlockHeight = unlockHeight
	this.Nonce = nonce
	this.BalanceProof = balanceProof
	this.NonceProof = nonceProof
	return &this
}

// NewGetAccountDataschemaWithDefaults instantiates a new GetAccountDataschema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountDataschemaWithDefaults() *GetAccountDataschema {
	this := GetAccountDataschema{}
	return &this
}

// GetBalance returns the Balance field value
func (o *GetAccountDataschema) GetBalance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *GetAccountDataschema) GetBalanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *GetAccountDataschema) SetBalance(v string) {
	o.Balance = v
}

// GetLocked returns the Locked field value
func (o *GetAccountDataschema) GetLocked() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locked
}

// GetLockedOk returns a tuple with the Locked field value
// and a boolean to check if the value has been set.
func (o *GetAccountDataschema) GetLockedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locked, true
}

// SetLocked sets field value
func (o *GetAccountDataschema) SetLocked(v string) {
	o.Locked = v
}

// GetUnlockHeight returns the UnlockHeight field value
func (o *GetAccountDataschema) GetUnlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UnlockHeight
}

// GetUnlockHeightOk returns a tuple with the UnlockHeight field value
// and a boolean to check if the value has been set.
func (o *GetAccountDataschema) GetUnlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnlockHeight, true
}

// SetUnlockHeight sets field value
func (o *GetAccountDataschema) SetUnlockHeight(v int32) {
	o.UnlockHeight = v
}

// GetNonce returns the Nonce field value
func (o *GetAccountDataschema) GetNonce() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *GetAccountDataschema) GetNonceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *GetAccountDataschema) SetNonce(v int32) {
	o.Nonce = v
}

// GetBalanceProof returns the BalanceProof field value
func (o *GetAccountDataschema) GetBalanceProof() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BalanceProof
}

// GetBalanceProofOk returns a tuple with the BalanceProof field value
// and a boolean to check if the value has been set.
func (o *GetAccountDataschema) GetBalanceProofOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BalanceProof, true
}

// SetBalanceProof sets field value
func (o *GetAccountDataschema) SetBalanceProof(v string) {
	o.BalanceProof = v
}

// GetNonceProof returns the NonceProof field value
func (o *GetAccountDataschema) GetNonceProof() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NonceProof
}

// GetNonceProofOk returns a tuple with the NonceProof field value
// and a boolean to check if the value has been set.
func (o *GetAccountDataschema) GetNonceProofOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NonceProof, true
}

// SetNonceProof sets field value
func (o *GetAccountDataschema) SetNonceProof(v string) {
	o.NonceProof = v
}

func (o GetAccountDataschema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccountDataschema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["balance"] = o.Balance
	toSerialize["locked"] = o.Locked
	toSerialize["unlock_height"] = o.UnlockHeight
	toSerialize["nonce"] = o.Nonce
	toSerialize["balance_proof"] = o.BalanceProof
	toSerialize["nonce_proof"] = o.NonceProof
	return toSerialize, nil
}

func (o *GetAccountDataschema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"balance",
		"locked",
		"unlock_height",
		"nonce",
		"balance_proof",
		"nonce_proof",
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

	varGetAccountDataschema := _GetAccountDataschema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetAccountDataschema)

	if err != nil {
		return err
	}

	*o = GetAccountDataschema(varGetAccountDataschema)

	return err
}

type NullableGetAccountDataschema struct {
	value *GetAccountDataschema
	isSet bool
}

func (v NullableGetAccountDataschema) Get() *GetAccountDataschema {
	return v.value
}

func (v *NullableGetAccountDataschema) Set(val *GetAccountDataschema) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountDataschema) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountDataschema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountDataschema(val *GetAccountDataschema) *NullableGetAccountDataschema {
	return &NullableGetAccountDataschema{value: val, isSet: true}
}

func (v NullableGetAccountDataschema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountDataschema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


