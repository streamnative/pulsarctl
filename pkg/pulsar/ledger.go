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
//

package pulsar

import (
	"strconv"
)

type Ledger interface {
	// Delete the specified ledger
	DeleteLedger(int64) error
}

type ledger struct {
	client   *bookieClient
	request  *client
	basePath string
	params   map[string]string
}

func (c *bookieClient) Ledger() Ledger {
	return &ledger{
		client:   c,
		request:  c.client,
		basePath: "/delete",
		params:   make(map[string]string),
	}
}

func (c *ledger) DeleteLedger(ledgerID int64) error {
	endpoint := c.client.bookieEndpoint(c.basePath)
	c.params["ledger_id"] = strconv.FormatInt(ledgerID, 10)
	return c.request.deleteWithQueryParams(endpoint, nil, c.params)
}
