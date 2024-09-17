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

// checks if the FungibleTokenAssetTransactionEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FungibleTokenAssetTransactionEvent{}

// FungibleTokenAssetTransactionEvent struct for FungibleTokenAssetTransactionEvent
type FungibleTokenAssetTransactionEvent struct {
	EventIndex int32 `json:"event_index"`
	EventType string `json:"event_type"`
	TxId string `json:"tx_id"`
	Asset FungibleTokenAssetTransactionEventAllOfAsset `json:"asset"`
}

type _FungibleTokenAssetTransactionEvent FungibleTokenAssetTransactionEvent

// NewFungibleTokenAssetTransactionEvent instantiates a new FungibleTokenAssetTransactionEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFungibleTokenAssetTransactionEvent(eventIndex int32, eventType string, txId string, asset FungibleTokenAssetTransactionEventAllOfAsset) *FungibleTokenAssetTransactionEvent {
	this := FungibleTokenAssetTransactionEvent{}
	this.EventIndex = eventIndex
	this.EventType = eventType
	this.TxId = txId
	this.Asset = asset
	return &this
}

// NewFungibleTokenAssetTransactionEventWithDefaults instantiates a new FungibleTokenAssetTransactionEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFungibleTokenAssetTransactionEventWithDefaults() *FungibleTokenAssetTransactionEvent {
	this := FungibleTokenAssetTransactionEvent{}
	return &this
}

// GetEventIndex returns the EventIndex field value
func (o *FungibleTokenAssetTransactionEvent) GetEventIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventIndex
}

// GetEventIndexOk returns a tuple with the EventIndex field value
// and a boolean to check if the value has been set.
func (o *FungibleTokenAssetTransactionEvent) GetEventIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventIndex, true
}

// SetEventIndex sets field value
func (o *FungibleTokenAssetTransactionEvent) SetEventIndex(v int32) {
	o.EventIndex = v
}

// GetEventType returns the EventType field value
func (o *FungibleTokenAssetTransactionEvent) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *FungibleTokenAssetTransactionEvent) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *FungibleTokenAssetTransactionEvent) SetEventType(v string) {
	o.EventType = v
}

// GetTxId returns the TxId field value
func (o *FungibleTokenAssetTransactionEvent) GetTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *FungibleTokenAssetTransactionEvent) GetTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *FungibleTokenAssetTransactionEvent) SetTxId(v string) {
	o.TxId = v
}

// GetAsset returns the Asset field value
func (o *FungibleTokenAssetTransactionEvent) GetAsset() FungibleTokenAssetTransactionEventAllOfAsset {
	if o == nil {
		var ret FungibleTokenAssetTransactionEventAllOfAsset
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *FungibleTokenAssetTransactionEvent) GetAssetOk() (*FungibleTokenAssetTransactionEventAllOfAsset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *FungibleTokenAssetTransactionEvent) SetAsset(v FungibleTokenAssetTransactionEventAllOfAsset) {
	o.Asset = v
}

func (o FungibleTokenAssetTransactionEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FungibleTokenAssetTransactionEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event_index"] = o.EventIndex
	toSerialize["event_type"] = o.EventType
	toSerialize["tx_id"] = o.TxId
	toSerialize["asset"] = o.Asset
	return toSerialize, nil
}

func (o *FungibleTokenAssetTransactionEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event_index",
		"event_type",
		"tx_id",
		"asset",
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

	varFungibleTokenAssetTransactionEvent := _FungibleTokenAssetTransactionEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFungibleTokenAssetTransactionEvent)

	if err != nil {
		return err
	}

	*o = FungibleTokenAssetTransactionEvent(varFungibleTokenAssetTransactionEvent)

	return err
}

type NullableFungibleTokenAssetTransactionEvent struct {
	value *FungibleTokenAssetTransactionEvent
	isSet bool
}

func (v NullableFungibleTokenAssetTransactionEvent) Get() *FungibleTokenAssetTransactionEvent {
	return v.value
}

func (v *NullableFungibleTokenAssetTransactionEvent) Set(val *FungibleTokenAssetTransactionEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableFungibleTokenAssetTransactionEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableFungibleTokenAssetTransactionEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFungibleTokenAssetTransactionEvent(val *FungibleTokenAssetTransactionEvent) *NullableFungibleTokenAssetTransactionEvent {
	return &NullableFungibleTokenAssetTransactionEvent{value: val, isSet: true}
}

func (v NullableFungibleTokenAssetTransactionEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFungibleTokenAssetTransactionEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


