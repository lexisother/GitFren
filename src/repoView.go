package src

import (
	"strconv"

	"git.nova-vps.ml/lexisother/GitFren/middle"
	"git.nova-vps.ml/lexisother/GitFren/middle/api"
	"github.com/20kdc/CCUpdaterUI/design"
	"github.com/20kdc/CCUpdaterUI/frenyard"
	"github.com/20kdc/CCUpdaterUI/frenyard/framework"
	"github.com/20kdc/CCUpdaterUI/frenyard/integration"
)

// Clearly inspired packageView. It's not like I really had any other option to display metadata properly.
// TODO: Overhaul? Someday?
func (app *UpApplication) ShowRepoView(back framework.ButtonBehavior, repo api.Repository) {
	annotations := "\n    ID: " + strconv.Itoa(repo.ID)

	releases := middle.GetReleases(baseURL, repo)
	if len(releases) > 0 {
		annotations += "\n    Latest Version: " + releases[0].TagName
	}

	chunks := []integration.TypeChunk{
		integration.NewColouredTextTypeChunk(repo.Name, design.GlobalFont, design.ThemeText),
		integration.NewColouredTextTypeChunk(annotations, design.ListItemSubTextFont, design.ThemeSubText),
	}

	descriptionText := repo.Description + "\n"

	detail := framework.NewUIFlexboxContainerPtr(framework.FlexboxContainer{
		DirVertical: false,
		Slots: []framework.FlexboxSlot{
			// framework.FlexboxSlot{
			// 	Element: design.NewIconPtr(0xFFFFFFFF, middle.PackageIcon(latestPkg), 36),
			// },
			{
				Basis: design.SizeMarginAroundEverything,
			},
			{
				Element: framework.NewUILabelPtr(integration.NewCompoundTypeChunk(chunks), 0xFFFFFFFF, 0, frenyard.Alignment2i{X: frenyard.AlignStart, Y: frenyard.AlignStart}),
			},
		},
	})

	fullPanel := framework.NewUIFlexboxContainerPtr(framework.FlexboxContainer{
		DirVertical: true,
		Slots: []framework.FlexboxSlot{
			{
				Element: detail,
			},
			{
				Basis: design.SizeMarginAroundEverything,
			},
			{
				Element: framework.NewUILabelPtr(integration.NewTextTypeChunk(descriptionText, design.GlobalFont), design.ThemeText, 0, frenyard.Alignment2i{X: frenyard.AlignStart, Y: frenyard.AlignStart}),
				Shrink: 1,
			},
			{
				Grow: 1,
				Shrink: 1,
			},
		},
	})

	app.Teleport(design.LayoutDocument(design.Header{
		Title: repo.Name,
		Back: back,
	}, fullPanel, true))
}
