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

// LaunchpadPresetTheme struct for LaunchpadPresetTheme
type LaunchpadPresetTheme struct {
	ButtonTextColor *string `json:"button_text_color,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	AvatarTextColor *string `json:"avatar_text_color,omitempty"`
	LaunchpadTextColor *string `json:"launchpad_text_color,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	Background *string `json:"background,omitempty"`
	AvatarBackground *string `json:"avatar_background,omitempty"`
	ThemeId *int32 `json:"theme_id,omitempty"`
	ButtonBackground *string `json:"button_background,omitempty"`
}

// NewLaunchpadPresetTheme instantiates a new LaunchpadPresetTheme object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLaunchpadPresetTheme() *LaunchpadPresetTheme {
	this := LaunchpadPresetTheme{}
	return &this
}

// NewLaunchpadPresetThemeWithDefaults instantiates a new LaunchpadPresetTheme object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLaunchpadPresetThemeWithDefaults() *LaunchpadPresetTheme {
	this := LaunchpadPresetTheme{}
	return &this
}

// GetButtonTextColor returns the ButtonTextColor field value if set, zero value otherwise.
func (o *LaunchpadPresetTheme) GetButtonTextColor() string {
	if o == nil || o.ButtonTextColor == nil {
		var ret string
		return ret
	}
	return *o.ButtonTextColor
}

// GetButtonTextColorOk returns a tuple with the ButtonTextColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadPresetTheme) GetButtonTextColorOk() (*string, bool) {
	if o == nil || o.ButtonTextColor == nil {
		return nil, false
	}
	return o.ButtonTextColor, true
}

// HasButtonTextColor returns a boolean if a field has been set.
func (o *LaunchpadPresetTheme) HasButtonTextColor() bool {
	if o != nil && o.ButtonTextColor != nil {
		return true
	}

	return false
}

// SetButtonTextColor gets a reference to the given string and assigns it to the ButtonTextColor field.
func (o *LaunchpadPresetTheme) SetButtonTextColor(v string) {
	o.ButtonTextColor = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *LaunchpadPresetTheme) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadPresetTheme) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *LaunchpadPresetTheme) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *LaunchpadPresetTheme) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetAvatarTextColor returns the AvatarTextColor field value if set, zero value otherwise.
func (o *LaunchpadPresetTheme) GetAvatarTextColor() string {
	if o == nil || o.AvatarTextColor == nil {
		var ret string
		return ret
	}
	return *o.AvatarTextColor
}

// GetAvatarTextColorOk returns a tuple with the AvatarTextColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadPresetTheme) GetAvatarTextColorOk() (*string, bool) {
	if o == nil || o.AvatarTextColor == nil {
		return nil, false
	}
	return o.AvatarTextColor, true
}

// HasAvatarTextColor returns a boolean if a field has been set.
func (o *LaunchpadPresetTheme) HasAvatarTextColor() bool {
	if o != nil && o.AvatarTextColor != nil {
		return true
	}

	return false
}

// SetAvatarTextColor gets a reference to the given string and assigns it to the AvatarTextColor field.
func (o *LaunchpadPresetTheme) SetAvatarTextColor(v string) {
	o.AvatarTextColor = &v
}

// GetLaunchpadTextColor returns the LaunchpadTextColor field value if set, zero value otherwise.
func (o *LaunchpadPresetTheme) GetLaunchpadTextColor() string {
	if o == nil || o.LaunchpadTextColor == nil {
		var ret string
		return ret
	}
	return *o.LaunchpadTextColor
}

// GetLaunchpadTextColorOk returns a tuple with the LaunchpadTextColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadPresetTheme) GetLaunchpadTextColorOk() (*string, bool) {
	if o == nil || o.LaunchpadTextColor == nil {
		return nil, false
	}
	return o.LaunchpadTextColor, true
}

// HasLaunchpadTextColor returns a boolean if a field has been set.
func (o *LaunchpadPresetTheme) HasLaunchpadTextColor() bool {
	if o != nil && o.LaunchpadTextColor != nil {
		return true
	}

	return false
}

// SetLaunchpadTextColor gets a reference to the given string and assigns it to the LaunchpadTextColor field.
func (o *LaunchpadPresetTheme) SetLaunchpadTextColor(v string) {
	o.LaunchpadTextColor = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *LaunchpadPresetTheme) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadPresetTheme) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *LaunchpadPresetTheme) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *LaunchpadPresetTheme) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetBackground returns the Background field value if set, zero value otherwise.
func (o *LaunchpadPresetTheme) GetBackground() string {
	if o == nil || o.Background == nil {
		var ret string
		return ret
	}
	return *o.Background
}

// GetBackgroundOk returns a tuple with the Background field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadPresetTheme) GetBackgroundOk() (*string, bool) {
	if o == nil || o.Background == nil {
		return nil, false
	}
	return o.Background, true
}

