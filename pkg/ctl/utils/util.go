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
	"fmt"
	"github.com/pkg/errors"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"os"
	"strconv"
	"strings"
	"time"
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

func ValidateSizeString(s string) (int64, error) {
	end := s[len(s)-1:]
	value := s[:len(s)-1]
	switch end {
	case "k":
		fallthrough
	case "K":
		v, err := strconv.ParseInt(value, 10, 64)
		return v * 1024, err
	case "m":
		fallthrough
	case "M":
		v, err := strconv.ParseInt(value, 10, 64)
		return v * 1024 * 1024, err
	case "g":
		fallthrough
	case "G":
		v, err := strconv.ParseInt(value, 10, 64)
		return v * 1024 * 1024 * 1024, err
	case "t":
		fallthrough
	case "T":
		v, err := strconv.ParseInt(value, 10, 64)
		return v * 1024 * 1024 * 1024 * 1024, err
	default:
		return strconv.ParseInt(s, 10, 64)
	}
}

func ParseRelativeTimeInSeconds(relativeTime string) (time.Duration, error) {
	if relativeTime == "" {
		return -1, errors.New("Time can not be empty.")
	}

	unitTime := relativeTime[len(relativeTime)-1:]
	t := relativeTime[:len(relativeTime)-1]
	timeValue, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		return -1, errors.Errorf("Invalid time '%s'", t)
	}

	switch strings.ToLower(unitTime) {
	case "s":
		return time.Duration(timeValue) * time.Second, nil
	case "m":
		return time.Duration(timeValue) * time.Minute, nil
	case "h":
		return time.Duration(timeValue) * time.Hour, nil
	case "d":
		return time.Duration(timeValue) * time.Hour * 24, nil
	case "w":
		return time.Duration(timeValue) * time.Hour * 24 * 7, nil
	case "y":
		return time.Duration(timeValue) * time.Hour * 24 * 7 * 365, nil
	default:
		return -1, errors.Errorf("Invalid time unit '%s'", unitTime)
	}
}
