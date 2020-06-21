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
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

type Config struct {
	Url string `json:"MappingUrl"`
}

func getMapDir() string {
	usr, err := user.Current()

	if err != nil {

		log.Fatalln("Failed to get mat")
	}

	mapDir := usr.HomeDir + "/.mapcli"

	if _, err := os.Stat(mapDir); os.IsNotExist(err) {
		os.Mkdir(mapDir, os.ModePerm)
	}

	return mapDir
}

func getConfig() Config {
	confDir := getMapDir() + "/config.json"
	var conf Config

	if _, err := os.Stat(confDir); err != nil {
		conf = Config{
			Url: "BLARG",
		}

		file, err := json.MarshalIndent(conf, "", "")
		if err != nil {
			log.Fatalln(err)
		}

		err = ioutil.WriteFile(confDir, file, 0664)

		if err != nil {
			log.Fatalln(err)
		}
	}

	file, err := ioutil.ReadFile(confDir)

	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal([]byte(file), &conf)

	if err != nil {
		log.Fatalln(err)
	}

	if conf.Url == "BLARG" {
		log.Fatalln("Please enter the MCP MAPPINGS URL in the config file found in ~/.mapcli")
	}

	return conf
}
