package main

import (
	"github.com/20kdc/CCUpdaterUI/frenyard/framework"
	"github.com/20kdc/CCUpdaterUI/design"
)

func (app *upApplication) ShowPrimaryView() {
	slots := []framework.FlexboxSlot{}

	slots = append(slots, framework.FlexboxSlot{
		Element: design.ListItem(design.ListItemDetails{
			Text: "Yes",
			Subtext: "Uh huh",
		}),
	})

	app.Teleport(design.LayoutDocument(design.Header{
		Title: "GitFren",
	}, framework.NewUIFlexboxContainerPtr(framework.FlexboxContainer{
		DirVertical: true,
		Slots: slots,
	}), true))
}
