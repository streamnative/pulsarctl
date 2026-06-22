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

package functions

import (
	"errors"
	"testing"

	util "github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/stretchr/testify/assert"
)

func TestApplyFunctionRuntimeOptionsUsesCurrentValue(t *testing.T) {
	funcData := &util.FunctionData{FuncConf: &util.FunctionConfig{CustomRuntimeOptions: `{"base":true}`}}
	runtimeOptions := cmdutils.RuntimeOptionsConfig{
		CustomRuntimeOptionsInjector: func(ctx cmdutils.CustomRuntimeOptionsContext) (string, error) {
			assert.Equal(t, cmdutils.RuntimeResourceFunction, ctx.Resource)
			assert.Equal(t, cmdutils.RuntimeOperationCreate, ctx.Operation)
			assert.Equal(t, `{"base":true}`, ctx.Current)
			return `{"base":true,"extra":true}`, nil
		},
	}

	err := applyFunctionRuntimeOptions(funcData, runtimeOptions, cmdutils.RuntimeOperationCreate)

	assert.NoError(t, err)
	assert.Equal(t, `{"base":true,"extra":true}`, funcData.FuncConf.CustomRuntimeOptions)
}

func TestApplyFunctionRuntimeOptionsCanClearValue(t *testing.T) {
	funcData := &util.FunctionData{FuncConf: &util.FunctionConfig{CustomRuntimeOptions: `{"base":true}`}}
	runtimeOptions := cmdutils.RuntimeOptionsConfig{
		CustomRuntimeOptionsInjector: func(cmdutils.CustomRuntimeOptionsContext) (string, error) {
			return "", nil
		},
	}

	err := applyFunctionRuntimeOptions(funcData, runtimeOptions, cmdutils.RuntimeOperationUpdate)

	assert.NoError(t, err)
	assert.Empty(t, funcData.FuncConf.CustomRuntimeOptions)
}

func TestApplyFunctionRuntimeOptionsWrapsInjectorError(t *testing.T) {
	funcData := &util.FunctionData{FuncConf: &util.FunctionConfig{CustomRuntimeOptions: "current"}}
	runtimeOptions := cmdutils.RuntimeOptionsConfig{
		CustomRuntimeOptionsInjector: func(cmdutils.CustomRuntimeOptionsContext) (string, error) {
			return "", errors.New("boom")
		},
	}

	err := applyFunctionRuntimeOptions(funcData, runtimeOptions, cmdutils.RuntimeOperationUpdate)

	assert.EqualError(t, err, "inject function update custom runtime options: boom")
	assert.Equal(t, "current", funcData.FuncConf.CustomRuntimeOptions)
}

func TestApplyFunctionRuntimeOptionsNoInjectorLeavesValue(t *testing.T) {
	funcData := &util.FunctionData{FuncConf: &util.FunctionConfig{CustomRuntimeOptions: `{"base":true}`}}

	err := applyFunctionRuntimeOptions(funcData, cmdutils.RuntimeOptionsConfig{}, cmdutils.RuntimeOperationCreate)

	assert.NoError(t, err)
	assert.Equal(t, `{"base":true}`, funcData.FuncConf.CustomRuntimeOptions)
}
