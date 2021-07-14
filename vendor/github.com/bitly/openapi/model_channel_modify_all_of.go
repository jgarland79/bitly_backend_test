/*
 * Bitly API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.0.0
 * Contact: api@bitly.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ChannelModifyAllOf struct for ChannelModifyAllOf
type ChannelModifyAllOf struct {
	Bitlinks *[]BaseChannelBitlink `json:"bitlinks,omitempty"`
}

// NewChannelModifyAllOf instantiates a new ChannelModifyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelModifyAllOf() *ChannelModifyAllOf {
	this := ChannelModifyAllOf{}
	return &this
}

// NewChannelModifyAllOfWithDefaults instantiates a new ChannelModifyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelModifyAllOfWithDefaults() *ChannelModifyAllOf {
	this := ChannelModifyAllOf{}
	return &this
}

// GetBitlinks returns the Bitlinks field value if set, zero value otherwise.
func (o *ChannelModifyAllOf) GetBitlinks() []BaseChannelBitlink {
	if o == nil || o.Bitlinks == nil {
		var ret []BaseChannelBitlink
		return ret
	}
	return *o.Bitlinks
}

// GetBitlinksOk returns a tuple with the Bitlinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelModifyAllOf) GetBitlinksOk() (*[]BaseChannelBitlink, bool) {
	if o == nil || o.Bitlinks == nil {
		return nil, false
	}
	return o.Bitlinks, true
}

// HasBitlinks returns a boolean if a field has been set.
func (o *ChannelModifyAllOf) HasBitlinks() bool {
	if o != nil && o.Bitlinks != nil {
		return true
	}

	return false
}

// SetBitlinks gets a reference to the given []BaseChannelBitlink and assigns it to the Bitlinks field.
func (o *ChannelModifyAllOf) SetBitlinks(v []BaseChannelBitlink) {
	o.Bitlinks = &v
}

func (o ChannelModifyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bitlinks != nil {
		toSerialize["bitlinks"] = o.Bitlinks
	}
	return json.Marshal(toSerialize)
}

type NullableChannelModifyAllOf struct {
	value *ChannelModifyAllOf
	isSet bool
}

func (v NullableChannelModifyAllOf) Get() *ChannelModifyAllOf {
	return v.value
}

func (v *NullableChannelModifyAllOf) Set(val *ChannelModifyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelModifyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelModifyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelModifyAllOf(val *ChannelModifyAllOf) *NullableChannelModifyAllOf {
	return &NullableChannelModifyAllOf{value: val, isSet: true}
}

func (v NullableChannelModifyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelModifyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

