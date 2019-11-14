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

package bookie

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListDiskFileArgError(t *testing.T) {
	args := []string{"listdiskfile"}
	_, _, nameErr, _ := testBookieCommands(listDiskFileCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "the file type is not specified or the file type is specified more than one",
		nameErr.Error())

	args = []string{"listdiskfile", "invalid"}
	_, execErr, _, _ := testBookieCommands(listDiskFileCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, fmt.Sprintf("invalid file type %s, the file type only can be specified as 'journal', "+
		"'entrylog', 'index'", "invalid"), execErr.Error())
}
