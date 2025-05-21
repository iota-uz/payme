# StatementTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Time** | **int64** |  | 
**Amount** | **float64** |  | 
**Account** | **map[string]interface{}** |  | 
**CreateTime** | **int64** |  | 
**PerformTime** | **int64** |  | 
**CancelTime** | **int64** |  | 
**Transaction** | **string** |  | 
**State** | **int32** |  | 
**Reason** | **int32** |  | 
**Receivers** | Pointer to [**[]Receiver**](Receiver.md) |  | [optional] 

## Methods

### NewStatementTransaction

`func NewStatementTransaction(id string, time int64, amount float64, account map[string]interface{}, createTime int64, performTime int64, cancelTime int64, transaction string, state int32, reason int32, ) *StatementTransaction`

NewStatementTransaction instantiates a new StatementTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementTransactionWithDefaults

`func NewStatementTransactionWithDefaults() *StatementTransaction`

NewStatementTransactionWithDefaults instantiates a new StatementTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StatementTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StatementTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StatementTransaction) SetId(v string)`

SetId sets Id field to given value.


### GetTime

`func (o *StatementTransaction) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *StatementTransaction) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *StatementTransaction) SetTime(v int64)`

SetTime sets Time field to given value.


### GetAmount

`func (o *StatementTransaction) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *StatementTransaction) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *StatementTransaction) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetAccount

`func (o *StatementTransaction) GetAccount() map[string]interface{}`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *StatementTransaction) GetAccountOk() (*map[string]interface{}, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *StatementTransaction) SetAccount(v map[string]interface{})`

SetAccount sets Account field to given value.


### GetCreateTime

`func (o *StatementTransaction) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *StatementTransaction) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *StatementTransaction) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.


### GetPerformTime

`func (o *StatementTransaction) GetPerformTime() int64`

GetPerformTime returns the PerformTime field if non-nil, zero value otherwise.

### GetPerformTimeOk

`func (o *StatementTransaction) GetPerformTimeOk() (*int64, bool)`

GetPerformTimeOk returns a tuple with the PerformTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformTime

`func (o *StatementTransaction) SetPerformTime(v int64)`

SetPerformTime sets PerformTime field to given value.


### GetCancelTime

`func (o *StatementTransaction) GetCancelTime() int64`

GetCancelTime returns the CancelTime field if non-nil, zero value otherwise.

### GetCancelTimeOk

`func (o *StatementTransaction) GetCancelTimeOk() (*int64, bool)`

GetCancelTimeOk returns a tuple with the CancelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelTime

`func (o *StatementTransaction) SetCancelTime(v int64)`

SetCancelTime sets CancelTime field to given value.


### GetTransaction

`func (o *StatementTransaction) GetTransaction() string`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *StatementTransaction) GetTransactionOk() (*string, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *StatementTransaction) SetTransaction(v string)`

SetTransaction sets Transaction field to given value.


### GetState

`func (o *StatementTransaction) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StatementTransaction) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StatementTransaction) SetState(v int32)`

SetState sets State field to given value.


### GetReason

`func (o *StatementTransaction) GetReason() int32`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *StatementTransaction) GetReasonOk() (*int32, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *StatementTransaction) SetReason(v int32)`

SetReason sets Reason field to given value.


### GetReceivers

`func (o *StatementTransaction) GetReceivers() []Receiver`

GetReceivers returns the Receivers field if non-nil, zero value otherwise.

### GetReceiversOk

`func (o *StatementTransaction) GetReceiversOk() (*[]Receiver, bool)`

GetReceiversOk returns a tuple with the Receivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivers

`func (o *StatementTransaction) SetReceivers(v []Receiver)`

SetReceivers sets Receivers field to given value.

### HasReceivers

`func (o *StatementTransaction) HasReceivers() bool`

HasReceivers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


