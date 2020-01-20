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

var testParseBookieTypeData = []struct {
	bookieType       string
	parsedBookieType BookieType
	errString        string
}{
	{"rw", rw, ""},
	{"ro", ro, ""},
	{"", "", bookieTypeErrStr("")},
	{"r", "", bookieTypeErrStr("r")},
}

func TestParseBookieType(t *testing.T) {
	for _, data := range testParseBookieTypeData {
		bkt, err := ParseBookieType(data.bookieType)
		if data.errString != "" {
			assert.NotNil(t, err)
			assert.Equal(t, data.errString, err.Error())
			continue
		}
		assert.Equal(t, data.parsedBookieType, bkt)
	}
}

func bookieTypeErrStr(bookieType string) string {
	return errors.Errorf(
		"invalid bookie type %s, the bookie type only can be specified as 'rw' or 'ro'", bookieType).Error()
}
