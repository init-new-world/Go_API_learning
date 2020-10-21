package errno

var (
	// Common errors
	OK                   = &Errno{Code: 0, Message: "OK"}
	InternalServerError  = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind              = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	ConnectDatabaseError = &Errno{Code: 10003, Message: "Connect database failed."}

	// User errors
	// module 00 : user auth errors
	ErrValidation        = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase          = &Errno{Code: 20002, Message: "Database error."}
	ErrToken             = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}
	ErrTokenInvalid      = &Errno{Code: 20004, Message: "The token was invalid."}
	ErrPasswordIncorrect = &Errno{Code: 20005, Message: "The password was incorrect."}
	ErrLoginFirst        = &Errno{Code: 20006, Message: "You need to login first."}

	// module 01 : user create errors
	ErrUserAlreadyExist = &Errno{Code: 20101, Message: "The username already exist."}
	ErrEncrypt          = &Errno{Code: 20102, Message: "Error occurred while encrypting the user password."}

	// module 02: user update/delete/get errors
	ErrUserNotFound = &Errno{Code: 20201, Message: "The user cannot found in database."}
)
