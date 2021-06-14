// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package main

import (
	"fmt"
	"os"

	"github.com/streamnative/pulsarctl/pkg"
	"github.com/streamnative/pulsarctl/pkg/plugin"
)

func main() {
	rootCmd := pkg.NewPulsarctlCmd()
	handler := plugin.NewDefaultPluginHandler(plugin.ValidPluginFilenamePrefixes)

	_, _, err := rootCmd.Find(os.Args)
	if err != nil {
		err = plugin.HandlePluginCommand(handler, getArgs())
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	}

	if err = rootCmd.Execute(); err != nil {
		fmt.Println(err) // outputs cobra errors
		os.Exit(-1)
	}
}

func getArgs() []string {
	args := os.Args
	if len(args) > 1 {
		return args[1:]
	}
	return []string{}
}