// HasBackground returns a boolean if a field has been set.
func (o *LaunchpadPresetTheme) HasBackground() bool {
	if o != nil && o.Background != nil {
		return true
	}

	return false
}

// SetBackground gets a reference to the given string and assigns it to the Background field.
func (o *LaunchpadPresetTheme) SetBackground(v string) {
	o.Background = &v
}

// GetAvatarBackground returns the AvatarBackground field value if set, zero value otherwise.
func (o *LaunchpadPresetTheme) GetAvatarBackground() string {
	if o == nil || o.AvatarBackground == nil {
		var ret string
		return ret
	}
	return *o.AvatarBackground
}

// GetAvatarBackgroundOk returns a tuple with the AvatarBackground field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadPresetTheme) GetAvatarBackgroundOk() (*string, bool) {
	if o == nil || o.AvatarBackground == nil {
		return nil, false
	}
	return o.AvatarBackground, true
}

// HasAvatarBackground returns a boolean if a field has been set.
func (o *LaunchpadPresetTheme) HasAvatarBackground() bool {
	if o != nil && o.AvatarBackground != nil {
		return true
	}

	return false
}

// SetAvatarBackground gets a reference to the given string and assigns it to the AvatarBackground field.
func (o *LaunchpadPresetTheme) SetAvatarBackground(v string) {
	o.AvatarBackground = &v
}

// GetThemeId returns the ThemeId field value if set, zero value otherwise.
func (o *LaunchpadPresetTheme) GetThemeId() int32 {
	if o == nil || o.ThemeId == nil {
		var ret int32
		return ret
	}
	return *o.ThemeId
}

// GetThemeIdOk returns a tuple with the ThemeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadPresetTheme) GetThemeIdOk() (*int32, bool) {
	if o == nil || o.ThemeId == nil {
		return nil, false
	}
	return o.ThemeId, true
}

// HasThemeId returns a boolean if a field has been set.
func (o *LaunchpadPresetTheme) HasThemeId() bool {
	if o != nil && o.ThemeId != nil {
		return true
	}

	return false
}

// SetThemeId gets a reference to the given int32 and assigns it to the ThemeId field.
func (o *LaunchpadPresetTheme) SetThemeId(v int32) {
	o.ThemeId = &v
}

// GetButtonBackground returns the ButtonBackground field value if set, zero value otherwise.
func (o *LaunchpadPresetTheme) GetButtonBackground() string {
	if o == nil || o.ButtonBackground == nil {
		var ret string
		return ret
	}
	return *o.ButtonBackground
}

// GetButtonBackgroundOk returns a tuple with the ButtonBackground field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadPresetTheme) GetButtonBackgroundOk() (*string, bool) {
	if o == nil || o.ButtonBackground == nil {
		return nil, false
	}
	return o.ButtonBackground, true
}

// HasButtonBackground returns a boolean if a field has been set.
func (o *LaunchpadPresetTheme) HasButtonBackground() bool {
	if o != nil && o.ButtonBackground != nil {
		return true
	}

	return false
}

// SetButtonBackground gets a reference to the given string and assigns it to the ButtonBackground field.
func (o *LaunchpadPresetTheme) SetButtonBackground(v string) {
	o.ButtonBackground = &v
}

func (o LaunchpadPresetTheme) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ButtonTextColor != nil {
		toSerialize["button_text_color"] = o.ButtonTextColor
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.AvatarTextColor != nil {
		toSerialize["avatar_text_color"] = o.AvatarTextColor
	}
	if o.LaunchpadTextColor != nil {
		toSerialize["launchpad_text_color"] = o.LaunchpadTextColor
	}
	if o.IsActive != nil {
		toSerialize["is_active"] = o.IsActive
	}
	if o.Background != nil {
		toSerialize["background"] = o.Background
	}
	if o.AvatarBackground != nil {
		toSerialize["avatar_background"] = o.AvatarBackground
	}
	if o.ThemeId != nil {
		toSerialize["theme_id"] = o.ThemeId
	}
	if o.ButtonBackground != nil {
		toSerialize["button_background"] = o.ButtonBackground
	}
	return json.Marshal(toSerialize)
}

type NullableLaunchpadPresetTheme struct {
	value *LaunchpadPresetTheme
	isSet bool
}

func (v NullableLaunchpadPresetTheme) Get() *LaunchpadPresetTheme {
	return v.value
}

func (v *NullableLaunchpadPresetTheme) Set(val *LaunchpadPresetTheme) {
	v.value = val
	v.isSet = true
}

func (v NullableLaunchpadPresetTheme) IsSet() bool {
	return v.isSet
}

func (v *NullableLaunchpadPresetTheme) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLaunchpadPresetTheme(val *LaunchpadPresetTheme) *NullableLaunchpadPresetTheme {
	return &NullableLaunchpadPresetTheme{value: val, isSet: true}
}

func (v NullableLaunchpadPresetTheme) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLaunchpadPresetTheme) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


