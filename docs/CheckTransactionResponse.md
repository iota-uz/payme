# CheckTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | **int64** |  | 
**PerformTime** | **int64** |  | 
**CancelTime** | **int64** |  | 
**Transaction** | **string** |  | 
**State** | **int32** |  | 
**Reason** | **NullableInt32** |  | 

## Methods

### NewCheckTransactionResponse

`func NewCheckTransactionResponse(createTime int64, performTime int64, cancelTime int64, transaction string, state int32, reason NullableInt32, ) *CheckTransactionResponse`

NewCheckTransactionResponse instantiates a new CheckTransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckTransactionResponseWithDefaults

`func NewCheckTransactionResponseWithDefaults() *CheckTransactionResponse`

NewCheckTransactionResponseWithDefaults instantiates a new CheckTransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *CheckTransactionResponse) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *CheckTransactionResponse) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *CheckTransactionResponse) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.


### GetPerformTime

`func (o *CheckTransactionResponse) GetPerformTime() int64`

GetPerformTime returns the PerformTime field if non-nil, zero value otherwise.

### GetPerformTimeOk

`func (o *CheckTransactionResponse) GetPerformTimeOk() (*int64, bool)`

GetPerformTimeOk returns a tuple with the PerformTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformTime

`func (o *CheckTransactionResponse) SetPerformTime(v int64)`

SetPerformTime sets PerformTime field to given value.


### GetCancelTime

`func (o *CheckTransactionResponse) GetCancelTime() int64`

GetCancelTime returns the CancelTime field if non-nil, zero value otherwise.

### GetCancelTimeOk

`func (o *CheckTransactionResponse) GetCancelTimeOk() (*int64, bool)`

GetCancelTimeOk returns a tuple with the CancelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelTime

`func (o *CheckTransactionResponse) SetCancelTime(v int64)`

SetCancelTime sets CancelTime field to given value.


### GetTransaction

`func (o *CheckTransactionResponse) GetTransaction() string`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *CheckTransactionResponse) GetTransactionOk() (*string, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *CheckTransactionResponse) SetTransaction(v string)`

SetTransaction sets Transaction field to given value.


### GetState

`func (o *CheckTransactionResponse) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CheckTransactionResponse) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CheckTransactionResponse) SetState(v int32)`

SetState sets State field to given value.


### GetReason

`func (o *CheckTransactionResponse) GetReason() int32`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CheckTransactionResponse) GetReasonOk() (*int32, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CheckTransactionResponse) SetReason(v int32)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *CheckTransactionResponse) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *CheckTransactionResponse) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


