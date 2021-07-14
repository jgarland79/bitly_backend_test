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

// ClickMetricsAllOf struct for ClickMetricsAllOf
type ClickMetricsAllOf struct {
	Metrics *[]ClickMetric `json:"metrics,omitempty"`
}

// NewClickMetricsAllOf instantiates a new ClickMetricsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClickMetricsAllOf() *ClickMetricsAllOf {
	this := ClickMetricsAllOf{}
	return &this
}

// NewClickMetricsAllOfWithDefaults instantiates a new ClickMetricsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClickMetricsAllOfWithDefaults() *ClickMetricsAllOf {
	this := ClickMetricsAllOf{}
	return &this
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *ClickMetricsAllOf) GetMetrics() []ClickMetric {
	if o == nil || o.Metrics == nil {
		var ret []ClickMetric
		return ret
	}
	return *o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickMetricsAllOf) GetMetricsOk() (*[]ClickMetric, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *ClickMetricsAllOf) HasMetrics() bool {
	if o != nil && o.Metrics != nil {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given []ClickMetric and assigns it to the Metrics field.
func (o *ClickMetricsAllOf) SetMetrics(v []ClickMetric) {
	o.Metrics = &v
}

func (o ClickMetricsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Metrics != nil {
		toSerialize["metrics"] = o.Metrics
	}
	return json.Marshal(toSerialize)
}

type NullableClickMetricsAllOf struct {
	value *ClickMetricsAllOf
	isSet bool
}

func (v NullableClickMetricsAllOf) Get() *ClickMetricsAllOf {
	return v.value
}

func (v *NullableClickMetricsAllOf) Set(val *ClickMetricsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableClickMetricsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableClickMetricsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClickMetricsAllOf(val *ClickMetricsAllOf) *NullableClickMetricsAllOf {
	return &NullableClickMetricsAllOf{value: val, isSet: true}
}

func (v NullableClickMetricsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClickMetricsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

