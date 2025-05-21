# JSONRPCRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float64** |  | 
**Account** | **map[string]interface{}** |  | 
**Id** | **string** |  | 
**Time** | **int64** |  | 
**Reason** | **int32** |  | 
**From** | **int64** |  | 
**To** | **int64** |  | 
**Type** | **string** |  | 
**FiscalData** | [**FiscalData**](FiscalData.md) |  | 

## Methods

### NewJSONRPCRequestParams

`func NewJSONRPCRequestParams(amount float64, account map[string]interface{}, id string, time int64, reason int32, from int64, to int64, type_ string, fiscalData FiscalData, ) *JSONRPCRequestParams`

NewJSONRPCRequestParams instantiates a new JSONRPCRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSONRPCRequestParamsWithDefaults

`func NewJSONRPCRequestParamsWithDefaults() *JSONRPCRequestParams`

NewJSONRPCRequestParamsWithDefaults instantiates a new JSONRPCRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *JSONRPCRequestParams) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *JSONRPCRequestParams) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *JSONRPCRequestParams) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetAccount

`func (o *JSONRPCRequestParams) GetAccount() map[string]interface{}`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *JSONRPCRequestParams) GetAccountOk() (*map[string]interface{}, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *JSONRPCRequestParams) SetAccount(v map[string]interface{})`

SetAccount sets Account field to given value.


### GetId

`func (o *JSONRPCRequestParams) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JSONRPCRequestParams) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JSONRPCRequestParams) SetId(v string)`

SetId sets Id field to given value.


### GetTime

`func (o *JSONRPCRequestParams) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *JSONRPCRequestParams) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *JSONRPCRequestParams) SetTime(v int64)`

SetTime sets Time field to given value.


### GetReason

`func (o *JSONRPCRequestParams) GetReason() int32`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *JSONRPCRequestParams) GetReasonOk() (*int32, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *JSONRPCRequestParams) SetReason(v int32)`

SetReason sets Reason field to given value.


### GetFrom

`func (o *JSONRPCRequestParams) GetFrom() int64`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *JSONRPCRequestParams) GetFromOk() (*int64, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *JSONRPCRequestParams) SetFrom(v int64)`

SetFrom sets From field to given value.


### GetTo

`func (o *JSONRPCRequestParams) GetTo() int64`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *JSONRPCRequestParams) GetToOk() (*int64, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *JSONRPCRequestParams) SetTo(v int64)`

SetTo sets To field to given value.


### GetType

`func (o *JSONRPCRequestParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JSONRPCRequestParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JSONRPCRequestParams) SetType(v string)`

SetType sets Type field to given value.


### GetFiscalData

`func (o *JSONRPCRequestParams) GetFiscalData() FiscalData`

GetFiscalData returns the FiscalData field if non-nil, zero value otherwise.

### GetFiscalDataOk

`func (o *JSONRPCRequestParams) GetFiscalDataOk() (*FiscalData, bool)`

GetFiscalDataOk returns a tuple with the FiscalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalData

`func (o *JSONRPCRequestParams) SetFiscalData(v FiscalData)`

SetFiscalData sets FiscalData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


