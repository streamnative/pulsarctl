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
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/admin"
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/admin/config"
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/rest"
	"github.com/kris-nova/logger"
	"github.com/spf13/cobra"

	"github.com/streamnative/pulsarctl/pkg/bookkeeper"
)

// Client defines all the interfaces of pulsarctl
type Client interface {
	admin.Client

	Token() Token
}

type client struct {
	admin admin.Client
}

func (c *client) Clusters() admin.Clusters {
	return c.admin.Clusters()
}

func (c *client) Functions() admin.Functions {
	return c.admin.Functions()
}

func (c *client) Tenants() admin.Tenants {
	return c.admin.Tenants()
}

func (c *client) Topics() admin.Topics {
	return c.admin.Topics()
}

func (c *client) Subscriptions() admin.Subscriptions {
	return c.admin.Subscriptions()
}

func (c *client) Sources() admin.Sources {
	return c.admin.Sources()
}

func (c *client) Sinks() admin.Sinks {
	return c.admin.Sinks()
}

func (c *client) Namespaces() admin.Namespaces {
	return c.admin.Namespaces()
}

func (c *client) Schemas() admin.Schema {
	return c.admin.Schemas()
}

func (c *client) NsIsolationPolicy() admin.NsIsolationPolicy {
	return c.admin.NsIsolationPolicy()
}

func (c *client) Brokers() admin.Brokers {
	return c.admin.Brokers()
}

func (c *client) BrokerStats() admin.BrokerStats {
	return c.admin.BrokerStats()
}

func (c *client) ResourceQuotas() admin.ResourceQuotas {
	return c.admin.ResourceQuotas()
}

func (c *client) FunctionsWorker() admin.FunctionsWorker {
	return c.admin.FunctionsWorker()
}

func (c *client) Packages() admin.Packages {
	return c.admin.Packages()
}

func (c *client) Token() Token {
	return &token{}
}

const IncompatibleFlags = "cannot be used at the same time"

// NewVerbCmd defines a standard resource command
func NewResourceCmd(use, short, long string, aliases ...string) *cobra.Command {
	return &cobra.Command{
		Use:     use,
		Short:   short,
		Long:    long,
		Aliases: aliases,
		Run: func(cmd *cobra.Command, _ []string) {
			if err := cmd.Help(); err != nil {
				logger.Debug("ignoring error %q", err.Error())
			}
		},
	}
}

var CheckNameArgError = defaultNameArgsError

var defaultNameArgsError = func(err error) {
	os.Exit(1)
}

// GetNameArg tests to ensure there is only 1 name argument
func GetNameArg(args []string, errMsg string) (string, error) {
	if len(args) > 1 || len(args) == 0 {
		logger.Critical(errMsg)
		err := errors.New(errMsg)
		CheckNameArgError(err)
		return "", err
	}
	if len(args) == 1 {
		return strings.TrimSpace(args[0]), nil
	}
	return "", nil
}

func GetNameArgs(args []string, check func(args []string) error) ([]string, error) {
	err := check(args)
	if err != nil {
		logger.Critical(err.Error())
		CheckNameArgError(err)
		// for testing
		return nil, err
	}
	return args, nil
}

func NewPulsarClient() Client {
	return PulsarCtlConfig.Client(config.V2)
}

func NewPulsarClientWithAPIVersion(version config.APIVersion) Client {
	return PulsarCtlConfig.Client(version)
}

func NewBookieClient() bookkeeper.Client {
	return PulsarCtlConfig.BookieClient()
}

func PrintJSON(w io.Writer, obj interface{}) {
	b, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		_, _ = fmt.Fprintf(w, "unexpected response type: %v\n", err)
		return
	}
	_, _ = fmt.Fprintln(w, string(b))
}

func PrintError(w io.Writer, err error) {
	msg := err.Error()
	if rest.IsAdminError(err) {
		ae, _ := err.(rest.Error)
		msg = ae.Reason
	}
	_, _ = fmt.Fprintln(w, "error:", msg)
}

func RunFuncWithTimeout(task func([]string, interface{}) bool, condition bool, timeout time.Duration,
	args []string, obj interface{}) error {

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	result := !condition
	timer := time.After(timeout)

	for condition != result {
		select {
		case <-timer:
			return errors.New("task timeout")
		case <-ticker.C:
			result = task(args, obj)
		}
	}

	return nil
}
