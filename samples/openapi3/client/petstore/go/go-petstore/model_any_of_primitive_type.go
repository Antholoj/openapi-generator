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


// AnyOfPrimitiveType struct for AnyOfPrimitiveType
type AnyOfPrimitiveType struct {
	OneOfPrimitiveTypeChild *OneOfPrimitiveTypeChild
	ArrayOfString *[]string
	Int32 *int32
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AnyOfPrimitiveType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into OneOfPrimitiveTypeChild
	err = json.Unmarshal(data, &dst.OneOfPrimitiveTypeChild);
	if err == nil {
		jsonOneOfPrimitiveTypeChild, _ := json.Marshal(dst.OneOfPrimitiveTypeChild)
		if string(jsonOneOfPrimitiveTypeChild) == "{}" { // empty struct
			dst.OneOfPrimitiveTypeChild = nil
		} else {
			return nil // data stored in dst.OneOfPrimitiveTypeChild, return on the first match
		}
	} else {
		dst.OneOfPrimitiveTypeChild = nil
	}

	// try to unmarshal JSON data into ArrayOfString
	err = json.Unmarshal(data, &dst.ArrayOfString);
	if err == nil {
		jsonArrayOfString, _ := json.Marshal(dst.ArrayOfString)
		if string(jsonArrayOfString) == "{}" { // empty struct
			dst.ArrayOfString = nil
		} else {
			return nil // data stored in dst.ArrayOfString, return on the first match
		}
	} else {
		dst.ArrayOfString = nil
	}

	// try to unmarshal JSON data into Int32
	err = json.Unmarshal(data, &dst.Int32);
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			return nil // data stored in dst.Int32, return on the first match
		}
	} else {
		dst.Int32 = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AnyOfPrimitiveType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AnyOfPrimitiveType) MarshalJSON() ([]byte, error) {
	if src.OneOfPrimitiveTypeChild != nil {
		return json.Marshal(&src.OneOfPrimitiveTypeChild)
	}

	if src.ArrayOfString != nil {
		return json.Marshal(&src.ArrayOfString)
	}

	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAnyOfPrimitiveType struct {
	value *AnyOfPrimitiveType
	isSet bool
}

func (v NullableAnyOfPrimitiveType) Get() *AnyOfPrimitiveType {
	return v.value
}

func (v *NullableAnyOfPrimitiveType) Set(val *AnyOfPrimitiveType) {
	v.value = val
	v.isSet = true
}

func (v NullableAnyOfPrimitiveType) IsSet() bool {
	return v.isSet
}

func (v *NullableAnyOfPrimitiveType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnyOfPrimitiveType(val *AnyOfPrimitiveType) *NullableAnyOfPrimitiveType {
	return &NullableAnyOfPrimitiveType{value: val, isSet: true}
}

func (v NullableAnyOfPrimitiveType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnyOfPrimitiveType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

