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
	"github.com/spf13/pflag"
)

var GlobalOutputConfig = OutputConfig{}

type OutputConfig struct {
	// the output format (Table, Plain, Json)
	Format string
}

func (c *OutputConfig) FlagSet() *pflag.FlagSet {
	flags := pflag.NewFlagSet(
		"Output Config",
		pflag.ContinueOnError)

	flags.StringVarP(
		&c.Format,
		"output",
		"o",
		string(Table),
		"The output format")

	return flags
}

func (c *OutputConfig) IsTableFormat() bool {
	return c.Format == string(Table)
}

func (c *OutputConfig) IsJsonFormat() bool {
	return c.Format == string(Json)
}

func (c *OutputConfig) IsPlainFormat() bool {
	return c.Format == string(Plain)
}

type OutputFormat string

const (
	Table OutputFormat = "table"
	Plain OutputFormat = "plain"
	Json  OutputFormat = "json"
)

func (fmt OutputFormat) String() string {
	return string(fmt)
}

func validateOutputFormat(v string) OutputFormat {
	switch v {
	case string(Table):
		return Table
	case string(Plain):
		return Plain
	case string(Json):
		return Json
	default:
		return ""
	}
}
