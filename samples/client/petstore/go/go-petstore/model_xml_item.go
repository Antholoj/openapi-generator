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

// checks if the XmlItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &XmlItem{}

// XmlItem struct for XmlItem
type XmlItem struct {
	AttributeString *string `json:"attribute_string,omitempty"`
	AttributeNumber *float32 `json:"attribute_number,omitempty"`
	AttributeInteger *int32 `json:"attribute_integer,omitempty"`
	AttributeBoolean *bool `json:"attribute_boolean,omitempty"`
	WrappedArray []int32 `json:"wrapped_array,omitempty"`
	NameString *string `json:"name_string,omitempty"`
	NameNumber *float32 `json:"name_number,omitempty"`
	NameInteger *int32 `json:"name_integer,omitempty"`
	NameBoolean *bool `json:"name_boolean,omitempty"`
	NameArray []int32 `json:"name_array,omitempty"`
	NameWrappedArray []int32 `json:"name_wrapped_array,omitempty"`
	PrefixString *string `json:"prefix_string,omitempty"`
	PrefixNumber *float32 `json:"prefix_number,omitempty"`
	PrefixInteger *int32 `json:"prefix_integer,omitempty"`
	PrefixBoolean *bool `json:"prefix_boolean,omitempty"`
	PrefixArray []int32 `json:"prefix_array,omitempty"`
	PrefixWrappedArray []int32 `json:"prefix_wrapped_array,omitempty"`
	NamespaceString *string `json:"namespace_string,omitempty"`
	NamespaceNumber *float32 `json:"namespace_number,omitempty"`
	NamespaceInteger *int32 `json:"namespace_integer,omitempty"`
	NamespaceBoolean *bool `json:"namespace_boolean,omitempty"`
	NamespaceArray []int32 `json:"namespace_array,omitempty"`
	NamespaceWrappedArray []int32 `json:"namespace_wrapped_array,omitempty"`
	PrefixNsString *string `json:"prefix_ns_string,omitempty"`
	PrefixNsNumber *float32 `json:"prefix_ns_number,omitempty"`
	PrefixNsInteger *int32 `json:"prefix_ns_integer,omitempty"`
	PrefixNsBoolean *bool `json:"prefix_ns_boolean,omitempty"`
	PrefixNsArray []int32 `json:"prefix_ns_array,omitempty"`
	PrefixNsWrappedArray []int32 `json:"prefix_ns_wrapped_array,omitempty"`
}

// NewXmlItem instantiates a new XmlItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXmlItem() *XmlItem {
	this := XmlItem{}
	return &this
}

// NewXmlItemWithDefaults instantiates a new XmlItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXmlItemWithDefaults() *XmlItem {
	this := XmlItem{}
	return &this
}

// GetAttributeString returns the AttributeString field value if set, zero value otherwise.
func (o *XmlItem) GetAttributeString() string {
	if o == nil || IsNil(o.AttributeString) {
		var ret string
		return ret
	}
	return *o.AttributeString
}

// GetAttributeStringOk returns a tuple with the AttributeString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlItem) GetAttributeStringOk() (*string, bool) {
	if o == nil || IsNil(o.AttributeString) {
		return nil, false
	}
	return o.AttributeString, true
}

// HasAttributeString returns a boolean if a field has been set.
func (o *XmlItem) HasAttributeString() bool {
	if o != nil && !IsNil(o.AttributeString) {
		return true
	}

	return false
}

// SetAttributeString gets a reference to the given string and assigns it to the AttributeString field.
func (o *XmlItem) SetAttributeString(v string) {
	o.AttributeString = &v
}

// GetAttributeNumber returns the AttributeNumber field value if set, zero value otherwise.
func (o *XmlItem) GetAttributeNumber() float32 {
	if o == nil || IsNil(o.AttributeNumber) {
		var ret float32
		return ret
	}
	return *o.AttributeNumber
}

// GetAttributeNumberOk returns a tuple with the AttributeNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlItem) GetAttributeNumberOk() (*float32, bool) {
	if o == nil || IsNil(o.AttributeNumber) {
		return nil, false
	}
	return o.AttributeNumber, true
}

// HasAttributeNumber returns a boolean if a field has been set.
func (o *XmlItem) HasAttributeNumber() bool {
	if o != nil && !IsNil(o.AttributeNumber) {
		return true
	}

	return false
}

