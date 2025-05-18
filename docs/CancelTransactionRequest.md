# CancelTransactionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Reason** | **int32** |  | 

## Methods

### NewCancelTransactionRequest

`func NewCancelTransactionRequest(id string, reason int32, ) *CancelTransactionRequest`

NewCancelTransactionRequest instantiates a new CancelTransactionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelTransactionRequestWithDefaults

`func NewCancelTransactionRequestWithDefaults() *CancelTransactionRequest`

NewCancelTransactionRequestWithDefaults instantiates a new CancelTransactionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CancelTransactionRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CancelTransactionRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CancelTransactionRequest) SetId(v string)`

SetId sets Id field to given value.


### GetReason

`func (o *CancelTransactionRequest) GetReason() int32`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CancelTransactionRequest) GetReasonOk() (*int32, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CancelTransactionRequest) SetReason(v int32)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


