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

// checks if the BlockSearchResultBlockData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockSearchResultBlockData{}

// BlockSearchResultBlockData struct for BlockSearchResultBlockData
type BlockSearchResultBlockData struct {
	Canonical bool `json:"canonical"`
	Hash string `json:"hash"`
	ParentBlockHash string `json:"parent_block_hash"`
	BurnBlockTime int32 `json:"burn_block_time"`
	Height int32 `json:"height"`
}

type _BlockSearchResultBlockData BlockSearchResultBlockData

// NewBlockSearchResultBlockData instantiates a new BlockSearchResultBlockData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockSearchResultBlockData(canonical bool, hash string, parentBlockHash string, burnBlockTime int32, height int32) *BlockSearchResultBlockData {
	this := BlockSearchResultBlockData{}
	this.Canonical = canonical
	this.Hash = hash
	this.ParentBlockHash = parentBlockHash
	this.BurnBlockTime = burnBlockTime
	this.Height = height
	return &this
}

// NewBlockSearchResultBlockDataWithDefaults instantiates a new BlockSearchResultBlockData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockSearchResultBlockDataWithDefaults() *BlockSearchResultBlockData {
	this := BlockSearchResultBlockData{}
	return &this
}

// GetCanonical returns the Canonical field value
func (o *BlockSearchResultBlockData) GetCanonical() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Canonical
}

// GetCanonicalOk returns a tuple with the Canonical field value
// and a boolean to check if the value has been set.
func (o *BlockSearchResultBlockData) GetCanonicalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Canonical, true
}

// SetCanonical sets field value
func (o *BlockSearchResultBlockData) SetCanonical(v bool) {
	o.Canonical = v
}

// GetHash returns the Hash field value
func (o *BlockSearchResultBlockData) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *BlockSearchResultBlockData) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *BlockSearchResultBlockData) SetHash(v string) {
	o.Hash = v
}

// GetParentBlockHash returns the ParentBlockHash field value
func (o *BlockSearchResultBlockData) GetParentBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentBlockHash
}

// GetParentBlockHashOk returns a tuple with the ParentBlockHash field value
// and a boolean to check if the value has been set.
func (o *BlockSearchResultBlockData) GetParentBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentBlockHash, true
}

// SetParentBlockHash sets field value
func (o *BlockSearchResultBlockData) SetParentBlockHash(v string) {
	o.ParentBlockHash = v
}

// GetBurnBlockTime returns the BurnBlockTime field value
func (o *BlockSearchResultBlockData) GetBurnBlockTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BurnBlockTime
}

// GetBurnBlockTimeOk returns a tuple with the BurnBlockTime field value
// and a boolean to check if the value has been set.
func (o *BlockSearchResultBlockData) GetBurnBlockTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockTime, true
}

// SetBurnBlockTime sets field value
func (o *BlockSearchResultBlockData) SetBurnBlockTime(v int32) {
	o.BurnBlockTime = v
}

// GetHeight returns the Height field value
func (o *BlockSearchResultBlockData) GetHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *BlockSearchResultBlockData) GetHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *BlockSearchResultBlockData) SetHeight(v int32) {
	o.Height = v
}

func (o BlockSearchResultBlockData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockSearchResultBlockData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["canonical"] = o.Canonical
	toSerialize["hash"] = o.Hash
	toSerialize["parent_block_hash"] = o.ParentBlockHash
	toSerialize["burn_block_time"] = o.BurnBlockTime
	toSerialize["height"] = o.Height
	return toSerialize, nil
}

func (o *BlockSearchResultBlockData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"canonical",
		"hash",
		"parent_block_hash",
		"burn_block_time",
		"height",
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

	varBlockSearchResultBlockData := _BlockSearchResultBlockData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBlockSearchResultBlockData)

	if err != nil {
		return err
	}

	*o = BlockSearchResultBlockData(varBlockSearchResultBlockData)

	return err
}

type NullableBlockSearchResultBlockData struct {
	value *BlockSearchResultBlockData
	isSet bool
}

func (v NullableBlockSearchResultBlockData) Get() *BlockSearchResultBlockData {
	return v.value
}

func (v *NullableBlockSearchResultBlockData) Set(val *BlockSearchResultBlockData) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockSearchResultBlockData) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockSearchResultBlockData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockSearchResultBlockData(val *BlockSearchResultBlockData) *NullableBlockSearchResultBlockData {
	return &NullableBlockSearchResultBlockData{value: val, isSet: true}
}

func (v NullableBlockSearchResultBlockData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockSearchResultBlockData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


