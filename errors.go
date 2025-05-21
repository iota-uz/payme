package paymeapi

func MethodNotPOSTError() JSONRPCErrorResponseError {
	return JSONRPCErrorResponseError{
		Code: ErrorMethodNotPOST,
		Message: JSONRPCErrorResponseErrorMessage{
			Ru: p("Метод запроса должен быть POST"),
			Uz: p("So‘rov usuli POST bo‘lishi kerak"),
			En: p("Request method must be POST"),
		},
	}
}

func JSONParseError() JSONRPCErrorResponseError {
	return JSONRPCErrorResponseError{
		Code: ErrorJSONParse,
		Message: JSONRPCErrorResponseErrorMessage{
			Ru: p("Ошибка при разборе JSON"),
			Uz: p("JSON ni tahlil qilishda xatolik"),
			En: p("Error parsing JSON"),
		},
	}
}

func InvalidRPCFieldsError() JSONRPCErrorResponseError {
	return JSONRPCErrorResponseError{
		Code: ErrorInvalidRPCFields,
		Message: JSONRPCErrorResponseErrorMessage{
			Ru: p("Некорректные поля запроса"),
			Uz: p("So‘rovdagi maydonlar noto‘g‘ri"),
			En: p("Invalid RPC fields"),
		},
	}
}

func MethodNotFoundInDataError() JSONRPCErrorResponseError {
	return JSONRPCErrorResponseError{
		Code: ErrorMethodNotFoundInData,
		Message: JSONRPCErrorResponseErrorMessage{
			Ru: p("Метод не найден в поле data"),
			Uz: p("Data maydonida metod topilmadi"),
			En: p("Method not found in data field"),
		},
	}
}

func InsufficientPrivilegesError() JSONRPCErrorResponseError {
	return JSONRPCErrorResponseError{
		Code: ErrorInsufficientPrivileges,
		Message: JSONRPCErrorResponseErrorMessage{
			Ru: p("Недостаточно прав для выполнения операции"),
			Uz: p("Amalni bajarish uchun huquqlar yetarli emas"),
			En: p("Insufficient privileges to perform operation"),
		},
	}
}

func InternalSystemError() JSONRPCErrorResponseError {
	return JSONRPCErrorResponseError{
		Code: ErrorInternalSystem,
		Message: JSONRPCErrorResponseErrorMessage{
			Ru: p("Внутренняя ошибка системы"),
			Uz: p("Tizim ichki xatosi"),
			En: p("Internal system error"),
		},
	}
}

func InvalidAmountError() JSONRPCErrorResponseError {
	return JSONRPCErrorResponseError{
		Code: ErrorInvalidAmount,
		Message: JSONRPCErrorResponseErrorMessage{
			Ru: p("Некорректная сумма"),
			Uz: p("Noto‘g‘ri miqdor"),
			En: p("Invalid amount"),
		},
	}
}

func TransactionNotFoundError() JSONRPCErrorResponseError {
	return JSONRPCErrorResponseError{
		Code: ErrorTransactionNotFound,
		Message: JSONRPCErrorResponseErrorMessage{
			Ru: p("Транзакция не найдена"),
			Uz: p("Tranzaksiya topilmadi"),
			En: p("Transaction not found"),
		},
	}
}

func UnableToCancelTransactionError() JSONRPCErrorResponseError {
	return JSONRPCErrorResponseError{
		Code: ErrorUnableToCancelTransaction,
		Message: JSONRPCErrorResponseErrorMessage{
			Ru: p("Невозможно отменить транзакцию — услуга уже оказана"),
			Uz: p("Tranzaksiyani bekor qilib bo‘lmaydi — xizmat ko‘rsatilgan"),
			En: p("Unable to cancel transaction — service already provided"),
		},
	}
}

func InvalidTransactionStateError() JSONRPCErrorResponseError {
	return JSONRPCErrorResponseError{
		Code: ErrorInvalidTransactionState,
		Message: JSONRPCErrorResponseErrorMessage{
			Ru: p("Невозможно выполнить операцию — неправильное состояние транзакции"),
			Uz: p("Amalni bajarib bo‘lmaydi — tranzaksiya holati noto‘g‘ri"),
			En: p("Cannot perform operation — invalid transaction state"),
		},
	}
}