// SetAttributeNumber gets a reference to the given float32 and assigns it to the AttributeNumber field.
func (o *XmlItem) SetAttributeNumber(v float32) {
	o.AttributeNumber = &v
}

// GetAttributeInteger returns the AttributeInteger field value if set, zero value otherwise.
func (o *XmlItem) GetAttributeInteger() int32 {
	if o == nil || IsNil(o.AttributeInteger) {
		var ret int32
		return ret
	}
	return *o.AttributeInteger
}

// GetAttributeIntegerOk returns a tuple with the AttributeInteger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlItem) GetAttributeIntegerOk() (*int32, bool) {
	if o == nil || IsNil(o.AttributeInteger) {
		return nil, false
	}
	return o.AttributeInteger, true
}

// HasAttributeInteger returns a boolean if a field has been set.
func (o *XmlItem) HasAttributeInteger() bool {
	if o != nil && !IsNil(o.AttributeInteger) {
		return true
	}

	return false
}

// SetAttributeInteger gets a reference to the given int32 and assigns it to the AttributeInteger field.
func (o *XmlItem) SetAttributeInteger(v int32) {
	o.AttributeInteger = &v
}

// GetAttributeBoolean returns the AttributeBoolean field value if set, zero value otherwise.
func (o *XmlItem) GetAttributeBoolean() bool {
	if o == nil || IsNil(o.AttributeBoolean) {
		var ret bool
		return ret
	}
	return *o.AttributeBoolean
}

// GetAttributeBooleanOk returns a tuple with the AttributeBoolean field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlItem) GetAttributeBooleanOk() (*bool, bool) {
	if o == nil || IsNil(o.AttributeBoolean) {
		return nil, false
	}
	return o.AttributeBoolean, true
}

// HasAttributeBoolean returns a boolean if a field has been set.
func (o *XmlItem) HasAttributeBoolean() bool {
	if o != nil && !IsNil(o.AttributeBoolean) {
		return true
	}

	return false
}

// SetAttributeBoolean gets a reference to the given bool and assigns it to the AttributeBoolean field.
func (o *XmlItem) SetAttributeBoolean(v bool) {
	o.AttributeBoolean = &v
}

// GetWrappedArray returns the WrappedArray field value if set, zero value otherwise.
func (o *XmlItem) GetWrappedArray() []int32 {
	if o == nil || IsNil(o.WrappedArray) {
		var ret []int32
		return ret
	}
	return o.WrappedArray
}

// GetWrappedArrayOk returns a tuple with the WrappedArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlItem) GetWrappedArrayOk() ([]int32, bool) {
	if o == nil || IsNil(o.WrappedArray) {
		return nil, false
	}
	return o.WrappedArray, true
}

// HasWrappedArray returns a boolean if a field has been set.
func (o *XmlItem) HasWrappedArray() bool {
	if o != nil && !IsNil(o.WrappedArray) {
		return true
	}

	return false
}

// SetWrappedArray gets a reference to the given []int32 and assigns it to the WrappedArray field.
func (o *XmlItem) SetWrappedArray(v []int32) {
	o.WrappedArray = v
}

// GetNameString returns the NameString field value if set, zero value otherwise.
func (o *XmlItem) GetNameString() string {
	if o == nil || IsNil(o.NameString) {
		var ret string
		return ret
	}
	return *o.NameString
}

// GetNameStringOk returns a tuple with the NameString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlItem) GetNameStringOk() (*string, bool) {
	if o == nil || IsNil(o.NameString) {
		return nil, false
	}
	return o.NameString, true
}

// HasNameString returns a boolean if a field has been set.
func (o *XmlItem) HasNameString() bool {
	if o != nil && !IsNil(o.NameString) {
		return true
	}

	return false
}

// SetNameString gets a reference to the given string and assigns it to the NameString field.
func (o *XmlItem) SetNameString(v string) {
	o.NameString = &v
}

// GetNameNumber returns the NameNumber field value if set, zero value otherwise.
func (o *XmlItem) GetNameNumber() float32 {
	if o == nil || IsNil(o.NameNumber) {
		var ret float32
		return ret
	}
	return *o.NameNumber
}

// GetNameNumberOk returns a tuple with the NameNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlItem) GetNameNumberOk() (*float32, bool) {
	if o == nil || IsNil(o.NameNumber) {
		return nil, false
	}
	return o.NameNumber, true
}

