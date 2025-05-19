# CheckPerformTransactionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** |  | 
**Account** | **map[string]interface{}** |  | 

## Methods

### NewCheckPerformTransactionRequest

`func NewCheckPerformTransactionRequest(amount int32, account map[string]interface{}, ) *CheckPerformTransactionRequest`

NewCheckPerformTransactionRequest instantiates a new CheckPerformTransactionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckPerformTransactionRequestWithDefaults

`func NewCheckPerformTransactionRequestWithDefaults() *CheckPerformTransactionRequest`

NewCheckPerformTransactionRequestWithDefaults instantiates a new CheckPerformTransactionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CheckPerformTransactionRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CheckPerformTransactionRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CheckPerformTransactionRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetAccount

`func (o *CheckPerformTransactionRequest) GetAccount() map[string]interface{}`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CheckPerformTransactionRequest) GetAccountOk() (*map[string]interface{}, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CheckPerformTransactionRequest) SetAccount(v map[string]interface{})`

SetAccount sets Account field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


