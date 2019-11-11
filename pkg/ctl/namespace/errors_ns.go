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

package namespace

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

var ArgError = cmdutils.Output{
	Desc: "the namespace name is not specified or the namespace name is specified more than one",
	Out:  "[✖]  the namespace name is not specified or the namespace name is specified more than one",
}

var AuthNotEnable = cmdutils.Output{
	Desc: "the authorization is not enabled",
	Out:  "[✖]  code: 501 reason: Authorization is not enabled",
}

var NsNotExistError = cmdutils.Output{
	Desc: "the specified namespace name does not exist",
	Out:  "[✖]  code: 404 reason: Namespace does not exist",
}

var NsErrors = []cmdutils.Output{
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
