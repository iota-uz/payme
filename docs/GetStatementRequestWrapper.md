# GetStatementRequestWrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jsonrpc** | **string** |  | 
**Id** | **int64** |  | 
**Method** | **string** |  | 
**Params** | [**GetStatementRequest**](GetStatementRequest.md) |  | 

## Methods

### NewGetStatementRequestWrapper

`func NewGetStatementRequestWrapper(jsonrpc string, id int64, method string, params GetStatementRequest, ) *GetStatementRequestWrapper`

NewGetStatementRequestWrapper instantiates a new GetStatementRequestWrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStatementRequestWrapperWithDefaults

`func NewGetStatementRequestWrapperWithDefaults() *GetStatementRequestWrapper`

NewGetStatementRequestWrapperWithDefaults instantiates a new GetStatementRequestWrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonrpc

`func (o *GetStatementRequestWrapper) GetJsonrpc() string`

GetJsonrpc returns the Jsonrpc field if non-nil, zero value otherwise.

### GetJsonrpcOk

`func (o *GetStatementRequestWrapper) GetJsonrpcOk() (*string, bool)`

GetJsonrpcOk returns a tuple with the Jsonrpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonrpc

`func (o *GetStatementRequestWrapper) SetJsonrpc(v string)`

SetJsonrpc sets Jsonrpc field to given value.


### GetId

`func (o *GetStatementRequestWrapper) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetStatementRequestWrapper) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetStatementRequestWrapper) SetId(v int64)`

SetId sets Id field to given value.


### GetMethod

`func (o *GetStatementRequestWrapper) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *GetStatementRequestWrapper) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *GetStatementRequestWrapper) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetParams

`func (o *GetStatementRequestWrapper) GetParams() GetStatementRequest`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *GetStatementRequestWrapper) GetParamsOk() (*GetStatementRequest, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *GetStatementRequestWrapper) SetParams(v GetStatementRequest)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


