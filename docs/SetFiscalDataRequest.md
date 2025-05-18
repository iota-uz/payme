# SetFiscalDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **string** |  | 
**FiscalData** | [**FiscalData**](FiscalData.md) |  | 

## Methods

### NewSetFiscalDataRequest

`func NewSetFiscalDataRequest(id string, type_ string, fiscalData FiscalData, ) *SetFiscalDataRequest`

NewSetFiscalDataRequest instantiates a new SetFiscalDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetFiscalDataRequestWithDefaults

`func NewSetFiscalDataRequestWithDefaults() *SetFiscalDataRequest`

NewSetFiscalDataRequestWithDefaults instantiates a new SetFiscalDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SetFiscalDataRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SetFiscalDataRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SetFiscalDataRequest) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *SetFiscalDataRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SetFiscalDataRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SetFiscalDataRequest) SetType(v string)`

SetType sets Type field to given value.


### GetFiscalData

`func (o *SetFiscalDataRequest) GetFiscalData() FiscalData`

GetFiscalData returns the FiscalData field if non-nil, zero value otherwise.

### GetFiscalDataOk

`func (o *SetFiscalDataRequest) GetFiscalDataOk() (*FiscalData, bool)`

GetFiscalDataOk returns a tuple with the FiscalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalData

`func (o *SetFiscalDataRequest) SetFiscalData(v FiscalData)`

SetFiscalData sets FiscalData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


