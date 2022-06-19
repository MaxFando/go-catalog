package starter

import (
	"context"
	"github.com/MaxFando/go-catalog/internal/logic/app/repos/communityrepo"
	"github.com/MaxFando/go-catalog/internal/logic/app/repos/grouprepo"
	"github.com/MaxFando/go-catalog/internal/logic/app/repos/organizationrepo"
	"github.com/MaxFando/go-catalog/internal/logic/app/repos/projectrepo"
	"github.com/MaxFando/go-catalog/internal/logic/app/repos/userrepo"
	"sync"
)

type App struct {
	users         *userrepo.Users
	organizations *organizationrepo.Organizations
	groups        *grouprepo.Groups
	projects      *projectrepo.Projects
	communities   *communityrepo.Communities
}

func NewApp(
	userStore userrepo.UserStore,
	organizationStore organizationrepo.OrganizationStore,
	groupStore grouprepo.GroupStore,
	projectStore projectrepo.ProjectStore,
	communityStore communityrepo.CommunityStore,
) *App {
	a := &App{
		users:         userrepo.New(userStore),
		organizations: organizationrepo.New(organizationStore),
		groups:        grouprepo.New(groupStore),
		projects:      projectrepo.New(projectStore),
		communities:   communityrepo.New(communityStore),
	}
	return a
}

type APIServer interface {
	Start(users *userrepo.Users, organizations *organizationrepo.Organizations, groups *grouprepo.Groups, projects *projectrepo.Projects, communities *communityrepo.Communities)
	Stop()
}

func (a *App) Serve(ctx context.Context, wg *sync.WaitGroup, hs APIServer) {
	defer wg.Done()
	hs.Start(a.users, a.organizations, a.groups, a.projects, a.communities)
	<-ctx.Done()
	hs.Stop()
}
