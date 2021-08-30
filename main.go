package main

import (
  "github.com/20kdc/CCUpdaterUI/design"
  "github.com/20kdc/CCUpdaterUI/frenyard"
  "github.com/20kdc/CCUpdaterUI/frenyard/framework"
	"git.nova-vps.ml/lexisother/GitFren/src"
  "git.nova-vps.ml/lexisother/GitFren/middle"
)

func main() {
	frenyard.TargetFrameTime = 0.016
	slideContainer := framework.NewUISlideTransitionContainerPtr(nil)
	slideContainer.FyEResize(design.SizeWindowInit)
	wnd, err := framework.CreateBoundWindow("GitFren", true, design.ThemeBackground, slideContainer)
	if err != nil {
		panic(err)
	}
	design.Setup(frenyard.InferScale(wnd))
	wnd.SetSize(design.SizeWindow)
	// Ok, now get it ready.
	app := (&src.UpApplication{
		Config: middle.ReadUpdaterConfig(),
		MainContainer: slideContainer,
		Window: wnd,
		UpQueued: make(chan func(), 16),
		TeleportSettings: framework.SlideTransition{},
	})
	app.ShowPreface()
	// Started!
	frenyard.GlobalBackend.Run(func(frameTime float64) {
		select {
			case fn := <- app.UpQueued:

				fn()
			default:
		}
	})
}
