# \DefaultAPI

All URIs are relative to *https://checkout.paycom.uz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Canceltransaction**](DefaultAPI.md#Canceltransaction) | **Post** /CancelTransaction | CancelTransaction operation
[**Checkperformtransaction**](DefaultAPI.md#Checkperformtransaction) | **Post** /CheckPerformTransaction | CheckPerformTransaction operation
[**Checktransaction**](DefaultAPI.md#Checktransaction) | **Post** /CheckTransaction | CheckTransaction operation
[**Createtransaction**](DefaultAPI.md#Createtransaction) | **Post** /CreateTransaction | CreateTransaction operation
[**Getstatement**](DefaultAPI.md#Getstatement) | **Post** /GetStatement | GetStatement operation
[**Performtransaction**](DefaultAPI.md#Performtransaction) | **Post** /PerformTransaction | PerformTransaction operation
[**Setfiscaldata**](DefaultAPI.md#Setfiscaldata) | **Post** /SetFiscalData | SetFiscalData operation



## Canceltransaction

> CancelTransactionResponse Canceltransaction(ctx).CancelTransactionRequest(cancelTransactionRequest).Execute()

CancelTransaction operation

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/payme/paymeapi"
)

func main() {
	cancelTransactionRequest := *openapiclient.NewCancelTransactionRequest("Id_example", int32(123)) // CancelTransactionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Canceltransaction(context.Background()).CancelTransactionRequest(cancelTransactionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Canceltransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Canceltransaction`: CancelTransactionResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Canceltransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCanceltransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cancelTransactionRequest** | [**CancelTransactionRequest**](CancelTransactionRequest.md) |  | 

### Return type

[**CancelTransactionResponse**](CancelTransactionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Checkperformtransaction

> CheckPerformTransactionResponse Checkperformtransaction(ctx).CheckPerformTransactionRequest(checkPerformTransactionRequest).Execute()

CheckPerformTransaction operation

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/payme/paymeapi"
)

func main() {
	checkPerformTransactionRequest := *openapiclient.NewCheckPerformTransactionRequest(int32(123), *openapiclient.NewCheckPerformTransactionRequestAccount("Phone_example")) // CheckPerformTransactionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Checkperformtransaction(context.Background()).CheckPerformTransactionRequest(checkPerformTransactionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Checkperformtransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Checkperformtransaction`: CheckPerformTransactionResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Checkperformtransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckperformtransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkPerformTransactionRequest** | [**CheckPerformTransactionRequest**](CheckPerformTransactionRequest.md) |  | 

### Return type

[**CheckPerformTransactionResponse**](CheckPerformTransactionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Checktransaction

> CheckTransactionResponse Checktransaction(ctx).CheckTransactionRequest(checkTransactionRequest).Execute()

CheckTransaction operation

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/payme/paymeapi"
)

func main() {
	checkTransactionRequest := *openapiclient.NewCheckTransactionRequest("Id_example") // CheckTransactionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Checktransaction(context.Background()).CheckTransactionRequest(checkTransactionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Checktransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Checktransaction`: CheckTransactionResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Checktransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChecktransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkTransactionRequest** | [**CheckTransactionRequest**](CheckTransactionRequest.md) |  | 

### Return type

[**CheckTransactionResponse**](CheckTransactionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Createtransaction

> CreateTransactionResponse Createtransaction(ctx).CreateTransactionRequest(createTransactionRequest).Execute()

CreateTransaction operation

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/payme/paymeapi"
)

func main() {
	createTransactionRequest := *openapiclient.NewCreateTransactionRequest("Id_example", int64(123), int32(123), *openapiclient.NewAccount()) // CreateTransactionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Createtransaction(context.Background()).CreateTransactionRequest(createTransactionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Createtransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Createtransaction`: CreateTransactionResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Createtransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatetransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTransactionRequest** | [**CreateTransactionRequest**](CreateTransactionRequest.md) |  | 

### Return type

[**CreateTransactionResponse**](CreateTransactionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Getstatement

> GetStatementResponse Getstatement(ctx).GetStatementRequest(getStatementRequest).Execute()

GetStatement operation

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/payme/paymeapi"
)

func main() {
	getStatementRequest := *openapiclient.NewGetStatementRequest(int64(123), int64(123)) // GetStatementRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Getstatement(context.Background()).GetStatementRequest(getStatementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Getstatement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Getstatement`: GetStatementResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Getstatement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetstatementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getStatementRequest** | [**GetStatementRequest**](GetStatementRequest.md) |  | 

### Return type

[**GetStatementResponse**](GetStatementResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Performtransaction

> PerformTransactionResponse Performtransaction(ctx).PerformTransactionRequest(performTransactionRequest).Execute()

PerformTransaction operation

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/payme/paymeapi"
)

func main() {
	performTransactionRequest := *openapiclient.NewPerformTransactionRequest("Id_example") // PerformTransactionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Performtransaction(context.Background()).PerformTransactionRequest(performTransactionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Performtransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Performtransaction`: PerformTransactionResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Performtransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPerformtransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **performTransactionRequest** | [**PerformTransactionRequest**](PerformTransactionRequest.md) |  | 

### Return type

[**PerformTransactionResponse**](PerformTransactionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Setfiscaldata

> SetFiscalDataResponse Setfiscaldata(ctx).SetFiscalDataRequest(setFiscalDataRequest).Execute()

SetFiscalData operation

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/payme/paymeapi"
)

func main() {
	setFiscalDataRequest := *openapiclient.NewSetFiscalDataRequest("Id_example", "Type_example", *openapiclient.NewFiscalData(int32(123), int32(123), "Message_example", "TerminalId_example", "FiscalSign_example", "QrCodeUrl_example", "Date_example")) // SetFiscalDataRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Setfiscaldata(context.Background()).SetFiscalDataRequest(setFiscalDataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Setfiscaldata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Setfiscaldata`: SetFiscalDataResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Setfiscaldata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetfiscaldataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setFiscalDataRequest** | [**SetFiscalDataRequest**](SetFiscalDataRequest.md) |  | 

### Return type

[**SetFiscalDataResponse**](SetFiscalDataResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