// HasNameNumber returns a boolean if a field has been set.
func (o *XmlItem) HasNameNumber() bool {
	if o != nil && !IsNil(o.NameNumber) {
		return true
	}

	return false
}

// SetNameNumber gets a reference to the given float32 and assigns it to the NameNumber field.
func (o *XmlItem) SetNameNumber(v float32) {
	o.NameNumber = &v
}

// GetNameInteger returns the NameInteger field value if set, zero value otherwise.
func (o *XmlItem) GetNameInteger() int32 {
	if o == nil || IsNil(o.NameInteger) {
		var ret int32
		return ret
	}
	return *o.NameInteger
}

// GetNameIntegerOk returns a tuple with the NameInteger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlItem) GetNameIntegerOk() (*int32, bool) {
	if o == nil || IsNil(o.NameInteger) {
		return nil, false
	}
	return o.NameInteger, true
}

// HasNameInteger returns a boolean if a field has been set.
func (o *XmlItem) HasNameInteger() bool {
	if o != nil && !IsNil(o.NameInteger) {
		return true
	}

	return false
}

// SetNameInteger gets a reference to the given int32 and assigns it to the NameInteger field.
func (o *XmlItem) SetNameInteger(v int32) {
	o.NameInteger = &v
}

// GetNameBoolean returns the NameBoolean field value if set, zero value otherwise.
func (o *XmlItem) GetNameBoolean() bool {
	if o == nil || IsNil(o.NameBoolean) {
		var ret bool
		return ret
	}
	return *o.NameBoolean
}

// GetNameBooleanOk returns a tuple with the NameBoolean field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlItem) GetNameBooleanOk() (*bool, bool) {
	if o == nil || IsNil(o.NameBoolean) {
		return nil, false
	}
	return o.NameBoolean, true
}

// HasNameBoolean returns a boolean if a field has been set.
func (o *XmlItem) HasNameBoolean() bool {
	if o != nil && !IsNil(o.NameBoolean) {
		return true
	}

	return false
}

// SetNameBoolean gets a reference to the given bool and assigns it to the NameBoolean field.
func (o *XmlItem) SetNameBoolean(v bool) {
	o.NameBoolean = &v
}

// GetNameArray returns the NameArray field value if set, zero value otherwise.
func (o *XmlItem) GetNameArray() []int32 {
	if o == nil || IsNil(o.NameArray) {
		var ret []int32
		return ret
	}
	return o.NameArray
}

// GetNameArrayOk returns a tuple with the NameArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlItem) GetNameArrayOk() ([]int32, bool) {
	if o == nil || IsNil(o.NameArray) {
		return nil, false
	}
	return o.NameArray, true
}

// HasNameArray returns a boolean if a field has been set.
func (o *XmlItem) HasNameArray() bool {
	if o != nil && !IsNil(o.NameArray) {
		return true
	}

	return false
}

// SetNameArray gets a reference to the given []int32 and assigns it to the NameArray field.
func (o *XmlItem) SetNameArray(v []int32) {
	o.NameArray = v
}

// GetNameWrappedArray returns the NameWrappedArray field value if set, zero value otherwise.
func (o *XmlItem) GetNameWrappedArray() []int32 {
	if o == nil || IsNil(o.NameWrappedArray) {
		var ret []int32
		return ret
	}
	return o.NameWrappedArray
}

// GetNameWrappedArrayOk returns a tuple with the NameWrappedArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlItem) GetNameWrappedArrayOk() ([]int32, bool) {
	if o == nil || IsNil(o.NameWrappedArray) {
		return nil, false
	}
	return o.NameWrappedArray, true
}

// HasNameWrappedArray returns a boolean if a field has been set.
func (o *XmlItem) HasNameWrappedArray() bool {
	if o != nil && !IsNil(o.NameWrappedArray) {
		return true
	}

	return false
}

// SetNameWrappedArray gets a reference to the given []int32 and assigns it to the NameWrappedArray field.
func (o *XmlItem) SetNameWrappedArray(v []int32) {
	o.NameWrappedArray = v
}

// GetPrefixString returns the PrefixString field value if set, zero value otherwise.
func (o *XmlItem) GetPrefixString() string {
	if o == nil || IsNil(o.PrefixString) {
		var ret string
		return ret
	}
	return *o.PrefixString
}

// GetPrefixStringOk returns a tuple with the PrefixString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlItem) GetPrefixStringOk() (*string, bool) {
	if o == nil || IsNil(o.PrefixString) {
		return nil, false
	}
	return o.PrefixString, true
}

