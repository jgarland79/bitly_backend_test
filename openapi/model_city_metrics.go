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

// CityMetrics struct for CityMetrics
type CityMetrics struct {
	Units *int32 `json:"units,omitempty"`
	Facet *string `json:"facet,omitempty"`
	UnitReference *string `json:"unit_reference,omitempty"`
	Unit *string `json:"unit,omitempty"`
	Metrics *[]CityMetric `json:"metrics,omitempty"`
	OtherMetrics *OtherMetrics `json:"other_metrics,omitempty"`
}

// NewCityMetrics instantiates a new CityMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCityMetrics() *CityMetrics {
	this := CityMetrics{}
	return &this
}

// NewCityMetricsWithDefaults instantiates a new CityMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCityMetricsWithDefaults() *CityMetrics {
	this := CityMetrics{}
	return &this
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *CityMetrics) GetUnits() int32 {
	if o == nil || o.Units == nil {
		var ret int32
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CityMetrics) GetUnitsOk() (*int32, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *CityMetrics) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given int32 and assigns it to the Units field.
func (o *CityMetrics) SetUnits(v int32) {
	o.Units = &v
}

// GetFacet returns the Facet field value if set, zero value otherwise.
func (o *CityMetrics) GetFacet() string {
	if o == nil || o.Facet == nil {
		var ret string
		return ret
	}
	return *o.Facet
}

// GetFacetOk returns a tuple with the Facet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CityMetrics) GetFacetOk() (*string, bool) {
	if o == nil || o.Facet == nil {
		return nil, false
	}
	return o.Facet, true
}

// HasFacet returns a boolean if a field has been set.
func (o *CityMetrics) HasFacet() bool {
	if o != nil && o.Facet != nil {
		return true
	}

	return false
}

// SetFacet gets a reference to the given string and assigns it to the Facet field.
func (o *CityMetrics) SetFacet(v string) {
	o.Facet = &v
}

// GetUnitReference returns the UnitReference field value if set, zero value otherwise.
func (o *CityMetrics) GetUnitReference() string {
	if o == nil || o.UnitReference == nil {
		var ret string
		return ret
	}
	return *o.UnitReference
}

// GetUnitReferenceOk returns a tuple with the UnitReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CityMetrics) GetUnitReferenceOk() (*string, bool) {
	if o == nil || o.UnitReference == nil {
		return nil, false
	}
	return o.UnitReference, true
}

// HasUnitReference returns a boolean if a field has been set.
func (o *CityMetrics) HasUnitReference() bool {
	if o != nil && o.UnitReference != nil {
		return true
	}

	return false
}

// SetUnitReference gets a reference to the given string and assigns it to the UnitReference field.
func (o *CityMetrics) SetUnitReference(v string) {
	o.UnitReference = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *CityMetrics) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CityMetrics) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *CityMetrics) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *CityMetrics) SetUnit(v string) {
	o.Unit = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *CityMetrics) GetMetrics() []CityMetric {
	if o == nil || o.Metrics == nil {
		var ret []CityMetric
		return ret
	}
	return *o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CityMetrics) GetMetricsOk() (*[]CityMetric, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *CityMetrics) HasMetrics() bool {
	if o != nil && o.Metrics != nil {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given []CityMetric and assigns it to the Metrics field.
func (o *CityMetrics) SetMetrics(v []CityMetric) {
	o.Metrics = &v
}

// GetOtherMetrics returns the OtherMetrics field value if set, zero value otherwise.
func (o *CityMetrics) GetOtherMetrics() OtherMetrics {
	if o == nil || o.OtherMetrics == nil {
		var ret OtherMetrics
		return ret
	}
	return *o.OtherMetrics
}

// GetOtherMetricsOk returns a tuple with the OtherMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CityMetrics) GetOtherMetricsOk() (*OtherMetrics, bool) {
	if o == nil || o.OtherMetrics == nil {
		return nil, false
	}
	return o.OtherMetrics, true
}

// HasOtherMetrics returns a boolean if a field has been set.
func (o *CityMetrics) HasOtherMetrics() bool {
	if o != nil && o.OtherMetrics != nil {
		return true
	}

	return false
}

// SetOtherMetrics gets a reference to the given OtherMetrics and assigns it to the OtherMetrics field.
func (o *CityMetrics) SetOtherMetrics(v OtherMetrics) {
	o.OtherMetrics = &v
}

func (o CityMetrics) MarshalJSON() ([]byte, error) {
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
	if o.OtherMetrics != nil {
		toSerialize["other_metrics"] = o.OtherMetrics
	}
	return json.Marshal(toSerialize)
}

type NullableCityMetrics struct {
	value *CityMetrics
	isSet bool
}

func (v NullableCityMetrics) Get() *CityMetrics {
	return v.value
}

func (v *NullableCityMetrics) Set(val *CityMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableCityMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableCityMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCityMetrics(val *CityMetrics) *NullableCityMetrics {
	return &NullableCityMetrics{value: val, isSet: true}
}

func (v NullableCityMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCityMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


