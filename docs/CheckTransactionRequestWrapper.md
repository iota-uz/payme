# CheckTransactionRequestWrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jsonrpc** | **string** |  | 
**Id** | **int64** |  | 
**Method** | **string** |  | 
**Params** | [**CheckTransactionRequest**](CheckTransactionRequest.md) |  | 

## Methods

### NewCheckTransactionRequestWrapper

`func NewCheckTransactionRequestWrapper(jsonrpc string, id int64, method string, params CheckTransactionRequest, ) *CheckTransactionRequestWrapper`

NewCheckTransactionRequestWrapper instantiates a new CheckTransactionRequestWrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckTransactionRequestWrapperWithDefaults

`func NewCheckTransactionRequestWrapperWithDefaults() *CheckTransactionRequestWrapper`

NewCheckTransactionRequestWrapperWithDefaults instantiates a new CheckTransactionRequestWrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonrpc

`func (o *CheckTransactionRequestWrapper) GetJsonrpc() string`

GetJsonrpc returns the Jsonrpc field if non-nil, zero value otherwise.

### GetJsonrpcOk

`func (o *CheckTransactionRequestWrapper) GetJsonrpcOk() (*string, bool)`

GetJsonrpcOk returns a tuple with the Jsonrpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonrpc

`func (o *CheckTransactionRequestWrapper) SetJsonrpc(v string)`

SetJsonrpc sets Jsonrpc field to given value.


### GetId

`func (o *CheckTransactionRequestWrapper) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckTransactionRequestWrapper) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckTransactionRequestWrapper) SetId(v int64)`

SetId sets Id field to given value.


### GetMethod

`func (o *CheckTransactionRequestWrapper) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *CheckTransactionRequestWrapper) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *CheckTransactionRequestWrapper) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetParams

`func (o *CheckTransactionRequestWrapper) GetParams() CheckTransactionRequest`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CheckTransactionRequestWrapper) GetParamsOk() (*CheckTransactionRequest, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CheckTransactionRequestWrapper) SetParams(v CheckTransactionRequest)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


