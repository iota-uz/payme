# CancelTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transaction** | **string** |  | 
**CancelTime** | **int64** |  | 
**State** | **int32** |  | 

## Methods

### NewCancelTransactionResponse

`func NewCancelTransactionResponse(transaction string, cancelTime int64, state int32, ) *CancelTransactionResponse`

NewCancelTransactionResponse instantiates a new CancelTransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelTransactionResponseWithDefaults

`func NewCancelTransactionResponseWithDefaults() *CancelTransactionResponse`

NewCancelTransactionResponseWithDefaults instantiates a new CancelTransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransaction

`func (o *CancelTransactionResponse) GetTransaction() string`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *CancelTransactionResponse) GetTransactionOk() (*string, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *CancelTransactionResponse) SetTransaction(v string)`

SetTransaction sets Transaction field to given value.


### GetCancelTime

`func (o *CancelTransactionResponse) GetCancelTime() int64`

GetCancelTime returns the CancelTime field if non-nil, zero value otherwise.

### GetCancelTimeOk

`func (o *CancelTransactionResponse) GetCancelTimeOk() (*int64, bool)`

GetCancelTimeOk returns a tuple with the CancelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelTime

`func (o *CancelTransactionResponse) SetCancelTime(v int64)`

SetCancelTime sets CancelTime field to given value.


### GetState

`func (o *CancelTransactionResponse) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CancelTransactionResponse) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CancelTransactionResponse) SetState(v int32)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


