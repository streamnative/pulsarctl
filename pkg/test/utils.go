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
	"math/rand"
	"os/exec"
	"time"

	"github.com/testcontainers/testcontainers-go"
)

// NewNetwork creates a network.
func NewNetwork(name string) (testcontainers.Network, error) {
	ctx := context.Background()
	dp, err := testcontainers.NewDockerProvider()
	if err != nil {
		return nil, err
	}

	net, err := dp.CreateNetwork(ctx, testcontainers.NetworkRequest{
		Name:           name,
		CheckDuplicate: true,
	})
	return net, err
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomSuffix() string {
	b := make([]rune, 6)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func ExecCmd(containerID string, cmd []string) (string, error) {
	args := []string{"exec", containerID}
	args = append(args, cmd...)
	out, err := exec.Command("docker", args...).Output()
	return string(out), err
}
