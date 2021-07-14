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

// InvitationsError struct for InvitationsError
type InvitationsError struct {
	Message *string `json:"message,omitempty"`
	Emails *[]string `json:"emails,omitempty"`
}

// NewInvitationsError instantiates a new InvitationsError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitationsError() *InvitationsError {
	this := InvitationsError{}
	return &this
}

// NewInvitationsErrorWithDefaults instantiates a new InvitationsError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitationsErrorWithDefaults() *InvitationsError {
	this := InvitationsError{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InvitationsError) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationsError) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InvitationsError) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *InvitationsError) SetMessage(v string) {
	o.Message = &v
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *InvitationsError) GetEmails() []string {
	if o == nil || o.Emails == nil {
		var ret []string
		return ret
	}
	return *o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationsError) GetEmailsOk() (*[]string, bool) {
	if o == nil || o.Emails == nil {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *InvitationsError) HasEmails() bool {
	if o != nil && o.Emails != nil {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []string and assigns it to the Emails field.
func (o *InvitationsError) SetEmails(v []string) {
	o.Emails = &v
}

func (o InvitationsError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Emails != nil {
		toSerialize["emails"] = o.Emails
	}
	return json.Marshal(toSerialize)
}

type NullableInvitationsError struct {
	value *InvitationsError
	isSet bool
}

func (v NullableInvitationsError) Get() *InvitationsError {
	return v.value
}

func (v *NullableInvitationsError) Set(val *InvitationsError) {
	v.value = val
	v.isSet = true
}

func (v NullableInvitationsError) IsSet() bool {
	return v.isSet
}

func (v *NullableInvitationsError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvitationsError(val *InvitationsError) *NullableInvitationsError {
	return &NullableInvitationsError{value: val, isSet: true}
}

func (v NullableInvitationsError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvitationsError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

