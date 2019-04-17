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

package core

import (
	"github.com/zmarcantel/arduino-cli/commands"
	"github.com/spf13/cobra"
)

// InitCommand prepares the command.
func InitCommand() *cobra.Command {
	coreCommand := &cobra.Command{
		Use:     "core",
		Short:   "Arduino Core operations.",
		Long:    "Arduino Core operations.",
		Example: "  " + commands.AppName + " core update-index",
	}
	coreCommand.AddCommand(initDownloadCommand())
	coreCommand.AddCommand(initInstallCommand())
	coreCommand.AddCommand(initListCommand())
	coreCommand.AddCommand(initUpdateIndexCommand())
	coreCommand.AddCommand(initUpgradeCommand())
	coreCommand.AddCommand(initUninstallCommand())
	coreCommand.AddCommand(initSearchCommand())
	return coreCommand
}
