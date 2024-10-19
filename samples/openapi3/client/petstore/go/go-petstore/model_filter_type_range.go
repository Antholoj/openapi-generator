/*
OpenAPI Petstore

This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore

import (
	"encoding/json"
	"fmt"
)

// checks if the FilterTypeRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilterTypeRange{}

// FilterTypeRange struct for FilterTypeRange
type FilterTypeRange struct {
	Type string `json:"type"`
	Data []string `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FilterTypeRange FilterTypeRange

// NewFilterTypeRange instantiates a new FilterTypeRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilterTypeRange(type_ string) *FilterTypeRange {
	this := FilterTypeRange{}
	this.Type = type_
	return &this
}

// NewFilterTypeRangeWithDefaults instantiates a new FilterTypeRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilterTypeRangeWithDefaults() *FilterTypeRange {
	this := FilterTypeRange{}
	return &this
}

// GetType returns the Type field value
func (o *FilterTypeRange) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FilterTypeRange) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FilterTypeRange) SetType(v string) {
	o.Type = v
}


// GetData returns the Data field value if set, zero value otherwise.
func (o *FilterTypeRange) GetData() []string {
	if o == nil || IsNil(o.Data) {
		var ret []string
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterTypeRange) GetDataOk() ([]string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *FilterTypeRange) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []string and assigns it to the Data field.
func (o *FilterTypeRange) SetData(v []string) {
	o.Data = v
}

func (o FilterTypeRange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilterTypeRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FilterTypeRange) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{} {
	}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == ""{
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil{
			return err
		}
	}
	varFilterTypeRange := _FilterTypeRange{}

	err = json.Unmarshal(data, &varFilterTypeRange)

	if err != nil {
		return err
	}

	*o = FilterTypeRange(varFilterTypeRange)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFilterTypeRange struct {
	value *FilterTypeRange
	isSet bool
}

func (v NullableFilterTypeRange) Get() *FilterTypeRange {
	return v.value
}

func (v *NullableFilterTypeRange) Set(val *FilterTypeRange) {
	v.value = val
	v.isSet = true
}

func (v NullableFilterTypeRange) IsSet() bool {
	return v.isSet
}

func (v *NullableFilterTypeRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilterTypeRange(val *FilterTypeRange) *NullableFilterTypeRange {
	return &NullableFilterTypeRange{value: val, isSet: true}
}

func (v NullableFilterTypeRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilterTypeRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

