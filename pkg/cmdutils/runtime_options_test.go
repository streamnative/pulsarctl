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

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuntimeOptionsConfigApplyCustomRuntimeOptionsNoInjector(t *testing.T) {
	value, err := (RuntimeOptionsConfig{}).ApplyCustomRuntimeOptions(
		RuntimeResourceFunction,
		RuntimeOperationCreate,
		`{"base":true}`,
	)

	assert.NoError(t, err)
	assert.Equal(t, `{"base":true}`, value)
}

func TestResolveRuntimeOptionsUsesLastInjector(t *testing.T) {
	first := RuntimeOptionsConfig{CustomRuntimeOptionsInjector: func(CustomRuntimeOptionsContext) (string, error) {
		return "first", nil
	}}
	second := RuntimeOptionsConfig{CustomRuntimeOptionsInjector: func(CustomRuntimeOptionsContext) (string, error) {
		return "second", nil
	}}

	value, err := ResolveRuntimeOptions(first, second).ApplyCustomRuntimeOptions(
		RuntimeResourceSource,
		RuntimeOperationUpdate,
		"current",
	)

	assert.NoError(t, err)
	assert.Equal(t, "second", value)
}

func TestAddVerbCmdWithRuntimeOptionsAssignsVerbConfig(t *testing.T) {
	flagGrouping := NewGrouping()
	parent := NewResourceCmd("resource", "short", "long")
	captured := RuntimeOptionsConfig{}
	expected := RuntimeOptionsConfig{CustomRuntimeOptionsInjector: func(CustomRuntimeOptionsContext) (string, error) {
		return "injected", nil
	}}

	AddVerbCmdWithRuntimeOptions(flagGrouping, parent, expected, func(vc *VerbCmd) {
		captured = vc.RuntimeOptions
		vc.SetDescription("create", "create", "create", "")
	})

	value, err := captured.ApplyCustomRuntimeOptions(RuntimeResourceSink, RuntimeOperationCreate, "current")
	assert.NoError(t, err)
	assert.Equal(t, "injected", value)
	assert.Len(t, parent.Commands(), 1)
}
