package helpers

import (
	"encoding/json"

	logger "github.com/tiket-dev/tiket-microservice-configuration/helpers/logger"
)

type Language struct {
	Label string `json:"label"`
}

type SuccessResponse struct {
	Status    int        `json:"status"`
	Languages []Language `json:"languages,omitempty"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// func getResponse(status int) *Response {
// 	return &Response{
// 		Status: status,
// 	}
// }

func SendSuccess(status int, languages []Language /* FIXME: `languages` should be general */) []byte {
	_successResponse := SuccessResponse{
		Status:    status,
		Languages: languages,
	}

	successResponse, err := json.Marshal(_successResponse)

	logger.Error(err, "Failed to marshal response.")

	return successResponse
}

func SendError(status int, code string, msg string) []byte {
	_errorResponse := ErrorResponse{
		Status:  status,
		Code:    code,
		Message: msg,
	}

	errorResponse, err := json.Marshal(_errorResponse)

	logger.Error(err, "Failed to marshal response.")

	return errorResponse
}

func Http200Ok() int {
	return 200
}

func Http201Created() int {
	return 201
}

func Http400BadRequest() int {
	return 400
}

func Http401Unauthorized() int {
	return 401
}

func Http404NotFound() int {
	return 404
}

func htpp500InternalServerError() int {
	return 500
}

func Http503ServiceUnavailable() int {
	return 503
}
