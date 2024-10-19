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

// checks if the Animal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Animal{}

// Animal struct for Animal
type Animal struct {
	ClassName string `json:"className"`
	Color *string `json:"color,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Animal Animal

// NewAnimal instantiates a new Animal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnimal(className string) *Animal {
	this := Animal{}
	this.ClassName = className
	var color string = "red"
	this.Color = &color
	return &this
}

// NewAnimalWithDefaults instantiates a new Animal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnimalWithDefaults() *Animal {
	this := Animal{}
	var color string = "red"
	this.Color = &color
	return &this
}

// GetClassName returns the ClassName field value
func (o *Animal) GetClassName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassName
}

// GetClassNameOk returns a tuple with the ClassName field value
// and a boolean to check if the value has been set.
func (o *Animal) GetClassNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassName, true
}

// SetClassName sets field value
func (o *Animal) SetClassName(v string) {
	o.ClassName = v
}


// GetColor returns the Color field value if set, zero value otherwise.
func (o *Animal) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Animal) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *Animal) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *Animal) SetColor(v string) {
	o.Color = &v
}

func (o Animal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Animal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["className"] = o.ClassName
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Animal) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"className",
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
	varAnimal := _Animal{}

	err = json.Unmarshal(data, &varAnimal)

	if err != nil {
		return err
	}

	*o = Animal(varAnimal)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "className")
		delete(additionalProperties, "color")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAnimal struct {
	value *Animal
	isSet bool
}

func (v NullableAnimal) Get() *Animal {
	return v.value
}

func (v *NullableAnimal) Set(val *Animal) {
	v.value = val
	v.isSet = true
}

func (v NullableAnimal) IsSet() bool {
	return v.isSet
}

func (v *NullableAnimal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnimal(val *Animal) *NullableAnimal {
	return &NullableAnimal{value: val, isSet: true}
}

func (v NullableAnimal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnimal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

