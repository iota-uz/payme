# JSONRPCSuccessResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allow** | **bool** |  | 
**Additional** | Pointer to **map[string]string** |  | [optional] 
**Detail** | Pointer to [**TransactionDetail**](TransactionDetail.md) |  | [optional] 
**CreateTime** | **int64** |  | 
**Transaction** | **string** |  | 
**State** | **int32** |  | 
**Receivers** | Pointer to [**[]Receiver**](Receiver.md) |  | [optional] 
**PerformTime** | **int64** |  | 
**CancelTime** | **int64** |  | 
**Reason** | **int32** |  | 
**Transactions** | [**[]StatementTransaction**](StatementTransaction.md) |  | 
**Success** | **bool** |  | 

## Methods

### NewJSONRPCSuccessResponseResult

`func NewJSONRPCSuccessResponseResult(allow bool, createTime int64, transaction string, state int32, performTime int64, cancelTime int64, reason int32, transactions []StatementTransaction, success bool, ) *JSONRPCSuccessResponseResult`

NewJSONRPCSuccessResponseResult instantiates a new JSONRPCSuccessResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSONRPCSuccessResponseResultWithDefaults

`func NewJSONRPCSuccessResponseResultWithDefaults() *JSONRPCSuccessResponseResult`

NewJSONRPCSuccessResponseResultWithDefaults instantiates a new JSONRPCSuccessResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllow

`func (o *JSONRPCSuccessResponseResult) GetAllow() bool`

GetAllow returns the Allow field if non-nil, zero value otherwise.

### GetAllowOk

`func (o *JSONRPCSuccessResponseResult) GetAllowOk() (*bool, bool)`

GetAllowOk returns a tuple with the Allow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllow

`func (o *JSONRPCSuccessResponseResult) SetAllow(v bool)`

SetAllow sets Allow field to given value.


### GetAdditional

`func (o *JSONRPCSuccessResponseResult) GetAdditional() map[string]string`

GetAdditional returns the Additional field if non-nil, zero value otherwise.

### GetAdditionalOk

`func (o *JSONRPCSuccessResponseResult) GetAdditionalOk() (*map[string]string, bool)`

GetAdditionalOk returns a tuple with the Additional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditional

`func (o *JSONRPCSuccessResponseResult) SetAdditional(v map[string]string)`

SetAdditional sets Additional field to given value.

### HasAdditional

`func (o *JSONRPCSuccessResponseResult) HasAdditional() bool`

HasAdditional returns a boolean if a field has been set.

### GetDetail

`func (o *JSONRPCSuccessResponseResult) GetDetail() TransactionDetail`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *JSONRPCSuccessResponseResult) GetDetailOk() (*TransactionDetail, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *JSONRPCSuccessResponseResult) SetDetail(v TransactionDetail)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *JSONRPCSuccessResponseResult) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetCreateTime

`func (o *JSONRPCSuccessResponseResult) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *JSONRPCSuccessResponseResult) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *JSONRPCSuccessResponseResult) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.


### GetTransaction

`func (o *JSONRPCSuccessResponseResult) GetTransaction() string`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *JSONRPCSuccessResponseResult) GetTransactionOk() (*string, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *JSONRPCSuccessResponseResult) SetTransaction(v string)`

SetTransaction sets Transaction field to given value.


### GetState

`func (o *JSONRPCSuccessResponseResult) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *JSONRPCSuccessResponseResult) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *JSONRPCSuccessResponseResult) SetState(v int32)`

SetState sets State field to given value.


### GetReceivers

`func (o *JSONRPCSuccessResponseResult) GetReceivers() []Receiver`

GetReceivers returns the Receivers field if non-nil, zero value otherwise.

### GetReceiversOk

`func (o *JSONRPCSuccessResponseResult) GetReceiversOk() (*[]Receiver, bool)`

GetReceiversOk returns a tuple with the Receivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivers

`func (o *JSONRPCSuccessResponseResult) SetReceivers(v []Receiver)`

SetReceivers sets Receivers field to given value.

### HasReceivers

`func (o *JSONRPCSuccessResponseResult) HasReceivers() bool`

HasReceivers returns a boolean if a field has been set.

### GetPerformTime

`func (o *JSONRPCSuccessResponseResult) GetPerformTime() int64`

GetPerformTime returns the PerformTime field if non-nil, zero value otherwise.

### GetPerformTimeOk

`func (o *JSONRPCSuccessResponseResult) GetPerformTimeOk() (*int64, bool)`

GetPerformTimeOk returns a tuple with the PerformTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformTime

`func (o *JSONRPCSuccessResponseResult) SetPerformTime(v int64)`

SetPerformTime sets PerformTime field to given value.


### GetCancelTime

`func (o *JSONRPCSuccessResponseResult) GetCancelTime() int64`

GetCancelTime returns the CancelTime field if non-nil, zero value otherwise.

### GetCancelTimeOk

`func (o *JSONRPCSuccessResponseResult) GetCancelTimeOk() (*int64, bool)`

GetCancelTimeOk returns a tuple with the CancelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelTime

`func (o *JSONRPCSuccessResponseResult) SetCancelTime(v int64)`

SetCancelTime sets CancelTime field to given value.


### GetReason

`func (o *JSONRPCSuccessResponseResult) GetReason() int32`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *JSONRPCSuccessResponseResult) GetReasonOk() (*int32, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *JSONRPCSuccessResponseResult) SetReason(v int32)`

SetReason sets Reason field to given value.


### GetTransactions

`func (o *JSONRPCSuccessResponseResult) GetTransactions() []StatementTransaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *JSONRPCSuccessResponseResult) GetTransactionsOk() (*[]StatementTransaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *JSONRPCSuccessResponseResult) SetTransactions(v []StatementTransaction)`

SetTransactions sets Transactions field to given value.


### GetSuccess

`func (o *JSONRPCSuccessResponseResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *JSONRPCSuccessResponseResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *JSONRPCSuccessResponseResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


