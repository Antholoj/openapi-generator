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

// checks if the AppleReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppleReq{}

// AppleReq struct for AppleReq
type AppleReq struct {
	Cultivar string `json:"cultivar"`
	Mealy *bool `json:"mealy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppleReq AppleReq

// NewAppleReq instantiates a new AppleReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppleReq(cultivar string) *AppleReq {
	this := AppleReq{}
	this.Cultivar = cultivar
	return &this
}

// NewAppleReqWithDefaults instantiates a new AppleReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppleReqWithDefaults() *AppleReq {
	this := AppleReq{}
	return &this
}

// GetCultivar returns the Cultivar field value
func (o *AppleReq) GetCultivar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cultivar
}

// GetCultivarOk returns a tuple with the Cultivar field value
// and a boolean to check if the value has been set.
func (o *AppleReq) GetCultivarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cultivar, true
}

// SetCultivar sets field value
func (o *AppleReq) SetCultivar(v string) {
	o.Cultivar = v
}


// GetMealy returns the Mealy field value if set, zero value otherwise.
func (o *AppleReq) GetMealy() bool {
	if o == nil || IsNil(o.Mealy) {
		var ret bool
		return ret
	}
	return *o.Mealy
}

// GetMealyOk returns a tuple with the Mealy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleReq) GetMealyOk() (*bool, bool) {
	if o == nil || IsNil(o.Mealy) {
		return nil, false
	}
	return o.Mealy, true
}

// HasMealy returns a boolean if a field has been set.
func (o *AppleReq) HasMealy() bool {
	if o != nil && !IsNil(o.Mealy) {
		return true
	}

	return false
}

// SetMealy gets a reference to the given bool and assigns it to the Mealy field.
func (o *AppleReq) SetMealy(v bool) {
	o.Mealy = &v
}

func (o AppleReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppleReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cultivar"] = o.Cultivar
	if !IsNil(o.Mealy) {
		toSerialize["mealy"] = o.Mealy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppleReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cultivar",
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
	varAppleReq := _AppleReq{}

	err = json.Unmarshal(data, &varAppleReq)

	if err != nil {
		return err
	}

	*o = AppleReq(varAppleReq)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cultivar")
		delete(additionalProperties, "mealy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppleReq struct {
	value *AppleReq
	isSet bool
}

func (v NullableAppleReq) Get() *AppleReq {
	return v.value
}

func (v *NullableAppleReq) Set(val *AppleReq) {
	v.value = val
	v.isSet = true
}

func (v NullableAppleReq) IsSet() bool {
	return v.isSet
}

func (v *NullableAppleReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppleReq(val *AppleReq) *NullableAppleReq {
	return &NullableAppleReq{value: val, isSet: true}
}

func (v NullableAppleReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppleReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

