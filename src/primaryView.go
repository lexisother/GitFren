package src

import (
	"git.nova-vps.ml/lexisother/GitFren/middle"
	"github.com/20kdc/CCUpdaterUI/design"
	"github.com/20kdc/CCUpdaterUI/frenyard/framework"
)

var baseURL = "https://git.nova-vps.ml/api/v1"

func (app *UpApplication) getPrimaryViewRepoList() []middle.Repository {
	repos := []middle.Repository{}
	repoNames := middle.GetRepoNames()

	for _, name := range repoNames {
		repos = append(repos, middle.GetRepoData(baseURL, name))
	}

	return repos
}

func (app *UpApplication) ShowPrimaryView() {
	slots := []framework.FlexboxSlot{}

	repoList := app.getPrimaryViewRepoList()
	repoListItems := []design.ListItemDetails{}
	for _, repo := range repoList {
		repoListItems = append(repoListItems, design.ListItemDetails{
			Text: repo.Name,
			Subtext: repo.Description,
		})
	}

	slots = append(slots, framework.FlexboxSlot{
		Element: design.NewUISearchBoxPtr("Search...", repoListItems),
		Grow: 1,
	})

	app.Teleport(design.LayoutDocument(design.Header{
		Title: "GitFren",
	}, framework.NewUIFlexboxContainerPtr(framework.FlexboxContainer{
		DirVertical: true,
		Slots: slots,
	}), true))
}
