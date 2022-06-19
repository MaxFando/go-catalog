package handler

import (
	"github.com/MaxFando/go-catalog/internal/logic/app/repos/communityrepo"
	"github.com/MaxFando/go-catalog/internal/logic/app/repos/grouprepo"
	"github.com/MaxFando/go-catalog/internal/logic/app/repos/organizationrepo"
	"github.com/MaxFando/go-catalog/internal/logic/app/repos/projectrepo"
	"github.com/MaxFando/go-catalog/internal/logic/app/repos/userrepo"
	"github.com/google/uuid"
	"net/http"
)

type Router struct {
	*http.ServeMux
	us  *userrepo.Users
	org *organizationrepo.Organizations
	gr  *grouprepo.Groups
	pr  *projectrepo.Projects
	co  *communityrepo.Communities
}

func NewRouter(
	us *userrepo.Users,
	org *organizationrepo.Organizations,
	gr *grouprepo.Groups,
	pr *projectrepo.Projects,
	co *communityrepo.Communities,
) *Router {
	r := &Router{
		ServeMux: http.NewServeMux(),
		us:       us,
		org:      org,
		gr:       gr,
		pr:       pr,
		co:       co,
	}
	r.Handle("/create", r.AuthMiddleware(http.HandlerFunc(r.CreateUser)))
	r.Handle("/read", r.AuthMiddleware(http.HandlerFunc(r.ReadUser)))
	r.Handle("/delete", r.AuthMiddleware(http.HandlerFunc(r.DeleteUser)))
	r.Handle("/search", r.AuthMiddleware(http.HandlerFunc(r.SearchUser)))
	return r
}

// DTO object
type User struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Data       string    `json:"data"`
	Permission int       `json:"perms"`
}

func (rt *Router) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if u, p, ok := r.BasicAuth(); !ok || !(u == "admin" && p == "admin") {
				http.Error(w, "unautorized", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		},
	)
}

func (rt *Router) CreateUser(w http.ResponseWriter, r *http.Request) {

}

// read?uid=...
func (rt *Router) ReadUser(w http.ResponseWriter, r *http.Request) {

}

func (rt *Router) DeleteUser(w http.ResponseWriter, r *http.Request) {

}

// /search?q=...
func (rt *Router) SearchUser(w http.ResponseWriter, r *http.Request) {

}
