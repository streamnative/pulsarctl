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

package containers

import (
	"fmt"

	"github.com/streamnative/pulsarctl/pkg/test"
)

const (
	ZookeeperName               = "zookeeper"
	DefaultZookeeperServicePort = 2181
)

func NewZookeeperContainer(image, network string) *test.BaseContainer {
	zookeeper := test.NewContainer(image)
	zookeeper.WithNetwork([]string{network})
	zookeeper.WithNetworkAliases(map[string][]string{network: {ZookeeperName}})
	zookeeper.WithCmd([]string{"zookeeper"})
	zookeeper.WaitForLog("Started AdminServer on address 0.0.0.0, port 8080 and command URL /commands")
	return zookeeper
}

func DefaultZookeeperServiceString() string {
	return fmt.Sprintf("%s:%d", ZookeeperName, DefaultZookeeperServicePort)
}
