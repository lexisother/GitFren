/*
 * CCUpdaterUI/middle
 * Written starting in 2019 by 20kdc
 * This work is licensed under the terms of the MIT license.
 * For a copy, see <https://opensource.org/licenses/MIT>.
*/


package middle

import (
	"os/exec"
	"fmt"
)

// OpenURL opens a URL. It *MAY* be susceptible to bad URLs depending on OS.
func OpenURL(url string) error {
	cmd := exec.Command("xdg-open", url) // Linux
	if cmd.Run() == nil {
		return nil
	}
	cmd = exec.Command("cmd", "/c", "start", url) // Windows (thanks 2767mr)
	if cmd.Run() == nil {
		return nil
	}
	cmd = exec.Command("open", url) // Mac OS X
	if cmd.Run() == nil {
		return nil
	}
	return fmt.Errorf("all methods failed")
}
