package payload

import "github.com/eolexe/campaigner/server/httperror"

type ResponseHttpErrors struct {
	Errors []*httperror.HttpError `json:"errors"`
}
