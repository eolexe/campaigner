package httperror

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var (
	HttpErrInternalError = NewHttpError(
		"Internal Server Error",
		"internal",
		http.StatusInternalServerError,
	)

	HttpErrInvalidMethod = NewHttpError(
		"Invalid request method",
		"invalid_request_method",
		http.StatusMethodNotAllowed,
	)
)

func NewHttpErrorInvalidQueryParam(key string, err error) *HttpError {
	return NewHttpError(
		fmt.Sprintf("Invalid query parameter %s. Details: %s", key, err.Error()),
		"bad_query_param",
		http.StatusBadRequest,
	)
}

func NewHttpErrorGenerateCampaignFailed(err error) *HttpError {
	return NewHttpError(
		err.Error(),
		"campaign_generator_failed_unexpectedly",
		http.StatusInternalServerError,
	)
}

func NewHttpErrorGenerateUserFailed(err error) *HttpError {
	return NewHttpError(
		err.Error(),
		"user_generator_failed_unexpectedly",
		http.StatusInternalServerError,
	)
}

func NewHttpErrorInvalidJson(err error) *HttpError {
	message := "Invalid JSON data, "

	switch e := err.(type) {
	case *json.SyntaxError:
		message = message + e.Error() + " At offset: " + strconv.FormatInt(e.Offset, 10)
	case *json.UnmarshalTypeError:
		message = message + e.Error() + " At offset: " + strconv.FormatInt(e.Offset, 10)
	default:

		message = message + e.Error()
	}

	return NewHttpError(
		message,
		"invalid_json",
		http.StatusBadRequest,
	)
}
