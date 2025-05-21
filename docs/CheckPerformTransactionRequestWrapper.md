# CheckPerformTransactionRequestWrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jsonrpc** | **string** |  | 
**Id** | **int64** |  | 
**Method** | **string** |  | 
**Params** | [**CheckPerformTransactionRequest**](CheckPerformTransactionRequest.md) |  | 

## Methods

### NewCheckPerformTransactionRequestWrapper

`func NewCheckPerformTransactionRequestWrapper(jsonrpc string, id int64, method string, params CheckPerformTransactionRequest, ) *CheckPerformTransactionRequestWrapper`

NewCheckPerformTransactionRequestWrapper instantiates a new CheckPerformTransactionRequestWrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckPerformTransactionRequestWrapperWithDefaults

`func NewCheckPerformTransactionRequestWrapperWithDefaults() *CheckPerformTransactionRequestWrapper`

NewCheckPerformTransactionRequestWrapperWithDefaults instantiates a new CheckPerformTransactionRequestWrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonrpc

`func (o *CheckPerformTransactionRequestWrapper) GetJsonrpc() string`

GetJsonrpc returns the Jsonrpc field if non-nil, zero value otherwise.

### GetJsonrpcOk

`func (o *CheckPerformTransactionRequestWrapper) GetJsonrpcOk() (*string, bool)`

GetJsonrpcOk returns a tuple with the Jsonrpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonrpc

`func (o *CheckPerformTransactionRequestWrapper) SetJsonrpc(v string)`

SetJsonrpc sets Jsonrpc field to given value.


### GetId

`func (o *CheckPerformTransactionRequestWrapper) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckPerformTransactionRequestWrapper) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckPerformTransactionRequestWrapper) SetId(v int64)`

SetId sets Id field to given value.


### GetMethod

`func (o *CheckPerformTransactionRequestWrapper) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *CheckPerformTransactionRequestWrapper) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *CheckPerformTransactionRequestWrapper) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetParams

`func (o *CheckPerformTransactionRequestWrapper) GetParams() CheckPerformTransactionRequest`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CheckPerformTransactionRequestWrapper) GetParamsOk() (*CheckPerformTransactionRequest, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CheckPerformTransactionRequestWrapper) SetParams(v CheckPerformTransactionRequest)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


