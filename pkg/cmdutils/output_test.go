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

package cmdutils

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutputConfig(t *testing.T) {
	c := &OutputConfig{Format: "json"}
	sb := &strings.Builder{}

	// unsupported format
	err := c.WriteOutput(sb, OutputNegotiableFunc(func(format OutputFormat) OutputWritable {
		return nil
	}))
	assert.Error(t, err)

	// write error
	err = c.WriteOutput(sb, OutputNegotiableFunc(func(format OutputFormat) OutputWritable {
		return OutputWritableFunc(func(w io.Writer) error {
			return errors.New("expected")
		})
	}))
	assert.Error(t, err)

	// success
	err = c.WriteOutput(sb, OutputNegotiableFunc(func(format OutputFormat) OutputWritable {
		return OutputWritableFunc(func(w io.Writer) error {
			return nil
		})
	}))
	assert.NoError(t, err)
}

func TestOutputContent(t *testing.T) {
	// a test matrix of content, expected formats, and specific output
	contents := map[string]struct {
		tests map[OutputFormat]string
		init  func(oc *OutputContent)
	}{
		"WithText": {
			init: func(oc *OutputContent) {
				oc.WithText("foobar")
			},
			tests: map[OutputFormat]string{
				OutputFormat("bad"): "",
				TextOutputFormat:    "foobar",
				JSONOutputFormat:    "",
				YAMLOutputFormat:    "",
			},
		},
		"WithTextFunc": {
			init: func(oc *OutputContent) {
				oc.WithTextFunc(func(w io.Writer) error {
					_, _ = fmt.Fprint(w, "foobar")
					return nil
				})
			},
			tests: map[OutputFormat]string{
				OutputFormat("bad"): "",
				TextOutputFormat:    "foobar",
				JSONOutputFormat:    "",
				YAMLOutputFormat:    "",
			},
		},
		"WithObject": {
			init: func(oc *OutputContent) {
				oc.WithObject([]string{"foo", "bar"})
			},
			tests: map[OutputFormat]string{
				OutputFormat("bad"): "",
				TextOutputFormat: `[
  "foo",
  "bar"
]`,
				JSONOutputFormat: `[
  "foo",
  "bar"
]`,
				YAMLOutputFormat: `- foo
- bar
`,
			},
		},
		"WithObjectFunc": {
			init: func(oc *OutputContent) {
				oc.WithObjectFunc(func() interface{} { return []string{"foo", "bar"} })
			},
			tests: map[OutputFormat]string{
				OutputFormat("bad"): "",
				TextOutputFormat: `[
  "foo",
  "bar"
]`,
				JSONOutputFormat: `[
  "foo",
  "bar"
]`,
				YAMLOutputFormat: `- foo
- bar
`,
			},
		},
	}

	for content, c := range contents {
		t.Run(content, func(t *testing.T) {
			oc := NewOutputContent()
			c.init(oc)
			for format, expected := range c.tests {
				t.Run(string(format), func(t *testing.T) {
					ow := oc.Negotiate(format)
					if expected == "" {
						assert.Nil(t, ow)
					} else {
						assert.NotNil(t, ow)
						sb := &strings.Builder{}
						assert.NoError(t, ow.WriteTo(sb))
						assert.Equal(t, expected, sb.String())
					}
				})
			}
		})
	}
}
