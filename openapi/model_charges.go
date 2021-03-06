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

// Charges struct for Charges
type Charges struct {
	ChargeAmount *float32 `json:"charge_amount,omitempty"`
	ProcessingType *string `json:"processing_type,omitempty"`
	TaxAmount *float32 `json:"tax_amount,omitempty"`
	ChargeName *string `json:"charge_name,omitempty"`
}

// NewCharges instantiates a new Charges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCharges() *Charges {
	this := Charges{}
	return &this
}

// NewChargesWithDefaults instantiates a new Charges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargesWithDefaults() *Charges {
	this := Charges{}
	return &this
}

// GetChargeAmount returns the ChargeAmount field value if set, zero value otherwise.
func (o *Charges) GetChargeAmount() float32 {
	if o == nil || o.ChargeAmount == nil {
		var ret float32
		return ret
	}
	return *o.ChargeAmount
}

// GetChargeAmountOk returns a tuple with the ChargeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Charges) GetChargeAmountOk() (*float32, bool) {
	if o == nil || o.ChargeAmount == nil {
		return nil, false
	}
	return o.ChargeAmount, true
}

// HasChargeAmount returns a boolean if a field has been set.
func (o *Charges) HasChargeAmount() bool {
	if o != nil && o.ChargeAmount != nil {
		return true
	}

	return false
}

// SetChargeAmount gets a reference to the given float32 and assigns it to the ChargeAmount field.
func (o *Charges) SetChargeAmount(v float32) {
	o.ChargeAmount = &v
}

// GetProcessingType returns the ProcessingType field value if set, zero value otherwise.
func (o *Charges) GetProcessingType() string {
	if o == nil || o.ProcessingType == nil {
		var ret string
		return ret
	}
	return *o.ProcessingType
}

// GetProcessingTypeOk returns a tuple with the ProcessingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Charges) GetProcessingTypeOk() (*string, bool) {
	if o == nil || o.ProcessingType == nil {
		return nil, false
	}
	return o.ProcessingType, true
}

// HasProcessingType returns a boolean if a field has been set.
func (o *Charges) HasProcessingType() bool {
	if o != nil && o.ProcessingType != nil {
		return true
	}

	return false
}

// SetProcessingType gets a reference to the given string and assigns it to the ProcessingType field.
func (o *Charges) SetProcessingType(v string) {
	o.ProcessingType = &v
}

// GetTaxAmount returns the TaxAmount field value if set, zero value otherwise.
func (o *Charges) GetTaxAmount() float32 {
	if o == nil || o.TaxAmount == nil {
		var ret float32
		return ret
	}
	return *o.TaxAmount
}

// GetTaxAmountOk returns a tuple with the TaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Charges) GetTaxAmountOk() (*float32, bool) {
	if o == nil || o.TaxAmount == nil {
		return nil, false
	}
	return o.TaxAmount, true
}

// HasTaxAmount returns a boolean if a field has been set.
func (o *Charges) HasTaxAmount() bool {
	if o != nil && o.TaxAmount != nil {
		return true
	}

	return false
}

// SetTaxAmount gets a reference to the given float32 and assigns it to the TaxAmount field.
func (o *Charges) SetTaxAmount(v float32) {
	o.TaxAmount = &v
}

// GetChargeName returns the ChargeName field value if set, zero value otherwise.
func (o *Charges) GetChargeName() string {
	if o == nil || o.ChargeName == nil {
		var ret string
		return ret
	}
	return *o.ChargeName
}

// GetChargeNameOk returns a tuple with the ChargeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Charges) GetChargeNameOk() (*string, bool) {
	if o == nil || o.ChargeName == nil {
		return nil, false
	}
	return o.ChargeName, true
}

// HasChargeName returns a boolean if a field has been set.
func (o *Charges) HasChargeName() bool {
	if o != nil && o.ChargeName != nil {
		return true
	}

	return false
}

// SetChargeName gets a reference to the given string and assigns it to the ChargeName field.
func (o *Charges) SetChargeName(v string) {
	o.ChargeName = &v
}

func (o Charges) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChargeAmount != nil {
		toSerialize["charge_amount"] = o.ChargeAmount
	}
	if o.ProcessingType != nil {
		toSerialize["processing_type"] = o.ProcessingType
	}
	if o.TaxAmount != nil {
		toSerialize["tax_amount"] = o.TaxAmount
	}
	if o.ChargeName != nil {
		toSerialize["charge_name"] = o.ChargeName
	}
	return json.Marshal(toSerialize)
}

type NullableCharges struct {
	value *Charges
	isSet bool
}

func (v NullableCharges) Get() *Charges {
	return v.value
}

func (v *NullableCharges) Set(val *Charges) {
	v.value = val
	v.isSet = true
}

func (v NullableCharges) IsSet() bool {
	return v.isSet
}

func (v *NullableCharges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCharges(val *Charges) *NullableCharges {
	return &NullableCharges{value: val, isSet: true}
}

func (v NullableCharges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCharges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


