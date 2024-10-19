/*
OpenAPI Petstore

This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore

import (
	"encoding/json"
)

// checks if the FooGetDefaultResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FooGetDefaultResponse{}

// FooGetDefaultResponse struct for FooGetDefaultResponse
type FooGetDefaultResponse struct {
	String *Foo `json:"string,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FooGetDefaultResponse FooGetDefaultResponse

// NewFooGetDefaultResponse instantiates a new FooGetDefaultResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFooGetDefaultResponse() *FooGetDefaultResponse {
	this := FooGetDefaultResponse{}
	return &this
}

// NewFooGetDefaultResponseWithDefaults instantiates a new FooGetDefaultResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFooGetDefaultResponseWithDefaults() *FooGetDefaultResponse {
	this := FooGetDefaultResponse{}
	return &this
}

// GetString returns the String field value if set, zero value otherwise.
func (o *FooGetDefaultResponse) GetString() Foo {
	if o == nil || IsNil(o.String) {
		var ret Foo
		return ret
	}
	return *o.String
}

// GetStringOk returns a tuple with the String field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FooGetDefaultResponse) GetStringOk() (*Foo, bool) {
	if o == nil || IsNil(o.String) {
		return nil, false
	}
	return o.String, true
}

// HasString returns a boolean if a field has been set.
func (o *FooGetDefaultResponse) HasString() bool {
	if o != nil && !IsNil(o.String) {
		return true
	}

	return false
}

// SetString gets a reference to the given Foo and assigns it to the String field.
func (o *FooGetDefaultResponse) SetString(v Foo) {
	o.String = &v
}

func (o FooGetDefaultResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FooGetDefaultResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.String) {
		toSerialize["string"] = o.String
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FooGetDefaultResponse) UnmarshalJSON(data []byte) (err error) {
	varFooGetDefaultResponse := _FooGetDefaultResponse{}

	err = json.Unmarshal(data, &varFooGetDefaultResponse)

	if err != nil {
		return err
	}

	*o = FooGetDefaultResponse(varFooGetDefaultResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "string")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFooGetDefaultResponse struct {
	value *FooGetDefaultResponse
	isSet bool
}

func (v NullableFooGetDefaultResponse) Get() *FooGetDefaultResponse {
	return v.value
}

func (v *NullableFooGetDefaultResponse) Set(val *FooGetDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFooGetDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFooGetDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFooGetDefaultResponse(val *FooGetDefaultResponse) *NullableFooGetDefaultResponse {
	return &NullableFooGetDefaultResponse{value: val, isSet: true}
}

func (v NullableFooGetDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFooGetDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

