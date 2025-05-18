# CreateTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | **int64** |  | 
**Transaction** | **string** |  | 
**State** | **int32** |  | 
**Receivers** | Pointer to [**[]Receiver**](Receiver.md) |  | [optional] 

## Methods

### NewCreateTransactionResponse

`func NewCreateTransactionResponse(createTime int64, transaction string, state int32, ) *CreateTransactionResponse`

NewCreateTransactionResponse instantiates a new CreateTransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransactionResponseWithDefaults

`func NewCreateTransactionResponseWithDefaults() *CreateTransactionResponse`

NewCreateTransactionResponseWithDefaults instantiates a new CreateTransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *CreateTransactionResponse) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *CreateTransactionResponse) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *CreateTransactionResponse) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.


### GetTransaction

`func (o *CreateTransactionResponse) GetTransaction() string`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *CreateTransactionResponse) GetTransactionOk() (*string, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *CreateTransactionResponse) SetTransaction(v string)`

SetTransaction sets Transaction field to given value.


### GetState

`func (o *CreateTransactionResponse) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CreateTransactionResponse) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CreateTransactionResponse) SetState(v int32)`

SetState sets State field to given value.


### GetReceivers

`func (o *CreateTransactionResponse) GetReceivers() []Receiver`

GetReceivers returns the Receivers field if non-nil, zero value otherwise.

### GetReceiversOk

`func (o *CreateTransactionResponse) GetReceiversOk() (*[]Receiver, bool)`

GetReceiversOk returns a tuple with the Receivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivers

`func (o *CreateTransactionResponse) SetReceivers(v []Receiver)`

SetReceivers sets Receivers field to given value.

### HasReceivers

`func (o *CreateTransactionResponse) HasReceivers() bool`

HasReceivers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


