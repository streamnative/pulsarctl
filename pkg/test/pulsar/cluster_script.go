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

package pulsar

import (
	"fmt"

	"github.com/streamnative/pulsarctl/pkg/test"
)

// InitConf is a configuration for the initialize the pulsar cluster.
type InitConf struct {
	ClusterName        string
	ConfigurationStore string
	Zookeeper          string
	Broker             string
}

// InitCluster returns a container for executing init pulsar cluster.
func InitCluster(conf *InitConf, image, network string) *test.BaseContainer {
	pulsarInit := test.NewContainer(image)
	pulsarInit.WithNetwork([]string{network})
	pulsarInit.WaitForLog(fmt.Sprintf("Cluster metadata for '%s' setup correctly", conf.ClusterName))
	pulsarInit.WithCmd([]string{
		"bash", "-c",
		fmt.Sprintf("bin/pulsar initialize-cluster-metadata  -c %s -cs %s -uw  %s -zk %s",
			conf.ClusterName, conf.ConfigurationStore, conf.Broker, conf.Zookeeper),
	})
	return pulsarInit
}
