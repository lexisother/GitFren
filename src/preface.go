package src

import "git.nova-vps.ml/lexisother/GitFren/middle"

func (app *UpApplication) ShowPreface() {
	app.ShowWaiter("Starting...", func(progress func(string)) {
		progress("Fetching remote repositories...")
		middle.GetUserNames()
	}, func () {
		app.CachedPrimaryView = nil
		app.ShowPrimaryView()
	})
}