// HasPrefixString returns a boolean if a field has been set.
func (o *XmlItem) HasPrefixString() bool {
	if o != nil && !IsNil(o.PrefixString) {
		return true
	}

	return false
}

// SetPrefixString gets a reference to the given string and assigns it to the PrefixString field.
func (o *XmlItem) SetPrefixString(v string) {
	o.PrefixString = &v
}

// GetPrefixNumber returns the PrefixNumber field value if set, zero value otherwise.
func (o *XmlItem) GetPrefixNumber() float32 {
	if o == nil || IsNil(o.PrefixNumber) {
		var ret float32
		return ret
	}
	return *o.PrefixNumber
}

// GetPrefixNumberOk returns a tuple with the PrefixNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlItem) GetPrefixNumberOk() (*float32, bool) {
	if o == nil || IsNil(o.PrefixNumber) {
		return nil, false
	}
	return o.PrefixNumber, true
}

// HasPrefixNumber returns a boolean if a field has been set.
func (o *XmlItem) HasPrefixNumber() bool {
	if o != nil && !IsNil(o.PrefixNumber) {
		return true
	}

	return false
}

// SetPrefixNumber gets a reference to the given float32 and assigns it to the PrefixNumber field.
func (o *XmlItem) SetPrefixNumber(v float32) {
	o.PrefixNumber = &v
}

// GetPrefixInteger returns the PrefixInteger field value if set, zero value otherwise.
func (o *XmlItem) GetPrefixInteger() int32 {
	if o == nil || IsNil(o.PrefixInteger) {
		var ret int32
		return ret
	}
	return *o.PrefixInteger
}

// GetPrefixIntegerOk returns a tuple with the PrefixInteger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlItem) GetPrefixIntegerOk() (*int32, bool) {
	if o == nil || IsNil(o.PrefixInteger) {
		return nil, false
	}
	return o.PrefixInteger, true
}

// HasPrefixInteger returns a boolean if a field has been set.
func (o *XmlItem) HasPrefixInteger() bool {
	if o != nil && !IsNil(o.PrefixInteger) {
		return true
	}

	return false
}

// SetPrefixInteger gets a reference to the given int32 and assigns it to the PrefixInteger field.
func (o *XmlItem) SetPrefixInteger(v int32) {
	o.PrefixInteger = &v
}

// GetPrefixBoolean returns the PrefixBoolean field value if set, zero value otherwise.
func (o *XmlItem) GetPrefixBoolean() bool {
	if o == nil || IsNil(o.PrefixBoolean) {
		var ret bool
		return ret
	}
	return *o.PrefixBoolean
}

// GetPrefixBooleanOk returns a tuple with the PrefixBoolean field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlItem) GetPrefixBooleanOk() (*bool, bool) {
	if o == nil || IsNil(o.PrefixBoolean) {
		return nil, false
	}
	return o.PrefixBoolean, true
}

// HasPrefixBoolean returns a boolean if a field has been set.
func (o *XmlItem) HasPrefixBoolean() bool {
	if o != nil && !IsNil(o.PrefixBoolean) {
		return true
	}

	return false
}

// SetPrefixBoolean gets a reference to the given bool and assigns it to the PrefixBoolean field.
func (o *XmlItem) SetPrefixBoolean(v bool) {
	o.PrefixBoolean = &v
}

// GetPrefixArray returns the PrefixArray field value if set, zero value otherwise.
func (o *XmlItem) GetPrefixArray() []int32 {
	if o == nil || IsNil(o.PrefixArray) {
		var ret []int32
		return ret
	}
	return o.PrefixArray
}

// GetPrefixArrayOk returns a tuple with the PrefixArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlItem) GetPrefixArrayOk() ([]int32, bool) {
	if o == nil || IsNil(o.PrefixArray) {
		return nil, false
	}
	return o.PrefixArray, true
}

// HasPrefixArray returns a boolean if a field has been set.
func (o *XmlItem) HasPrefixArray() bool {
	if o != nil && !IsNil(o.PrefixArray) {
		return true
	}

	return false
}

// SetPrefixArray gets a reference to the given []int32 and assigns it to the PrefixArray field.
func (o *XmlItem) SetPrefixArray(v []int32) {
	o.PrefixArray = v
}

