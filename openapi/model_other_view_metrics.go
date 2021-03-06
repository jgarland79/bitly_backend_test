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

// OtherViewMetrics struct for OtherViewMetrics
type OtherViewMetrics struct {
	OtherCityViews *int32 `json:"other_city_views,omitempty"`
	NoCityViews *int32 `json:"no_city_views,omitempty"`
}

// NewOtherViewMetrics instantiates a new OtherViewMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOtherViewMetrics() *OtherViewMetrics {
	this := OtherViewMetrics{}
	return &this
}

// NewOtherViewMetricsWithDefaults instantiates a new OtherViewMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOtherViewMetricsWithDefaults() *OtherViewMetrics {
	this := OtherViewMetrics{}
	return &this
}

// GetOtherCityViews returns the OtherCityViews field value if set, zero value otherwise.
func (o *OtherViewMetrics) GetOtherCityViews() int32 {
	if o == nil || o.OtherCityViews == nil {
		var ret int32
		return ret
	}
	return *o.OtherCityViews
}

// GetOtherCityViewsOk returns a tuple with the OtherCityViews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OtherViewMetrics) GetOtherCityViewsOk() (*int32, bool) {
	if o == nil || o.OtherCityViews == nil {
		return nil, false
	}
	return o.OtherCityViews, true
}

// HasOtherCityViews returns a boolean if a field has been set.
func (o *OtherViewMetrics) HasOtherCityViews() bool {
	if o != nil && o.OtherCityViews != nil {
		return true
	}

	return false
}

// SetOtherCityViews gets a reference to the given int32 and assigns it to the OtherCityViews field.
func (o *OtherViewMetrics) SetOtherCityViews(v int32) {
	o.OtherCityViews = &v
}

// GetNoCityViews returns the NoCityViews field value if set, zero value otherwise.
func (o *OtherViewMetrics) GetNoCityViews() int32 {
	if o == nil || o.NoCityViews == nil {
		var ret int32
		return ret
	}
	return *o.NoCityViews
}

// GetNoCityViewsOk returns a tuple with the NoCityViews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OtherViewMetrics) GetNoCityViewsOk() (*int32, bool) {
	if o == nil || o.NoCityViews == nil {
		return nil, false
	}
	return o.NoCityViews, true
}

// HasNoCityViews returns a boolean if a field has been set.
func (o *OtherViewMetrics) HasNoCityViews() bool {
	if o != nil && o.NoCityViews != nil {
		return true
	}

	return false
}

// SetNoCityViews gets a reference to the given int32 and assigns it to the NoCityViews field.
func (o *OtherViewMetrics) SetNoCityViews(v int32) {
	o.NoCityViews = &v
}

func (o OtherViewMetrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OtherCityViews != nil {
		toSerialize["other_city_views"] = o.OtherCityViews
	}
	if o.NoCityViews != nil {
		toSerialize["no_city_views"] = o.NoCityViews
	}
	return json.Marshal(toSerialize)
}

type NullableOtherViewMetrics struct {
	value *OtherViewMetrics
	isSet bool
}

func (v NullableOtherViewMetrics) Get() *OtherViewMetrics {
	return v.value
}

func (v *NullableOtherViewMetrics) Set(val *OtherViewMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableOtherViewMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableOtherViewMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOtherViewMetrics(val *OtherViewMetrics) *NullableOtherViewMetrics {
	return &NullableOtherViewMetrics{value: val, isSet: true}
}

func (v NullableOtherViewMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOtherViewMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


