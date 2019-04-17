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

package sketch

import (
	"os"

	"github.com/zmarcantel/arduino-cli/commands"
	"github.com/zmarcantel/arduino-cli/common/formatter"
	"github.com/spf13/cobra"
)

func initNewCommand() *cobra.Command {
	newCommand := &cobra.Command{
		Use:     "new",
		Short:   "Create a new Sketch",
		Long:    "Create a new Sketch",
		Example: "  " + commands.AppName + " sketch new MultiBlinker",
		Args:    cobra.ExactArgs(1),
		Run:     runNewCommand,
	}
	return newCommand
}

var emptySketch = []byte(`
void setup() {
}

void loop() {
}
`)

func runNewCommand(cmd *cobra.Command, args []string) {
	sketchDir := commands.Config.SketchbookDir.Join(args[0])
	if err := sketchDir.MkdirAll(); err != nil {
		formatter.PrintError(err, "Could not create sketch directory.")
		os.Exit(commands.ErrGeneric)
	}

	sketchFile := sketchDir.Join(args[0] + ".ino")
	if err := sketchFile.WriteFile(emptySketch); err != nil {
		formatter.PrintError(err, "Error creating sketch.")
		os.Exit(commands.ErrGeneric)
	}

	formatter.Print("Sketch created in: " + sketchDir.String())
}
