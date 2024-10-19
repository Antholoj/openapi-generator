/*
OpenAPI Petstore

This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// OneOfWithComplexType - struct for OneOfWithComplexType
type OneOfWithComplexType struct {
	ArrayOfString *[]string
	String *string
}

// []stringAsOneOfWithComplexType is a convenience function that returns []string wrapped in OneOfWithComplexType
func ArrayOfStringAsOneOfWithComplexType(v *[]string) OneOfWithComplexType {
	return OneOfWithComplexType{
		ArrayOfString: v,
	}
}

// stringAsOneOfWithComplexType is a convenience function that returns string wrapped in OneOfWithComplexType
func StringAsOneOfWithComplexType(v *string) OneOfWithComplexType {
	return OneOfWithComplexType{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *OneOfWithComplexType) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfString
	err = newStrictDecoder(data).Decode(&dst.ArrayOfString)
	if err == nil {
		jsonArrayOfString, _ := json.Marshal(dst.ArrayOfString)
		if string(jsonArrayOfString) == "{}" { // empty struct
			dst.ArrayOfString = nil
		} else {
			if err = validator.Validate(dst.ArrayOfString); err != nil {
				dst.ArrayOfString = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfString = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			if err = validator.Validate(dst.String); err != nil {
				dst.String = nil
			} else {
				match++
			}
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfString = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(OneOfWithComplexType)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(OneOfWithComplexType)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src OneOfWithComplexType) MarshalJSON() ([]byte, error) {
	if src.ArrayOfString != nil {
		return json.Marshal(&src.ArrayOfString)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *OneOfWithComplexType) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfString != nil {
		return obj.ArrayOfString
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableOneOfWithComplexType struct {
	value *OneOfWithComplexType
	isSet bool
}

func (v NullableOneOfWithComplexType) Get() *OneOfWithComplexType {
	return v.value
}

func (v *NullableOneOfWithComplexType) Set(val *OneOfWithComplexType) {
	v.value = val
	v.isSet = true
}

func (v NullableOneOfWithComplexType) IsSet() bool {
	return v.isSet
}

func (v *NullableOneOfWithComplexType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOneOfWithComplexType(val *OneOfWithComplexType) *NullableOneOfWithComplexType {
	return &NullableOneOfWithComplexType{value: val, isSet: true}
}

func (v NullableOneOfWithComplexType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOneOfWithComplexType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

