package namespace

import . "github.com/streamnative/pulsarctl/pkg/pulsar"

var ArgError = Output{
	Desc: "the namespace name is not specified",
	Out:  "[✖]  only one argument is allowed to be used as a name",
}

var NsNotExistError = Output{
	Desc: "the specified namespace name does not exist",
	Out: "[✖]  code: 404 reason: Namespace does not exist",
}

var NsErrors = []Output{
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