// GetPrefixWrappedArray returns the PrefixWrappedArray field value if set, zero value otherwise.
func (o *XmlItem) GetPrefixWrappedArray() []int32 {
	if o == nil || IsNil(o.PrefixWrappedArray) {
		var ret []int32
		return ret
	}
	return o.PrefixWrappedArray
}

// GetPrefixWrappedArrayOk returns a tuple with the PrefixWrappedArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlItem) GetPrefixWrappedArrayOk() ([]int32, bool) {
	if o == nil || IsNil(o.PrefixWrappedArray) {
		return nil, false
	}
	return o.PrefixWrappedArray, true
}

// HasPrefixWrappedArray returns a boolean if a field has been set.
func (o *XmlItem) HasPrefixWrappedArray() bool {
	if o != nil && !IsNil(o.PrefixWrappedArray) {
		return true
	}

	return false
}

// SetPrefixWrappedArray gets a reference to the given []int32 and assigns it to the PrefixWrappedArray field.
func (o *XmlItem) SetPrefixWrappedArray(v []int32) {
	o.PrefixWrappedArray = v
}

// GetNamespaceString returns the NamespaceString field value if set, zero value otherwise.
func (o *XmlItem) GetNamespaceString() string {
	if o == nil || IsNil(o.NamespaceString) {
		var ret string
		return ret
	}
	return *o.NamespaceString
}

// GetNamespaceStringOk returns a tuple with the NamespaceString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlItem) GetNamespaceStringOk() (*string, bool) {
	if o == nil || IsNil(o.NamespaceString) {
		return nil, false
	}
	return o.NamespaceString, true
}

// HasNamespaceString returns a boolean if a field has been set.
func (o *XmlItem) HasNamespaceString() bool {
	if o != nil && !IsNil(o.NamespaceString) {
		return true
	}

	return false
}

// SetNamespaceString gets a reference to the given string and assigns it to the NamespaceString field.
func (o *XmlItem) SetNamespaceString(v string) {
	o.NamespaceString = &v
}

// GetNamespaceNumber returns the NamespaceNumber field value if set, zero value otherwise.
func (o *XmlItem) GetNamespaceNumber() float32 {
	if o == nil || IsNil(o.NamespaceNumber) {
		var ret float32
		return ret
	}
	return *o.NamespaceNumber
}

// GetNamespaceNumberOk returns a tuple with the NamespaceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlItem) GetNamespaceNumberOk() (*float32, bool) {
	if o == nil || IsNil(o.NamespaceNumber) {
		return nil, false
	}
	return o.NamespaceNumber, true
}

// HasNamespaceNumber returns a boolean if a field has been set.
func (o *XmlItem) HasNamespaceNumber() bool {
	if o != nil && !IsNil(o.NamespaceNumber) {
		return true
	}

	return false
}

// SetNamespaceNumber gets a reference to the given float32 and assigns it to the NamespaceNumber field.
func (o *XmlItem) SetNamespaceNumber(v float32) {
	o.NamespaceNumber = &v
}

// GetNamespaceInteger returns the NamespaceInteger field value if set, zero value otherwise.
func (o *XmlItem) GetNamespaceInteger() int32 {
	if o == nil || IsNil(o.NamespaceInteger) {
		var ret int32
		return ret
	}
	return *o.NamespaceInteger
}

// GetNamespaceIntegerOk returns a tuple with the NamespaceInteger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlItem) GetNamespaceIntegerOk() (*int32, bool) {
	if o == nil || IsNil(o.NamespaceInteger) {
		return nil, false
	}
	return o.NamespaceInteger, true
}

// HasNamespaceInteger returns a boolean if a field has been set.
func (o *XmlItem) HasNamespaceInteger() bool {
	if o != nil && !IsNil(o.NamespaceInteger) {
		return true
	}

	return false
}

// SetNamespaceInteger gets a reference to the given int32 and assigns it to the NamespaceInteger field.
func (o *XmlItem) SetNamespaceInteger(v int32) {
	o.NamespaceInteger = &v
}

// GetNamespaceBoolean returns the NamespaceBoolean field value if set, zero value otherwise.
func (o *XmlItem) GetNamespaceBoolean() bool {
	if o == nil || IsNil(o.NamespaceBoolean) {
		var ret bool
		return ret
	}
	return *o.NamespaceBoolean
}

// GetNamespaceBooleanOk returns a tuple with the NamespaceBoolean field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlItem) GetNamespaceBooleanOk() (*bool, bool) {
	if o == nil || IsNil(o.NamespaceBoolean) {
		return nil, false
	}
	return o.NamespaceBoolean, true
}

