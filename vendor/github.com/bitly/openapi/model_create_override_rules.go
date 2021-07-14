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

// CreateOverrideRules struct for CreateOverrideRules
type CreateOverrideRules struct {
	Rules string `json:"rules"`
	GroupGuid *string `json:"group_guid,omitempty"`
}

// NewCreateOverrideRules instantiates a new CreateOverrideRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOverrideRules(rules string) *CreateOverrideRules {
	this := CreateOverrideRules{}
	this.Rules = rules
	return &this
}

// NewCreateOverrideRulesWithDefaults instantiates a new CreateOverrideRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOverrideRulesWithDefaults() *CreateOverrideRules {
	this := CreateOverrideRules{}
	return &this
}

// GetRules returns the Rules field value
func (o *CreateOverrideRules) GetRules() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *CreateOverrideRules) GetRulesOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Rules, true
}

// SetRules sets field value
func (o *CreateOverrideRules) SetRules(v string) {
	o.Rules = v
}

// GetGroupGuid returns the GroupGuid field value if set, zero value otherwise.
func (o *CreateOverrideRules) GetGroupGuid() string {
	if o == nil || o.GroupGuid == nil {
		var ret string
		return ret
	}
	return *o.GroupGuid
}

// GetGroupGuidOk returns a tuple with the GroupGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOverrideRules) GetGroupGuidOk() (*string, bool) {
	if o == nil || o.GroupGuid == nil {
		return nil, false
	}
	return o.GroupGuid, true
}

// HasGroupGuid returns a boolean if a field has been set.
func (o *CreateOverrideRules) HasGroupGuid() bool {
	if o != nil && o.GroupGuid != nil {
		return true
	}

	return false
}

// SetGroupGuid gets a reference to the given string and assigns it to the GroupGuid field.
func (o *CreateOverrideRules) SetGroupGuid(v string) {
	o.GroupGuid = &v
}

func (o CreateOverrideRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rules"] = o.Rules
	}
	if o.GroupGuid != nil {
		toSerialize["group_guid"] = o.GroupGuid
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOverrideRules struct {
	value *CreateOverrideRules
	isSet bool
}

func (v NullableCreateOverrideRules) Get() *CreateOverrideRules {
	return v.value
}

func (v *NullableCreateOverrideRules) Set(val *CreateOverrideRules) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOverrideRules) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOverrideRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOverrideRules(val *CreateOverrideRules) *NullableCreateOverrideRules {
	return &NullableCreateOverrideRules{value: val, isSet: true}
}

func (v NullableCreateOverrideRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOverrideRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

