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

// ThirdPartyAppData struct for ThirdPartyAppData
type ThirdPartyAppData struct {
	ThirdPartyAppId *string `json:"third_party_app_id,omitempty"`
	InstallUrl *string `json:"install_url,omitempty"`
	Os *string `json:"os,omitempty"`
	Name *string `json:"name,omitempty"`
	IconUrl *string `json:"icon_url,omitempty"`
}

// NewThirdPartyAppData instantiates a new ThirdPartyAppData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyAppData() *ThirdPartyAppData {
	this := ThirdPartyAppData{}
	return &this
}

// NewThirdPartyAppDataWithDefaults instantiates a new ThirdPartyAppData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyAppDataWithDefaults() *ThirdPartyAppData {
	this := ThirdPartyAppData{}
	return &this
}

// GetThirdPartyAppId returns the ThirdPartyAppId field value if set, zero value otherwise.
func (o *ThirdPartyAppData) GetThirdPartyAppId() string {
	if o == nil || o.ThirdPartyAppId == nil {
		var ret string
		return ret
	}
	return *o.ThirdPartyAppId
}

// GetThirdPartyAppIdOk returns a tuple with the ThirdPartyAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAppData) GetThirdPartyAppIdOk() (*string, bool) {
	if o == nil || o.ThirdPartyAppId == nil {
		return nil, false
	}
	return o.ThirdPartyAppId, true
}

// HasThirdPartyAppId returns a boolean if a field has been set.
func (o *ThirdPartyAppData) HasThirdPartyAppId() bool {
	if o != nil && o.ThirdPartyAppId != nil {
		return true
	}

	return false
}

// SetThirdPartyAppId gets a reference to the given string and assigns it to the ThirdPartyAppId field.
func (o *ThirdPartyAppData) SetThirdPartyAppId(v string) {
	o.ThirdPartyAppId = &v
}

// GetInstallUrl returns the InstallUrl field value if set, zero value otherwise.
func (o *ThirdPartyAppData) GetInstallUrl() string {
	if o == nil || o.InstallUrl == nil {
		var ret string
		return ret
	}
	return *o.InstallUrl
}

// GetInstallUrlOk returns a tuple with the InstallUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAppData) GetInstallUrlOk() (*string, bool) {
	if o == nil || o.InstallUrl == nil {
		return nil, false
	}
	return o.InstallUrl, true
}

// HasInstallUrl returns a boolean if a field has been set.
func (o *ThirdPartyAppData) HasInstallUrl() bool {
	if o != nil && o.InstallUrl != nil {
		return true
	}

	return false
}

// SetInstallUrl gets a reference to the given string and assigns it to the InstallUrl field.
func (o *ThirdPartyAppData) SetInstallUrl(v string) {
	o.InstallUrl = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *ThirdPartyAppData) GetOs() string {
	if o == nil || o.Os == nil {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAppData) GetOsOk() (*string, bool) {
	if o == nil || o.Os == nil {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *ThirdPartyAppData) HasOs() bool {
	if o != nil && o.Os != nil {
		return true
	}

	return false
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *ThirdPartyAppData) SetOs(v string) {
	o.Os = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ThirdPartyAppData) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAppData) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ThirdPartyAppData) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ThirdPartyAppData) SetName(v string) {
	o.Name = &v
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise.
func (o *ThirdPartyAppData) GetIconUrl() string {
	if o == nil || o.IconUrl == nil {
		var ret string
		return ret
	}
	return *o.IconUrl
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAppData) GetIconUrlOk() (*string, bool) {
	if o == nil || o.IconUrl == nil {
		return nil, false
	}
	return o.IconUrl, true
}

// HasIconUrl returns a boolean if a field has been set.
func (o *ThirdPartyAppData) HasIconUrl() bool {
	if o != nil && o.IconUrl != nil {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given string and assigns it to the IconUrl field.
func (o *ThirdPartyAppData) SetIconUrl(v string) {
	o.IconUrl = &v
}

func (o ThirdPartyAppData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ThirdPartyAppId != nil {
		toSerialize["third_party_app_id"] = o.ThirdPartyAppId
	}
	if o.InstallUrl != nil {
		toSerialize["install_url"] = o.InstallUrl
	}
	if o.Os != nil {
		toSerialize["os"] = o.Os
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.IconUrl != nil {
		toSerialize["icon_url"] = o.IconUrl
	}
	return json.Marshal(toSerialize)
}

type NullableThirdPartyAppData struct {
	value *ThirdPartyAppData
	isSet bool
}

func (v NullableThirdPartyAppData) Get() *ThirdPartyAppData {
	return v.value
}

func (v *NullableThirdPartyAppData) Set(val *ThirdPartyAppData) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyAppData) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyAppData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyAppData(val *ThirdPartyAppData) *NullableThirdPartyAppData {
	return &NullableThirdPartyAppData{value: val, isSet: true}
}

func (v NullableThirdPartyAppData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyAppData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


