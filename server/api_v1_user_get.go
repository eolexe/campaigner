package server

import (
	"net/http"

	"github.com/eolexe/campaigner/model"
	"github.com/eolexe/campaigner/server/httperror"
)

func (s *Server) ApiV1UserGet(w http.ResponseWriter, r *http.Request) {
	s.env.UserCounter = s.env.UserCounter + 1
	user, err := model.NewUser(s.env.UserCounter)

	if err != nil {
		s.SendError(w, httperror.NewHttpErrorGenerateCampaignFailed(err))
		return
	}

	s.JSON(w, http.StatusOK, user)
}