// HasNamespaceBoolean returns a boolean if a field has been set.
func (o *XmlItem) HasNamespaceBoolean() bool {
	if o != nil && !IsNil(o.NamespaceBoolean) {
		return true
	}

	return false
}

// SetNamespaceBoolean gets a reference to the given bool and assigns it to the NamespaceBoolean field.
func (o *XmlItem) SetNamespaceBoolean(v bool) {
	o.NamespaceBoolean = &v
}

// GetNamespaceArray returns the NamespaceArray field value if set, zero value otherwise.
func (o *XmlItem) GetNamespaceArray() []int32 {
	if o == nil || IsNil(o.NamespaceArray) {
		var ret []int32
		return ret
	}
	return o.NamespaceArray
}

// GetNamespaceArrayOk returns a tuple with the NamespaceArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlItem) GetNamespaceArrayOk() ([]int32, bool) {
	if o == nil || IsNil(o.NamespaceArray) {
		return nil, false
	}
	return o.NamespaceArray, true
}

// HasNamespaceArray returns a boolean if a field has been set.
func (o *XmlItem) HasNamespaceArray() bool {
	if o != nil && !IsNil(o.NamespaceArray) {
		return true
	}

	return false
}

// SetNamespaceArray gets a reference to the given []int32 and assigns it to the NamespaceArray field.
func (o *XmlItem) SetNamespaceArray(v []int32) {
	o.NamespaceArray = v
}

// GetNamespaceWrappedArray returns the NamespaceWrappedArray field value if set, zero value otherwise.
func (o *XmlItem) GetNamespaceWrappedArray() []int32 {
	if o == nil || IsNil(o.NamespaceWrappedArray) {
		var ret []int32
		return ret
	}
	return o.NamespaceWrappedArray
}

// GetNamespaceWrappedArrayOk returns a tuple with the NamespaceWrappedArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlItem) GetNamespaceWrappedArrayOk() ([]int32, bool) {
	if o == nil || IsNil(o.NamespaceWrappedArray) {
		return nil, false
	}
	return o.NamespaceWrappedArray, true
}

// HasNamespaceWrappedArray returns a boolean if a field has been set.
func (o *XmlItem) HasNamespaceWrappedArray() bool {
	if o != nil && !IsNil(o.NamespaceWrappedArray) {
		return true
	}

	return false
}

// SetNamespaceWrappedArray gets a reference to the given []int32 and assigns it to the NamespaceWrappedArray field.
func (o *XmlItem) SetNamespaceWrappedArray(v []int32) {
	o.NamespaceWrappedArray = v
}

// GetPrefixNsString returns the PrefixNsString field value if set, zero value otherwise.
func (o *XmlItem) GetPrefixNsString() string {
	if o == nil || IsNil(o.PrefixNsString) {
		var ret string
		return ret
	}
	return *o.PrefixNsString
}

// GetPrefixNsStringOk returns a tuple with the PrefixNsString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlItem) GetPrefixNsStringOk() (*string, bool) {
	if o == nil || IsNil(o.PrefixNsString) {
		return nil, false
	}
	return o.PrefixNsString, true
}

// HasPrefixNsString returns a boolean if a field has been set.
func (o *XmlItem) HasPrefixNsString() bool {
	if o != nil && !IsNil(o.PrefixNsString) {
		return true
	}

	return false
}

// SetPrefixNsString gets a reference to the given string and assigns it to the PrefixNsString field.
func (o *XmlItem) SetPrefixNsString(v string) {
	o.PrefixNsString = &v
}

// GetPrefixNsNumber returns the PrefixNsNumber field value if set, zero value otherwise.
func (o *XmlItem) GetPrefixNsNumber() float32 {
	if o == nil || IsNil(o.PrefixNsNumber) {
		var ret float32
		return ret
	}
	return *o.PrefixNsNumber
}

// GetPrefixNsNumberOk returns a tuple with the PrefixNsNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlItem) GetPrefixNsNumberOk() (*float32, bool) {
	if o == nil || IsNil(o.PrefixNsNumber) {
		return nil, false
	}
	return o.PrefixNsNumber, true
}

// HasPrefixNsNumber returns a boolean if a field has been set.
func (o *XmlItem) HasPrefixNsNumber() bool {
	if o != nil && !IsNil(o.PrefixNsNumber) {
		return true
	}

	return false
}

