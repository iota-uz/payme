package paymeapi

// Methods
const (
	CreateTransactionMethod       = "CreateTransaction"
	CancelTransactionMethod       = "CancelTransaction"
	CheckPerformTransactionMethod = "CheckPerformTransaction"
	CheckTransactionMethod        = "CheckTransaction"
	GetStatementMethod            = "GetStatement"
	SetFiscalDataMethod           = "SetFiscalData"
)

// Transaction State
const (
	// TransactionStateCreated Transaction successfully created, waiting for confirmation (initial state 0)
	TransactionStateCreated = 1

	// TransactionStateCompleted Transaction successfully completed (initial state 1)
	TransactionStateCompleted = 2

	// TransactionStateCancelledBeforeCompletion Transaction cancelled before completion (initial state 1)
	TransactionStateCancelledBeforeCompletion = -1

	// TransactionStateCancelledAfterCompletion Transaction cancelled after completion (initial state 2)
	TransactionStateCancelledAfterCompletion = -2
)

// Transaction Cancel Reason
const (
	// CancelReasonRecipientNotFound One or more recipients not found or inactive in Payme Business
	CancelReasonRecipientNotFound = 1

	// CancelReasonDebitError Error during debit operation in processing center
	CancelReasonDebitError = 2

	// CancelReasonExecutionError Transaction execution error
	CancelReasonExecutionError = 3

	// CancelReasonTimeout Transaction cancelled due to timeout
	CancelReasonTimeout = 4

	// CancelReasonRefund Refund issued
	CancelReasonRefund = 5

	// CancelReasonUnknown Unknown error
	CancelReasonUnknown = 10
)

// Receipt (Check) Status
const (
	// CheckStatusCreated Check created. Waiting for payment confirmation.
	CheckStatusCreated = 0

	// CheckStatusValidationStarted First validation stage. Creating a transaction in the merchant billing system.
	CheckStatusValidationStarted = 1

	// CheckStatusDebited Funds withdrawn from the card.
	CheckStatusDebited = 2

	// CheckStatusClosedInBilling Transaction closed in the merchant billing system.
	CheckStatusClosedInBilling = 3

	// CheckStatusPaid Check paid.
	CheckStatusPaid = 4

	// CheckStatusAuthorized Check is authorized (funds are held).
	CheckStatusAuthorized = 5

	// CheckStatusAuthorizationCommandReceived Received command to authorize, will move to status 5. If stuck — contact Payme Business support.
	CheckStatusAuthorizationCommandReceived = 6

	// CheckStatusManualIntervention Check is paused for manual intervention.
	CheckStatusManualIntervention = 20

	// CheckStatusQueuedForCancel Check is in cancel queue.
	CheckStatusQueuedForCancel = 21

	// CheckStatusQueuedForBillingClose Check is in queue for closing transaction in merchant billing.
	CheckStatusQueuedForBillingClose = 30

	// CheckStatusCancelled Check is cancelled.
	CheckStatusCancelled = 50
)

// Common Payme API Errors
const (
	// ErrorMethodNotPOST Error occurs if the request method is not POST.
	ErrorMethodNotPOST = -32300

	// ErrorJSONParse JSON parsing error.
	ErrorJSONParse = -32700

	// ErrorInvalidRPCFields Required fields are missing in the RPC request, or field types do not match the specification.
	ErrorInvalidRPCFields = -32600

	// ErrorMethodNotFoundInData Requested method not found. In RPC, the method name is in the data field.
	ErrorMethodNotFoundInData = -32601

	// ErrorInsufficientPrivileges Insufficient privileges to execute the method.
	ErrorInsufficientPrivileges = -32504

	// ErrorInternalSystem System (internal) error. Used in cases of DB failure, filesystem issues, undefined behavior, etc.
	ErrorInternalSystem = -32400
)

// Merchant Server Response Errors
const (
	// ErrorInvalidAmount Invalid amount. Happens when the transaction amount does not match the order amount.
	ErrorInvalidAmount = -31001

	// ErrorTransactionNotFound Transaction not found.
	ErrorTransactionNotFound = -31003

	// ErrorUnableToCancelTransaction Unable to cancel transaction. The product or service has already been fully provided.
	ErrorUnableToCancelTransaction = -31007

	// ErrorInvalidTransactionState Unable to perform the operation. Transaction state does not allow it.
	ErrorInvalidTransactionState = -31008

	// ErrorInvalidAccount Account-related error. For example, user login or phone number not found.
	// The localized error must contain "message", and "data" must include the failed "account" subfield.
	ErrorInvalidAccount = -31050 // NOTE: Range is -31050 to -31099; adjust dynamically as needed
)

// Errors specific to CheckPerformTransaction method
const (
	// CheckPerformTransactionErrorInvalidAmount Invalid amount.
	CheckPerformTransactionErrorInvalidAmount = -31001

	// CheckPerformTransactionErrorInvalidAccount Buyer account input errors.
	// For example: login not found, phone number not found, etc.
	// Localized field "message" is required, and "data" must include "account".
	CheckPerformTransactionErrorInvalidAccount = -31050 // range: -31050 to -31099
)

// Errors specific to PerformTransaction method
const (
	// PerformTransactionErrorTransactionNotFound Transaction not found.
	PerformTransactionErrorTransactionNotFound = -31003

	// PerformTransactionErrorOperationNotAllowed Operation cannot be performed due to transaction state or other reasons.
	PerformTransactionErrorOperationNotAllowed = -31008

	// PerformTransactionErrorInvalidAccount Buyer account input errors.
	// For example: login not found, phone number not found, etc.
	// Localized "message" field is required. "data" must include the "account" field name.
	PerformTransactionErrorInvalidAccount = -31050 // range: -31050 to -31099
)

// Errors specific to CancelTransaction method
const (
	// CancelTransactionErrorTransactionNotFound Transaction not found.
	CancelTransactionErrorTransactionNotFound = -31003

	// CancelTransactionErrorCannotCancelCompleted Order completed. Cannot cancel the transaction because goods or services have been fully delivered.
	CancelTransactionErrorCannotCancelCompleted = -31007
)

// Errors specific to CheckTransaction method
const (
	// CheckTransactionErrorTransactionNotFound Transaction not found.
	CheckTransactionErrorTransactionNotFound = -31003
)
