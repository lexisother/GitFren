package src

import (
	"git.nova-vps.ml/lexisother/GitFren/middle"
	"github.com/20kdc/CCUpdaterUI/design"
	"github.com/20kdc/CCUpdaterUI/frenyard/framework"
)

var baseURL = "https://git.nova-vps.ml/api/v1"

func (app *UpApplication) ShowPrimaryView() {
	slots := []framework.FlexboxSlot{}

	repoNames := middle.GetRepoNames()
	for _, name := range repoNames {
		repo := middle.GetRepoData(baseURL, name)
		slots = append(slots, framework.FlexboxSlot{
			Element: design.ListItem(design.ListItemDetails{
				Text: repo.Name,
			}),
		})
	}

	app.Teleport(design.LayoutDocument(design.Header{
		Title: "GitFren",
	}, framework.NewUIFlexboxContainerPtr(framework.FlexboxContainer{
		DirVertical: true,
		Slots: slots,
	}), true))
}
