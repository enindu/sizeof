// sizeof is a simple command line tool to view size of local or remote files.
// Copyright (C) 2026 Enindu Alahapperuma
//
// sizeof is free software: you can redistribute it and/or modify it under the
// terms of the GNU General Public License as published by the Free Software
// Foundation, either version 3 of the License, or (at your option) any later
// version.
//
// sizeof is distributed in the hope that it will be useful, but WITHOUT ANY
// WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
// A PARTICULAR PURPOSE. See the GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along with
// sizeof. If not, see <https://www.gnu.org/licenses/>.

// sizeof is a simple command line tool to view size of local or remote files.
//
// Usage:
//
//	sizeof <command>:<subcommand> [arguments]
//	sizeof [flags]
//
// Available commands:
//
//	file
//
// Available flags:
//
//	-v, --version # Display version message.
//	-h, --help    # Display help message.
//
// Use "sizeof <command>:help" to see more information about commands.
package main

import (
	"errors"
	"os"

	"github.com/enindu/palette"
	"github.com/enindu/sizeof/commands/file"
)

var (
	errInstructionNotFound error = errors.New("Instruction is not found, use \"-h\" or \"--help\" to see help message")
	errCommandInvalid      error = errors.New("Command is invalid, use \"-h\" or \"--help\" to see help message")
)

var (
	reguPrinter *palette.Printer = palette.NewPrinterRegu()
	erroPrinter *palette.Printer = palette.NewPrinterErro()
)

func main() {
	dispatchers := map[string]func([]string){
		"file:remote": file.Remote,
		"file:help":   file.Help,
	}

	inputs := os.Args

	if len(inputs) < 2 {
		erroPrinter.Print("%s\n", errInstructionNotFound.Error())
		return
	}

	instruction := inputs[1]

	switch instruction {
	case "-v", "--version":
		version()
		return
	case "-h", "--help":
		help()
		return
	}

	execute, ok := dispatchers[instruction]

	if !ok {
		erroPrinter.Print("%s\n", errCommandInvalid)
		return
	}

	arguments := inputs[2:]

	execute(arguments)
}
