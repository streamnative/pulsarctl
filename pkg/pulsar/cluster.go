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
	"github.com/streamnative/pulsarctl/pkg/cli"
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

// Clusters is admin interface for clusters management
type Clusters interface {
	// List returns the list of clusters
	List() ([]string, error)

	// Get the configuration data for the specified cluster
	Get(string) (utils.ClusterData, error)

	// Create a new cluster
	Create(utils.ClusterData) error

	// Delete an existing cluster
	Delete(string) error

	// Update the configuration for a cluster
	Update(utils.ClusterData) error

	// UpdatePeerClusters updates peer cluster names.
	UpdatePeerClusters(string, []string) error

	// GetPeerClusters returns peer-cluster names
	GetPeerClusters(string) ([]string, error)

	// CreateFailureDomain creates a domain into cluster
	CreateFailureDomain(utils.FailureDomainData) error

	// GetFailureDomain returns the domain registered into a cluster
	GetFailureDomain(clusterName, domainName string) (utils.FailureDomainData, error)

	// ListFailureDomains returns all registered domains in cluster
	ListFailureDomains(string) (utils.FailureDomainMap, error)

	// DeleteFailureDomain deletes a domain in cluster
	DeleteFailureDomain(utils.FailureDomainData) error

	// UpdateFailureDomain updates a domain into cluster
	UpdateFailureDomain(utils.FailureDomainData) error
}

type clusters struct {
	client   *pulsarClient
	request  *cli.Client
	basePath string
}

// Clusters is used to access the cluster endpoints.
func (c *pulsarClient) Clusters() Clusters {
	return &clusters{
		client:   c,
		request:  c.Client,
		basePath: "/clusters",
	}
}

func (c *clusters) List() ([]string, error) {
	var clusters []string
	err := c.request.Get(c.client.endpoint(c.basePath), &clusters)
	return clusters, err
}

func (c *clusters) Get(name string) (utils.ClusterData, error) {
	cdata := utils.ClusterData{}
	endpoint := c.client.endpoint(c.basePath, name)
	err := c.request.Get(endpoint, &cdata)
	return cdata, err
}

func (c *clusters) Create(cdata utils.ClusterData) error {
	endpoint := c.client.endpoint(c.basePath, cdata.Name)
	return c.request.Put(endpoint, &cdata)
}

func (c *clusters) Delete(name string) error {
	endpoint := c.client.endpoint(c.basePath, name)
	return c.request.Delete(endpoint)
}

func (c *clusters) Update(cdata utils.ClusterData) error {
	endpoint := c.client.endpoint(c.basePath, cdata.Name)
	return c.request.Post(endpoint, &cdata)
}

func (c *clusters) GetPeerClusters(name string) ([]string, error) {
	var peerClusters []string
	endpoint := c.client.endpoint(c.basePath, name, "peers")
	err := c.request.Get(endpoint, &peerClusters)
	return peerClusters, err
}

func (c *clusters) UpdatePeerClusters(cluster string, peerClusters []string) error {
	endpoint := c.client.endpoint(c.basePath, cluster, "peers")
	return c.request.Post(endpoint, peerClusters)
}

func (c *clusters) CreateFailureDomain(data utils.FailureDomainData) error {
	endpoint := c.client.endpoint(c.basePath, data.ClusterName, "failureDomains", data.DomainName)
	return c.request.Post(endpoint, &data)
}

func (c *clusters) GetFailureDomain(clusterName string, domainName string) (utils.FailureDomainData, error) {
	var res utils.FailureDomainData
	endpoint := c.client.endpoint(c.basePath, clusterName, "failureDomains", domainName)
	err := c.request.Get(endpoint, &res)
	return res, err
}

func (c *clusters) ListFailureDomains(clusterName string) (utils.FailureDomainMap, error) {
	var domainData utils.FailureDomainMap
	endpoint := c.client.endpoint(c.basePath, clusterName, "failureDomains")
	err := c.request.Get(endpoint, &domainData)
	return domainData, err
}

func (c *clusters) DeleteFailureDomain(data utils.FailureDomainData) error {
	endpoint := c.client.endpoint(c.basePath, data.ClusterName, "failureDomains", data.DomainName)
	return c.request.Delete(endpoint)
}
func (c *clusters) UpdateFailureDomain(data utils.FailureDomainData) error {
	endpoint := c.client.endpoint(c.basePath, data.ClusterName, "failureDomains", data.DomainName)
	return c.request.Post(endpoint, &data)
}
