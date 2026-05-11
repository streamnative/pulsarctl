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

package schemas

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/stretchr/testify/assert"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func withAdminURLForTest(t *testing.T, webServiceURL string) {
	t.Helper()
	oldURL := cmdutils.PulsarCtlConfig.WebServiceURL
	cmdutils.PulsarCtlConfig.WebServiceURL = webServiceURL
	t.Cleanup(func() {
		cmdutils.PulsarCtlConfig.WebServiceURL = oldURL
	})
}

func TestGetSchemaAllVersion(t *testing.T) {
	topic := "persistent://public/default/test-schema"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/admin/v2/schemas/public/default/test-schema/schemas", r.URL.Path)
		_, _ = w.Write([]byte(`{
  "getSchemaResponses": [
    {
      "version": 2,
      "type": "AVRO",
      "timestamp": 1,
      "data": "{\"type\":\"record\",\"name\":\"Test\",\"fields\":[]}",
      "properties": {}
    }
  ]
}`))
	}))
	defer srv.Close()
	withAdminURLForTest(t, srv.URL)

	args := []string{"get", topic, "--all-version"}
	out, execErr, err := TestSchemasCommands(getSchema, args)
	assert.Nil(t, err)
	assert.Nil(t, execErr)

	var infos []utils.SchemaInfoWithVersion
	err = json.Unmarshal(out.Bytes(), &infos)
	assert.Nil(t, err)
	assert.Len(t, infos, 1)
	assert.Equal(t, int64(2), infos[0].Version)
	assert.Equal(t, "test-schema", infos[0].SchemaInfo.Name)
}

func TestGetSchemaVersionConflict(t *testing.T) {
	args := []string{"get", "persistent://public/default/test-schema", "--version", "1", "--all-version"}
	_, execErr, err := TestSchemasCommands(getSchema, args)
	assert.Nil(t, err)
	assert.NotNil(t, execErr)
	assert.Contains(t, execErr.Error(), "--version and --all-version")
}

func TestDeleteSchemaWithForce(t *testing.T) {
	topic := "persistent://public/default/test-schema"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/admin/v2/schemas/public/default/test-schema/schema", r.URL.Path)
		assert.Equal(t, "true", r.URL.Query().Get("force"))
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()
	withAdminURLForTest(t, srv.URL)

	args := []string{"delete", topic, "--force"}
	out, execErr, err := TestSchemasCommands(deleteSchema, args)
	assert.Nil(t, err)
	assert.Nil(t, execErr)
	assert.Equal(t, "Deleted persistent://public/default/test-schema successfully\n", out.String())
}

func TestSchemaCompatibility(t *testing.T) {
	topic := "persistent://public/default/test-schema"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/admin/v2/schemas/public/default/test-schema/compatibility", r.URL.Path)
		var payload utils.PostSchemaPayload
		err := json.NewDecoder(r.Body).Decode(&payload)
		assert.Nil(t, err)
		assert.Equal(t, "AVRO", payload.SchemaType)
		_, _ = w.Write([]byte(`{"compatibility":true,"schemaCompatibilityStrategy":"FULL"}`))
	}))
	defer srv.Close()
	withAdminURLForTest(t, srv.URL)

	tmpFile := filepath.Join(t.TempDir(), "schema.json")
	err := os.WriteFile(tmpFile, []byte(`{
  "type":"AVRO",
  "schema":"{\"type\":\"record\",\"name\":\"Test\",\"fields\":[]}",
  "properties":{}
}`), 0o644)
	assert.Nil(t, err)

	args := []string{"compatibility", topic, "-f", tmpFile}
	out, execErr, err := TestSchemasCommands(testCompatibility, args)
	assert.Nil(t, err)
	assert.Nil(t, execErr)

	var result utils.IsCompatibility
	err = json.Unmarshal(out.Bytes(), &result)
	assert.Nil(t, err)
	assert.True(t, result.IsCompatibility)
	assert.Equal(t, utils.SchemaCompatibilityStrategyFull, result.SchemaCompatibilityStrategy)
}

func TestSchemaCompatibilityMissingFile(t *testing.T) {
	args := []string{"compatibility", "persistent://public/default/test-schema", "-f", "not-exist.json"}
	_, execErr, err := TestSchemasCommands(testCompatibility, args)
	assert.Nil(t, err)
	assert.NotNil(t, execErr)
	assert.Contains(t, execErr.Error(), "no such file or directory")
}
