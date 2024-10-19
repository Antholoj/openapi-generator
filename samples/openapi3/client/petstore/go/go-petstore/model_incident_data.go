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

// IncidentData - struct for IncidentData
type IncidentData struct {
	ArrayOfMapmapOfStringAny *[]map[string]interface{}
	MapmapOfStringAny *map[string]interface{}
}

// []map[string]interface{}AsIncidentData is a convenience function that returns []map[string]interface{} wrapped in IncidentData
func ArrayOfMapmapOfStringAnyAsIncidentData(v *[]map[string]interface{}) IncidentData {
	return IncidentData{
		ArrayOfMapmapOfStringAny: v,
	}
}

// map[string]interface{}AsIncidentData is a convenience function that returns map[string]interface{} wrapped in IncidentData
func MapmapOfStringAnyAsIncidentData(v *map[string]interface{}) IncidentData {
	return IncidentData{
		MapmapOfStringAny: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *IncidentData) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into ArrayOfMapmapOfStringAny
	err = newStrictDecoder(data).Decode(&dst.ArrayOfMapmapOfStringAny)
	if err == nil {
		jsonArrayOfMapmapOfStringAny, _ := json.Marshal(dst.ArrayOfMapmapOfStringAny)
		if string(jsonArrayOfMapmapOfStringAny) == "{}" { // empty struct
			dst.ArrayOfMapmapOfStringAny = nil
		} else {
			if err = validator.Validate(dst.ArrayOfMapmapOfStringAny); err != nil {
				dst.ArrayOfMapmapOfStringAny = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfMapmapOfStringAny = nil
	}

	// try to unmarshal data into MapmapOfStringAny
	err = newStrictDecoder(data).Decode(&dst.MapmapOfStringAny)
	if err == nil {
		jsonMapmapOfStringAny, _ := json.Marshal(dst.MapmapOfStringAny)
		if string(jsonMapmapOfStringAny) == "{}" { // empty struct
			dst.MapmapOfStringAny = nil
		} else {
			if err = validator.Validate(dst.MapmapOfStringAny); err != nil {
				dst.MapmapOfStringAny = nil
			} else {
				match++
			}
		}
	} else {
		dst.MapmapOfStringAny = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfMapmapOfStringAny = nil
		dst.MapmapOfStringAny = nil

		return fmt.Errorf("data matches more than one schema in oneOf(IncidentData)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(IncidentData)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IncidentData) MarshalJSON() ([]byte, error) {
	if src.ArrayOfMapmapOfStringAny != nil {
		return json.Marshal(&src.ArrayOfMapmapOfStringAny)
	}

	if src.MapmapOfStringAny != nil {
		return json.Marshal(&src.MapmapOfStringAny)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IncidentData) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfMapmapOfStringAny != nil {
		return obj.ArrayOfMapmapOfStringAny
	}

	if obj.MapmapOfStringAny != nil {
		return obj.MapmapOfStringAny
	}

	// all schemas are nil
	return nil
}

type NullableIncidentData struct {
	value *IncidentData
	isSet bool
}

func (v NullableIncidentData) Get() *IncidentData {
	return v.value
}

func (v *NullableIncidentData) Set(val *IncidentData) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentData) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentData(val *IncidentData) *NullableIncidentData {
	return &NullableIncidentData{value: val, isSet: true}
}

func (v NullableIncidentData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

