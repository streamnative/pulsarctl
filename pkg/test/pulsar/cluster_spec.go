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

// ClusterSpec is to build a pulsar cluster.
type ClusterSpec struct {
	Image                 string
	ClusterName           string
	BookiePort            int
	NumBookies            int
	BrokerServicePort     int
	BrokerHTTPServicePort int
	NumBrokers            int
	ProxyServicePort      int
	ProxyHTTPServicePort  int
	NumProxies            int
}

// DefaultClusterSpec returns default configuration of a cluster.
func DefaultClusterSpec() *ClusterSpec {
	return &ClusterSpec{
		Image:                 LatestImage,
		ClusterName:           "default-cluster",
		BookiePort:            DefaultBookiePort,
		NumBookies:            2,
		BrokerServicePort:     DefaultBrokerPort,
		BrokerHTTPServicePort: DefaultBrokerHTTPPort,
		NumBrokers:            2,
		ProxyServicePort:      DefaultBrokerPort,
		ProxyHTTPServicePort:  DefaultBrokerHTTPPort,
		NumProxies:            1,
	}
}
