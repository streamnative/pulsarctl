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

package topic

import "github.com/streamnative/pulsarctl/pkg/cmdutils"

var ArgError = cmdutils.Output{
	Desc: "the topic name is not specified or the topic name is specified more than one",
	Out:  "[✖]  the topic name is not specified or the topic name is specified more than one",
}

var ArgsError = cmdutils.Output{
	Desc: "the topic name and(or) the partitions is not specified",
	Out:  "[✖]  need to specified the topic name and the partitions",
}

var TopicAlreadyExistError = cmdutils.Output{
	Desc: "the topic has been created",
	Out:  "[✖]  code: 409 reason: Partitioned topic already exists",
}

var TopicNotFoundError = cmdutils.Output{
	Desc: "the specified topic does not exist",
	Out:  "[✖]  code: 404 reason: Topic not found",
}

var TenantNotExistError = cmdutils.Output{
	Desc: "the tenant of the namespace does not exist",
	Out:  "[✖]  code: 404 reason: Tenant does not exist",
}

var NamespaceNotExistError = cmdutils.Output{
	Desc: "the namespace does not exist",
	Out:  "[✖]  code: 404 reason: Namespace does not exist",
}

var InvalidPartitionsNumberError = cmdutils.Output{
	Desc: "the partitions number is invalid",
	Out:  "[✖]  invalid partition number '<number>'",
}

var TopicNameErrors = []cmdutils.Output{
	{
		Desc: "the topic name is not in the format of <tenant>/<namespace>/<topic> or <topic>",
		Out: "[✖]  Invalid short topic name '<topic-name>', it should be " +
			"in the format of <tenant>/<namespace>/<topic> or <topic>",
	},
	{
		Desc: "the topic name is not in the format of <domain>://<tenant>/<namespace>/<topic>",
		Out: "[✖]  Invalid complete topic name '<topic-name>', it should be in " +
			"the format of <domain>://<tenant>/<namespace>/<topic>",
	},
	{
		Desc: "the topic name is not in the format of <tenant>/<namespace>/<topic>",
		Out: "[✖]  Invalid topic name '<topic-name>', it should be in the format of" +
			"<tenant>/<namespace>/<topic>",
	},
}

var NamespaceErrors = []cmdutils.Output{
	{
		Desc: "the namespace name is not in the format of <tenant>/<namespace>",
		Out:  "[✖]  The complete name of namespace is invalid. complete name : <namespace-complete-name>",
	},
	{
		Desc: "the tenant name and(or) namespace name is empty",
		Out:  "[✖]  Invalid tenant or namespace. [<tenant>/<namespace>]",
	},
	{
		Desc: "the tenant name contains unsupported special chars. " +
			"the alphanumeric (a-zA-Z0-9) and the special chars (-=:.%)  is allowed",
		Out: "[✖]  Tenant name include unsupported special chars. tenant : [<namespace>]",
	},
	{
		Desc: "the namespace name contains unsupported special chars. " +
			"the  alphanumeric (a-zA-Z0-9) and the special chars (-=:.%) is allowed",
		Out: "[✖]  Namespace name include unsupported special chars. namespace : [<namespace>]",
	},
}
