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

# 如何添加一个新命令

pulsarctl 是使用 go 语言编写的一个命令行工具，可以帮助管理员和用户管理 clusters、tenants、namespaces、topics、schemas、source、sink、functions 等相关的命令。  

## 工程结构

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

- `pkg` 用来存放 pulsarctl 相关的 libraries, 有四个子目录, 详情如下：
    - `auth` 用来存储加密相关的代码
    - `cmdutils` 对 cobra 进行了简单封装，
    - `ctl` 用来存放 pulsarctl 相关的命令
    - `pulsar` 是 pulsarctl 的公共包
- `test` 用来存放和 test 相关的资源
- `site` 是 pulsarctl 的 website 相关的 code，方便用户查看和快速定位相关命令的使用和注意事项等
- `docs` 用来存放 pulsarctl 相关的文档内容

> 为了避免循环引用, 其中 `auth` 和 `cmdutils` 是两个独立的包, 不会引用 `ctl` 和 `pulsar` 这两个包, 彼此之间也不会相互引用。 
`pulsar` 作为公共的包，会引用到 `auth`, 但是不会引用到 `ctl` 和 `cmdutils`。 
`ctl` 作为实现 pulsarctl 的核心 pkg, 会同时引用到 `auth`, `cmdutils`, `pulsar` 这三个包, 但是其它包不会引用到它。

## 添加一个新命令

pulsarctl 使用的命令格式如下：

```bash
pulsarctl [commands] [sub commands] [flags]
```

其中 `[command]` 的内容与 `ctl` 下的文件目录保持一致，当你想要创建一个新的 command 时，请在 `ctl` 目录下创建一个新的文件夹，
并将该文件夹的名字命名为该 command 的名字。同时在 `pulsar` 的目录下创建一个 `command-name.go` 的文件，用来编写该 command 相关的接口函数。

`[sub commands]` 属于 `[commands]` 下的子命令，如果你想要在当前现有的 `[commands]` 下为其添加一个 sub command，
请在该 command 的目录下创建一个 `sub-command-name.go` 的文件，添加你相关的代码逻辑。在完成代码的编写之后，请为你的代码添加相关的测试代码，
确保该测试可以覆盖到你的代码逻辑。

下面，以 `pulsarctl topics create (topic name) 0` 为例，详细阐述如何快速为 pulsarctl 添加一个新的 command。

1. 在 ctl 目录下创建一个名字为 topic 的文件夹

```bash
mkdir topic
```

2. 在 topic 目录下创建名为 `create.go` 和 `topic.go` 的文件, 其中 `create.go` 用来编写 create topic 相关的内容，
`topic.go` 用来存放和 topic 相关的所有命令。

```bash
cd pkg/ctl/topic
touch create.go topic.go
```

### 编写 create.go

正常情况下，每一个 command 文件由两个函数构成，`CreateTopicCmd` 和 `doCreateTopic`，具体如下：

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

如上所示：在 `CreateTopicCmd` 中需要包含以下几部分信息：

- 描述信息
    - CommandUsedFor // 用来描述该命令的使用场景
    - CommandPermission // 描述该命令的权限信息
    - CommandExamples // 描述该命令的所有使用示例
    - CommandOutput // 描述该命令的输出信息，包括正确输出与错误输出

- 参数信息 (pulsarctl 支持如下两种形式的 command：)

    - `Pulsarctl command sub-command name-arg-1 name-arg-2 ...`
        - 针对该场景，pulsarctl 提供了如下函数：
            - SetRunFuncWithMultiNameArgs // 设置多个 name args
            - SetRunFuncWithNameArg // 设置单个 name arg
    - `pulsarctl command sub-command --flag xxx --flag yyy ...`
        - 针对该场景，pulsarctl 提供了如下函数：
            - SetRunFunc // 不设置 name args
        - 当有 flag 需要指定时，你可以在 `pulsar/data.go` 文件下，创建 `TopicData` 的结构体，在该结构体中添加你需要的参数列表。
        
        > 为了确保请求参数的正确性，请使用 json 标签格式化参数名字，如果你想要该参数列表在 yaml 文件中指定，请为其添加 yaml 标签。
        
        你可以使用如下示例，为该命令指定具体的参数列表：
        
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
        
        > 如果该参数为必须请求的参数，请使用 `cobra.MarkFlagRequired(set, "flag-name")` 进行标记
        
在 `doCreateTopic` 中，编写创建 topic 相关的逻辑，具体如下：

1) 创建 pulsar client

Pulsar 目前支持三种版本的 Api 接口, 你可参照该命令目前在 Pulsar broker 中使用的版本信息，选择具体的版本号，为此，pulsarctl 提供了如下函数：

- NewPulsarClient() // 默认情况，使用 V2 版本
- NewPulsarClientWithApiVersion(version pulsar.ApiVersion) // 自定义版本号

2) 调用接口函数

在 `pulsar/admin.go`, 抽象了 `Client` 的接口，以 topic command 为例，代码如下：

```
type Client interface {
	Topics() Topics
}
```

为了统一 topic 相关的 sub command，你可以在 pulsar 目录下创建 `topic.go` 文件，具体如下：

```
type Topics interface {
	Create(TopicName, int) error
}

// 定义 topics struct，并实现 Topics 接口
func (t *topics) Create(topic TopicName, partitions int) error {
	endpoint := t.client.endpoint(t.basePath, topic.GetRestPath(), "partitions")
	if partitions == 0 {
		endpoint = t.client.endpoint(t.basePath, topic.GetRestPath())
	}
	return t.client.put(endpoint, partitions, nil)
}
```

根据不同的请求类型，pulsarctl 封装了如下请求方法：

- put
- get
- delete
- post

可根据具体情况，选择使用不同的请求类型。

在完成上述调用后，对于输出到命令终端的信息，pulsarctl 提供了如下几种打印形式：

- PrintJson // 按照 json 格式打印
- PrintError // 当有错误时，按照封装好的错误信息输出

- 如果需要按照表格形式打印，可以使用 `tablewriter` lib，具体如下：

    ```
    table := tablewriter.NewWriter(vc.Command.OutOrStdout())
    table.SetHeader([]string{"Pulsar Function Name"})
    
    for _, f := range functions {
        table.Append([]string{f})
    }
    
    table.Render()
    ```

- 如果输出的信息是单行文本的提示信息，具体如下：

    ```
    vc.Command.Printf("Create topic %s with %d partitions successfully\n", topic.String(), partitions)
    ```

### 编写 `ctl/topic/topic.go`

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

如上所示，在 `topic.go` 中，主要逻辑就是调用 `AddVerbCmds` 将上述编写好的 command 添加进来。
在上述操作完成之后，在 `main.go` 中将相关的 command group 添加进去。

### 编写 test

为了方便测试，pulsarctl 对 error 信息进行了拦截处理，默认情况下，当遇到错误时，程序会调用 `os.Exit(1)` 将进程资源释放，
在编写测试时可以将错误重定向到指定位置输出，具体如下：

```
var execError error
cmdutils.ExecErrorHandler = func(err error) {
	execError = err
}
```

在进行测试时，你需要 mock 一个 test runner，如果该 test 有需要用到相关的辅助函数，请将该文件命名为 `test_help.go`，并在该文件中编写相关的代码。
