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
	"context"
	"fmt"
	"strconv"

	"github.com/streamnative/pulsarctl/pkg/test"
	"github.com/streamnative/pulsarctl/pkg/test/pulsar/containers"

	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
)

var (
	InvalidPort           = -1
	DefaultZKPort         = 2181
	DefaultBookiePort     = 3181
	DefaultBrokerPort     = 6650
	DefaultBrokerHTTPPort = 8080

	LatestImage = "apachepulsar/pulsar:latest"
)

type ClusterDef struct {
	clusterSpec      *ClusterSpec
	networkName      string
	network          *testcontainers.DockerNetwork
	zkContainer      *test.BaseContainer
	proxyContainer   *test.BaseContainer
	bookieContainers map[string]*test.BaseContainer
	brokerContainers map[string]*test.BaseContainer
}

type Cluster interface {
	// Start a pulsar cluster.
	Start(ctx context.Context) error

	// Stop a pulsar cluster.
	Stop(ctx context.Context) error

	// GetPlainTextServiceURL gets the pulsar service connect string.
	GetPlainTextServiceURL(ctx context.Context) (string, error)

	// GetHTTPServiceURL gets the pulsar HTTP service connect string.
	GetHTTPServiceURL(ctx context.Context) (string, error)

	// Close closes resources used for starting the cluster.
	Close(ctx context.Context)
}

// DefaultPulsarCluster creates a pulsar cluster using the default cluster spec.
func DefaultPulsarCluster() (test.Cluster, error) {
	return NewPulsarCluster(DefaultClusterSpec())
}

// NewPulsarCluster creates a pulsar cluster using the spec.
func NewPulsarCluster(spec *ClusterSpec) (test.Cluster, error) {
	c := &ClusterDef{clusterSpec: spec}
	c.networkName = spec.ClusterName + test.RandomSuffix()
	network, err := test.NewNetwork(c.networkName)
	if err != nil {
		return c, err
	}
	c.network = network

	c.zkContainer = containers.NewZookeeperContainer(LatestImage, c.networkName)
	c.bookieContainers = getBookieContainers(c.networkName, spec.NumBookies)
	brokers := getBrokerContainers(spec.ClusterName, c.networkName, spec.NumBrokers)
	broker := getABrokerNetAlias(brokers)
	c.brokerContainers = brokers
	c.proxyContainer = containers.NewProxyContainer(LatestImage, c.networkName).WithEnv(map[string]string{
		"webServicePort":      strconv.Itoa(spec.ProxyHTTPServicePort),
		"servicePort":         strconv.Itoa(spec.ProxyServicePort),
		"brokerServiceURL":    fmt.Sprintf("pulsar://%s:%d", broker, spec.BrokerServicePort),
		"brokerWebServiceURL": fmt.Sprintf("http://%s:%d", broker, spec.BrokerHTTPServicePort),
	})

	return c, nil
}

func getBookieContainers(network string, num int) map[string]*test.BaseContainer {
	bookies := make(map[string]*test.BaseContainer)
	for i := 0; i < num; i++ {
		name := fmt.Sprintf("%s-%d", containers.BookieName, i)
		bookies[name] = containers.NewBookieContainer(LatestImage, network).WithEnv(map[string]string{
			"zkServers": containers.ZookeeperName,
		}).WithNetworkAliases(map[string][]string{
			network: {name},
		})
	}
	return bookies
}

func getBrokerContainers(clusterName, network string, num int) map[string]*test.BaseContainer {
	brokers := make(map[string]*test.BaseContainer)
	for i := 0; i < num; i++ {
		name := fmt.Sprintf("%s-%d", containers.BrokerName, i)
		brokers[name] = containers.NewBrokerContainer(LatestImage, network).WithEnv(map[string]string{
			"zookeeperServers": containers.ZookeeperName,
			"clusterName":      clusterName,
		}).WithNetworkAliases(map[string][]string{
			network: {name},
		})
	}
	return brokers
}

func getABrokerNetAlias(brokers map[string]*test.BaseContainer) string {
	for k := range brokers {
		return k
	}
	return containers.BrokerName
}

func (c *ClusterDef) Start(ctx context.Context) error {
	err := c.zkContainer.Start(ctx)
	if err != nil {
		return errors.WithMessage(err, "encountered errors when starting the zookeeper")
	}
	fmt.Printf("Zookeeper %s:%s started.\n", containers.ZookeeperName, c.zkContainer.GetContainerID())

	init := InitCluster(&InitConf{
		ClusterName:        c.clusterSpec.ClusterName,
		ConfigurationStore: fmt.Sprintf("%s:%d", containers.ZookeeperName, DefaultZKPort),
		Zookeeper:          fmt.Sprintf("%s:%d", containers.ZookeeperName, DefaultZKPort),
		Broker: fmt.Sprintf("%s:%d",
			getABrokerNetAlias(c.brokerContainers), c.clusterSpec.BrokerHTTPServicePort),
	}, LatestImage, c.networkName)
	err = init.Start(ctx)
	if err != nil {
		return errors.WithMessage(err, "encountered errors when initializing the pulsar cluster")
	}
	fmt.Printf("Initialize pulsar cluster %s successfully.\n", c.clusterSpec.ClusterName)

	for k, v := range c.bookieContainers {
		err = v.Start(ctx)
		if err != nil {
			return errors.WithMessagef(err, "encountered errors when starting the bookie %s", k)
		}
		fmt.Printf("Bookie %s:%s started.\n", k, v.GetContainerID())
	}

	for k, v := range c.brokerContainers {
		err = v.Start(ctx)
		if err != nil {
			return errors.WithMessagef(err, "encountered errors when starting the bookie %s", k)
		}
		fmt.Printf("Broker %s:%s started.\n", k, v.GetContainerID())
	}

	err = c.proxyContainer.Start(ctx)
	if err != nil {
		return errors.WithMessage(err, "encountered errors when starting the proxy")
	}
	fmt.Printf("Proxy %s:%s started.\n", containers.ProxyName, c.proxyContainer.GetContainerID())

	return nil
}

func (c *ClusterDef) Stop(ctx context.Context) error {
	if c.zkContainer != nil {
		err := c.zkContainer.Stop(ctx)
		if err != nil {
			return err
		}
	}

	if c.bookieContainers != nil {
		for _, v := range c.bookieContainers {
			err := v.Stop(ctx)
			if err != nil {
				return err
			}
		}
	}

	if c.brokerContainers != nil {
		for _, v := range c.brokerContainers {
			err := v.Stop(ctx)
			if err != nil {
				return err
			}
		}
	}

	if c.proxyContainer != nil {
		err := c.proxyContainer.Stop(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *ClusterDef) GetPlainTextServiceURL(ctx context.Context) (string, error) {
	port, err := c.proxyContainer.MappedPort(ctx, strconv.Itoa(c.clusterSpec.BrokerHTTPServicePort))
	if err != nil {
		return "", err
	}
	return "pulsar://localhost:" + port.Port(), nil
}

func (c *ClusterDef) GetHTTPServiceURL(ctx context.Context) (string, error) {
	port, err := c.proxyContainer.MappedPort(ctx, strconv.Itoa(c.clusterSpec.ProxyHTTPServicePort))
	if err != nil {
		return "", err
	}
	return "http://localhost:" + port.Port(), nil
}

func (c *ClusterDef) Close(ctx context.Context) {
	if c.network != nil {
		_ = c.network.Remove(ctx)
	}
}
