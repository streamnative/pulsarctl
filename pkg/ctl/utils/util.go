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

package utils

import (
	`fmt`
	"github.com/streamnative/pulsarctl/pkg/pulsar"
	`os`
	"strings"
)

const (
	HTTP    = "http"
	FILE    = "file"
	BUILTIN = "builtin"

	PublicTenant     = "public"
	DefaultNamespace = "default"
)

func IsPackageUrlSupported(functionPkgUrl string) bool {
	return functionPkgUrl != "" && strings.HasPrefix(functionPkgUrl, HTTP) ||
		strings.HasPrefix(functionPkgUrl, FILE)
}

func InferMissingFunctionName(funcConf *pulsar.FunctionConfig) {
	className := funcConf.ClassName
	domains := strings.Split(className, "\\.")

	if len(domains) == 0 {
		funcConf.Name = funcConf.ClassName
	} else {
		funcConf.Name = domains[len(domains)-1]
	}
}

func InferMissingTenant(funcConf *pulsar.FunctionConfig) {
	funcConf.Tenant = PublicTenant
}

func InferMissingNamespace(funcConf *pulsar.FunctionConfig) {
	funcConf.Namespace = DefaultNamespace
}

func InferMissingSourceArguments(sourceConf *pulsar.SourceConfig) {
	if sourceConf.Tenant == "" {
		sourceConf.Tenant = PublicTenant
	}

	if sourceConf.Namespace == "" {
		sourceConf.Namespace = DefaultNamespace
	}

	if sourceConf.Parallelism == 0 {
		sourceConf.Parallelism = 1
	}
}

func InferMissingSinkeArguments(sinkConf *pulsar.SinkConfig) {
	if sinkConf.Tenant == "" {
		sinkConf.Tenant = PublicTenant
	}

	if sinkConf.Namespace == "" {
		sinkConf.Namespace = DefaultNamespace
	}

	if sinkConf.Parallelism == 0 {
		sinkConf.Parallelism = 1
	}
}

func IsFileExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	fmt.Println("exists", info.Name(), info.Size(), info.ModTime())
	return true
}
