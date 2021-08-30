package src

import (
	"git.nova-vps.ml/lexisother/GitFren/middle"
	"git.nova-vps.ml/lexisother/GitFren/middle/api"
	"github.com/20kdc/CCUpdaterUI/design"
	"github.com/20kdc/CCUpdaterUI/frenyard/framework"
)

var baseURL = "https://git.nova-vps.ml/api/v1"

func (app *UpApplication) getPrimaryViewRepoList() []api.Repository {
	repos := []api.Repository{}
	users := middle.GetUserNames()

	for _, name := range users {
		userRepos := middle.GetUserRepos(baseURL, name)
		for _, repo := range userRepos {
			repos = append(repos, repo)
		}
	}

	return repos
}

func (app *UpApplication) ShowPrimaryView() {
	if app.CachedPrimaryView != nil {
		app.Teleport(app.CachedPrimaryView)
		return
	}

	slots := []framework.FlexboxSlot{}

	repoList := app.getPrimaryViewRepoList()
	repoListItems := []design.ListItemDetails{}
	for _, repo := range repoList {
		localRepo := repo
		repoListItems = append(repoListItems, design.ListItemDetails{
			Text: repo.Name,
			Subtext: repo.Description,
			Click: func () {
				app.GSRightwards()
				app.ShowRepoView(func () {
					app.GSLeftwards()
					app.ShowPrimaryView()
				}, localRepo)
			},
		})
	}

	slots = append(slots, framework.FlexboxSlot{
		Element: design.NewUISearchBoxPtr("Search...", repoListItems),
		Grow: 1,
	})

	app.CachedPrimaryView = design.LayoutDocument(design.Header{
		Title: "Repositories",
	}, framework.NewUIFlexboxContainerPtr(framework.FlexboxContainer{
		DirVertical: true,
		Slots: slots,
	}), true)

	app.Teleport(app.CachedPrimaryView)
}
