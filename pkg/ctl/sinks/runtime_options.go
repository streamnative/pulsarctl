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

package sinks

import (
	"fmt"

	util "github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func applySinkRuntimeOptions(
	sinkData *util.SinkData,
	runtimeOptions cmdutils.RuntimeOptionsConfig,
	operation cmdutils.RuntimeOperation,
) error {
	if runtimeOptions.CustomRuntimeOptionsInjector == nil || sinkData.SinkConf == nil {
		return nil
	}

	value, err := runtimeOptions.ApplyCustomRuntimeOptions(
		cmdutils.RuntimeResourceSink,
		operation,
		sinkData.SinkConf.CustomRuntimeOptions,
	)
	if err != nil {
		return fmt.Errorf("inject sink %s custom runtime options: %w", operation, err)
	}
	sinkData.SinkConf.CustomRuntimeOptions = value
	return nil
}
