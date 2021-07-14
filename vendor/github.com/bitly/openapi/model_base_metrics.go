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

// BaseMetrics struct for BaseMetrics
type BaseMetrics struct {
	Units *int32 `json:"units,omitempty"`
	Facet *string `json:"facet,omitempty"`
	UnitReference *string `json:"unit_reference,omitempty"`
	Unit *string `json:"unit,omitempty"`
}

// NewBaseMetrics instantiates a new BaseMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseMetrics() *BaseMetrics {
	this := BaseMetrics{}
	return &this
}

// NewBaseMetricsWithDefaults instantiates a new BaseMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseMetricsWithDefaults() *BaseMetrics {
	this := BaseMetrics{}
	return &this
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *BaseMetrics) GetUnits() int32 {
	if o == nil || o.Units == nil {
		var ret int32
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseMetrics) GetUnitsOk() (*int32, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *BaseMetrics) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given int32 and assigns it to the Units field.
func (o *BaseMetrics) SetUnits(v int32) {
	o.Units = &v
}

// GetFacet returns the Facet field value if set, zero value otherwise.
func (o *BaseMetrics) GetFacet() string {
	if o == nil || o.Facet == nil {
		var ret string
		return ret
	}
	return *o.Facet
}

// GetFacetOk returns a tuple with the Facet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseMetrics) GetFacetOk() (*string, bool) {
	if o == nil || o.Facet == nil {
		return nil, false
	}
	return o.Facet, true
}

// HasFacet returns a boolean if a field has been set.
func (o *BaseMetrics) HasFacet() bool {
	if o != nil && o.Facet != nil {
		return true
	}

	return false
}

// SetFacet gets a reference to the given string and assigns it to the Facet field.
func (o *BaseMetrics) SetFacet(v string) {
	o.Facet = &v
}

// GetUnitReference returns the UnitReference field value if set, zero value otherwise.
func (o *BaseMetrics) GetUnitReference() string {
	if o == nil || o.UnitReference == nil {
		var ret string
		return ret
	}
	return *o.UnitReference
}

// GetUnitReferenceOk returns a tuple with the UnitReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseMetrics) GetUnitReferenceOk() (*string, bool) {
	if o == nil || o.UnitReference == nil {
		return nil, false
	}
	return o.UnitReference, true
}

// HasUnitReference returns a boolean if a field has been set.
func (o *BaseMetrics) HasUnitReference() bool {
	if o != nil && o.UnitReference != nil {
		return true
	}

	return false
}

// SetUnitReference gets a reference to the given string and assigns it to the UnitReference field.
func (o *BaseMetrics) SetUnitReference(v string) {
	o.UnitReference = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *BaseMetrics) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseMetrics) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *BaseMetrics) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *BaseMetrics) SetUnit(v string) {
	o.Unit = &v
}

func (o BaseMetrics) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableBaseMetrics struct {
	value *BaseMetrics
	isSet bool
}

func (v NullableBaseMetrics) Get() *BaseMetrics {
	return v.value
}

func (v *NullableBaseMetrics) Set(val *BaseMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseMetrics(val *BaseMetrics) *NullableBaseMetrics {
	return &NullableBaseMetrics{value: val, isSet: true}
}

func (v NullableBaseMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

