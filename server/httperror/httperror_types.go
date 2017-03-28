package httperror

import (
	"fmt"
	"net/http"
)

var (
	HttpErrInternalError = NewHttpError(
		"Internal Server Error",
		"internal",
		http.StatusInternalServerError,
	)

	HttpErrInvalidJSON = NewHttpError(
		"Invalid JSON data",
		"bad_json",
		http.StatusBadRequest,
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