// SetPrefixNsNumber gets a reference to the given float32 and assigns it to the PrefixNsNumber field.
func (o *XmlItem) SetPrefixNsNumber(v float32) {
	o.PrefixNsNumber = &v
}

// GetPrefixNsInteger returns the PrefixNsInteger field value if set, zero value otherwise.
func (o *XmlItem) GetPrefixNsInteger() int32 {
	if o == nil || IsNil(o.PrefixNsInteger) {
		var ret int32
		return ret
	}
	return *o.PrefixNsInteger
}

// GetPrefixNsIntegerOk returns a tuple with the PrefixNsInteger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlItem) GetPrefixNsIntegerOk() (*int32, bool) {
	if o == nil || IsNil(o.PrefixNsInteger) {
		return nil, false
	}
	return o.PrefixNsInteger, true
}

// HasPrefixNsInteger returns a boolean if a field has been set.
func (o *XmlItem) HasPrefixNsInteger() bool {
	if o != nil && !IsNil(o.PrefixNsInteger) {
		return true
	}

	return false
}

// SetPrefixNsInteger gets a reference to the given int32 and assigns it to the PrefixNsInteger field.
func (o *XmlItem) SetPrefixNsInteger(v int32) {
	o.PrefixNsInteger = &v
}

// GetPrefixNsBoolean returns the PrefixNsBoolean field value if set, zero value otherwise.
func (o *XmlItem) GetPrefixNsBoolean() bool {
	if o == nil || IsNil(o.PrefixNsBoolean) {
		var ret bool
		return ret
	}
	return *o.PrefixNsBoolean
}

// GetPrefixNsBooleanOk returns a tuple with the PrefixNsBoolean field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlItem) GetPrefixNsBooleanOk() (*bool, bool) {
	if o == nil || IsNil(o.PrefixNsBoolean) {
		return nil, false
	}
	return o.PrefixNsBoolean, true
}

// HasPrefixNsBoolean returns a boolean if a field has been set.
func (o *XmlItem) HasPrefixNsBoolean() bool {
	if o != nil && !IsNil(o.PrefixNsBoolean) {
		return true
	}

	return false
}

// SetPrefixNsBoolean gets a reference to the given bool and assigns it to the PrefixNsBoolean field.
func (o *XmlItem) SetPrefixNsBoolean(v bool) {
	o.PrefixNsBoolean = &v
}

// GetPrefixNsArray returns the PrefixNsArray field value if set, zero value otherwise.
func (o *XmlItem) GetPrefixNsArray() []int32 {
	if o == nil || IsNil(o.PrefixNsArray) {
		var ret []int32
		return ret
	}
	return o.PrefixNsArray
}

// GetPrefixNsArrayOk returns a tuple with the PrefixNsArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlItem) GetPrefixNsArrayOk() ([]int32, bool) {
	if o == nil || IsNil(o.PrefixNsArray) {
		return nil, false
	}
	return o.PrefixNsArray, true
}

// HasPrefixNsArray returns a boolean if a field has been set.
func (o *XmlItem) HasPrefixNsArray() bool {
	if o != nil && !IsNil(o.PrefixNsArray) {
		return true
	}

	return false
}

// SetPrefixNsArray gets a reference to the given []int32 and assigns it to the PrefixNsArray field.
func (o *XmlItem) SetPrefixNsArray(v []int32) {
	o.PrefixNsArray = v
}

// GetPrefixNsWrappedArray returns the PrefixNsWrappedArray field value if set, zero value otherwise.
func (o *XmlItem) GetPrefixNsWrappedArray() []int32 {
	if o == nil || IsNil(o.PrefixNsWrappedArray) {
		var ret []int32
		return ret
	}
	return o.PrefixNsWrappedArray
}

// GetPrefixNsWrappedArrayOk returns a tuple with the PrefixNsWrappedArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlItem) GetPrefixNsWrappedArrayOk() ([]int32, bool) {
	if o == nil || IsNil(o.PrefixNsWrappedArray) {
		return nil, false
	}
	return o.PrefixNsWrappedArray, true
}

// HasPrefixNsWrappedArray returns a boolean if a field has been set.
func (o *XmlItem) HasPrefixNsWrappedArray() bool {
	if o != nil && !IsNil(o.PrefixNsWrappedArray) {
		return true
	}

	return false
}

