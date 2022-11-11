// Copyright 2022 Jeremy Edwards
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"log"
	"time"

	"github.com/jeremyje/coretemp-exporter/internal"
)

var (
	endpoint = flag.String("endpoint", ":8080", "Endpoint to serve metrics via HTTP.")
	interval = flag.Duration("interval", time.Second, "Polling interval for temperature information")
	logFile  = flag.String("log", "", "JSON logfile")
	console  = flag.Bool("console", true, "Indicates that records should be printed to console.")
)

func main() {
	flag.Parse()

	if err := internal.Run(&internal.Args{
		Endpoint: *endpoint,
		Interval: *interval,
		Log:      *logFile,
		Console:  *console,
	}); err != nil {
		log.Printf("%s", err)
	}
}