func InvalidAccountError() JSONRPCErrorResponseError {
	return JSONRPCErrorResponseError{
		Code: ErrorInvalidAccount,
		Message: JSONRPCErrorResponseErrorMessage{
			Ru: p("Некорректные данные учетной записи"),
			Uz: p("Hisob ma’lumotlari noto‘g‘ri"),
			En: p("Invalid account data"),
		},
		Data: p("account"),
	}
}

func CheckPerformTransactionInvalidAmountError() JSONRPCErrorResponseError {
	return JSONRPCErrorResponseError{
		Code: CheckPerformTransactionErrorInvalidAmount,
		Message: JSONRPCErrorResponseErrorMessage{
			Ru: p("Некорректная сумма"),
			Uz: p("Noto‘g‘ri miqdor"),
			En: p("Invalid amount"),
		},
	}
}

func CheckPerformTransactionInvalidAccountError() JSONRPCErrorResponseError {
	return JSONRPCErrorResponseError{
		Code: CheckPerformTransactionErrorInvalidAccount,
		Message: JSONRPCErrorResponseErrorMessage{
			Ru: p("Некорректные данные учетной записи"),
			Uz: p("Hisob ma’lumotlari noto‘g‘ri"),
			En: p("Invalid account data"),
		},
		Data: p("account"),
	}
}

func PerformTransactionTransactionNotFoundError() JSONRPCErrorResponseError {
	return JSONRPCErrorResponseError{
		Code: PerformTransactionErrorTransactionNotFound,
		Message: JSONRPCErrorResponseErrorMessage{
			Ru: p("Транзакция не найдена"),
			Uz: p("Tranzaksiya topilmadi"),
			En: p("Transaction not found"),
		},
	}
}

func PerformTransactionOperationNotAllowedError() JSONRPCErrorResponseError {
	return JSONRPCErrorResponseError{
		Code: PerformTransactionErrorOperationNotAllowed,
		Message: JSONRPCErrorResponseErrorMessage{
			Ru: p("Невозможно выполнить операцию — неправильное состояние транзакции"),
			Uz: p("Amalni bajarib bo‘lmaydi — tranzaksiya holati noto‘g‘ri"),
			En: p("Cannot perform operation — invalid transaction state"),
		},
	}
}

func PerformTransactionInvalidAccountError() JSONRPCErrorResponseError {
	return JSONRPCErrorResponseError{
		Code: PerformTransactionErrorInvalidAccount,
		Message: JSONRPCErrorResponseErrorMessage{
			Ru: p("Некорректные данные учетной записи"),
			Uz: p("Hisob ma’lumotlari noto‘g‘ri"),
			En: p("Invalid account data"),
		},
		Data: p("account"),
	}
}

func CancelTransactionTransactionNotFoundError() JSONRPCErrorResponseError {
	return JSONRPCErrorResponseError{
		Code: CancelTransactionErrorTransactionNotFound,
		Message: JSONRPCErrorResponseErrorMessage{
			Ru: p("Транзакция не найдена"),
			Uz: p("Tranzaksiya topilmadi"),
			En: p("Transaction not found"),
		},
	}
}

func CancelTransactionCannotCancelCompletedError() JSONRPCErrorResponseError {
	return JSONRPCErrorResponseError{
		Code: CancelTransactionErrorCannotCancelCompleted,
		Message: JSONRPCErrorResponseErrorMessage{
			Ru: p("Невозможно отменить транзакцию — услуга уже оказана"),
			Uz: p("Tranzaksiyani bekor qilib bo‘lmaydi — xizmat ko‘rsatilgan"),
			En: p("Unable to cancel transaction — service already provided"),
		},
	}
}

func CheckTransactionTransactionNotFoundError() JSONRPCErrorResponseError {
	return JSONRPCErrorResponseError{
		Code: CheckTransactionErrorTransactionNotFound,
		Message: JSONRPCErrorResponseErrorMessage{
			Ru: p("Транзакция не найдена"),
			Uz: p("Tranzaksiya topilmadi"),
			En: p("Transaction not found"),
		},
	}
}

func p(s string) *string {
	return &s
}
