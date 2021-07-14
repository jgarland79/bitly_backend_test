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

// GroupClicksByFacetRollup struct for GroupClicksByFacetRollup
type GroupClicksByFacetRollup struct {
	Units *int32 `json:"units,omitempty"`
	Data *[]FacetCountItem `json:"data,omitempty"`
	UnitReference *string `json:"unit_reference,omitempty"`
	Unit *string `json:"unit,omitempty"`
}

// NewGroupClicksByFacetRollup instantiates a new GroupClicksByFacetRollup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupClicksByFacetRollup() *GroupClicksByFacetRollup {
	this := GroupClicksByFacetRollup{}
	return &this
}

// NewGroupClicksByFacetRollupWithDefaults instantiates a new GroupClicksByFacetRollup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupClicksByFacetRollupWithDefaults() *GroupClicksByFacetRollup {
	this := GroupClicksByFacetRollup{}
	return &this
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *GroupClicksByFacetRollup) GetUnits() int32 {
	if o == nil || o.Units == nil {
		var ret int32
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupClicksByFacetRollup) GetUnitsOk() (*int32, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *GroupClicksByFacetRollup) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given int32 and assigns it to the Units field.
func (o *GroupClicksByFacetRollup) SetUnits(v int32) {
	o.Units = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GroupClicksByFacetRollup) GetData() []FacetCountItem {
	if o == nil || o.Data == nil {
		var ret []FacetCountItem
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupClicksByFacetRollup) GetDataOk() (*[]FacetCountItem, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GroupClicksByFacetRollup) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []FacetCountItem and assigns it to the Data field.
func (o *GroupClicksByFacetRollup) SetData(v []FacetCountItem) {
	o.Data = &v
}

// GetUnitReference returns the UnitReference field value if set, zero value otherwise.
func (o *GroupClicksByFacetRollup) GetUnitReference() string {
	if o == nil || o.UnitReference == nil {
		var ret string
		return ret
	}
	return *o.UnitReference
}

// GetUnitReferenceOk returns a tuple with the UnitReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupClicksByFacetRollup) GetUnitReferenceOk() (*string, bool) {
	if o == nil || o.UnitReference == nil {
		return nil, false
	}
	return o.UnitReference, true
}

// HasUnitReference returns a boolean if a field has been set.
func (o *GroupClicksByFacetRollup) HasUnitReference() bool {
	if o != nil && o.UnitReference != nil {
		return true
	}

	return false
}

// SetUnitReference gets a reference to the given string and assigns it to the UnitReference field.
func (o *GroupClicksByFacetRollup) SetUnitReference(v string) {
	o.UnitReference = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *GroupClicksByFacetRollup) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupClicksByFacetRollup) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *GroupClicksByFacetRollup) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *GroupClicksByFacetRollup) SetUnit(v string) {
	o.Unit = &v
}

func (o GroupClicksByFacetRollup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Units != nil {
		toSerialize["units"] = o.Units
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.UnitReference != nil {
		toSerialize["unit_reference"] = o.UnitReference
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableGroupClicksByFacetRollup struct {
	value *GroupClicksByFacetRollup
	isSet bool
}

func (v NullableGroupClicksByFacetRollup) Get() *GroupClicksByFacetRollup {
	return v.value
}

func (v *NullableGroupClicksByFacetRollup) Set(val *GroupClicksByFacetRollup) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupClicksByFacetRollup) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupClicksByFacetRollup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupClicksByFacetRollup(val *GroupClicksByFacetRollup) *NullableGroupClicksByFacetRollup {
	return &NullableGroupClicksByFacetRollup{value: val, isSet: true}
}

func (v NullableGroupClicksByFacetRollup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupClicksByFacetRollup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

