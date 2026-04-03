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

import "os"

func Local(a []string) {
	if len(a) != 1 {
		Help([]string{errArgumentsInvalid.Error()})
		return
	}

	path := a[0]
	file, err := os.Open(path)

	if err != nil {
		erroPrinter.Print("%s\n", err.Error())
		return
	}

	defer file.Close()

	info, err := file.Stat()

	if err != nil {
		erroPrinter.Print("%s\n", err.Error())
		return
	}

	size := float64(info.Size())

	printSize(size)
}
