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

package bkdata

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

var testParseFileTypeData = []struct {
	fileType       string
	parsedFileType FileType
	errString      string
}{
	{"journal", journal, ""},
	{"entrylog", entryLog, ""},
	{"index", index, ""},
	{"", "", fileTypeErrStr("")},
	{"invalid", "", fileTypeErrStr("invalid")},
}

func TestParseFileType(t *testing.T) {
	for _, data := range testParseFileTypeData {
		ft, err := ParseFileType(data.fileType)
		if data.errString != "" {
			assert.NotNil(t, err)
			assert.Equal(t, data.errString, err.Error())
			continue
		}
		assert.Equal(t, data.parsedFileType, ft)
	}
}

func fileTypeErrStr(fileType string) string {
	return errors.Errorf(
		"invalid file type %s, the file type only can be specified as 'journal', "+
			"'entrylog', 'index'", fileType).Error()
}
