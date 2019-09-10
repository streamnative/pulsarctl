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
    `github.com/stretchr/testify/assert`
    `os`
    `strings`
    `testing`
)

func TestCreateSinks(t *testing.T) {
    basePath, err := getDirHelp()
    if basePath == "" || err != nil {
        t.Error(err)
    }
    t.Logf("base path: %s", basePath)

    args := []string{"create",
        "--tenant", "public",
        "--namespace", "default",
        "--name", "test-sink-delete",
        "--inputs", "test-topic",
        "--archive", basePath + "/test/sinks/pulsar-io-jdbc-2.4.0.nar",
        "--sink-config-file", basePath + "/test/sinks/mysql-jdbc-sink.yaml",
        "--parallelism", "1",
    }
    _, _, err = TestSinksCommands(createSinksCmd, args)
    assert.Nil(t, err)
}

func TestFailureCreateSinks(t *testing.T) {
    basePath, err := getDirHelp()
    if basePath == "" || err != nil {
        t.Error(err)
    }
    t.Logf("base path: %s", basePath)

    narName := "dummy-pulsar-io-mysql.nar"
    _, err = os.Create(narName)
    assert.Nil(t, err)

    defer os.Remove(narName)

    failArgs := []string{"create",
        "--tenant", "public",
        "--namespace", "default",
        "--name", "test-sink-delete",
        "--inputs", "test-topic",
        "--archive", basePath + "/test/sinks/pulsar-io-jdbc-2.4.0.nar",
        "--sink-config-file", basePath + "/test/sinks/mysql-jdbc-sink.yaml",
    }

    exceptedErr := "Sink test-sink-create already exists"
    out, execErr, _ := TestSinksCommands(createSinksCmd, failArgs)
    assert.True(t, strings.Contains(out.String(), exceptedErr))
    assert.NotNil(t, execErr)

    narFailArgs := []string{"create",
       "--tenant", "public",
       "--namespace", "default",
       "--name", "test-sink-create-nar-fail",
       "--inputs", "my-topic",
       "--archive", narName,
    }

    narErrInfo := "Sink class org.apache.pulsar.io.jdbc.JdbcAutoSchemaSink must be in class path"
    narOut, execErr, _ := TestSinksCommands(createSinksCmd, narFailArgs)
    assert.True(t, strings.Contains(narOut.String(), narErrInfo))
    assert.NotNil(t, execErr)
}

