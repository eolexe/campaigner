package server

import (
	"encoding/json"
	"net/http"

	"errors"
	"fmt"
	"strconv"

	"github.com/eolexe/campaigner/server/httperror"
	"github.com/eolexe/campaigner/server/payload"
	"github.com/eolexe/campaigner/shared/env"
)

const (
	HeaderContentType              = "Content-Type"
	MIMEApplicationJSONCharsetUTF8 = "application/json; charset=utf-8"
)

type Server struct {
	env *env.Environment
}

func NewServer(env *env.Environment) *Server {
	return &Server{
		env: env,
	}
}

func (s *Server) Run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", s.ApiHealth)
	mux.HandleFunc("/campaign", s.ApiV1CampaignsGet)
	mux.HandleFunc("/user", s.ApiV1UserGet)
	mux.HandleFunc("/import_camp", s.ApiV1CampaignsImportPost)
	mux.HandleFunc("/search", s.ApiV1CampaignSearchPost)
	mux.HandleFunc("/search_auto", s.ApiV1CampaignSearchAutoGet)

	return http.ListenAndServe(
		s.env.Config.Server.String(),
		mux,
	)
}

func (s *Server) JSON(w http.ResponseWriter, code int, i interface{}) error {
	data, err := json.Marshal(i)
	if s.env.IsDebug {
		data, err = json.MarshalIndent(i, "", "  ")
	}
	if err != nil {
		return err
	}

	w.Header().Set(HeaderContentType, MIMEApplicationJSONCharsetUTF8)
	w.WriteHeader(code)
	w.Write(data)

	return nil
}

func (s *Server) SendError(w http.ResponseWriter, err *httperror.HttpError) error {
	w.Header().Set("X-Content-Type-Options", "nosniff")
	return s.JSON(w, err.StatusCode, payload.ResponseHttpErrors{[]*httperror.HttpError{err}})
}

func (s *Server) ParamInt(r *http.Request, key string, required bool, defaultValue ...int64) (int64, error) {
	var defaultVal int64
	if len(defaultValue) > 0 {
		defaultVal = defaultValue[0]
	}

	strVal, err := s.ParamString(r, key, required, "")
	if err != nil {
		return 0, err
	}

	if strVal == "" {
		return defaultVal, nil
	}

	if value, err := strconv.ParseInt(strVal, 10, 64); err != nil {
		return 0, errors.New(fmt.Sprintf("expected integer in %s, but got: %s", key, strVal))
	} else {
		return value, nil
	}
}

func (s *Server) ParamString(r *http.Request, key string, required bool, defaultValue ...string) (string, error) {
	var defaultVal string
	if len(defaultValue) > 0 {
		defaultVal = defaultValue[0]
	}

	value := r.URL.Query().Get(key)
	if value != "" {
		return value, nil
	}

	if required {
		return "", errors.New(fmt.Sprintf("query parameter: " + key + " must be specified"))
	}

	return defaultVal, nil
}