// SetPrefixNsWrappedArray gets a reference to the given []int32 and assigns it to the PrefixNsWrappedArray field.
func (o *XmlItem) SetPrefixNsWrappedArray(v []int32) {
	o.PrefixNsWrappedArray = v
}

func (o XmlItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o XmlItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttributeString) {
		toSerialize["attribute_string"] = o.AttributeString
	}
	if !IsNil(o.AttributeNumber) {
		toSerialize["attribute_number"] = o.AttributeNumber
	}
	if !IsNil(o.AttributeInteger) {
		toSerialize["attribute_integer"] = o.AttributeInteger
	}
	if !IsNil(o.AttributeBoolean) {
		toSerialize["attribute_boolean"] = o.AttributeBoolean
	}
	if !IsNil(o.WrappedArray) {
		toSerialize["wrapped_array"] = o.WrappedArray
	}
	if !IsNil(o.NameString) {
		toSerialize["name_string"] = o.NameString
	}
	if !IsNil(o.NameNumber) {
		toSerialize["name_number"] = o.NameNumber
	}
	if !IsNil(o.NameInteger) {
		toSerialize["name_integer"] = o.NameInteger
	}
	if !IsNil(o.NameBoolean) {
		toSerialize["name_boolean"] = o.NameBoolean
	}
	if !IsNil(o.NameArray) {
		toSerialize["name_array"] = o.NameArray
	}
	if !IsNil(o.NameWrappedArray) {
		toSerialize["name_wrapped_array"] = o.NameWrappedArray
	}
	if !IsNil(o.PrefixString) {
		toSerialize["prefix_string"] = o.PrefixString
	}
	if !IsNil(o.PrefixNumber) {
		toSerialize["prefix_number"] = o.PrefixNumber
	}
	if !IsNil(o.PrefixInteger) {
		toSerialize["prefix_integer"] = o.PrefixInteger
	}
	if !IsNil(o.PrefixBoolean) {
		toSerialize["prefix_boolean"] = o.PrefixBoolean
	}
	if !IsNil(o.PrefixArray) {
		toSerialize["prefix_array"] = o.PrefixArray
	}
	if !IsNil(o.PrefixWrappedArray) {
		toSerialize["prefix_wrapped_array"] = o.PrefixWrappedArray
	}
	if !IsNil(o.NamespaceString) {
		toSerialize["namespace_string"] = o.NamespaceString
	}
	if !IsNil(o.NamespaceNumber) {
		toSerialize["namespace_number"] = o.NamespaceNumber
	}
	if !IsNil(o.NamespaceInteger) {
		toSerialize["namespace_integer"] = o.NamespaceInteger
	}
	if !IsNil(o.NamespaceBoolean) {
		toSerialize["namespace_boolean"] = o.NamespaceBoolean
	}
	if !IsNil(o.NamespaceArray) {
		toSerialize["namespace_array"] = o.NamespaceArray
	}
	if !IsNil(o.NamespaceWrappedArray) {
		toSerialize["namespace_wrapped_array"] = o.NamespaceWrappedArray
	}
	if !IsNil(o.PrefixNsString) {
		toSerialize["prefix_ns_string"] = o.PrefixNsString
	}
	if !IsNil(o.PrefixNsNumber) {
		toSerialize["prefix_ns_number"] = o.PrefixNsNumber
	}
	if !IsNil(o.PrefixNsInteger) {
		toSerialize["prefix_ns_integer"] = o.PrefixNsInteger
	}
	if !IsNil(o.PrefixNsBoolean) {
		toSerialize["prefix_ns_boolean"] = o.PrefixNsBoolean
	}
	if !IsNil(o.PrefixNsArray) {
		toSerialize["prefix_ns_array"] = o.PrefixNsArray
	}
	if !IsNil(o.PrefixNsWrappedArray) {
		toSerialize["prefix_ns_wrapped_array"] = o.PrefixNsWrappedArray
	}
	return toSerialize, nil
}

type NullableXmlItem struct {
	value *XmlItem
	isSet bool
}

func (v NullableXmlItem) Get() *XmlItem {
	return v.value
}

func (v *NullableXmlItem) Set(val *XmlItem) {
	v.value = val
	v.isSet = true
}

func (v NullableXmlItem) IsSet() bool {
	return v.isSet
}

func (v *NullableXmlItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXmlItem(val *XmlItem) *NullableXmlItem {
	return &NullableXmlItem{value: val, isSet: true}
}

func (v NullableXmlItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXmlItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

