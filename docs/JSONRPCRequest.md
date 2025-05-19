# JSONRPCRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jsonrpc** | Pointer to **string** |  | [optional] 
**Method** | **string** |  | 
**Params** | [**JSONRPCRequestParams**](JSONRPCRequestParams.md) |  | 
**Id** | **int64** |  | 

## Methods

### NewJSONRPCRequest

`func NewJSONRPCRequest(method string, params JSONRPCRequestParams, id int64, ) *JSONRPCRequest`

NewJSONRPCRequest instantiates a new JSONRPCRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSONRPCRequestWithDefaults

`func NewJSONRPCRequestWithDefaults() *JSONRPCRequest`

NewJSONRPCRequestWithDefaults instantiates a new JSONRPCRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonrpc

`func (o *JSONRPCRequest) GetJsonrpc() string`

GetJsonrpc returns the Jsonrpc field if non-nil, zero value otherwise.

### GetJsonrpcOk

`func (o *JSONRPCRequest) GetJsonrpcOk() (*string, bool)`

GetJsonrpcOk returns a tuple with the Jsonrpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonrpc

`func (o *JSONRPCRequest) SetJsonrpc(v string)`

SetJsonrpc sets Jsonrpc field to given value.

### HasJsonrpc

`func (o *JSONRPCRequest) HasJsonrpc() bool`

HasJsonrpc returns a boolean if a field has been set.

### GetMethod

`func (o *JSONRPCRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *JSONRPCRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *JSONRPCRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetParams

`func (o *JSONRPCRequest) GetParams() JSONRPCRequestParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *JSONRPCRequest) GetParamsOk() (*JSONRPCRequestParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *JSONRPCRequest) SetParams(v JSONRPCRequestParams)`

SetParams sets Params field to given value.


### GetId

`func (o *JSONRPCRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JSONRPCRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JSONRPCRequest) SetId(v int64)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


