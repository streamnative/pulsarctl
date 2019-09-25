package errors

import "github.com/streamnative/pulsarctl/pkg/pulsar"

var ArgError = pulsar.Output{
	Desc: "the topic name is not specified",
	Out:  "[✖]  only one argument is allowed to be used as a name",
}

var ArgsError = pulsar.Output{
	Desc: "the topic name and(or) the partitions is not specified",
	Out:  "[✖]  need to specified the topic name and the partitions",
}

var TopicAlreadyExistError = pulsar.Output{
	Desc: "the topic has been created",
	Out:  "[✖]  code: 409 reason: Partitioned topic already exists",
}

var TenantNotExistError = pulsar.Output{
	Desc: "the tenant of the namespace is not exist",
	Out:  "[✖]  code: 404 reason: Tenant does not exist",
}

var NamespaceNotExistError = pulsar.Output{
	Desc: "the namespace is not exist",
	Out:  "[✖]  code: 404 reason: Namespace does not exist",
}

var InvalidPartitionsNumberError = pulsar.Output{
	Desc: "the partitions number is invalid",
	Out:  "[✖]  invalid partition number '<number>'",
}

var TopicNameErrors = []pulsar.Output{
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

var NamespaceErrors = []pulsar.Output{
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
