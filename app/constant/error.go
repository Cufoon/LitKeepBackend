package constant

import (
	"errors"
)

var (
	ErrAccountExist                = errors.New("account-exist")
	ErrAccountIDExist              = errors.New("account-id-exist")
	ErrLoginPassWrong              = errors.New("login-pass-wrong")
	ErrLoginEmailNotExist          = errors.New("login-email-no-exist")
	ErrBillRecordQueryNoUserID     = errors.New("bill-record-query-no-user-id")
	ErrBillRecordDeleteNoUserID    = errors.New("bill-record-delete-no-user-id")
	ErrBillRecordModifyParamsWrong = errors.New("bill-record-modify-params-wrong")
)
