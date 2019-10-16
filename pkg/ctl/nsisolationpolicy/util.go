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

package nsisolationpolicy

import (
	"fmt"
	"strings"
)

func convert(value string) (map[string]string, error) {
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
