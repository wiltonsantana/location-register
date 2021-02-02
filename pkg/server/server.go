package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wiltonsantana/location-register/pkg/data/controllers"
	"github.com/wiltonsantana/location-register/pkg/logging"

	"github.com/gorilla/mux"
)

// Server represents the HTTP server
type Server struct {
	port               int
	logger             logging.Logger
	locationController *controllers.LocationController
	srv                *http.Server
}

// Health represents the service's health status
type Health struct {
	Status string `json:"status"`
}

// NewServer creates a new server instance
func NewServer(port int, logger logging.Logger, locationController *controllers.LocationController) Server {
	return Server{port, logger, locationController, nil}
}

// Start starts the http server
func (s *Server) Start(started chan bool) {
	routers := s.createRouters()
	s.logger.Infof("listening on %d", s.port)
	started <- true
	s.srv = &http.Server{Addr: fmt.Sprintf(":%d", s.port), Handler: s.logRequest(routers)}
	err := s.srv.ListenAndServe()
	if err != nil {
		s.logger.Error(err)
		started <- false
	}
}

// Stop stops the server
func (s *Server) Stop() {
	err := s.srv.Shutdown(context.TODO())
	if err != nil {
		s.logger.Error(err)
	}
}

func (s *Server) createRouters() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", s.healthcheckHandler)
	r.HandleFunc("/location", s.locationController.Create).Methods("POST")
	return r
}

// Healthcheck godoc
// @Summary Verify the service health
// @Produce json
// @Success 200 {object} Health
// @Router /healthcheck [get]
func (s *Server) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(&Health{Status: "online"})
	_, err := w.Write(response)
	if err != nil {
		s.logger.Errorf("error sending response, %s\n", err)
	}
}
