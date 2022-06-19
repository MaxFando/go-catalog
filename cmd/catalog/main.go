package main

import (
	"context"
	"github.com/MaxFando/go-catalog/internal/infrastructure/api/handler"
	"github.com/MaxFando/go-catalog/internal/infrastructure/api/server"
	"github.com/MaxFando/go-catalog/internal/infrastructure/db/mem/communiitymemstore"
	"github.com/MaxFando/go-catalog/internal/infrastructure/db/mem/groupmemstore"
	"github.com/MaxFando/go-catalog/internal/infrastructure/db/mem/organizationememstore"
	"github.com/MaxFando/go-catalog/internal/infrastructure/db/mem/projectmemstore"
	"github.com/MaxFando/go-catalog/internal/infrastructure/db/mem/usermemstore"
	"github.com/MaxFando/go-catalog/internal/logic/app/repos/communityrepo"
	"github.com/MaxFando/go-catalog/internal/logic/app/repos/grouprepo"
	"github.com/MaxFando/go-catalog/internal/logic/app/repos/organizationrepo"
	"github.com/MaxFando/go-catalog/internal/logic/app/repos/projectrepo"
	"github.com/MaxFando/go-catalog/internal/logic/app/repos/userrepo"
	"github.com/MaxFando/go-catalog/internal/logic/app/starter"
	"os"
	"os/signal"
	"sync"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	usersStore := usermemstore.New()
	projectsStore := projectmemstore.New()
	groupsStore := groupmemstore.New()
	organizationsStore := organizationememstore.New()
	communitiesStore := communiitymemstore.New()

	app := starter.NewApp(usersStore, organizationsStore, groupsStore, projectsStore, communitiesStore)
	usersRepo := userrepo.New(usersStore)
	projectsRepo := projectrepo.New(projectsStore)
	organizationsRepo := organizationrepo.New(organizationsStore)
	groupsRepo := grouprepo.New(groupsStore)
	communitiesRepo := communityrepo.New(communitiesStore)
	h := handler.NewRouter(usersRepo, organizationsRepo, groupsRepo, projectsRepo, communitiesRepo)
	srv := server.NewServer(":8000", h)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go app.Serve(ctx, wg, srv)

	<-ctx.Done()
	cancel()
	wg.Wait()
}
