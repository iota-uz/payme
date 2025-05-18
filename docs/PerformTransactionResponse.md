# PerformTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transaction** | **string** |  | 
**PerformTime** | **int64** |  | 
**State** | **int32** |  | 

## Methods

### NewPerformTransactionResponse

`func NewPerformTransactionResponse(transaction string, performTime int64, state int32, ) *PerformTransactionResponse`

NewPerformTransactionResponse instantiates a new PerformTransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerformTransactionResponseWithDefaults

`func NewPerformTransactionResponseWithDefaults() *PerformTransactionResponse`

NewPerformTransactionResponseWithDefaults instantiates a new PerformTransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransaction

`func (o *PerformTransactionResponse) GetTransaction() string`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *PerformTransactionResponse) GetTransactionOk() (*string, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *PerformTransactionResponse) SetTransaction(v string)`

SetTransaction sets Transaction field to given value.


### GetPerformTime

`func (o *PerformTransactionResponse) GetPerformTime() int64`

GetPerformTime returns the PerformTime field if non-nil, zero value otherwise.

### GetPerformTimeOk

`func (o *PerformTransactionResponse) GetPerformTimeOk() (*int64, bool)`

GetPerformTimeOk returns a tuple with the PerformTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformTime

`func (o *PerformTransactionResponse) SetPerformTime(v int64)`

SetPerformTime sets PerformTime field to given value.


### GetState

`func (o *PerformTransactionResponse) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PerformTransactionResponse) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PerformTransactionResponse) SetState(v int32)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


