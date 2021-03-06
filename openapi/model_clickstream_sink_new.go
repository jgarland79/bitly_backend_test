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

// ClickstreamSinkNew struct for ClickstreamSinkNew
type ClickstreamSinkNew struct {
	// Location of the sink
	Destination *string `json:"destination,omitempty"`
	// Descriptive name for the sink
	Name *string `json:"name,omitempty"`
	// Auth key for syncing to destination
	AuthKey *string `json:"auth_key,omitempty"`
	// Provider of sink
	Provider *string `json:"provider,omitempty"`
}

// NewClickstreamSinkNew instantiates a new ClickstreamSinkNew object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClickstreamSinkNew() *ClickstreamSinkNew {
	this := ClickstreamSinkNew{}
	return &this
}

// NewClickstreamSinkNewWithDefaults instantiates a new ClickstreamSinkNew object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClickstreamSinkNewWithDefaults() *ClickstreamSinkNew {
	this := ClickstreamSinkNew{}
	return &this
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *ClickstreamSinkNew) GetDestination() string {
	if o == nil || o.Destination == nil {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickstreamSinkNew) GetDestinationOk() (*string, bool) {
	if o == nil || o.Destination == nil {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *ClickstreamSinkNew) HasDestination() bool {
	if o != nil && o.Destination != nil {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *ClickstreamSinkNew) SetDestination(v string) {
	o.Destination = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClickstreamSinkNew) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickstreamSinkNew) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClickstreamSinkNew) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClickstreamSinkNew) SetName(v string) {
	o.Name = &v
}

// GetAuthKey returns the AuthKey field value if set, zero value otherwise.
func (o *ClickstreamSinkNew) GetAuthKey() string {
	if o == nil || o.AuthKey == nil {
		var ret string
		return ret
	}
	return *o.AuthKey
}

// GetAuthKeyOk returns a tuple with the AuthKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickstreamSinkNew) GetAuthKeyOk() (*string, bool) {
	if o == nil || o.AuthKey == nil {
		return nil, false
	}
	return o.AuthKey, true
}

// HasAuthKey returns a boolean if a field has been set.
func (o *ClickstreamSinkNew) HasAuthKey() bool {
	if o != nil && o.AuthKey != nil {
		return true
	}

	return false
}

// SetAuthKey gets a reference to the given string and assigns it to the AuthKey field.
func (o *ClickstreamSinkNew) SetAuthKey(v string) {
	o.AuthKey = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *ClickstreamSinkNew) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickstreamSinkNew) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *ClickstreamSinkNew) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *ClickstreamSinkNew) SetProvider(v string) {
	o.Provider = &v
}

func (o ClickstreamSinkNew) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Destination != nil {
		toSerialize["destination"] = o.Destination
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.AuthKey != nil {
		toSerialize["auth_key"] = o.AuthKey
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	return json.Marshal(toSerialize)
}

type NullableClickstreamSinkNew struct {
	value *ClickstreamSinkNew
	isSet bool
}

func (v NullableClickstreamSinkNew) Get() *ClickstreamSinkNew {
	return v.value
}

func (v *NullableClickstreamSinkNew) Set(val *ClickstreamSinkNew) {
	v.value = val
	v.isSet = true
}

func (v NullableClickstreamSinkNew) IsSet() bool {
	return v.isSet
}

func (v *NullableClickstreamSinkNew) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClickstreamSinkNew(val *ClickstreamSinkNew) *NullableClickstreamSinkNew {
	return &NullableClickstreamSinkNew{value: val, isSet: true}
}

func (v NullableClickstreamSinkNew) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClickstreamSinkNew) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


