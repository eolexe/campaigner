package server

import (
	"net/http"

	"github.com/eolexe/campaigner/model"
	"github.com/eolexe/campaigner/server/httperror"
)

func (s *Server) ApiV1CampaignsGet(w http.ResponseWriter, r *http.Request) {
	x, err := s.ParamInt(r, "x", true)
	if err != nil {
		s.SendError(w, httperror.NewHttpErrorInvalidQueryParam("x", err))
		return
	}

	y, err := s.ParamInt(r, "y", true)
	if err != nil {
		s.SendError(w, httperror.NewHttpErrorInvalidQueryParam("y", err))
		return
	}

	z, err := s.ParamInt(r, "z", true)
	if err != nil {
		s.SendError(w, httperror.NewHttpErrorInvalidQueryParam("z", err))
		return
	}

	/*
		x - boundaries for target attributes length. Target attributes should be random and < x
		y - boundaries for target list length. Target list should be random and < y
		z - number of campaigns to generate
	*/
	campaigns, err := model.NewCampaigns(z, y, x)

	if err != nil {
		s.SendError(w, httperror.NewHttpErrorGenerateCampaignFailed(err))
		return
	}

	s.JSON(w, http.StatusOK, campaigns)
}
