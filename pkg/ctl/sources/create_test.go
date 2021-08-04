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

package sources

import (
	"path"
	"strings"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestCreateSourceFailedByClassname(t *testing.T)  {
	narName := path.Join(resourceDir(), "data-generator.nar")
	sourceName := "create-source-fail" + test.RandomSuffix()

	narFailArgs := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", sourceName,
		"--destination-topic-name", "my-topic",
		"--classname", "org.apache.pulsar.io.kafka.KafkaBytesSource",
		"--archive", narName,
	}

	narErrInfo := "Source class org.apache.pulsar.io.kafka.KafkaBytesSource must be in class path"
	narOut, execErr, _ := TestSourcesCommands(createSourcesCmd, narFailArgs)
	assert.True(t, strings.Contains(narOut.String(), narErrInfo))
	assert.NotNil(t, execErr)
}
