# CreateTransactionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Time** | **int64** |  | 
**Amount** | **int32** |  | 
**Account** | [**Account**](Account.md) |  | 

## Methods

### NewCreateTransactionRequest

`func NewCreateTransactionRequest(id string, time int64, amount int32, account Account, ) *CreateTransactionRequest`

NewCreateTransactionRequest instantiates a new CreateTransactionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransactionRequestWithDefaults

`func NewCreateTransactionRequestWithDefaults() *CreateTransactionRequest`

NewCreateTransactionRequestWithDefaults instantiates a new CreateTransactionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateTransactionRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateTransactionRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateTransactionRequest) SetId(v string)`

SetId sets Id field to given value.


### GetTime

`func (o *CreateTransactionRequest) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *CreateTransactionRequest) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *CreateTransactionRequest) SetTime(v int64)`

SetTime sets Time field to given value.


### GetAmount

`func (o *CreateTransactionRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateTransactionRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateTransactionRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetAccount

`func (o *CreateTransactionRequest) GetAccount() Account`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CreateTransactionRequest) GetAccountOk() (*Account, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CreateTransactionRequest) SetAccount(v Account)`

SetAccount sets Account field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


