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
	"context"
	"fmt"
	"strconv"

	"github.com/streamnative/pulsarctl/pkg/test"
	"github.com/streamnative/pulsarctl/pkg/test/bookkeeper/containers"

	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
)

var (
	LatestImage      = "apache/bookkeeper:latest"
	BookKeeper4_10_0 = "apache/bookkeeper:4.10.0"
)

type ClusterDef struct {
	test.Cluster
	clusterSpec *ClusterSpec
	networkName string
	//nolint:staticcheck
	network          testcontainers.Network
	zkContainer      *test.BaseContainer
	bookieContainers map[string]*test.BaseContainer
}

func DefaultCluster() (*ClusterDef, error) {
	return NewBookieCluster(DefaultClusterSpec())
}

func NewBookieCluster(spec *ClusterSpec) (*ClusterDef, error) {
	spec = GetClusterSpec(spec)
	c := &ClusterDef{clusterSpec: spec}
	c.networkName = spec.ClusterName + test.RandomSuffix()
	network, err := test.NewNetwork(c.networkName)
	if err != nil {
		return c, err
	}
	c.network = network

	c.zkContainer = containers.NewZookeeperContainer(spec.Image, c.networkName)
	c.bookieContainers = getBookieContainers(spec, c.networkName, containers.DefaultZookeeperServiceString())

	return c, nil
}

func getBookieContainers(c *ClusterSpec, networkName, zkServers string) map[string]*test.BaseContainer {
	bookies := make(map[string]*test.BaseContainer)
	for i := 0; i < c.NumBookies; i++ {
		name := fmt.Sprintf("%s-%d", containers.BookieName, i)
		bookie := containers.NewBookieContainer(c.Image, networkName)
		bookie.WithEnv(map[string]string{
			"BK_zkServers":         zkServers,
			"BK_httpServerEnabled": "true",
			"BK_httpServerPort":    strconv.Itoa(c.BookieHTTPServicePort),
			"BK_httpServerClass":   "org.apache.bookkeeper.http.vertx.VertxHttpServer",
			"BK_ledgerDirectories": "bk/ledgers",
			"BK_indexDirectories":  "bk/ledgers",
			"BK_journalDirectory":  "bk/journal",
		})
		bookie.WithEnv(c.BookieEnv)
		bookies[name] = bookie
	}
	return bookies
}

func (c *ClusterDef) Start(ctx context.Context) error {
	err := c.zkContainer.Start(ctx)
	if err != nil {
		return errors.WithMessage(err, "encountering errors when starting the zookeeper")
	}
	fmt.Printf("Zookeeper %s:%s started.\n",
		c.zkContainer.GetANetworkAlias(c.networkName), c.zkContainer.GetContainerID())
	zkName := c.zkContainer.GetANetworkAlias(c.networkName)
	init := InitBookieCluster(c.clusterSpec.Image, c.networkName, fmt.Sprintf("%s:2181", zkName))
	err = init.Start(ctx)
	if err != nil {
		return errors.WithMessage(err, "encountering errors when formatting metadata for bookkeeper")
	}
	fmt.Println("BookKeeper metadata initialized successfully.")

	for k, v := range c.bookieContainers {
		err = v.Start(ctx)
		if err != nil {
			return errors.WithMessagef(err, "encountering errors when starting the bookie %s\n", k)
		}
		fmt.Printf("Bookie %s:%s started.\n", k, v.GetContainerID())
	}

	return nil
}

func (c *ClusterDef) Stop(ctx context.Context) error {
	if c.bookieContainers != nil {
		for _, v := range c.bookieContainers {
			err := v.Stop(ctx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *ClusterDef) GetPlainTextServiceURL(ctx context.Context) (string, error) {
	return "", errors.New("unsupported operation")
}

func (c *ClusterDef) GetHTTPServiceURL(ctx context.Context) (string, error) {
	port, err := c.getABookie().MappedPort(ctx, strconv.Itoa(c.clusterSpec.BookieHTTPServicePort))
	if err != nil {
		return "", err
	}
	return "http://localhost:" + port.Port(), nil
}

func (c *ClusterDef) Close(ctx context.Context) {
	if c.network != nil {
		c.network.Remove(ctx)
	}
}

func (c *ClusterDef) getABookie() *test.BaseContainer {
	for _, v := range c.bookieContainers {
		return v
	}
	return nil
}

func (c *ClusterDef) GetAllBookieContainerID() []string {
	containerIDs := make([]string, 0)
	for _, v := range c.bookieContainers {
		containerIDs = append(containerIDs, v.GetContainerID())
	}
	return containerIDs
}
