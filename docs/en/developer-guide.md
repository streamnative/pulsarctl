<!--

    Licensed to the Apache Software Foundation (ASF) under one
    or more contributor license agreements.  See the NOTICE file
    distributed with this work for additional information
    regarding copyright ownership.  The ASF licenses this file
    to you under the Apache License, Version 2.0 (the
    "License"); you may not use this file except in compliance
    with the License.  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing,
    software distributed under the License is distributed on an
    "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
    KIND, either express or implied.  See the License for the
    specific language governing permissions and limitations
    under the License.

-->

# How to add a new command

The `pulsarctl` is a command-line tool written in the go language that helps administrators and users manage clusters, tenants, namespaces, topics, schemas, sources, sinks, functions, and more.

## Project structure

```
├── docs
├── pkg
│   ├── auth
│   ├── cmdutils
│   ├── ctl
│   │   ├── cluster
│   │   ├── completion
│   │   ├── functions
│   │   ├── namespace
│   │   ├── schemas
│   │   ├── sinks
│   │   ├── sources
│   │   ├── tenant
│   │   ├── topic
│   │   └── utils
│   └── pulsar
├── site
└── test
```

- `pkg` is used to store pulsarctl related libraries. There are four subdirectories, as follows:
    - `auth` is used to store encryption-related code
    - `cmdutils` has a simple wrapper for cobra
    - `ctl` is used to store pulsarctl related commands
    - `pulsar` is a public package of pulsarctl
- `test` is used to store resources related to test
- `site` is the website related code of pulsarctl, which is convenient for users to view and quickly locate the usage and precautions of related commands.
- `docs` is used to store pulsarctl related document content

> To avoid circular references, where `auth` and `cmdutils` are two separate packages, the two packages `ctl` and `pulsar` are not referenced and will not be referenced to each other. 
`pulsar` is a public package that references `auth` but does not reference `ctl` and `cmdutils`. 
`ctl` as the core pkg that implements pulsarctl, it will refer to the three packages `auth`, `cmdutils`, `pulsar`, but other packages will not reference it.

## Add a new command

The command format used by pulsarctl is as follows:

```bash
pulsarctl [commands] [sub commands] [flags]
```

The contents of `[command]` are consistent with the file directory under `ctl`. When you want to create a new command, create a new folder in the `ctl` directory.
Name the folder the name of the command. Also create a `command-name.go` file in the `pulsar` directory to write the interface function associated with the command.

`[sub commands]` belongs to the subcommand under `[commands]`, if you want to add a sub command to it under the current existing `[commands]`
Please create a `sub-command-name.go` file in the command directory and add your relevant code logic. After you've finished writing your code, add the relevant test code to your code.
Make sure the test can override your code logic.

Let's take a look at how to quickly add a new command to pulsarctl with `pulsarctl topics create (topic name) 0` as an example.

1. Create a folder named topic under the ctl directory

```bash
mkdir topic
```

2. Create files named `create.go` and `topic.go` under the topic directory, where `create.go` is used to write create topic related content.
   `topic.go` is used to store all commands related to topic.

```bash
cd pkg/ctl/topic
touch create.go topic.go
```

### Write create.go

Under normal circumstances, each command file consists of two functions, `CreateTopicCmd` and `doCreateTopic`, as follows:
```
func CreateTopicCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for creating topic."
	desc.CommandPermission = "This command requires namespace admin permissions."

	var examples []pulsar.Example
	createNonPartitions := pulsar.Example{
		Desc:    "Create a non-partitioned topic <topic-name>",
		Command: "Pulsarctl topics create <topic-name> 0",
	}
	examples = append(examples, createNonPartitions)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Create topic <topic-name> with <partition-num> partitions successfully",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"create",
		"Create a topic with n partitions",
		desc.ToString(),
		desc.ExampleToString(),
		"c")

	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doCreateTopic(vc)
	}, CheckTopicNameTwoArgs)
}

func doCreateTopic(vc *cmdutils.VerbCmd) error {
	admin := cmdutils.NewPulsarClient()
	err = admin.Topics().Create(*topic, partitions)
	if err == nil {
		vc.Command.Printf("Create topic %s with %d partitions successfully\n", topic.String(), partitions)
	}

	return err
}
```

As shown above: in the `CreateTopicCmd` you need to include the following information:

