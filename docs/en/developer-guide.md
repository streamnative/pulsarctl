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

The `pulsarctl` is a command-line tool written in the Go language that helps administrators and users manage clusters, tenants, namespaces, topics, schemas, sources, sinks, functions, and so on.

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

- `pkg` is used to store pulsarctl related libraries. There are four subdirectories as follows:
    - `auth` is used to store encryption related code.
    - `cmdutils` has a simple wrapper for cobra.
    - `ctl` is used to store pulsarctl related commands.
    - `pulsar` is a public package of pulsarctl.
- `test` is used to store resources related to a test.
- `site` is used to store website related code of pulsarctl, which is convenient for users to view and quickly locate the usage and precautions of related commands.
- `docs` is used to store pulsarctl document.

> **NOTE:**
> * To avoid circular references, where `auth` and `cmdutils` are two separate packages, the two packages `ctl` and `pulsar` are not referenced and will not be referenced to each other. 
> * `pulsar` is a public package that references `auth` but does not reference `ctl` and `cmdutils`. 
> * `ctl`, as the core package implementing pulsarctl, will reference `auth`, `cmdutils` and `pulsar` packages; however, it is not referenced by other packages.

## Add a new command

### Usage

```bash
pulsarctl [commands] [sub commands] [flags]
```
The contents of `[command]` are consistent with the file directory under the `ctl` directory. 

When you create a new command, create a new folder in the `ctl` directory, give that folder a name which is the same to the command name, 
and create a `command-name.go` file in the `pulsar` directory to write the interface function associated with the command.

`[sub commands]` belongs to `[commands]`. 

If you want to add a subcommand to `[commands]`, create a `sub-command-name.go` file in the command directory and add relevant code logic. 

After finishing writing code, add the relevant test code and make sure it covers your code logic.

This example illustrates how to add a new command to pulsarctl with `pulsarctl topics create (topic name) 0`.

1. Under the `ctl` directory, create a `topic` folder. 

```bash
mkdir topic
```

2. Under the `topic` directory, create the following files:

* `create.go`, which is used to write contents related to create topics.
* `topic.go`, which is used to store all commands related to topics. 

```bash
cd pkg/ctl/topic
touch create.go topic.go
```

### Write create.go

In general, each command file consists of two functions, `CreateTopicCmd` and `doCreateTopic`. Details are as follows:
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

As shown above, you need to include the following information for `CreateTopicCmd`:

- Description
    - CommandUsedFor // describe the usage scenario of the command.
    - CommandPermission // describe the permission information of the command.
    - CommandExamples // describe all usage examples of the command.
    - CommandOutput // describe the output of the command, including correct output and error output.

- Args information (pulsarctl supports the following two commands:)

    - `Pulsarctl command sub-command name-arg-1 name-arg-2 ...`
        - For this scenario, pulsarctl provides the following functions:
            - SetRunFuncWithMultiNameArgs // set multiple name args.
            - SetRunFuncWithNameArg // set a single name arg.
    - `pulsarctl command sub-command --flag xxx --flag yyy ...`
        - For this scenario, pulsarctl provides the following functions:
            - SetRunFunc // no need to set name args.
        - If you need to specify a flag, you can create a structure of `TopicData` under the `pulsar/data.go` file and add a list of parameters.
        
        > NOTE: To ensure the parameter is correct, use the JSON tag to format the parameter name. 
        If you want to specify the parameter in the YAML file, add a YAML tag.
        
         To specify a specific parameter list for a command, you can use the following command.
        
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
        
        > NOTE: If a parameter is required, tag it with `cobra.MarkFlagRequired(set, "flag-name")`.
        
In `doCreateTopic`, write the logic to create a topic as follows:

1) Create a Pulsar client.

Currently, Pulsar supports three versions of the API interface. 

To get the version number, you can refer to the version information used by a command in a Pulsar broker. 
Consequently, pulsarctl provides the following functions:

- NewPulsarClient() // default value, use the V2 version
- NewPulsarClientWithApiVersion(version pulsar.ApiVersion) // custom version 

2) Call an interface function.

In `pulsar/admin.go`, abstract the interface of `Client`. Here takes `topic command` as an example.

```
type Client interface {
	Topics() Topics
}
```

To unify the topic related sub commands, you can create a `topic.go` file in the `pulsar` directory as follows:

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

Depending on the type of request, pulsarctl encapsulates the following request methods. 
You can choose a desired one based on your needs.

- put
- get
- delete
- post

For information printed to a terminal, `pulsarctl` provides the following ways:

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

### Write ctl/topic/topic.go

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

As shown above, in `topic.go`, the main logic is to call `AddVerbCmds` to add commands prepared before.

After finishing the steps above, add the relevant command group in `main.go`.

### Write a test

To simplify testing, pulsarctl intercepts the error information. By default, when an error is triggered, 
the program calls `os.Exit(1)` to release process resource.
You can specify an output location for an error when writing a test as follows:

```
var execError error
cmdutils.ExecErrorHandler = func(err error) {
	execError = err
}
```

When writing a test case, you need to mock a test runner. If the test needs to use an associated function, 
name the file as `test_help.go` and write relevant code in this file.

## Implementing Output Formats
The tool has a built-in framework for producing various output formats for a given command.

### Output Configuration
The output configuration (`cmdutils.OutputConfig`) defines possible output formats to be used by a command.
The available output formats are `text`, `json`, and `yaml`.  The default output format is `text`.

A flagset is defined for the output configuration, as seen in the help text for commands which support it:
```
Output flags:
  -o, --output string   The output format (text,json,yaml) (default "text")
``` 

### Writing a command with output format support
A command may opt-in to output formatting during the initialization stage of a `VerbCmd`.  Most commands
set a description, "run" function to be executed, and optionally define a flagset.  To enable the output flagset, 
do this last:
```
vc.EnableOutputFlagSet()
```

### Writing Output
When a command is ready to emit output, it calls `vc.OutputConfig.WriteOutput(...)` (where `vc` is 
the `VerbCmd` associated with the command).

The principle of output formatting is that the command _negotiates_ with the framework to supply the type of 
output that the user requested.  The core interface is `cmdutils.OutputNegotiable`, which yields 
a `cmdutils.OutputWritable` for the requested format (`text`,`json`,`yaml`).  The `OutputWritable` then emits 
the content to the given `io.Writer` (i.e. to `stdout`).

Since most of the relevant commands work with Go structs containing JSON tags, the framework makes it easy to 
use Go structs to provide both a JSON and YAML representation, and to serve as the default text representation. 
Meanwhile the framework makes it easy to develop a prettier text representation, such as a table layout.

These conveniences are provided with a built-in implementation of `cmdutils.OutputNegotiable` called `cmdutils.OutputContent`.
The `OutputContent` type implements the JSON and YAML formats using standard Go marshaling, and provides 
a convenient way to generate a text representation using a format string (see `WithText`) or using a function 
(see `WithTextFunc`).

### Caveats
Note that some commands emit JSON text for both the `text` and the `json` format.  Users should specify `-o json` for 
scripting purposes, because a given command's `text` representation may change at any time.

### Examples
Some commands which make non-trivial use of the output format framework:
- [`pulsarctl clusters list`](../../pkg/ctl/cluster/list.go) # uses `TextFunc` to produce a table representation
- [`pulsarctl schemas get`](../../pkg/ctl/schemas/get.go) # uses `TextFunc` to produce a text representation
- [`pulsarctl topic bundle-range`](../../pkg/ctl/topic/bundle_range.go) # uses `Text` with format strings
