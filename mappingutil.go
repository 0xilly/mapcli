// Copyright (c) 2020, Anthony Anderson
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
//  * Redistributions of source code must retain the above copyright notice,
//    this list of conditions and the following disclaimer.
//  * Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE AUTHOR AND CONTRIBUTORS ``AS IS'' AND ANY
// EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE AUTHOR OR CONTRIBUTORS BE LIABLE FOR ANY
// DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
// LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
// OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH
// DAMAGE.

package mapcli

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func cleanAndGet() {
	os.Remove(getMapDir() + "/mappings.zip")
	os.Remove(getMapDir() + "/params.csv")
	os.Remove(getMapDir() + "/fields.csv")
	os.Remove(getMapDir() + "/methods.csv")
	getMappings()
}

func getMappings() {
	conf := getConfig()
	mappings := getMapDir() + "/mappings.zip"

	if _, err := os.Stat(mappings); err != nil {
		downloadFile(mappings, conf.Url)
		unzip(mappings)
	}

}

func getMethod(method string) {
	file, err := os.Open(getMapDir() + "/methods.csv")
	m := make(map[string]string)

	defer file.Close()

	if err != nil {
		log.Fatalln(err)
	}
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		spl := strings.Split(scan.Text(), ",")
		m[spl[0]] = spl[1]
	}

	if _, ok := m[method]; !ok {
		fmt.Println("No mapping found")
		os.Exit(1)
	}

	fmt.Printf("%s\n", m[method])
}

func getFeild(field string) {
	file, err := os.Open(getMapDir() + "/fields.csv")
	m := make(map[string]string)

	defer file.Close()

	if err != nil {
		log.Fatalln(err)
	}
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		spl := strings.Split(scan.Text(), ",")
		m[spl[0]] = spl[1]
	}

	if _, ok := m[field]; !ok {
		fmt.Println("No mapping found")
		os.Exit(1)
	}

	fmt.Printf("%s\n", m[field])
}

func GetMapping(srg string) {
	getMappings()
	if strings.Contains(srg, "func_") {
		getMethod(srg)
	} else if strings.Contains(srg, "field_") {
		getFeild(srg)
	} else {
		fmt.Println("Unsupported Mapping")
	}
}

func ForceUpdate() {
	fmt.Println("Cleaning ~/.mapcli directory")
	cleanAndGet()
}
