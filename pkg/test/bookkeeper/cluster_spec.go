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

package bookkeeper

import (
	"strings"

	"github.com/streamnative/pulsarctl/pkg/test/bookkeeper/containers"
)

type ClusterSpec struct {
	Image                 string
	ClusterName           string
	NumBookies            int
	BookieServicePort     int
	BookieHTTPServicePort int
	ZookeeperServicePort  int
	BookieEnv             map[string]string
}

func DefaultClusterSpec() *ClusterSpec {
	return &ClusterSpec{
		Image:                 BookKeeper,
		ClusterName:           "default-bookie",
		NumBookies:            1,
		BookieServicePort:     containers.DefaultBookieServicePort,
		BookieHTTPServicePort: containers.DefaultBookieHTTPServicePort,
		ZookeeperServicePort:  containers.DefaultZookeeperServicePort,
	}
}

func GetClusterSpec(spec *ClusterSpec) *ClusterSpec {
	newSpec := DefaultClusterSpec()

	if spec.NumBookies > 0 {
		newSpec.NumBookies = spec.NumBookies
	}

	if strings.TrimSpace(spec.ClusterName) != "" {
		newSpec.ClusterName = spec.ClusterName
	}

	if strings.TrimSpace(spec.Image) != "" {
		newSpec.Image = spec.Image
	}

	if spec.ZookeeperServicePort > 0 {
		newSpec.ZookeeperServicePort = spec.ZookeeperServicePort
	}

	if spec.BookieServicePort > 0 {
		newSpec.BookieServicePort = spec.BookieServicePort
	}

	if spec.BookieHTTPServicePort > 0 {
		newSpec.BookieHTTPServicePort = spec.BookieHTTPServicePort
	}

	if spec.BookieEnv != nil {
		newSpec.BookieEnv = spec.BookieEnv
	}

	return newSpec
}
