# JSONRPCErrorResponseError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** |  | 
**Message** | [**JSONRPCErrorResponseErrorMessage**](JSONRPCErrorResponseErrorMessage.md) |  | 
**Data** | Pointer to **string** |  | [optional] 

## Methods

### NewJSONRPCErrorResponseError

`func NewJSONRPCErrorResponseError(code int32, message JSONRPCErrorResponseErrorMessage, ) *JSONRPCErrorResponseError`

NewJSONRPCErrorResponseError instantiates a new JSONRPCErrorResponseError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSONRPCErrorResponseErrorWithDefaults

`func NewJSONRPCErrorResponseErrorWithDefaults() *JSONRPCErrorResponseError`

NewJSONRPCErrorResponseErrorWithDefaults instantiates a new JSONRPCErrorResponseError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *JSONRPCErrorResponseError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *JSONRPCErrorResponseError) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *JSONRPCErrorResponseError) SetCode(v int32)`

SetCode sets Code field to given value.


### GetMessage

`func (o *JSONRPCErrorResponseError) GetMessage() JSONRPCErrorResponseErrorMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *JSONRPCErrorResponseError) GetMessageOk() (*JSONRPCErrorResponseErrorMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *JSONRPCErrorResponseError) SetMessage(v JSONRPCErrorResponseErrorMessage)`

SetMessage sets Message field to given value.


### GetData

`func (o *JSONRPCErrorResponseError) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *JSONRPCErrorResponseError) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *JSONRPCErrorResponseError) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *JSONRPCErrorResponseError) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


