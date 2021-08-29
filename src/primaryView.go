package src

import (
	"github.com/20kdc/CCUpdaterUI/design"
	"github.com/20kdc/CCUpdaterUI/frenyard/framework"
	"git.nova-vps.ml/lexisother/GitFren/middle"
)

func (app *UpApplication) ShowPrimaryView() {
	slots := []framework.FlexboxSlot{}

	repoNames := middle.GetRepoNames()
	for _, name := range repoNames {
		slots = append(slots, framework.FlexboxSlot{
			Element: design.ListItem(design.ListItemDetails{
				Text: name,
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
