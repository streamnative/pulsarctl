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

package pkg

import "github.com/streamnative/pulsarctl/pkg/cmdutils"

// ResourceKind identifies a resource that supports custom runtime options injection.
type ResourceKind = cmdutils.RuntimeResource

// RuntimeOperation identifies an operation that supports custom runtime options injection.
type RuntimeOperation = cmdutils.RuntimeOperation

const (
	// ResourceFunction identifies Pulsar Functions commands.
	ResourceFunction ResourceKind = cmdutils.RuntimeResourceFunction
	// ResourceSink identifies Pulsar IO sink commands.
	ResourceSink ResourceKind = cmdutils.RuntimeResourceSink
	// ResourceSource identifies Pulsar IO source commands.
	ResourceSource ResourceKind = cmdutils.RuntimeResourceSource

	// RuntimeOperationCreate identifies create operations.
	RuntimeOperationCreate RuntimeOperation = cmdutils.RuntimeOperationCreate
	// RuntimeOperationUpdate identifies update operations.
	RuntimeOperationUpdate RuntimeOperation = cmdutils.RuntimeOperationUpdate
)

// CustomRuntimeOptionsContext describes the current custom runtime options value
// before it is sent to the Pulsar admin API.
type CustomRuntimeOptionsContext = cmdutils.CustomRuntimeOptionsContext

// CustomRuntimeOptionsInjector returns the final custom runtime options string
// for functions, sinks, and sources create/update commands.
type CustomRuntimeOptionsInjector = cmdutils.CustomRuntimeOptionsInjector

// Option configures a pulsarctl root command.
type Option func(*options)

type options struct {
	customRuntimeOptionsInjector CustomRuntimeOptionsInjector
}

func defaultOptions() options {
	return options{}
}

// WithCustomRuntimeOptionsInjector configures a hook that can replace the final
// customRuntimeOptions value for functions, sinks, and sources create/update commands.
func WithCustomRuntimeOptionsInjector(injector CustomRuntimeOptionsInjector) Option {
	return func(opts *options) {
		opts.customRuntimeOptionsInjector = injector
	}
}

func (o options) runtimeOptionsConfig() cmdutils.RuntimeOptionsConfig {
	return cmdutils.RuntimeOptionsConfig{
		CustomRuntimeOptionsInjector: o.customRuntimeOptionsInjector,
	}
}
