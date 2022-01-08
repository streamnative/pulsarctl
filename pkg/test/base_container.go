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

package test

import (
	"context"

	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// BaseContainer provide the basic operations for a container.
type BaseContainer struct {
	containerRequest testcontainers.GenericContainerRequest
	container        testcontainers.Container
}

// NewContainer creates a container using the image.
func NewContainer(image string) *BaseContainer {
	gcr := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: image,
		},
	}

	return &BaseContainer{
		containerRequest: gcr,
	}
}

// WithNetwork uses a existent network for the container.
func (bc *BaseContainer) WithNetwork(network []string) *BaseContainer {
	bc.containerRequest.Networks = append(bc.containerRequest.Networks, network...)
	return bc
}

// WithNetworkAliases creates some aliases for the container.
func (bc *BaseContainer) WithNetworkAliases(aliases map[string][]string) *BaseContainer {
	if bc.containerRequest.NetworkAliases == nil {
		bc.containerRequest.NetworkAliases = make(map[string][]string)
	}
	for k, v := range aliases {
		bc.containerRequest.NetworkAliases[k] = append(bc.containerRequest.NetworkAliases[k], v...)
	}
	return bc
}

// GetANetworkAlias returns a network alias of the container.
func (bc *BaseContainer) GetANetworkAlias(network string) string {
	return bc.containerRequest.NetworkAliases[network][0]
}

// WithCmd sets the containers start up commands.
func (bc *BaseContainer) WithCmd(cmd []string) *BaseContainer {
	bc.containerRequest.Cmd = append(bc.containerRequest.Cmd, cmd...)
	return bc
}

// WithEnv sets the environment variable to the container.
func (bc *BaseContainer) WithEnv(env map[string]string) *BaseContainer {
	if bc.containerRequest.Env == nil {
		bc.containerRequest.Env = make(map[string]string)
	}
	for k, v := range env {
		bc.containerRequest.Env[k] = v
	}
	return bc
}

// ExposedPorts exposes the ports from the container.
func (bc *BaseContainer) ExposedPorts(ports []string) *BaseContainer {
	bc.containerRequest.ExposedPorts = append(bc.containerRequest.ExposedPorts, ports...)
	return bc
}

// WaitForPort waits for the container ports exposed.
func (bc *BaseContainer) WaitForPort(port string) *BaseContainer {
	bc.containerRequest.WaitingFor = wait.ForListeningPort(nat.Port(port))
	return bc
}

// WaitForLog waits for the log string appear.
func (bc *BaseContainer) WaitForLog(log string) *BaseContainer {
	bc.containerRequest.WaitingFor = wait.ForLog(log)
	return bc
}

// WaitForHTTPPath waits for the path can be used. The Default access port is 80.
// TODO: support the specified path with a port.
func (bc *BaseContainer) WaitForHTTPPath(path string) *BaseContainer {
	bc.containerRequest.WaitingFor = wait.ForHTTP(path)
	return bc
}

// Start starts the container.
func (bc *BaseContainer) Start(ctx context.Context) error {
	c, err := testcontainers.GenericContainer(ctx, bc.containerRequest)
	if err != nil {
		return err
	}
	bc.container = c
	err = c.Start(ctx)
	return err
}

// ExecCmd executes a command in the container.
func (bc *BaseContainer) ExecCmd(ctx context.Context, cmd []string) (int, error) {
	return bc.container.Exec(ctx, cmd)
}

// Stop stops the container.
func (bc *BaseContainer) Stop(ctx context.Context) error {
	if bc.container != nil {
		return bc.container.Terminate(ctx)
	}
	return nil
}

// GetContainerID gets the container ID.
func (bc *BaseContainer) GetContainerID() string {
	return bc.container.GetContainerID()
}

// MappedPort gets the outside port.
func (bc *BaseContainer) MappedPort(ctx context.Context, port string) (nat.Port, error) {
	return bc.container.MappedPort(ctx, nat.Port(port))
}

func (bc *BaseContainer) ContainerIP(ctx context.Context) (string, error) {
	c, err := client.NewClientWithOpts()
	if err != nil {
		return "", err
	}

	inspect, err := c.ContainerInspect(ctx, bc.container.GetContainerID())
	if err != nil {
		return "", err
	}

	ip := inspect.NetworkSettings.IPAddress
	if ip != "" {
		return ip, nil
	}

	for _, settings := range inspect.NetworkSettings.Networks {
		return settings.IPAddress, nil
	}

	return "", errors.New("cannot get container ip address")
}
