package errors

import . "github.com/streamnative/pulsarctl/pkg/pulsar"

var ArgError = Output{
	Desc: "the topic name is not specified",
	Out:  "[✖]  only one argument is allowed to be used as a name",
}

var ArgsError = Output{
	Desc: "the topic name and(or) the subscription name is not specified",
	Out:  "[✖]  need to specified the topic name and the subscription name",
}

var TopicNotFoundError = Output{
	Desc: "the specified topic is not exist",
	Out:  "[✖]  code: 404 reason: Topic not found",
}

var SubNotFoundError = Output{
	Desc: "the specified subscription is not exist",
	Out:  "[✖]  code: 404 reason: Subscription not found",
}

var TopicNameErrors = []Output{
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

var NamespaceErrors = []Output{
	{
		Desc: "the namespace name is not in the format of <tenant>/<namespace>",
		Out:  "[✖]  The complete name of namespace is invalid. complete name : <namespace-complete-name>",
	},
	{
		Desc: "the tenant name and(or) namespace name is empty",
		Out:  "[✖]  Invalid tenant or namespace. [<tenant>/<namespace>]",
	},
	{
		Desc: "the tenant name contains unsupported special chars" +
			"the alphanumeric (a-zA-Z0-9) and the special chars (-=:.%)  is allowed",
		Out: "[✖]  Tenant name include unsupported special chars. tenant : [<namespace>]",
	},
	{
		Desc: "the namespace name contains unsupported special chars" +
			"the  alphanumeric (a-zA-Z0-9) and the special chars (-=:.%) is allowed",
		Out: "[✖]  Namespace name include unsupported special chars. namespace : [<namespace>]",
	},
}

var MessageIdErrors = []Output{
	{
		Desc: "the split of message id is not valid",
		Out:  "[✖]  Invalid message id string. <message-id>",
	},
	{
		Desc: "the ledger id is not valid",
		Out:  "[✖]  Invalid ledger id string. <message-id>",
	},
	{
		Desc: "the entry id is not valid",
		Out:  "[✖]  Invalid entry id string. <message-id>",
	},
}
