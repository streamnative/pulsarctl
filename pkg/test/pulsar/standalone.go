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
	"strconv"

	"github.com/streamnative/pulsarctl/pkg/test"
)

type Standalone struct {
	*test.BaseContainer
	spec *ClusterSpec
}

func NewStandalone(spec *ClusterSpec) *Standalone {
	s := test.NewContainer(spec.Image)
	s.ExposedPorts([]string{strconv.Itoa(spec.BrokerHTTPServicePort), strconv.Itoa(spec.BrokerServicePort)})
	s.WithCmd([]string{
		"bin/pulsar", "standalone",
	})
	s.WaitForPort(strconv.Itoa(spec.BrokerHTTPServicePort))
	return &Standalone{s, spec}
}

func DefaultStandalone() *Standalone {
	return NewStandalone(DefaultClusterSpec())
}

func (s *Standalone) GetHTTPServiceURL(ctx context.Context) (string, error) {
	port, err := s.MappedPort(ctx, strconv.Itoa(s.spec.BrokerHTTPServicePort))
	if err != nil {
		return "", err
	}
	return "http://localhost:" + port.Port(), nil
}
