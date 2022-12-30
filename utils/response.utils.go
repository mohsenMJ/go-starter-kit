package utils

import (
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(status bool, message string, data interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Error:   nil,
		Data:    data,
	}
}

func BuildErrorResponse(message string, err error, data interface{}) Response {
	return Response{
		Status:  false,
		Message: message,
		Error:   strings.Split(err.Error(), "\n"),
		Data:    data,
	}
}
