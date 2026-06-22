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

package cmdutils

// RuntimeResource identifies a runtime-capable command resource.
type RuntimeResource string

// RuntimeOperation identifies a command operation that can apply runtime options.
type RuntimeOperation string

const (
	// RuntimeResourceFunction identifies Pulsar Functions commands.
	RuntimeResourceFunction RuntimeResource = "function"
	// RuntimeResourceSink identifies Pulsar IO sink commands.
	RuntimeResourceSink RuntimeResource = "sink"
	// RuntimeResourceSource identifies Pulsar IO source commands.
	RuntimeResourceSource RuntimeResource = "source"

	// RuntimeOperationCreate identifies create operations.
	RuntimeOperationCreate RuntimeOperation = "create"
	// RuntimeOperationUpdate identifies update operations.
	RuntimeOperationUpdate RuntimeOperation = "update"
)

// CustomRuntimeOptionsContext describes the current custom runtime options value
// before it is sent to the Pulsar admin API.
type CustomRuntimeOptionsContext struct {
	Resource  RuntimeResource
	Operation RuntimeOperation
	Current   string
}

// CustomRuntimeOptionsInjector returns the final custom runtime options string
// for a runtime-capable resource command.
type CustomRuntimeOptionsInjector func(CustomRuntimeOptionsContext) (string, error)

// RuntimeOptionsConfig contains immutable runtime option hooks for a command tree.
type RuntimeOptionsConfig struct {
	CustomRuntimeOptionsInjector CustomRuntimeOptionsInjector
}

// ResolveRuntimeOptions merges optional runtime option configs.
func ResolveRuntimeOptions(runtimeOptions ...RuntimeOptionsConfig) RuntimeOptionsConfig {
	var resolved RuntimeOptionsConfig
	for _, opts := range runtimeOptions {
		if opts.CustomRuntimeOptionsInjector != nil {
			resolved.CustomRuntimeOptionsInjector = opts.CustomRuntimeOptionsInjector
		}
	}
	return resolved
}

// ApplyCustomRuntimeOptions applies the configured custom runtime options injector.
func (c RuntimeOptionsConfig) ApplyCustomRuntimeOptions(
	resource RuntimeResource,
	operation RuntimeOperation,
	current string,
) (string, error) {
	if c.CustomRuntimeOptionsInjector == nil {
		return current, nil
	}
	return c.CustomRuntimeOptionsInjector(CustomRuntimeOptionsContext{
		Resource:  resource,
		Operation: operation,
		Current:   current,
	})
}
