package src

import (
	"git.nova-vps.ml/lexisother/GitFren/middle/api"
	"github.com/20kdc/CCUpdaterUI/design"
	"github.com/20kdc/CCUpdaterUI/frenyard/framework"
)

func (app *UpApplication) ShowRepoView(back framework.ButtonBehavior, repo api.Repository) {
	slots := []framework.FlexboxSlot{}

	slots = append(slots, framework.FlexboxSlot{
		Element: design.ListItem(design.ListItemDetails{
			Text: repo.Name,
		}),
	})

	app.Teleport(design.LayoutDocument(design.Header{
		Title: repo.Name,
		Back: back,
	}, framework.NewUIFlexboxContainerPtr(framework.FlexboxContainer{
		DirVertical: true,
		Slots: slots,
	}), true))
}
