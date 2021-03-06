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

// HasReferences struct for HasReferences
type HasReferences struct {
	References *map[string]string `json:"references,omitempty"`
}

// NewHasReferences instantiates a new HasReferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHasReferences() *HasReferences {
	this := HasReferences{}
	return &this
}

// NewHasReferencesWithDefaults instantiates a new HasReferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHasReferencesWithDefaults() *HasReferences {
	this := HasReferences{}
	return &this
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *HasReferences) GetReferences() map[string]string {
	if o == nil || o.References == nil {
		var ret map[string]string
		return ret
	}
	return *o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasReferences) GetReferencesOk() (*map[string]string, bool) {
	if o == nil || o.References == nil {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *HasReferences) HasReferences() bool {
	if o != nil && o.References != nil {
		return true
	}

	return false
}

// SetReferences gets a reference to the given map[string]string and assigns it to the References field.
func (o *HasReferences) SetReferences(v map[string]string) {
	o.References = &v
}

func (o HasReferences) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.References != nil {
		toSerialize["references"] = o.References
	}
	return json.Marshal(toSerialize)
}

type NullableHasReferences struct {
	value *HasReferences
	isSet bool
}

func (v NullableHasReferences) Get() *HasReferences {
	return v.value
}

func (v *NullableHasReferences) Set(val *HasReferences) {
	v.value = val
	v.isSet = true
}

func (v NullableHasReferences) IsSet() bool {
	return v.isSet
}

func (v *NullableHasReferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHasReferences(val *HasReferences) *NullableHasReferences {
	return &NullableHasReferences{value: val, isSet: true}
}

func (v NullableHasReferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHasReferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


