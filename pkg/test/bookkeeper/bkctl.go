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

package bookkeeper

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/streamnative/pulsarctl/pkg/test"
)

const bkctl = "/opt/bookkeeper/bin/bkctl"

func ListBookies(containerID string) ([]string, error) {
	bookies := make([]string, 0)
	bookieStr, err := test.ExecCmd(containerID, []string{bkctl, "bookies", "list"})
	if err != nil {
		return nil, err
	}
	fmt.Println(bookieStr)
	lines := strings.Split(bookieStr, "\n")
	for _, v := range lines {
		if strings.TrimSpace(v) == "ReadWrite Bookies :" ||
			strings.TrimSpace(v) == "All Bookies :" ||
			strings.TrimSpace(v) == "" {
			continue
		}
		bookieAddr := strings.Split(v, ":")
		if len(bookieAddr) != 2 {
			return nil, errors.Errorf("get bookies address '%s' encountered unexpected error", v)
		}
		bookieIP := strings.Split(bookieAddr[0], "(")
		if len(bookieIP) != 2 {
			return nil, errors.Errorf("get bookies address '%s' encountered unexpected error", v)
		}
		bookies = append(bookies, fmt.Sprintf("%s:%s", bookieIP[0], bookieAddr[1]))
	}
	return bookies, nil
}
