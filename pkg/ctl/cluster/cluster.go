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

package cluster

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var argsError = cmdutils.Output{
	Desc: "the cluster name is not specified or the cluster name is specified more than one",
	Out:  "[✖]  the cluster name is not specified or the cluster name is specified more than one",
}

var clusterNonExist = cmdutils.Output{
	Desc: "the specified cluster does not exist in the broker",
	Out:  "[✖]  code: 412 reason: Cluster (cluster-name) does not exist.",
}

var failureDomainArgsError = cmdutils.Output{
	Desc: "the cluster name and(or) failure domain name is not specified or the name is specified more than one",
	Out:  "[✖]  need to specified the cluster name and the failure domain name",
}

var checkFailureDomainArgs = func(args []string) error {
	if len(args) != 2 {
		return errors.New("need to specified the cluster name and the failure domain name")
	}
	return nil
}

func Command(flagGrouping *cmdutils.FlagGrouping) *cobra.Command {
	resourceCmd := cmdutils.NewResourceCmd(
		"clusters",
		"Operations about cluster(s)",
		"",
		"cluster")

	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, CreateClusterCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, ListClustersCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getClusterDataCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, deleteClusterCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, UpdateClusterCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, updatePeerClustersCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getPeerClustersCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, createFailureDomainCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getFailureDomainCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, listFailureDomainCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, deleteFailureDomainCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, updateFailureDomainCmd)

	return resourceCmd
}
