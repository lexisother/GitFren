/*
 * CCUpdaterUI/middle
 * Written starting in 2019 by 20kdc
 * This work is licensed under the terms of the MIT license.
 * For a copy, see <https://opensource.org/licenses/MIT>.
*/

package middle

import (
	"net/http"
	"io/ioutil"
)

// autoupdate alert stuff ; update this for new version names!!!
// version schedule:
// past: "lea\n" "emilie\n"
// present: "" (disable)
// future: ???
// note that the "\n" is important if this is a real ID
const localUIVersionID string = ""
var remoteUIVersionID string = localUIVersionID
var hasAlreadyCheckedRemoteUIVersionID bool = false

func updateAlertHook() {
	if localUIVersionID == "" {
		// version checks are disabled
		return
	}
	if !hasAlreadyCheckedRemoteUIVersionID {
		hasAlreadyCheckedRemoteUIVersionID = true
	} else {
		return
	}
	res, err := http.Get("https://20kdc.duckdns.org/ccmodloader/version")
	if err == nil {
		data, err := ioutil.ReadAll(res.Body)
		if err == nil {
			remoteUIVersionID = string(data)
		}
	}
}

// WarningID represents a kind of warning action.
type WarningID int

const (
	// NullActionWarningID cannot be automatically fixed.
	NullActionWarningID WarningID = iota
	// InstallOrUpdatePackageWarningID warnings can be solved by installing/updating the package Parameter.
	InstallOrUpdatePackageWarningID
	// URLAndCloseWarningID warnings can be solved manually by the user given navigation to a URL. The application closes.
	URLAndCloseWarningID
)

// Warning represents a warning to show the user on the primary view.
type Warning struct {
	Text string
	Action WarningID
	Parameter string
}
