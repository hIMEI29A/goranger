// Copyright 2018 hIMEI
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
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/hIMEI29A/goranger/libgoranger"
)

var (
	typeFlag    = flag.String("t", "", "type of request: by city, by country, by ISP")
	requestFlag = flag.String("r", "", "request: country code, city name, ISP single ip or url")
	outputFlag  = flag.String("o", "", "file for info output")

	version     = "0.1.0"
	versionFlag = flag.Bool("v", false, "print current version")

	helpCmd = flag.Bool("h", false, "help message")
)

// ErrFatal is the basic errors handler
func errFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func wrongArg() {
	errString := "Wrong argument"
	err := errors.New(errString)
	errFatal(err)
}

// ToFile saves results to given file.
func toFile(filepath string, ranges []string) {
	dir := path.Dir(filepath)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		errString := "Given path does not exist"
		newerr := errors.New(errString)
		errFatal(newerr)
	}

	if _, err := os.Stat(filepath); os.IsExist(err) {
		errString := "File already exist, we'll not rewrite it "
		newerr := errors.New(errString)
		errFatal(newerr)
	}

	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0666)
	errFatal(err)
	defer file.Close()

	for i := range ranges {
		file.WriteString(ranges[i] + "\n")
		errFatal(err)
	}
}

func main() {
	// Cli options parsing
	flag.Parse()

	if *versionFlag {
		fmt.Println(version)
		os.Exit(1)
	}

	if len(os.Args) < 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *typeFlag == "" || *requestFlag == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *typeFlag == libgoranger.ReqType[1] && libgoranger.ValidateCountry(*requestFlag) != true {
		wrongArg()
	}

	g := libgoranger.NewGoranger()

	ranges, err := g.GetRange(*typeFlag, *requestFlag)
	errFatal(err)

	for i := range ranges {
		fmt.Println(ranges[i])
	}

	if *outputFlag != "" {
		filepath := *outputFlag
		toFile(filepath, ranges)
	}
}
