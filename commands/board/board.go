/*
 * This file is part of arduino-cli.
 *
 * Copyright 2018 ARDUINO SA (http://www.arduino.cc/)
 *
 * This software is released under the GNU General Public License version 3,
 * which covers the main part of arduino-cli.
 * The terms of this license can be found at:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 * You can be released from the requirements of the above licenses by purchasing
 * a commercial license. Buying such a license is mandatory if you want to modify or
 * otherwise use the software for commercial activities involving the Arduino
 * software without disclosing the source code of your own applications. To purchase
 * a commercial license, send an email to license@arduino.cc.
 */

package board

import (
	"github.com/zmarcantel/arduino-cli/commands"
	"github.com/spf13/cobra"
)

// InitCommand prepares the command.
func InitCommand() *cobra.Command {
	boardCommand := &cobra.Command{
		Use:   "board",
		Short: "Arduino board commands.",
		Long:  "Arduino board commands.",
		Example: "  # Lists all connected boards.\n" +
			"  " + commands.AppName + " board list\n\n" +
			"  # Attaches a sketch to a board.\n" +
			"  " + commands.AppName + " board attach serial:///dev/tty/ACM0 mySketch",
	}
	boardCommand.AddCommand(initAttachCommand())
	boardCommand.AddCommand(initDetailsCommand())
	boardCommand.AddCommand(initListCommand())
	boardCommand.AddCommand(initListAllCommand())
	return boardCommand
}
