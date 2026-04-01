// sizeof is a simple command line tool to view remote file size without
// downloading it.
// Copyright (C) 2026  Enindu Alahapperuma
//
// This program is free software: you can redistribute it and/or modify it under
// the terms of the GNU General Public License as published by the Free Software
// Foundation, either version 3 of the License, or (at your option) any later
// version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
// FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more
// details.
//
// You should have received a copy of the GNU General Public License along with
// this program.  If not, see <https://www.gnu.org/licenses/>.

// sizeof is a simple command line tool to view remote file size without
// downloading it.
//
// Usage:
//
//	sizeof [flags] [arguments]
//
// Available flags:
//
//	-u	Define file URL. (required)
//	-f	Define output format. ["b", "kb", "mb", "gb"] (default "mb")
//	-v	Display version message.
//	-h	Display help message.
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Define command line flags
	var (
		url     string
		format  string
		version bool
	)

	flag.StringVar(&url, "u", "", "Define file URL.")
	flag.StringVar(&format, "f", "mb", "Define output format. [\"b\", \"kb\", \"mb\", \"gb\"]")
	flag.BoolVar(&version, "v", false, "Display version message.")
	flag.Parse()

	if version {
		printVersion()
		return
	}

	if url == "" || format == "" {
		flag.PrintDefaults()
		return
	}

	// Send request and receive response
	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	request.Header.Set("Range", "bytes=0-0")

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	defer response.Body.Close()

	// Read response and print size
	var size float64

	contentRange := response.Header.Get("Content-Range")

	fmt.Sscanf(contentRange, "bytes 0-0/%f", &size)

	switch format {
	case "b":
		fmt.Fprintf(os.Stdout, "%.2f B\n", size)
	case "kb":
		fmt.Fprintf(os.Stdout, "%.2f KB\n", size/1_000)
	case "mb":
		fmt.Fprintf(os.Stdout, "%.2f MB\n", size/(1_000*1_000))
	case "gb":
		fmt.Fprintf(os.Stdout, "%.2f GB\n", size/(1_000*1_000*1_000))
	default:
		flag.PrintDefaults()
	}
}
