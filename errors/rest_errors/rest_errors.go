package rest_errors

import (
	"net/http"
)

type RestErr struct {
	Message      string `json:"message,omitempty"`
	DebugMessage string `json:"debug_message,omitempty"`
	Status       int    `json:"status"`
	Error        string `json:"error"`
}

func NewBadRequestError(err error, message string) *RestErr {
	//var messageDebug string
	//if Config.Debug {
	//	messageDebug = err.Error()
	//}

	return &RestErr{
		Message: message,
		//s	DebugMessage: messageDebug,
		Status: http.StatusBadRequest,
		Error:  "bad_request",
	}
}

func NewNotFoundError(err error, message string) *RestErr {
	/*
		var messageDebug string
		if Config.Debug {
			messageDebug = err.Error()
		}
	*/

	return &RestErr{
		Message: message,
		//	DebugMessage: messageDebug,
		Status: http.StatusNotFound,
		Error:  "not_found",
	}
}

func NewInternalServerError(err error, message string) *RestErr {
	/*
		var messageDebug string
		if Config.Debug {
			messageDebug = err.Error()
		}
	*/

	//Logger.Error(message, err)

	return &RestErr{
		Message: message,
		//DebugMessage: messageDebug,
		Status: http.StatusInternalServerError,
		Error:  "internal_server_error",
	}
}

func NewUnauthorizedError(err error, message string) *RestErr {
	/*
		var messageDebug string
		if Config.Debug {
			messageDebug = err.Error()
		}
	*/

	return &RestErr{
		Message: message,
		//	DebugMessage: messageDebug,
		Status: http.StatusUnauthorized,
		Error:  "unauthorized",
	}
}
