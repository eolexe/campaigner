package server

import (
	"net/http"

	"github.com/eolexe/campaigner/server/httperror"
)

func (s *Server) ApiHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		s.SendError(w, httperror.HttpErrInvalidMethod)
		return
	}

	response := map[string]string{
		"version": s.env.Version,
		"name":    s.env.Config.Name,
	}

	s.JSON(w, http.StatusOK, response)
}
