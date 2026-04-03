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

package file

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func URL(a []string) {
	if len(a) != 1 {
		Help(nil)
		return
	}

	link, err := url.Parse(a[0])

	if err != nil {
		erroPrinter.Print("%s\n", err.Error())
		return
	}

	request, err := http.NewRequest(http.MethodGet, link.String(), nil)

	if err != nil {
		erroPrinter.Print("%s\n", err.Error())
		return
	}

	request.Header.Set("Range", "bytes=0-0")

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		erroPrinter.Print("%s\n", err.Error())
		return
	}

	defer response.Body.Close()

	contentRange := response.Header.Get("Content-Range")
	value, ok := strings.CutPrefix(contentRange, "bytes 0-0/")

	if !ok {
		erroPrinter.Print("%s\n", errSizeUnknown.Error())
		return
	}

	size, err := strconv.ParseFloat(value, 64)

	if err != nil {
		erroPrinter.Print("%s\n", err.Error())
		return
	}

	switch {
	case size > 1_000*1_000*1_000:
		infoPrinter.Print("%.2f GB\n", size/(1_000*1_000*1_000))
	case size > 1_000*1_000:
		infoPrinter.Print("%.2f MB\n", size/(1_000*1_000))
	case size > 1_000:
		infoPrinter.Print("%.2f KB\n", size/1_000)
	default:
		infoPrinter.Print("%.2f B\n", size)
	}
}
