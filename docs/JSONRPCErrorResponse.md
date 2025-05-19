# JSONRPCErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | [**JSONRPCErrorResponseError**](JSONRPCErrorResponseError.md) |  | 
**Id** | **int64** |  | 

## Methods

### NewJSONRPCErrorResponse

`func NewJSONRPCErrorResponse(error_ JSONRPCErrorResponseError, id int64, ) *JSONRPCErrorResponse`

NewJSONRPCErrorResponse instantiates a new JSONRPCErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSONRPCErrorResponseWithDefaults

`func NewJSONRPCErrorResponseWithDefaults() *JSONRPCErrorResponse`

NewJSONRPCErrorResponseWithDefaults instantiates a new JSONRPCErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *JSONRPCErrorResponse) GetError() JSONRPCErrorResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *JSONRPCErrorResponse) GetErrorOk() (*JSONRPCErrorResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *JSONRPCErrorResponse) SetError(v JSONRPCErrorResponseError)`

SetError sets Error field to given value.


### GetId

`func (o *JSONRPCErrorResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JSONRPCErrorResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JSONRPCErrorResponse) SetId(v int64)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


