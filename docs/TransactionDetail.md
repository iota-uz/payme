# TransactionDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReceiptType** | **int32** |  | 
**Shipping** | Pointer to [**TransactionDetailShipping**](TransactionDetailShipping.md) |  | [optional] 
**Items** | [**[]TransactionItem**](TransactionItem.md) |  | 

## Methods

### NewTransactionDetail

`func NewTransactionDetail(receiptType int32, items []TransactionItem, ) *TransactionDetail`

NewTransactionDetail instantiates a new TransactionDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionDetailWithDefaults

`func NewTransactionDetailWithDefaults() *TransactionDetail`

NewTransactionDetailWithDefaults instantiates a new TransactionDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceiptType

`func (o *TransactionDetail) GetReceiptType() int32`

GetReceiptType returns the ReceiptType field if non-nil, zero value otherwise.

### GetReceiptTypeOk

`func (o *TransactionDetail) GetReceiptTypeOk() (*int32, bool)`

GetReceiptTypeOk returns a tuple with the ReceiptType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptType

`func (o *TransactionDetail) SetReceiptType(v int32)`

SetReceiptType sets ReceiptType field to given value.


### GetShipping

`func (o *TransactionDetail) GetShipping() TransactionDetailShipping`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *TransactionDetail) GetShippingOk() (*TransactionDetailShipping, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *TransactionDetail) SetShipping(v TransactionDetailShipping)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *TransactionDetail) HasShipping() bool`

HasShipping returns a boolean if a field has been set.

### GetItems

`func (o *TransactionDetail) GetItems() []TransactionItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TransactionDetail) GetItemsOk() (*[]TransactionItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TransactionDetail) SetItems(v []TransactionItem)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


