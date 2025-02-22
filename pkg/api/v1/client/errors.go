package client

import "fmt"

// Error holds the cause of a client error and implements the Error interface.
type Error struct {
	Cause string
}

// Error returned for a client side problem.
func (c Error) Error() string {
	return "bom-service client error - " + c.Cause
}

// RequestError is returned when the client gets an error while performing a request.
type RequestError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

// Error returns the RequestError in string format
func (e RequestError) Error() string {
	return fmt.Sprintf("bom-service client request error, statusCode: %d, message: %s", e.StatusCode, e.Message)
}

// ClientError is returned when invalid arguments are provided to the client
type ClientError struct {
	Message string
}
