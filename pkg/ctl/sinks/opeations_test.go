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
	"encoding/json"
	"fmt"
	"path"
	"testing"
	"time"

	"github.com/streamnative/pulsar-admin-go/pkg/utils"
	"github.com/stretchr/testify/assert"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/test"
)

// This tests will test all sink operations
func TestSinksOperations(t *testing.T) {
	narFile := path.Join(resourceDir(), "data-generator.nar")
	sinkName := "test-sink-opt" + test.RandomSuffix()

	defaultArgs := []string{
		"--tenant", "public",
		"--namespace", "default",
		"--name", sinkName,
	}

	listArgs := []string{"list"}
	out, execErr, err := TestSinksCommands(listSinksCmd, listArgs)
	failImmediatelyIfErrorNotNil(t, execErr, err)
	assert.NotContains(t, out.String(), sinkName)

	createArgs := []string{
		"create",
		"--inputs", "sink-input",
		"--archive", narFile,
	}
	out, execErr, err = TestSinksCommands(createSinksCmd, append(createArgs, defaultArgs...))
	failImmediatelyIfErrorNotNil(t, execErr, err)
	assert.Equal(t, fmt.Sprintf("Created %s successfully\n", sinkName), out.String())

	out, execErr, err = TestSinksCommands(listSinksCmd, listArgs)
	failImmediatelyIfErrorNotNil(t, execErr, err)
	assert.Contains(t, out.String(), sinkName)

	getArgs := []string{"get"}
	out, execErr, err = TestSinksCommands(getSinksCmd, append(getArgs, defaultArgs...))
	failImmediatelyIfErrorNotNil(t, execErr, err)
	var sinkConf utils.SinkConfig
	err = json.Unmarshal(out.Bytes(), &sinkConf)
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(t, "public", sinkConf.Tenant)
	assert.Equal(t, "default", sinkConf.Namespace)
	assert.Equal(t, sinkName, sinkConf.Name)

	updateArgs := []string{"update", "--parallelism", "2"}
	_, execErr, err = TestSinksCommands(updateSinksCmd, append(updateArgs, defaultArgs...))
	failImmediatelyIfErrorNotNil(t, execErr, err)

	out, execErr, err = TestSinksCommands(getSinksCmd, append(getArgs, defaultArgs...))
	failImmediatelyIfErrorNotNil(t, execErr, err)
	err = json.Unmarshal(out.Bytes(), &sinkConf)
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(t, "public", sinkConf.Tenant)
	assert.Equal(t, "default", sinkConf.Namespace)
	assert.Equal(t, 2, sinkConf.Parallelism)
	assert.Equal(t, sinkName, sinkConf.Name)

	updateArgs = []string{"update", "--parallelism", "1"}
	_, execErr, err = TestSinksCommands(updateSinksCmd, append(updateArgs, defaultArgs...))
	failImmediatelyIfErrorNotNil(t, execErr, err)

	stopArgs := []string{"stop"}
	_, execErr, err = TestSinksCommands(stopSinksCmd, append(stopArgs, defaultArgs...))
	failImmediatelyIfErrorNotNil(t, execErr, err)

	startArgs := []string{"start"}
	_, execErr, err = TestSinksCommands(startSinksCmd, append(startArgs, defaultArgs...))
	failImmediatelyIfErrorNotNil(t, execErr, err)

	statusArgs := []string{"status"}
	var status utils.SinkStatus
	task := func(args []string, obj interface{}) bool {
		out, execErr, err := TestSinksCommands(statusSinksCmd, args)
		if err != nil {
			fmt.Println(err.Error())
			return false
		}
		if execErr != nil {
			fmt.Println(execErr.Error())
			return false
		}
		err = json.Unmarshal(out.Bytes(), &obj)
		if err != nil {
			fmt.Println(err.Error())
			return false
		}
		s := obj.(*utils.SinkStatus)
		return len(s.Instances) == 1 && s.Instances[0].Status.Running
	}
	err = cmdutils.RunFuncWithTimeout(task, true, 3*time.Minute,
		append(statusArgs, defaultArgs...), &status)
	failImmediatelyIfErrorNotNil(t, err)

	restartArgs := []string{"restart"}
	_, execErr, err = TestSinksCommands(restartSinksCmd, append(restartArgs, defaultArgs...))
	failImmediatelyIfErrorNotNil(t, execErr, err)

	err = cmdutils.RunFuncWithTimeout(task, true, 3*time.Minute,
		append(statusArgs, defaultArgs...), &status)
	failImmediatelyIfErrorNotNil(t, err)

	_, execErr, err = TestSinksCommands(stopSinksCmd, append(stopArgs, defaultArgs...))
	failImmediatelyIfErrorNotNil(t, execErr, err)

	task = func(args []string, obj interface{}) bool {
		out, execErr, err := TestSinksCommands(statusSinksCmd, args)
		if err != nil {
			return false
		}
		if execErr != nil {
			return false
		}
		err = json.Unmarshal(out.Bytes(), &obj)
		if err != nil {
			return false
		}
		s := obj.(*utils.SinkStatus)
		return len(s.Instances) == 1 && !s.Instances[0].Status.Running
	}
	err = cmdutils.RunFuncWithTimeout(task, true, 3*time.Minute,
		append(statusArgs, defaultArgs...), &status)
	failImmediatelyIfErrorNotNil(t, err)

	deleteArgs := []string{"delete"}
	_, execErr, err = TestSinksCommands(deleteSinksCmd, append(deleteArgs, defaultArgs...))
	failImmediatelyIfErrorNotNil(t, execErr, err)

	out, execErr, err = TestSinksCommands(listSinksCmd, listArgs)
	failImmediatelyIfErrorNotNil(t, execErr, err)
	assert.NotContains(t, out.String(), sinkName)
}
