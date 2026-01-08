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
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/spf13/cobra"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func Command(flagGrouping *cmdutils.FlagGrouping) *cobra.Command {
	resourceCmd := cmdutils.NewResourceCmd(
		"schemas",
		"Operations related to Schemas associated with Pulsar topics",
		"")

	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getSchema)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, deleteSchema)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, uploadSchema)

	return resourceCmd
}

func PrintSchema(w io.Writer, schema *utils.SchemaInfoWithVersion) {
	name, err := json.MarshalIndent(schema.SchemaInfo.Name, "", "  ")
	if err != nil {
		_, _ = fmt.Fprintf(w, "unexpected response type: %v\n", err)
		return
	}

	schemaType, err := json.MarshalIndent(schema.SchemaInfo.Type, "", "  ")
	if err != nil {
		_, _ = fmt.Fprintf(w, "unexpected response type: %v\n", err)
		return
	}

	properties, err := json.MarshalIndent(schema.SchemaInfo.Properties, "", "  ")
	if err != nil {
		_, _ = fmt.Fprintf(w, "unexpected response type: %v\n", err)
		return
	}
	s, _ := prettyPrint(schema.SchemaInfo.Schema)
	_, _ = fmt.Fprintf(w, "{\n  name: %s \n  schema: %s\n  type: %s \n  properties: %s\n}",
		string(name), string(s), string(schemaType), string(properties))
}

func prettyPrint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "    ")
	return out.Bytes(), err
}
