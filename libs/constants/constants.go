package consts

import "errors"

const (
	ErrorFieldsEmpty                  = "Fields are empty"
	ErrorUserNotFound                 = "User not found"
	ErrorSenderUserNotFound           = "Sender user not found"
	ErrorInsufficientFounds           = "Insufficient funds"
	ErrorSenderUserSameAsReceiverUser = "Sender user is the same as receiver user"
)

var RevisionChanged = errors.New("revision changed")
