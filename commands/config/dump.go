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

package config

import (
	"fmt"
	"os"

	"github.com/zmarcantel/arduino-cli/commands"
	"github.com/zmarcantel/arduino-cli/common/formatter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func initDumpCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "dump",
		Short:   "Prints the current configuration",
		Long:    "Prints the current configuration.",
		Example: "  " + commands.AppName + " config dump",
		Args:    cobra.NoArgs,
		Run:     runDumpCommand,
	}
}

func runDumpCommand(cmd *cobra.Command, args []string) {
	logrus.Info("Executing `arduino config dump`")

	data, err := commands.Config.SerializeToYAML()
	if err != nil {
		formatter.PrintError(err, "Error creating configuration")
		os.Exit(commands.ErrGeneric)
	}

	fmt.Println(string(data))
}
