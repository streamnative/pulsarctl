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
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const (
	HTTP    = "http"
	FILE    = "file"
	BUILTIN = "builtin"

	FUNCTION = "function"
	SINK     = "sink"
	SOURCE   = "source"

	PublicTenant     = "public"
	DefaultNamespace = "default"
)

func IsPackageURLSupported(functionPkgURL string) bool {
	return functionPkgURL != "" && (strings.HasPrefix(functionPkgURL, HTTP) ||
		strings.HasPrefix(functionPkgURL, FILE) ||
		strings.HasPrefix(functionPkgURL, FUNCTION) ||
		strings.HasPrefix(functionPkgURL, SINK) ||
		strings.HasPrefix(functionPkgURL, SOURCE))
}

func InferMissingFunctionName(funcConf *utils.FunctionConfig) {
	className := funcConf.ClassName
	domains := strings.Split(className, "\\.")

	if len(domains) == 0 {
		funcConf.Name = funcConf.ClassName
	} else {
		funcConf.Name = domains[len(domains)-1]
	}
}

func InferMissingTenant(funcConf *utils.FunctionConfig) {
	funcConf.Tenant = PublicTenant
}

func InferMissingNamespace(funcConf *utils.FunctionConfig) {
	funcConf.Namespace = DefaultNamespace
}

func InferMissingSourceArguments(sourceConf *utils.SourceConfig) {
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

func InferMissingSinkeArguments(sinkConf *utils.SinkConfig) {
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
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
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
		return -1, errors.New("time can not be empty")
	}

	if relativeTime == "-1" {
		return -1, nil
	}

	unitTime := relativeTime[len(relativeTime)-1:]
	t := relativeTime[:len(relativeTime)-1]
	timeValue, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		return -1, errors.Errorf("invalid time '%s'", t)
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
		return -1, errors.Errorf("invalid time unit '%s'", unitTime)
	}
}

func Convert(value string) (map[string]string, error) {
	err := false
	nvPairs := strings.Split(value, ",")

	tmpMap := make(map[string]string)

	for _, nvPair := range nvPairs {
		err = true
		if len(nvPair) != 0 {
			nv := strings.Split(nvPair, "=")
			if len(nv) == 2 {
				nv[0] = strings.TrimSpace(nv[0])
				nv[1] = strings.TrimSpace(nv[1])
				if len(nv[0]) != 0 && len(nv[1]) != 0 && !strings.HasPrefix(nv[0], "/") {
					tmpMap[nv[0]] = nv[1]
					err = false
				}
			}
		}
		if err {
			break
		}
	}

	if err {
		return nil, fmt.Errorf("unable to parse bad name=value parameter list: %v", value)
	}

	return tmpMap, nil
}

func NumProvidedStrings(sPointers ...*string) int {
	out := 0
	for _, sp := range sPointers {
		if sp != nil {
			out++
		}
	}
	return out
}
