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


// GmFruit struct for GmFruit
type GmFruit struct {
	Apple *Apple
	Banana *Banana
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *GmFruit) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into Apple
	err = json.Unmarshal(data, &dst.Apple);
	if err == nil {
		jsonApple, _ := json.Marshal(dst.Apple)
		if string(jsonApple) == "{}" { // empty struct
			dst.Apple = nil
		} else {
			return nil // data stored in dst.Apple, return on the first match
		}
	} else {
		dst.Apple = nil
	}

	// try to unmarshal JSON data into Banana
	err = json.Unmarshal(data, &dst.Banana);
	if err == nil {
		jsonBanana, _ := json.Marshal(dst.Banana)
		if string(jsonBanana) == "{}" { // empty struct
			dst.Banana = nil
		} else {
			return nil // data stored in dst.Banana, return on the first match
		}
	} else {
		dst.Banana = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(GmFruit)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *GmFruit) MarshalJSON() ([]byte, error) {
	if src.Apple != nil {
		return json.Marshal(&src.Apple)
	}

	if src.Banana != nil {
		return json.Marshal(&src.Banana)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableGmFruit struct {
	value *GmFruit
	isSet bool
}

func (v NullableGmFruit) Get() *GmFruit {
	return v.value
}

func (v *NullableGmFruit) Set(val *GmFruit) {
	v.value = val
	v.isSet = true
}

func (v NullableGmFruit) IsSet() bool {
	return v.isSet
}

func (v *NullableGmFruit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGmFruit(val *GmFruit) *NullableGmFruit {
	return &NullableGmFruit{value: val, isSet: true}
}

func (v NullableGmFruit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGmFruit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

