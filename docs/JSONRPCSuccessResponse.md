# JSONRPCSuccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | [**JSONRPCSuccessResponseResult**](JSONRPCSuccessResponseResult.md) |  | 
**Id** | **int64** |  | 

## Methods

### NewJSONRPCSuccessResponse

`func NewJSONRPCSuccessResponse(result JSONRPCSuccessResponseResult, id int64, ) *JSONRPCSuccessResponse`

NewJSONRPCSuccessResponse instantiates a new JSONRPCSuccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSONRPCSuccessResponseWithDefaults

`func NewJSONRPCSuccessResponseWithDefaults() *JSONRPCSuccessResponse`

NewJSONRPCSuccessResponseWithDefaults instantiates a new JSONRPCSuccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *JSONRPCSuccessResponse) GetResult() JSONRPCSuccessResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *JSONRPCSuccessResponse) GetResultOk() (*JSONRPCSuccessResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *JSONRPCSuccessResponse) SetResult(v JSONRPCSuccessResponseResult)`

SetResult sets Result field to given value.


### GetId

`func (o *JSONRPCSuccessResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JSONRPCSuccessResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JSONRPCSuccessResponse) SetId(v int64)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


