/*
OpenAPI Petstore

This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore

import (
	"encoding/json"
	"os"
	"time"
	"bytes"
	"fmt"
)

// checks if the FormatTest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormatTest{}

// FormatTest struct for FormatTest
type FormatTest struct {
	Integer *int32 `json:"integer,omitempty"`
	Int32 *int32 `json:"int32,omitempty"`
	Int64 *int64 `json:"int64,omitempty"`
	Number float32 `json:"number"`
	Float *float32 `json:"float,omitempty"`
	Double *float64 `json:"double,omitempty"`
	String *string `json:"string,omitempty" validate:"regexp=[a-z]/i"`
	Byte string `json:"byte" validate:"regexp=^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=)?$"`
	Binary **os.File `json:"binary,omitempty"`
	Date string `json:"date"`
	DateTime *time.Time `json:"dateTime,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	Password string `json:"password"`
	BigDecimal *float64 `json:"BigDecimal,omitempty"`
}

type _FormatTest FormatTest

// NewFormatTest instantiates a new FormatTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormatTest(number float32, byte_ string, date string, password string) *FormatTest {
	this := FormatTest{}
	this.Number = number
	this.Byte = byte_
	this.Date = date
	this.Password = password
	return &this
}

// NewFormatTestWithDefaults instantiates a new FormatTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormatTestWithDefaults() *FormatTest {
	this := FormatTest{}
	return &this
}

// GetInteger returns the Integer field value if set, zero value otherwise.
func (o *FormatTest) GetInteger() int32 {
	if o == nil || IsNil(o.Integer) {
		var ret int32
		return ret
	}
	return *o.Integer
}

// GetIntegerOk returns a tuple with the Integer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetIntegerOk() (*int32, bool) {
	if o == nil || IsNil(o.Integer) {
		return nil, false
	}
	return o.Integer, true
}

// HasInteger returns a boolean if a field has been set.
func (o *FormatTest) HasInteger() bool {
	if o != nil && !IsNil(o.Integer) {
		return true
	}

	return false
}

// SetInteger gets a reference to the given int32 and assigns it to the Integer field.
func (o *FormatTest) SetInteger(v int32) {
	o.Integer = &v
}

// GetInt32 returns the Int32 field value if set, zero value otherwise.
func (o *FormatTest) GetInt32() int32 {
	if o == nil || IsNil(o.Int32) {
		var ret int32
		return ret
	}
	return *o.Int32
}

// GetInt32Ok returns a tuple with the Int32 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetInt32Ok() (*int32, bool) {
	if o == nil || IsNil(o.Int32) {
		return nil, false
	}
	return o.Int32, true
}

// HasInt32 returns a boolean if a field has been set.
func (o *FormatTest) HasInt32() bool {
	if o != nil && !IsNil(o.Int32) {
		return true
	}

	return false
}

// SetInt32 gets a reference to the given int32 and assigns it to the Int32 field.
func (o *FormatTest) SetInt32(v int32) {
	o.Int32 = &v
}

// GetInt64 returns the Int64 field value if set, zero value otherwise.
func (o *FormatTest) GetInt64() int64 {
	if o == nil || IsNil(o.Int64) {
		var ret int64
		return ret
	}
	return *o.Int64
}

// GetInt64Ok returns a tuple with the Int64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetInt64Ok() (*int64, bool) {
	if o == nil || IsNil(o.Int64) {
		return nil, false
	}
	return o.Int64, true
}

// HasInt64 returns a boolean if a field has been set.
func (o *FormatTest) HasInt64() bool {
	if o != nil && !IsNil(o.Int64) {
		return true
	}

	return false
}

// SetInt64 gets a reference to the given int64 and assigns it to the Int64 field.
func (o *FormatTest) SetInt64(v int64) {
	o.Int64 = &v
}

// GetNumber returns the Number field value
func (o *FormatTest) GetNumber() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *FormatTest) GetNumberOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *FormatTest) SetNumber(v float32) {
	o.Number = v
}

// GetFloat returns the Float field value if set, zero value otherwise.
func (o *FormatTest) GetFloat() float32 {
	if o == nil || IsNil(o.Float) {
		var ret float32
		return ret
	}
	return *o.Float
}

// GetFloatOk returns a tuple with the Float field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetFloatOk() (*float32, bool) {
	if o == nil || IsNil(o.Float) {
		return nil, false
	}
	return o.Float, true
}

// HasFloat returns a boolean if a field has been set.
func (o *FormatTest) HasFloat() bool {
	if o != nil && !IsNil(o.Float) {
		return true
	}

	return false
}

// SetFloat gets a reference to the given float32 and assigns it to the Float field.
func (o *FormatTest) SetFloat(v float32) {
	o.Float = &v
}

// GetDouble returns the Double field value if set, zero value otherwise.
func (o *FormatTest) GetDouble() float64 {
	if o == nil || IsNil(o.Double) {
		var ret float64
		return ret
	}
	return *o.Double
}

// GetDoubleOk returns a tuple with the Double field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetDoubleOk() (*float64, bool) {
	if o == nil || IsNil(o.Double) {
		return nil, false
	}
	return o.Double, true
}

// HasDouble returns a boolean if a field has been set.
func (o *FormatTest) HasDouble() bool {
	if o != nil && !IsNil(o.Double) {
		return true
	}

	return false
}

// SetDouble gets a reference to the given float64 and assigns it to the Double field.
func (o *FormatTest) SetDouble(v float64) {
	o.Double = &v
}

// GetString returns the String field value if set, zero value otherwise.
func (o *FormatTest) GetString() string {
	if o == nil || IsNil(o.String) {
		var ret string
		return ret
	}
	return *o.String
}

// GetStringOk returns a tuple with the String field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetStringOk() (*string, bool) {
	if o == nil || IsNil(o.String) {
		return nil, false
	}
	return o.String, true
}

// HasString returns a boolean if a field has been set.
func (o *FormatTest) HasString() bool {
	if o != nil && !IsNil(o.String) {
		return true
	}

	return false
}

// SetString gets a reference to the given string and assigns it to the String field.
func (o *FormatTest) SetString(v string) {
	o.String = &v
}

// GetByte returns the Byte field value
func (o *FormatTest) GetByte() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Byte
}

// GetByteOk returns a tuple with the Byte field value
// and a boolean to check if the value has been set.
func (o *FormatTest) GetByteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Byte, true
}

// SetByte sets field value
func (o *FormatTest) SetByte(v string) {
	o.Byte = v
}

// GetBinary returns the Binary field value if set, zero value otherwise.
func (o *FormatTest) GetBinary() *os.File {
	if o == nil || IsNil(o.Binary) {
		var ret *os.File
		return ret
	}
	return *o.Binary
}

// GetBinaryOk returns a tuple with the Binary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetBinaryOk() (**os.File, bool) {
	if o == nil || IsNil(o.Binary) {
		return nil, false
	}
	return o.Binary, true
}

// HasBinary returns a boolean if a field has been set.
func (o *FormatTest) HasBinary() bool {
	if o != nil && !IsNil(o.Binary) {
		return true
	}

	return false
}

// SetBinary gets a reference to the given *os.File and assigns it to the Binary field.
func (o *FormatTest) SetBinary(v *os.File) {
	o.Binary = &v
}

// GetDate returns the Date field value
func (o *FormatTest) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *FormatTest) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *FormatTest) SetDate(v string) {
	o.Date = v
}

// GetDateTime returns the DateTime field value if set, zero value otherwise.
func (o *FormatTest) GetDateTime() time.Time {
	if o == nil || IsNil(o.DateTime) {
		var ret time.Time
		return ret
	}
	return *o.DateTime
}

// GetDateTimeOk returns a tuple with the DateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateTime) {
		return nil, false
	}
	return o.DateTime, true
}

// HasDateTime returns a boolean if a field has been set.
func (o *FormatTest) HasDateTime() bool {
	if o != nil && !IsNil(o.DateTime) {
		return true
	}

	return false
}

// SetDateTime gets a reference to the given time.Time and assigns it to the DateTime field.
func (o *FormatTest) SetDateTime(v time.Time) {
	o.DateTime = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *FormatTest) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *FormatTest) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *FormatTest) SetUuid(v string) {
	o.Uuid = &v
}

// GetPassword returns the Password field value
func (o *FormatTest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *FormatTest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *FormatTest) SetPassword(v string) {
	o.Password = v
}

// GetBigDecimal returns the BigDecimal field value if set, zero value otherwise.
func (o *FormatTest) GetBigDecimal() float64 {
	if o == nil || IsNil(o.BigDecimal) {
		var ret float64
		return ret
	}
	return *o.BigDecimal
}

// GetBigDecimalOk returns a tuple with the BigDecimal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetBigDecimalOk() (*float64, bool) {
	if o == nil || IsNil(o.BigDecimal) {
		return nil, false
	}
	return o.BigDecimal, true
}

// HasBigDecimal returns a boolean if a field has been set.
func (o *FormatTest) HasBigDecimal() bool {
	if o != nil && !IsNil(o.BigDecimal) {
		return true
	}

	return false
}

// SetBigDecimal gets a reference to the given float64 and assigns it to the BigDecimal field.
func (o *FormatTest) SetBigDecimal(v float64) {
	o.BigDecimal = &v
}

func (o FormatTest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormatTest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Integer) {
		toSerialize["integer"] = o.Integer
	}
	if !IsNil(o.Int32) {
		toSerialize["int32"] = o.Int32
	}
	if !IsNil(o.Int64) {
		toSerialize["int64"] = o.Int64
	}
	toSerialize["number"] = o.Number
	if !IsNil(o.Float) {
		toSerialize["float"] = o.Float
	}
	if !IsNil(o.Double) {
		toSerialize["double"] = o.Double
	}
	if !IsNil(o.String) {
		toSerialize["string"] = o.String
	}
	toSerialize["byte"] = o.Byte
	if !IsNil(o.Binary) {
		toSerialize["binary"] = o.Binary
	}
	toSerialize["date"] = o.Date
	if !IsNil(o.DateTime) {
		toSerialize["dateTime"] = o.DateTime
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	toSerialize["password"] = o.Password
	if !IsNil(o.BigDecimal) {
		toSerialize["BigDecimal"] = o.BigDecimal
	}
	return toSerialize, nil
}

func (o *FormatTest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"number",
		"byte",
		"date",
		"password",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFormatTest := _FormatTest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFormatTest)

	if err != nil {
		return err
	}

	*o = FormatTest(varFormatTest)

	return err
}

type NullableFormatTest struct {
	value *FormatTest
	isSet bool
}

func (v NullableFormatTest) Get() *FormatTest {
	return v.value
}

func (v *NullableFormatTest) Set(val *FormatTest) {
	v.value = val
	v.isSet = true
}

func (v NullableFormatTest) IsSet() bool {
	return v.isSet
}

func (v *NullableFormatTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormatTest(val *FormatTest) *NullableFormatTest {
	return &NullableFormatTest{value: val, isSet: true}
}

func (v NullableFormatTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormatTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

