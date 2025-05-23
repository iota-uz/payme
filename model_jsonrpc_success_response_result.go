/*
Payme Merchant API

API for managing Payme Merchant.

API version: 1.0.0
Contact: danil@iota.uz
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package paymeapi

import (
	"encoding/json"
	"fmt"
)

// JSONRPCSuccessResponseResult - struct for JSONRPCSuccessResponseResult
type JSONRPCSuccessResponseResult struct {
	CancelTransactionResponse       *CancelTransactionResponse
	CheckPerformTransactionResponse *CheckPerformTransactionResponse
	CheckTransactionResponse        *CheckTransactionResponse
	CreateTransactionResponse       *CreateTransactionResponse
	GetStatementResponse            *GetStatementResponse
	PerformTransactionResponse      *PerformTransactionResponse
	SetFiscalDataResponse           *SetFiscalDataResponse
}

// CancelTransactionResponseAsJSONRPCSuccessResponseResult is a convenience function that returns CancelTransactionResponse wrapped in JSONRPCSuccessResponseResult
func CancelTransactionResponseAsJSONRPCSuccessResponseResult(v *CancelTransactionResponse) JSONRPCSuccessResponseResult {
	return JSONRPCSuccessResponseResult{
		CancelTransactionResponse: v,
	}
}

// CheckPerformTransactionResponseAsJSONRPCSuccessResponseResult is a convenience function that returns CheckPerformTransactionResponse wrapped in JSONRPCSuccessResponseResult
func CheckPerformTransactionResponseAsJSONRPCSuccessResponseResult(v *CheckPerformTransactionResponse) JSONRPCSuccessResponseResult {
	return JSONRPCSuccessResponseResult{
		CheckPerformTransactionResponse: v,
	}
}

// CheckTransactionResponseAsJSONRPCSuccessResponseResult is a convenience function that returns CheckTransactionResponse wrapped in JSONRPCSuccessResponseResult
func CheckTransactionResponseAsJSONRPCSuccessResponseResult(v *CheckTransactionResponse) JSONRPCSuccessResponseResult {
	return JSONRPCSuccessResponseResult{
		CheckTransactionResponse: v,
	}
}

// CreateTransactionResponseAsJSONRPCSuccessResponseResult is a convenience function that returns CreateTransactionResponse wrapped in JSONRPCSuccessResponseResult
func CreateTransactionResponseAsJSONRPCSuccessResponseResult(v *CreateTransactionResponse) JSONRPCSuccessResponseResult {
	return JSONRPCSuccessResponseResult{
		CreateTransactionResponse: v,
	}
}

// GetStatementResponseAsJSONRPCSuccessResponseResult is a convenience function that returns GetStatementResponse wrapped in JSONRPCSuccessResponseResult
func GetStatementResponseAsJSONRPCSuccessResponseResult(v *GetStatementResponse) JSONRPCSuccessResponseResult {
	return JSONRPCSuccessResponseResult{
		GetStatementResponse: v,
	}
}

// PerformTransactionResponseAsJSONRPCSuccessResponseResult is a convenience function that returns PerformTransactionResponse wrapped in JSONRPCSuccessResponseResult
func PerformTransactionResponseAsJSONRPCSuccessResponseResult(v *PerformTransactionResponse) JSONRPCSuccessResponseResult {
	return JSONRPCSuccessResponseResult{
		PerformTransactionResponse: v,
	}
}

// SetFiscalDataResponseAsJSONRPCSuccessResponseResult is a convenience function that returns SetFiscalDataResponse wrapped in JSONRPCSuccessResponseResult
func SetFiscalDataResponseAsJSONRPCSuccessResponseResult(v *SetFiscalDataResponse) JSONRPCSuccessResponseResult {
	return JSONRPCSuccessResponseResult{
		SetFiscalDataResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *JSONRPCSuccessResponseResult) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CancelTransactionResponse
	err = json.Unmarshal(data, &dst.CancelTransactionResponse)
	if err == nil {
		jsonCancelTransactionResponse, _ := json.Marshal(dst.CancelTransactionResponse)
		if string(jsonCancelTransactionResponse) == "{}" { // empty struct
			dst.CancelTransactionResponse = nil
		} else {
			match++
		}
	} else {
		dst.CancelTransactionResponse = nil
	}

	// try to unmarshal data into CheckPerformTransactionResponse
	err = json.Unmarshal(data, &dst.CheckPerformTransactionResponse)
	if err == nil {
		jsonCheckPerformTransactionResponse, _ := json.Marshal(dst.CheckPerformTransactionResponse)
		if string(jsonCheckPerformTransactionResponse) == "{}" { // empty struct
			dst.CheckPerformTransactionResponse = nil
		} else {
			match++
		}
	} else {
		dst.CheckPerformTransactionResponse = nil
	}

	// try to unmarshal data into CheckTransactionResponse
	err = json.Unmarshal(data, &dst.CheckTransactionResponse)
	if err == nil {
		jsonCheckTransactionResponse, _ := json.Marshal(dst.CheckTransactionResponse)
		if string(jsonCheckTransactionResponse) == "{}" { // empty struct
			dst.CheckTransactionResponse = nil
		} else {
			match++
		}
	} else {
		dst.CheckTransactionResponse = nil
	}

	// try to unmarshal data into CreateTransactionResponse
	err = json.Unmarshal(data, &dst.CreateTransactionResponse)
	if err == nil {
		jsonCreateTransactionResponse, _ := json.Marshal(dst.CreateTransactionResponse)
		if string(jsonCreateTransactionResponse) == "{}" { // empty struct
			dst.CreateTransactionResponse = nil
		} else {
			match++
		}
	} else {
		dst.CreateTransactionResponse = nil
	}

	// try to unmarshal data into GetStatementResponse
	err = json.Unmarshal(data, &dst.GetStatementResponse)
	if err == nil {
		jsonGetStatementResponse, _ := json.Marshal(dst.GetStatementResponse)
		if string(jsonGetStatementResponse) == "{}" { // empty struct
			dst.GetStatementResponse = nil
		} else {
			match++
		}
	} else {
		dst.GetStatementResponse = nil
	}

	// try to unmarshal data into PerformTransactionResponse
	err = json.Unmarshal(data, &dst.PerformTransactionResponse)
	if err == nil {
		jsonPerformTransactionResponse, _ := json.Marshal(dst.PerformTransactionResponse)
		if string(jsonPerformTransactionResponse) == "{}" { // empty struct
			dst.PerformTransactionResponse = nil
		} else {
			match++
		}
	} else {
		dst.PerformTransactionResponse = nil
	}

	// try to unmarshal data into SetFiscalDataResponse
	err = json.Unmarshal(data, &dst.SetFiscalDataResponse)
	if err == nil {
		jsonSetFiscalDataResponse, _ := json.Marshal(dst.SetFiscalDataResponse)
		if string(jsonSetFiscalDataResponse) == "{}" { // empty struct
			dst.SetFiscalDataResponse = nil
		} else {
			match++
		}
	} else {
		dst.SetFiscalDataResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CancelTransactionResponse = nil
		dst.CheckPerformTransactionResponse = nil
		dst.CheckTransactionResponse = nil
		dst.CreateTransactionResponse = nil
		dst.GetStatementResponse = nil
		dst.PerformTransactionResponse = nil
		dst.SetFiscalDataResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(JSONRPCSuccessResponseResult)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(JSONRPCSuccessResponseResult)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src JSONRPCSuccessResponseResult) MarshalJSON() ([]byte, error) {
	if src.CancelTransactionResponse != nil {
		return json.Marshal(&src.CancelTransactionResponse)
	}

	if src.CheckPerformTransactionResponse != nil {
		return json.Marshal(&src.CheckPerformTransactionResponse)
	}

	if src.CheckTransactionResponse != nil {
		return json.Marshal(&src.CheckTransactionResponse)
	}

	if src.CreateTransactionResponse != nil {
		return json.Marshal(&src.CreateTransactionResponse)
	}

	if src.GetStatementResponse != nil {
		return json.Marshal(&src.GetStatementResponse)
	}

	if src.PerformTransactionResponse != nil {
		return json.Marshal(&src.PerformTransactionResponse)
	}

	if src.SetFiscalDataResponse != nil {
		return json.Marshal(&src.SetFiscalDataResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *JSONRPCSuccessResponseResult) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CancelTransactionResponse != nil {
		return obj.CancelTransactionResponse
	}

	if obj.CheckPerformTransactionResponse != nil {
		return obj.CheckPerformTransactionResponse
	}

	if obj.CheckTransactionResponse != nil {
		return obj.CheckTransactionResponse
	}

	if obj.CreateTransactionResponse != nil {
		return obj.CreateTransactionResponse
	}

	if obj.GetStatementResponse != nil {
		return obj.GetStatementResponse
	}

	if obj.PerformTransactionResponse != nil {
		return obj.PerformTransactionResponse
	}

	if obj.SetFiscalDataResponse != nil {
		return obj.SetFiscalDataResponse
	}

	// all schemas are nil
	return nil
}

type NullableJSONRPCSuccessResponseResult struct {
	value *JSONRPCSuccessResponseResult
	isSet bool
}

func (v NullableJSONRPCSuccessResponseResult) Get() *JSONRPCSuccessResponseResult {
	return v.value
}

func (v *NullableJSONRPCSuccessResponseResult) Set(val *JSONRPCSuccessResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableJSONRPCSuccessResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableJSONRPCSuccessResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJSONRPCSuccessResponseResult(val *JSONRPCSuccessResponseResult) *NullableJSONRPCSuccessResponseResult {
	return &NullableJSONRPCSuccessResponseResult{value: val, isSet: true}
}

func (v NullableJSONRPCSuccessResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJSONRPCSuccessResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
