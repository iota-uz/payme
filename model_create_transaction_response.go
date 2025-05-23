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

// checks if the CreateTransactionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTransactionResponse{}

// CreateTransactionResponse struct for CreateTransactionResponse
type CreateTransactionResponse struct {
	CreateTime  int64      `form:"create_time" json:"create_time"`
	Transaction string     `form:"transaction" json:"transaction"`
	State       int32      `form:"state" json:"state"`
	Receivers   []Receiver `form:"receivers" json:"receivers,omitempty"`
}

type _CreateTransactionResponse CreateTransactionResponse

// NewCreateTransactionResponse instantiates a new CreateTransactionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTransactionResponse(createTime int64, transaction string, state int32) *CreateTransactionResponse {
	this := CreateTransactionResponse{}
	this.CreateTime = createTime
	this.Transaction = transaction
	this.State = state
	return &this
}

// NewCreateTransactionResponseWithDefaults instantiates a new CreateTransactionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTransactionResponseWithDefaults() *CreateTransactionResponse {
	this := CreateTransactionResponse{}
	return &this
}

// GetCreateTime returns the CreateTime field value
func (o *CreateTransactionResponse) GetCreateTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value
// and a boolean to check if the value has been set.
func (o *CreateTransactionResponse) GetCreateTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreateTime, true
}

// SetCreateTime sets field value
func (o *CreateTransactionResponse) SetCreateTime(v int64) {
	o.CreateTime = v
}

// GetTransaction returns the Transaction field value
func (o *CreateTransactionResponse) GetTransaction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value
// and a boolean to check if the value has been set.
func (o *CreateTransactionResponse) GetTransactionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transaction, true
}

// SetTransaction sets field value
func (o *CreateTransactionResponse) SetTransaction(v string) {
	o.Transaction = v
}

// GetState returns the State field value
func (o *CreateTransactionResponse) GetState() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CreateTransactionResponse) GetStateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CreateTransactionResponse) SetState(v int32) {
	o.State = v
}

// GetReceivers returns the Receivers field value if set, zero value otherwise.
func (o *CreateTransactionResponse) GetReceivers() []Receiver {
	if o == nil || IsNil(o.Receivers) {
		var ret []Receiver
		return ret
	}
	return o.Receivers
}

// GetReceiversOk returns a tuple with the Receivers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTransactionResponse) GetReceiversOk() ([]Receiver, bool) {
	if o == nil || IsNil(o.Receivers) {
		return nil, false
	}
	return o.Receivers, true
}

// HasReceivers returns a boolean if a field has been set.
func (o *CreateTransactionResponse) HasReceivers() bool {
	if o != nil && !IsNil(o.Receivers) {
		return true
	}

	return false
}

// SetReceivers gets a reference to the given []Receiver and assigns it to the Receivers field.
func (o *CreateTransactionResponse) SetReceivers(v []Receiver) {
	o.Receivers = v
}

func (o CreateTransactionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTransactionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["create_time"] = o.CreateTime
	toSerialize["transaction"] = o.Transaction
	toSerialize["state"] = o.State
	if !IsNil(o.Receivers) {
		toSerialize["receivers"] = o.Receivers
	}
	return toSerialize, nil
}

func (o *CreateTransactionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"create_time",
		"transaction",
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

	varCreateTransactionResponse := _CreateTransactionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTransactionResponse)

	if err != nil {
		return err
	}

	*o = CreateTransactionResponse(varCreateTransactionResponse)

	return err
}

type NullableCreateTransactionResponse struct {
	value *CreateTransactionResponse
	isSet bool
}

func (v NullableCreateTransactionResponse) Get() *CreateTransactionResponse {
	return v.value
}

func (v *NullableCreateTransactionResponse) Set(val *CreateTransactionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTransactionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTransactionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTransactionResponse(val *CreateTransactionResponse) *NullableCreateTransactionResponse {
	return &NullableCreateTransactionResponse{value: val, isSet: true}
}

func (v NullableCreateTransactionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTransactionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
