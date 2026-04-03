// This file is part of sizeof.
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

package file

func Help(a []string) {
	if len(a) > 0 {
		erroPrinter.Print("%s\n\n", a[0])
	}

	message := `Usage:

	sizeof file:<subcommand> [arguments]

Available subcommands and arguments:

	local [path]  # Find size of a local file.
	remote [path] # Find size of a remote file.
	help          # Display help message.

Examples:

	sizeof file:local /etc/pacman.d/mirrorlist
	sizeof file:remote https://fastly.mirror.pkgbuild.com/iso/2026.04.01/archlinux-2026.04.01-x86_64.iso`

	reguPrinter.Print("%s\n", message)
}
