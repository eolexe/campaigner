package server

import (
	"net/http"

	"encoding/json"

	"github.com/eolexe/campaigner/model"
	"github.com/eolexe/campaigner/server/httperror"
)

func (s *Server) ApiV1CampaignsImportPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		s.SendError(w, httperror.HttpErrInvalidMethod)
		return
	}

	campaigns := model.Campaigns{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&campaigns)

	if err != nil {
		s.SendError(w, httperror.NewHttpErrorInvalidJson(err))
		return
	}

	s.env.InMemorySortedCampaigns = model.ImportCampaigns(campaigns)

	s.JSON(w, http.StatusOK, s.env.InMemorySortedCampaigns)
}
