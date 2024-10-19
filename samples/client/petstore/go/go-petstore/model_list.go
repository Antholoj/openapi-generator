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

// checks if the List type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &List{}

// List struct for List
type List struct {
	Var123List *string `json:"123-list,omitempty"`
}

// NewList instantiates a new List object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewList() *List {
	this := List{}
	return &this
}

// NewListWithDefaults instantiates a new List object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWithDefaults() *List {
	this := List{}
	return &this
}

// GetVar123List returns the Var123List field value if set, zero value otherwise.
func (o *List) GetVar123List() string {
	if o == nil || IsNil(o.Var123List) {
		var ret string
		return ret
	}
	return *o.Var123List
}

// GetVar123ListOk returns a tuple with the Var123List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *List) GetVar123ListOk() (*string, bool) {
	if o == nil || IsNil(o.Var123List) {
		return nil, false
	}
	return o.Var123List, true
}

// HasVar123List returns a boolean if a field has been set.
func (o *List) HasVar123List() bool {
	if o != nil && !IsNil(o.Var123List) {
		return true
	}

	return false
}

// SetVar123List gets a reference to the given string and assigns it to the Var123List field.
func (o *List) SetVar123List(v string) {
	o.Var123List = &v
}

func (o List) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o List) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Var123List) {
		toSerialize["123-list"] = o.Var123List
	}
	return toSerialize, nil
}

type NullableList struct {
	value *List
	isSet bool
}

func (v NullableList) Get() *List {
	return v.value
}

func (v *NullableList) Set(val *List) {
	v.value = val
	v.isSet = true
}

func (v NullableList) IsSet() bool {
	return v.isSet
}

func (v *NullableList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableList(val *List) *NullableList {
	return &NullableList{value: val, isSet: true}
}

func (v NullableList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

