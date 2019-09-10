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
    `encoding/json`
    `github.com/streamnative/pulsarctl/pkg/pulsar`
    `github.com/stretchr/testify/assert`
    `strings`
    `testing`
)

func TestUpdateSink(t *testing.T)  {
    basePath, err := getDirHelp()
    if basePath == "" || err != nil {
        t.Error(err)
    }
    t.Logf("base path: %s", basePath)

    args := []string{"create",
        "--tenant", "public",
        "--namespace", "default",
        "--name", "test-sink-update",
        "--inputs", "test-topic",
        "--archive", basePath + "/test/sinks/pulsar-io-jdbc-2.4.0.nar",
        "--sink-config-file", basePath + "/test/sinks/mysql-jdbc-sink.yaml",
    }

    createOut, _, err := TestSinksCommands(createSinksCmd, args)
    assert.Nil(t, err)
    assert.Equal(t, createOut.String(), "Created test-sink-update successfully")

    updateArgs := []string{"update",
        "--name", "test-sink-update",
        "--cpu", "5.0",
    }

    _, _, err = TestSinksCommands(updateSinksCmd, updateArgs)
    assert.Nil(t, err)
    getArgs := []string{"get",
        "--tenant", "public",
        "--namespace", "default",
        "--name", "test-sink-update",
    }

    out, _, err := TestSinksCommands(updateSinksCmd, getArgs)
    assert.Nil(t, err)

    var sinkConf pulsar.SinkConfig
    err = json.Unmarshal(out.Bytes(), &sinkConf)
    assert.Nil(t, err)

    assert.Equal(t, sinkConf.Resources.CPU, 5.0)

    // test the sink name not exist
    failureUpdateArgs := []string{"update",
        "--name", "not-exist",
    }
    _, err, _ = TestSinksCommands(updateSinksCmd, failureUpdateArgs)
    assert.NotNil(t, err)
    failMsg := "Sink not-exist doesn't exist"
    assert.True(t, strings.Contains(err.Error(), failMsg))
}
