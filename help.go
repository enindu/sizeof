//  This file is part of sizeof.
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

package main

func help() {
	message := `Usage:
	
	sizeof <command>:<subcommand> [arguments]
	sizeof [flags]
	
Available commands:

	file
	
Available flags:

	-v, --version # Display version message.
	-h, --help    # Display help message.
	
Use "sizeof <command>:help" to see more information about commands.`

	reguPrinter.Print("%s\n", message)
}
