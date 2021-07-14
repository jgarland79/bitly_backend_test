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

// NSQProducer struct for NSQProducer
type NSQProducer struct {
	HttpPort *int32 `json:"http_port,omitempty"`
	BroadcastAddress *string `json:"broadcast_address,omitempty"`
	TcpPort *int32 `json:"tcp_port,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
	RemoteAddress *string `json:"remote_address,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewNSQProducer instantiates a new NSQProducer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNSQProducer() *NSQProducer {
	this := NSQProducer{}
	return &this
}

// NewNSQProducerWithDefaults instantiates a new NSQProducer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNSQProducerWithDefaults() *NSQProducer {
	this := NSQProducer{}
	return &this
}

// GetHttpPort returns the HttpPort field value if set, zero value otherwise.
func (o *NSQProducer) GetHttpPort() int32 {
	if o == nil || o.HttpPort == nil {
		var ret int32
		return ret
	}
	return *o.HttpPort
}

// GetHttpPortOk returns a tuple with the HttpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NSQProducer) GetHttpPortOk() (*int32, bool) {
	if o == nil || o.HttpPort == nil {
		return nil, false
	}
	return o.HttpPort, true
}

// HasHttpPort returns a boolean if a field has been set.
func (o *NSQProducer) HasHttpPort() bool {
	if o != nil && o.HttpPort != nil {
		return true
	}

	return false
}

// SetHttpPort gets a reference to the given int32 and assigns it to the HttpPort field.
func (o *NSQProducer) SetHttpPort(v int32) {
	o.HttpPort = &v
}

// GetBroadcastAddress returns the BroadcastAddress field value if set, zero value otherwise.
func (o *NSQProducer) GetBroadcastAddress() string {
	if o == nil || o.BroadcastAddress == nil {
		var ret string
		return ret
	}
	return *o.BroadcastAddress
}

// GetBroadcastAddressOk returns a tuple with the BroadcastAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NSQProducer) GetBroadcastAddressOk() (*string, bool) {
	if o == nil || o.BroadcastAddress == nil {
		return nil, false
	}
	return o.BroadcastAddress, true
}

// HasBroadcastAddress returns a boolean if a field has been set.
func (o *NSQProducer) HasBroadcastAddress() bool {
	if o != nil && o.BroadcastAddress != nil {
		return true
	}

	return false
}

// SetBroadcastAddress gets a reference to the given string and assigns it to the BroadcastAddress field.
func (o *NSQProducer) SetBroadcastAddress(v string) {
	o.BroadcastAddress = &v
}

// GetTcpPort returns the TcpPort field value if set, zero value otherwise.
func (o *NSQProducer) GetTcpPort() int32 {
	if o == nil || o.TcpPort == nil {
		var ret int32
		return ret
	}
	return *o.TcpPort
}

// GetTcpPortOk returns a tuple with the TcpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NSQProducer) GetTcpPortOk() (*int32, bool) {
	if o == nil || o.TcpPort == nil {
		return nil, false
	}
	return o.TcpPort, true
}

// HasTcpPort returns a boolean if a field has been set.
func (o *NSQProducer) HasTcpPort() bool {
	if o != nil && o.TcpPort != nil {
		return true
	}

	return false
}

// SetTcpPort gets a reference to the given int32 and assigns it to the TcpPort field.
func (o *NSQProducer) SetTcpPort(v int32) {
	o.TcpPort = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *NSQProducer) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NSQProducer) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *NSQProducer) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *NSQProducer) SetHostname(v string) {
	o.Hostname = &v
}

// GetRemoteAddress returns the RemoteAddress field value if set, zero value otherwise.
func (o *NSQProducer) GetRemoteAddress() string {
	if o == nil || o.RemoteAddress == nil {
		var ret string
		return ret
	}
	return *o.RemoteAddress
}

// GetRemoteAddressOk returns a tuple with the RemoteAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NSQProducer) GetRemoteAddressOk() (*string, bool) {
	if o == nil || o.RemoteAddress == nil {
		return nil, false
	}
	return o.RemoteAddress, true
}

// HasRemoteAddress returns a boolean if a field has been set.
func (o *NSQProducer) HasRemoteAddress() bool {
	if o != nil && o.RemoteAddress != nil {
		return true
	}

	return false
}

// SetRemoteAddress gets a reference to the given string and assigns it to the RemoteAddress field.
func (o *NSQProducer) SetRemoteAddress(v string) {
	o.RemoteAddress = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *NSQProducer) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NSQProducer) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *NSQProducer) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *NSQProducer) SetVersion(v string) {
	o.Version = &v
}

func (o NSQProducer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HttpPort != nil {
		toSerialize["http_port"] = o.HttpPort
	}
	if o.BroadcastAddress != nil {
		toSerialize["broadcast_address"] = o.BroadcastAddress
	}
	if o.TcpPort != nil {
		toSerialize["tcp_port"] = o.TcpPort
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.RemoteAddress != nil {
		toSerialize["remote_address"] = o.RemoteAddress
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableNSQProducer struct {
	value *NSQProducer
	isSet bool
}

func (v NullableNSQProducer) Get() *NSQProducer {
	return v.value
}

func (v *NullableNSQProducer) Set(val *NSQProducer) {
	v.value = val
	v.isSet = true
}

func (v NullableNSQProducer) IsSet() bool {
	return v.isSet
}

func (v *NullableNSQProducer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNSQProducer(val *NSQProducer) *NullableNSQProducer {
	return &NullableNSQProducer{value: val, isSet: true}
}

func (v NullableNSQProducer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNSQProducer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

