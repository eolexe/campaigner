package server

import (
	"net/http"

	"encoding/json"

	"github.com/eolexe/campaigner/model"
	"github.com/eolexe/campaigner/server/httperror"
)

func (s *Server) ApiV1CampaignSearchPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		s.SendError(w, httperror.HttpErrInvalidMethod)
		return
	}

	user := &model.User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)

	if err != nil {
		s.SendError(w, httperror.NewHttpErrorInvalidJson(err))
		return
	}

	s.env.SearchCounter++
	campaign := s.env.InMemorySortedCampaigns.SearchByUser(user)

	winnerName := "none"

	if campaign != nil {
		winnerName = campaign.Name
	}

	s.JSON(w, http.StatusOK, struct {
		Winner  string `json:"winner"`
		Counter int64  `json:"counter"`
	}{
		Winner:  winnerName,
		Counter: s.env.SearchCounter,
	})
}
