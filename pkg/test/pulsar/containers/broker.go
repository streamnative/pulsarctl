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

import "github.com/streamnative/pulsarctl/pkg/test"

const BrokerName = "broker"

func NewBrokerContainer(image, network string) *test.BaseContainer {
	broker := test.NewContainer(image)
	broker.WithNetwork([]string{network})
	broker.WithNetworkAliases(map[string][]string{network: {BrokerName}})
	broker.ExposedPorts([]string{"8080", "6650"})
	broker.WithCmd([]string{
		"bash", "-c",
		"bin/apply-config-from-env.py conf/broker.conf && bin/pulsar broker",
	})
	broker.WaitForPort("8080")
	return broker
}