- Description
    - CommandUsedFor // used to describe the usage scenario of the command
    - CommandPermission // describe the permission information for the command
    - CommandExamples // describe all usage examples for this command
    - CommandOutput // describe the output of the command, including correct output and error output

- Args info (pulsarctl supports the following two forms of command:)

    - `Pulsarctl command sub-command name-arg-1 name-arg-2 ...`
        - For this scenario, pulsarctl provides the following functions:
            - SetRunFuncWithMultiNameArgs // set multiple name args
            - SetRunFuncWithNameArg // set a single name arg
    - `pulsarctl command sub-command --flag xxx --flag yyy ...`
        - For this scenario, pulsarctl provides the following functions:
            - SetRunFunc // no need to set name args
        - When there is a flag to specify, you can create a structure of `TopicData` under the `pulsar/data.go` file, and add the list of parameters you need in the structure.
        
        > To ensure the correctness of the request parameters, use the json tag to format the parameter name. If you want the parameter list to be specified in the yaml file, add a yaml tag to it.
        
        You can use the following example to specify a specific parameter list for the command:
        
        ```
        vc.FlagSetGroup.InFlagSet("Topic", func(set *pflag.FlagSet) {
            set.BoolVarP(&partition, 
            "partitioned-topic", 
            "p", 
            false,
            "Get the partitioned topic stats")
            
            // other flags    
        })
        ```
        
        > If the parameter is a parameter that must be requested, mark it with `cobra.MarkFlagRequired(set, "flag-name")`
        
In `doCreateTopic`, write the logic to create a topic, as follows:

1) Create pulsar client

Pulsar currently supports three versions of the Api interface. You can refer to the version information currently used by the command in the Pulsar broker to select the specific version number. To this end, pulsarctl provides the following functions:

- NewPulsarClient() // default value, use the V2 version
- NewPulsarClientWithApiVersion(version pulsar.ApiVersion) // custom version 

2) Call interface function

In `pulsar/admin.go`, abstract the interface of `Client`. Take the example command as an example. The code is as follows:

```
type Client interface {
	Topics() Topics
}
```

To unify the topic related sub commands, you can create a `topic.go` file in the pulsar directory as follows:

```
type Topics interface {
	Create(TopicName, int) error
}

// define topics struct and implement the Topics interface
func (t *topics) Create(topic TopicName, partitions int) error {
	endpoint := t.client.endpoint(t.basePath, topic.GetRestPath(), "partitions")
	if partitions == 0 {
		endpoint = t.client.endpoint(t.basePath, topic.GetRestPath())
	}
	return t.client.put(endpoint, partitions, nil)
}
```

Depending on the type of request, pulsarctl encapsulates the following request methods:

- put
- get
- delete
- post

different request types can be selected depending on the situation.

After completing the above call, for the information output to the command terminal, pulsarctl provides the following print forms:

- PrintJson // print in json format
- PrintError // when there is an error, output according to the packaged error message

- If you need to print in tabular form, you can use `tablewriter` lib as follows:

    ```
    table := tablewriter.NewWriter(vc.Command.OutOrStdout())
    table.SetHeader([]string{"Pulsar Function Name"})
    
    for _, f := range functions {
        table.Append([]string{f})
    }
    
    table.Render()
    ```

- If the output information is a single-line text prompt, the details are as follows:

    ```
    vc.Command.Printf("Create topic %s with %d partitions successfully\n", topic.String(), partitions)
    ```

### Write `ctl/topic/topic.go`

```
func Command(flagGrouping *cmdutils.FlagGrouping) *cobra.Command {
	resourceCmd := cmdutils.NewResourceCmd(
		"topics",
		"Operations about topic(s)",
		"",
		"topic")

	commands := []func(*cmdutils.VerbCmd){
		CreateTopicCmd,
	}

	cmdutils.AddVerbCmds(flagGrouping, resourceCmd, commands...)

	return resourceCmd
}
```

As shown above, in `topic.go`, the main logic is to call `AddVerbCmds` to add the above prepared command.

After the above operation is completed, add the relevant command group in `main.go`.

### Write test

In order to facilitate the test, pulsarctl intercepts the error information. By default, when an error is encountered, the program will call `os.Exit(1)` to release the process resource.
You can redirect the error to the specified location output when writing the test, as follows:

```
var execError error
cmdutils.ExecErrorHandler = func(err error) {
	execError = err
}
```

When writing test case, you need to mock a test runner. If the test needs to use the associated helper function, name the file `test_help.go` and write the relevant code in the file.