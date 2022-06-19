package server

import (
	"context"
	"github.com/MaxFando/go-catalog/internal/logic/app/repos/communityrepo"
	"github.com/MaxFando/go-catalog/internal/logic/app/repos/grouprepo"
	"github.com/MaxFando/go-catalog/internal/logic/app/repos/organizationrepo"
	"github.com/MaxFando/go-catalog/internal/logic/app/repos/projectrepo"
	"github.com/MaxFando/go-catalog/internal/logic/app/repos/userrepo"
	"github.com/MaxFando/go-catalog/internal/logic/app/starter"
	"net/http"
	"time"
)

var _ starter.APIServer = &Server{}

type Server struct {
	srv http.Server
	us  *userrepo.Users
}

func NewServer(addr string, h http.Handler) *Server {
	s := &Server{}

	s.srv = http.Server{
		Addr:              addr,
		Handler:           h,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}
	return s
}

func (s *Server) Start(users *userrepo.Users, organizations *organizationrepo.Organizations, groups *grouprepo.Groups, projects *projectrepo.Projects, communities *communityrepo.Communities) {
	s.us = users
	go s.srv.ListenAndServe()
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	s.srv.Shutdown(ctx)
	cancel()
}
