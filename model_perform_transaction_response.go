/*
Payme Merchant API

API for managing Payme Merchant.

API version: 1.0.0
Contact: danil@iota.uz
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package paymeapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PerformTransactionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PerformTransactionResponse{}

// PerformTransactionResponse struct for PerformTransactionResponse
type PerformTransactionResponse struct {
	Transaction string `form:"transaction" json:"transaction"`
	PerformTime int64  `form:"perform_time" json:"perform_time"`
	State       int32  `form:"state" json:"state"`
}

type _PerformTransactionResponse PerformTransactionResponse

// NewPerformTransactionResponse instantiates a new PerformTransactionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerformTransactionResponse(transaction string, performTime int64, state int32) *PerformTransactionResponse {
	this := PerformTransactionResponse{}
	this.Transaction = transaction
	this.PerformTime = performTime
	this.State = state
	return &this
}

// NewPerformTransactionResponseWithDefaults instantiates a new PerformTransactionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerformTransactionResponseWithDefaults() *PerformTransactionResponse {
	this := PerformTransactionResponse{}
	return &this
}

// GetTransaction returns the Transaction field value
func (o *PerformTransactionResponse) GetTransaction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value
// and a boolean to check if the value has been set.
func (o *PerformTransactionResponse) GetTransactionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transaction, true
}

// SetTransaction sets field value
func (o *PerformTransactionResponse) SetTransaction(v string) {
	o.Transaction = v
}

// GetPerformTime returns the PerformTime field value
func (o *PerformTransactionResponse) GetPerformTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PerformTime
}

// GetPerformTimeOk returns a tuple with the PerformTime field value
// and a boolean to check if the value has been set.
func (o *PerformTransactionResponse) GetPerformTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PerformTime, true
}

// SetPerformTime sets field value
func (o *PerformTransactionResponse) SetPerformTime(v int64) {
	o.PerformTime = v
}

// GetState returns the State field value
func (o *PerformTransactionResponse) GetState() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *PerformTransactionResponse) GetStateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *PerformTransactionResponse) SetState(v int32) {
	o.State = v
}

func (o PerformTransactionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PerformTransactionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transaction"] = o.Transaction
	toSerialize["perform_time"] = o.PerformTime
	toSerialize["state"] = o.State
	return toSerialize, nil
}

func (o *PerformTransactionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transaction",
		"perform_time",
		"state",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPerformTransactionResponse := _PerformTransactionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPerformTransactionResponse)

	if err != nil {
		return err
	}

	*o = PerformTransactionResponse(varPerformTransactionResponse)

	return err
}

type NullablePerformTransactionResponse struct {
	value *PerformTransactionResponse
	isSet bool
}

func (v NullablePerformTransactionResponse) Get() *PerformTransactionResponse {
	return v.value
}

func (v *NullablePerformTransactionResponse) Set(val *PerformTransactionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePerformTransactionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePerformTransactionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerformTransactionResponse(val *PerformTransactionResponse) *NullablePerformTransactionResponse {
	return &NullablePerformTransactionResponse{value: val, isSet: true}
}

func (v NullablePerformTransactionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerformTransactionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
