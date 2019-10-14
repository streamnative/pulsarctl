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

// Clusters  is used to access the cluster endpoints.

type Clusters interface {
	List() ([]string, error)
	Get(string) (ClusterData, error)
	Create(ClusterData) error
	Delete(string) error
	Update(ClusterData) error
	UpdatePeerClusters(string, []string) error
	GetPeerClusters(string) ([]string, error)
	CreateFailureDomain(FailureDomainData) error
	GetFailureDomain(clusterName, domainName string) (FailureDomainData, error)
	ListFailureDomains(string) (FailureDomainMap, error)
	DeleteFailureDomain(FailureDomainData) error
	UpdateFailureDomain(FailureDomainData) error
}

type clusters struct {
	client   *pulsarClient
	request  *client
	basePath string
}

func (c *pulsarClient) Clusters() Clusters {
	return &clusters{
		client:   c,
		request:  c.client,
		basePath: "/clusters",
	}
}

func (c *clusters) List() ([]string, error) {
	var clusters []string
	err := c.request.get(c.client.endpoint(c.basePath), &clusters)
	return clusters, err
}

func (c *clusters) Get(name string) (ClusterData, error) {
	cdata := ClusterData{}
	endpoint := c.client.endpoint(c.basePath, name)
	err := c.request.get(endpoint, &cdata)
	return cdata, err
}

func (c *clusters) Create(cdata ClusterData) error {
	endpoint := c.client.endpoint(c.basePath, cdata.Name)
	return c.request.put(endpoint, &cdata)
}

func (c *clusters) Delete(name string) error {
	endpoint := c.client.endpoint(c.basePath, name)
	return c.request.delete(endpoint)
}

func (c *clusters) Update(cdata ClusterData) error {
	endpoint := c.client.endpoint(c.basePath, cdata.Name)
	return c.request.post(endpoint, &cdata)
}

func (c *clusters) GetPeerClusters(name string) ([]string, error) {
	var peerClusters []string
	endpoint := c.client.endpoint(c.basePath, name, "peers")
	err := c.request.get(endpoint, &peerClusters)
	return peerClusters, err
}

func (c *clusters) UpdatePeerClusters(cluster string, peerClusters []string) error {
	endpoint := c.client.endpoint(c.basePath, cluster, "peers")
	return c.request.post(endpoint, peerClusters)
}

func (c *clusters) CreateFailureDomain(data FailureDomainData) error {
	endpoint := c.client.endpoint(c.basePath, data.ClusterName, "failureDomains", data.DomainName)
	return c.request.post(endpoint, &data)
}

func (c *clusters) GetFailureDomain(clusterName string, domainName string) (FailureDomainData, error) {
	var res FailureDomainData
	endpoint := c.client.endpoint(c.basePath, clusterName, "failureDomains", domainName)
	err := c.request.get(endpoint, &res)
	return res, err
}

func (c *clusters) ListFailureDomains(clusterName string) (FailureDomainMap, error) {
	var domainData FailureDomainMap
	endpoint := c.client.endpoint(c.basePath, clusterName, "failureDomains")
	err := c.request.get(endpoint, &domainData)
	return domainData, err
}

func (c *clusters) DeleteFailureDomain(data FailureDomainData) error {
	endpoint := c.client.endpoint(c.basePath, data.ClusterName, "failureDomains", data.DomainName)
	return c.request.delete(endpoint)
}
func (c *clusters) UpdateFailureDomain(data FailureDomainData) error {
	endpoint := c.client.endpoint(c.basePath, data.ClusterName, "failureDomains", data.DomainName)
	return c.request.post(endpoint, &data)
}
