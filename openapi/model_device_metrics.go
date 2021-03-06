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

// DeviceMetrics struct for DeviceMetrics
type DeviceMetrics struct {
	Units *int32 `json:"units,omitempty"`
	Facet *string `json:"facet,omitempty"`
	UnitReference *string `json:"unit_reference,omitempty"`
	Unit *string `json:"unit,omitempty"`
	Metrics *[]DeviceMetric `json:"metrics,omitempty"`
}

// NewDeviceMetrics instantiates a new DeviceMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceMetrics() *DeviceMetrics {
	this := DeviceMetrics{}
	return &this
}

// NewDeviceMetricsWithDefaults instantiates a new DeviceMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceMetricsWithDefaults() *DeviceMetrics {
	this := DeviceMetrics{}
	return &this
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *DeviceMetrics) GetUnits() int32 {
	if o == nil || o.Units == nil {
		var ret int32
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceMetrics) GetUnitsOk() (*int32, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *DeviceMetrics) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given int32 and assigns it to the Units field.
func (o *DeviceMetrics) SetUnits(v int32) {
	o.Units = &v
}

// GetFacet returns the Facet field value if set, zero value otherwise.
func (o *DeviceMetrics) GetFacet() string {
	if o == nil || o.Facet == nil {
		var ret string
		return ret
	}
	return *o.Facet
}

// GetFacetOk returns a tuple with the Facet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceMetrics) GetFacetOk() (*string, bool) {
	if o == nil || o.Facet == nil {
		return nil, false
	}
	return o.Facet, true
}

// HasFacet returns a boolean if a field has been set.
func (o *DeviceMetrics) HasFacet() bool {
	if o != nil && o.Facet != nil {
		return true
	}

	return false
}

// SetFacet gets a reference to the given string and assigns it to the Facet field.
func (o *DeviceMetrics) SetFacet(v string) {
	o.Facet = &v
}

// GetUnitReference returns the UnitReference field value if set, zero value otherwise.
func (o *DeviceMetrics) GetUnitReference() string {
	if o == nil || o.UnitReference == nil {
		var ret string
		return ret
	}
	return *o.UnitReference
}

// GetUnitReferenceOk returns a tuple with the UnitReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceMetrics) GetUnitReferenceOk() (*string, bool) {
	if o == nil || o.UnitReference == nil {
		return nil, false
	}
	return o.UnitReference, true
}

// HasUnitReference returns a boolean if a field has been set.
func (o *DeviceMetrics) HasUnitReference() bool {
	if o != nil && o.UnitReference != nil {
		return true
	}

	return false
}

// SetUnitReference gets a reference to the given string and assigns it to the UnitReference field.
func (o *DeviceMetrics) SetUnitReference(v string) {
	o.UnitReference = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *DeviceMetrics) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceMetrics) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *DeviceMetrics) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *DeviceMetrics) SetUnit(v string) {
	o.Unit = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *DeviceMetrics) GetMetrics() []DeviceMetric {
	if o == nil || o.Metrics == nil {
		var ret []DeviceMetric
		return ret
	}
	return *o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceMetrics) GetMetricsOk() (*[]DeviceMetric, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *DeviceMetrics) HasMetrics() bool {
	if o != nil && o.Metrics != nil {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given []DeviceMetric and assigns it to the Metrics field.
func (o *DeviceMetrics) SetMetrics(v []DeviceMetric) {
	o.Metrics = &v
}

func (o DeviceMetrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Units != nil {
		toSerialize["units"] = o.Units
	}
	if o.Facet != nil {
		toSerialize["facet"] = o.Facet
	}
	if o.UnitReference != nil {
		toSerialize["unit_reference"] = o.UnitReference
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if o.Metrics != nil {
		toSerialize["metrics"] = o.Metrics
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceMetrics struct {
	value *DeviceMetrics
	isSet bool
}

func (v NullableDeviceMetrics) Get() *DeviceMetrics {
	return v.value
}

func (v *NullableDeviceMetrics) Set(val *DeviceMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceMetrics(val *DeviceMetrics) *NullableDeviceMetrics {
	return &NullableDeviceMetrics{value: val, isSet: true}
}

func (v NullableDeviceMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


