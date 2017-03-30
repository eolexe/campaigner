package server

import (
	"net/http"

	"github.com/eolexe/campaigner/model"
	"github.com/eolexe/campaigner/server/httperror"
)

func (s *Server) ApiV1CampaignSearchAutoGet(w http.ResponseWriter, r *http.Request) {
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
