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

// checks if the NonFungibleTokenMintAnyOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NonFungibleTokenMintAnyOf{}

// NonFungibleTokenMintAnyOf struct for NonFungibleTokenMintAnyOf
type NonFungibleTokenMintAnyOf struct {
	Recipient *string `json:"recipient,omitempty"`
	EventIndex int32 `json:"event_index"`
	Value NonFungibleTokenHoldingWithTxIdValue `json:"value"`
	Tx GetTransactionList200ResponseResultsInner `json:"tx"`
}

type _NonFungibleTokenMintAnyOf NonFungibleTokenMintAnyOf

// NewNonFungibleTokenMintAnyOf instantiates a new NonFungibleTokenMintAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonFungibleTokenMintAnyOf(eventIndex int32, value NonFungibleTokenHoldingWithTxIdValue, tx GetTransactionList200ResponseResultsInner) *NonFungibleTokenMintAnyOf {
	this := NonFungibleTokenMintAnyOf{}
	this.EventIndex = eventIndex
	this.Value = value
	this.Tx = tx
	return &this
}

// NewNonFungibleTokenMintAnyOfWithDefaults instantiates a new NonFungibleTokenMintAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonFungibleTokenMintAnyOfWithDefaults() *NonFungibleTokenMintAnyOf {
	this := NonFungibleTokenMintAnyOf{}
	return &this
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *NonFungibleTokenMintAnyOf) GetRecipient() string {
	if o == nil || IsNil(o.Recipient) {
		var ret string
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenMintAnyOf) GetRecipientOk() (*string, bool) {
	if o == nil || IsNil(o.Recipient) {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *NonFungibleTokenMintAnyOf) HasRecipient() bool {
	if o != nil && !IsNil(o.Recipient) {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given string and assigns it to the Recipient field.
func (o *NonFungibleTokenMintAnyOf) SetRecipient(v string) {
	o.Recipient = &v
}

// GetEventIndex returns the EventIndex field value
func (o *NonFungibleTokenMintAnyOf) GetEventIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventIndex
}

// GetEventIndexOk returns a tuple with the EventIndex field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenMintAnyOf) GetEventIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventIndex, true
}

// SetEventIndex sets field value
func (o *NonFungibleTokenMintAnyOf) SetEventIndex(v int32) {
	o.EventIndex = v
}

// GetValue returns the Value field value
func (o *NonFungibleTokenMintAnyOf) GetValue() NonFungibleTokenHoldingWithTxIdValue {
	if o == nil {
		var ret NonFungibleTokenHoldingWithTxIdValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenMintAnyOf) GetValueOk() (*NonFungibleTokenHoldingWithTxIdValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *NonFungibleTokenMintAnyOf) SetValue(v NonFungibleTokenHoldingWithTxIdValue) {
	o.Value = v
}

// GetTx returns the Tx field value
func (o *NonFungibleTokenMintAnyOf) GetTx() GetTransactionList200ResponseResultsInner {
	if o == nil {
		var ret GetTransactionList200ResponseResultsInner
		return ret
	}

	return o.Tx
}

// GetTxOk returns a tuple with the Tx field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenMintAnyOf) GetTxOk() (*GetTransactionList200ResponseResultsInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tx, true
}

// SetTx sets field value
func (o *NonFungibleTokenMintAnyOf) SetTx(v GetTransactionList200ResponseResultsInner) {
	o.Tx = v
}

func (o NonFungibleTokenMintAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonFungibleTokenMintAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Recipient) {
		toSerialize["recipient"] = o.Recipient
	}
	toSerialize["event_index"] = o.EventIndex
	toSerialize["value"] = o.Value
	toSerialize["tx"] = o.Tx
	return toSerialize, nil
}

func (o *NonFungibleTokenMintAnyOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event_index",
		"value",
		"tx",
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

	varNonFungibleTokenMintAnyOf := _NonFungibleTokenMintAnyOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNonFungibleTokenMintAnyOf)

	if err != nil {
		return err
	}

	*o = NonFungibleTokenMintAnyOf(varNonFungibleTokenMintAnyOf)

	return err
}

type NullableNonFungibleTokenMintAnyOf struct {
	value *NonFungibleTokenMintAnyOf
	isSet bool
}

func (v NullableNonFungibleTokenMintAnyOf) Get() *NonFungibleTokenMintAnyOf {
	return v.value
}

func (v *NullableNonFungibleTokenMintAnyOf) Set(val *NonFungibleTokenMintAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNonFungibleTokenMintAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNonFungibleTokenMintAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonFungibleTokenMintAnyOf(val *NonFungibleTokenMintAnyOf) *NullableNonFungibleTokenMintAnyOf {
	return &NullableNonFungibleTokenMintAnyOf{value: val, isSet: true}
}

func (v NullableNonFungibleTokenMintAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonFungibleTokenMintAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


