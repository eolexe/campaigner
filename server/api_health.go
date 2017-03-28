package server

import "net/http"

func (s *Server) ApiHealth(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"version": s.env.Version,
		"name":    s.env.Config.Name,
	}

	s.JSON(w, http.StatusOK, response)
}
