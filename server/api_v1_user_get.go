package server

import (
	"net/http"

	"github.com/eolexe/campaigner/model"
	"github.com/eolexe/campaigner/server/httperror"
)

func (s *Server) ApiV1UserGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		s.SendError(w, httperror.HttpErrInvalidMethod)
		return
	}

	s.env.UserCounter++
	user, err := model.NewUser(s.env.UserCounter)

	if err != nil {
		s.SendError(w, httperror.NewHttpErrorGenerateUserFailed(err))
		return
	}

	s.JSON(w, http.StatusOK, user)
}
