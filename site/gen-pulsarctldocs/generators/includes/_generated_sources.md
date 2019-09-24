
------------

# sources





### Usage

`$ sources`



------------

## <em>create</em>

>bdocs-tab:example Create a Pulsar Source in cluster mode

```bdocs-tab:example_shell
pulsarctl sources create
--tenant public
--namespace default
--name (the name of Pulsar Sources)
--destination-topic-name kafka-topic
--classname org.apache.pulsar.io.kafka.KafkaBytesSource
--archive pulsar-io-kafka-2.4.0.nar
--source-config-file conf/kafkaSourceConfig.yaml
--parallelism 1
```

>bdocs-tab:example Create a Pulsar Source in cluster mode with pkg URL

```bdocs-tab:example_shell
pulsarctl source create
--tenant public
--namespace default
--name (the name of Pulsar Source)
--destination-topic-name kafka-topic
--classname org.apache.pulsar.io.kafka.KafkaBytesSource
--archive file://(or http://) + /examples/api-examples.nar
```

>bdocs-tab:example Create a Pulsar Source in cluster mode with schema type

```bdocs-tab:example_shell
pulsarctl source create
--schema-type schema.STRING
// Other source parameters
```

>bdocs-tab:example Create a Pulsar Source in cluster mode with parallelism

```bdocs-tab:example_shell
pulsarctl source create
--parallelism 1
// Other source parameters
```

>bdocs-tab:example Create a Pulsar Source in cluster mode with resource

```bdocs-tab:example_shell
pulsarctl source create
--ram 5656565656
--disk 8080808080808080
--cpu 5.0
// Other source parameters
```

>bdocs-tab:example Create a Pulsar Source in cluster mode with source config

```bdocs-tab:example_shell
pulsarctl source create
--source-config "{"publishTopic":"publishTopic", "key":"pulsar"}"
// Other source parameters
```

>bdocs-tab:example Create a Pulsar Source in cluster mode with processing guarantees

```bdocs-tab:example_shell
pulsarctl source create
--processing-guarantees EFFECTIVELY_ONCE
// Other source parameters
```


### Used For
 

 Submit a Pulsar IO source connector to run in a Pulsar cluster. 

  
### Required Permission
 

 This command requires namespace function permissions. 

  
### Output
 
 //normal output 

 Created (the name of a Pulsar Sources) successfully 

  
 //source archive not specified, please check --archive arg 

 [✖]  Source archive not specified 

  
 //Cannot specify both archive and source-type, please check --archive and --source-type args 

 [✖]  Cannot specify both archive and source-type 

  

 

### Usage

`$ create`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
archive | a |  | The path to the NAR archive for the Source. It also supports url-path [http/https/file (file protocol assumes that file already exists on worker host)] from which worker can download the package 
classname |  |  | The source's class name if archive is file-url-path (file://) 
cpu |  | 0 | The CPU (in cores) that needs to be allocated per source instance (applicable only to Docker runtime) 
deserialization-classname |  |  | The SerDe classname for the source 
destination-topic-name |  |  | The Pulsar topic to which data is sent 
disk |  | 0 | The disk (in bytes) that need to be allocated per source instance (applicable only to Docker runtime) 
name |  |  | The source's name 
namespace |  |  | The source's namespace 
parallelism |  | 0 | The source's parallelism factor (i.e. the number of source instances to run) 
processing-guarantees |  |  | The processing guarantees (aka delivery semantics) applied to the source 
ram |  | 0 | The RAM (in bytes) that need to be allocated per source instance (applicable only to the process and Docker runtimes) 
schema-type |  |  | The schema type (either a builtin schema like 'avro', 'json', etc.. or custom Schema class name to be used to encode messages emitted from the source 
source-config |  |  | Source config key/values 
source-config-file |  |  | he path to a YAML config file specifying the  
source-type | t |  | The source's connector provider 
tenant |  |  | The source's tenant 



------------

## <em>delete</em>

>bdocs-tab:example Delete a Pulsar IO source connector

```bdocs-tab:example_shell
pulsarctl source delete
--tenant public
--namespace default
--name (the name of Pulsar Source)
```


### Used For
 

 This command is used for deleting a Pulsar IO source connector. 

  
### Required Permission
 

 This command requires namespace function permissions. 

  
### Output
 
 //normal output 

 Deleted (the name of a Pulsar Source) successfully 

  
 //source doesn't exist 

 code: 404 reason: Source (the name of a Pulsar Source) doesn't exist 

  

 

### Usage

`$ delete`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
name |  |  | The source's name 
namespace |  |  | The source's namespace 
tenant |  |  | The source's tenant 



------------

## <em>get</em>

>bdocs-tab:example Gets the information about a Pulsar IO source connector

```bdocs-tab:example_shell
pulsarctl source get
--tenant public
--namespace default
--name (the name of Pulsar Source)
```


### Used For
 

 Gets the information about a Pulsar IO source connector 

  
### Required Permission
 

 This command requires namespace function permissions. 

  
### Output
 
 //normal output 

 { 

 "tenant": "public", 

 "namespace": "default", 

 "name": "kafka", 

 "className": "org.apache.pulsar.io.kafka.KafkaBytesSource", 

 "topicName": "my-topic", 

 "configs": { 

 "bootstrapServers": "pulsar-kafka:9092", 

 "groupId": "test-pulsar-io1", 

 "topic": "my-topic", 

 "sessionTimeoutMs": "10000", 

 "autoCommitEnabled": "false" 

 }, 

 "parallelism": 1, 

 "processingGuarantees": "ATLEAST_ONCE" 

 } 

  

  
 //source doesn't exist 

 code: 404 reason: Source (the name of a Pulsar Source) doesn't exist 

  

 

### Usage

`$ get`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
name |  |  | The source's name 
namespace |  |  | The source's namespace 
tenant |  |  | The source's tenant 



------------

## <em>list</em>

>bdocs-tab:example List all running Pulsar IO source connectors

```bdocs-tab:example_shell
pulsarctl source list
--tenant public
--namespace default
```


### Used For
 

 List all running Pulsar IO source connectors 

  
### Required Permission
 

 This command requires namespace function permissions. 

  
### Output
 
 //normal output 

 +--------------------+ 

 |   Source Name    | 

 +--------------------+ 

 | test_source_name | 

 +--------------------+ 

  

 

### Usage

`$ list`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
namespace |  |  | The source's namespace 
tenant |  |  | The source's tenant 



------------

## <em>restart</em>

>bdocs-tab:example Restart source instance

```bdocs-tab:example_shell
pulsarctl source restart
--tenant public
--namespace default
--name (the name of Pulsar Source)
```

>bdocs-tab:example Restart source instance with instance ID

```bdocs-tab:example_shell
pulsarctl source restart
--tenant public
--namespace default
--name (the name of Pulsar Source)
--instance-id 1
```


### Used For
 

 Restart source instance. 

  
### Required Permission
 

 This command requires namespace function permissions. 

  
### Output
 
 //normal output 

 Restarted (the name of a Pulsar Source) successfully 

  
 //source doesn't exist 

 code: 404 reason: Source (the name of a Pulsar Source) doesn't exist 

  

 

### Usage

`$ restart`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
instance-id |  |  | The source instanceId (stop all instances if instance-id is not provided) 
name |  |  | The source's name 
namespace |  |  | The source's namespace 
tenant |  |  | The source's tenant 



------------

## <em>start</em>

>bdocs-tab:example Start source instance

```bdocs-tab:example_shell
pulsarctl source start
--tenant public
--namespace default
--name (the name of Pulsar Source)
```

>bdocs-tab:example Starts a stopped source instance with instance ID

```bdocs-tab:example_shell
pulsarctl source start
--tenant public
--namespace default
--name (the name of Pulsar Source)
--instance-id 1
```


### Used For
 

 This command is used for starting a stopped source instance. 

  
### Required Permission
 

 This command requires namespace function permissions. 

  
### Output
 
 //normal output 

 Started (the name of a Pulsar Source) successfully 

  
 //source doesn't exist 

 code: 404 reason: Source (the name of a Pulsar Source) doesn't exist 

  

 

### Usage

`$ start`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
instance-id |  |  | The source instanceId (stop all instances if instance-id is not provided) 
name |  |  | The source's name 
namespace |  |  | The source's namespace 
tenant |  |  | The source's tenant 



------------

## <em>status</em>

>bdocs-tab:example Check the current status of a Pulsar Source

```bdocs-tab:example_shell
pulsarctl source status
--tenant public
--namespace default
--name (the name of Pulsar Source)
```


### Used For
 

 Check the current status of a Pulsar Source. 

  
### Required Permission
 

 This command requires namespace function permissions. 

  
### Output
 
 //normal output 

 { 

 "numInstances" : 1, 

 "numRunning" : 1, 

 "instances" : [ { 

 "instanceId" : 0, 

 "status" : { 

 "running" : true, 

 "error" : "", 

 "numRestarts" : 0, 

 "numReceivedFromSource" : 0, 

 "numSystemExceptions" : 0, 

 "latestSystemExceptions" : [ ], 

 "numSourceExceptions" : 0, 

 "latestSourceExceptions" : [ ], 

 "numWritten" : 0, 

 "lastReceivedTime" : 0, 

 "workerId" : "c-standalone-fw-7e0cf1b3bf9d-8080" 

 } 

 } ] 

 } 

  
 //Update contains no change 

 [✖]  code: 400 reason: Update contains no change 

  
 //The name of Pulsar Source doesn't exist, please check the --name args 

 [✖]  code: 404 reason: Source (your source name) doesn't exist 

  

 

### Usage

`$ status`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
instance-id |  |  | The source instanceId (stop all instances if instance-id is not provided) 
name |  |  | The source's name 
namespace |  |  | The source's namespace 
tenant |  |  | The source's tenant 



------------

## <em>stop</em>

>bdocs-tab:example Stops source instance

```bdocs-tab:example_shell
pulsarctl source stop
--tenant public
--namespace default
--name (the name of Pulsar Source)
```

>bdocs-tab:example Stops source instance with instance ID

```bdocs-tab:example_shell
pulsarctl source stop
--tenant public
--namespace default
--name (the name of Pulsar Source)
--instance-id 1
```


### Used For
 

 This command is used for stopping source instance. 

  
### Required Permission
 

 This command requires namespace function permissions. 

  
### Output
 
 //normal output 

 Stopped (the name of a Pulsar Source) successfully 

  
 //source doesn't exist 

 code: 404 reason: Source (the name of a Pulsar Source) doesn't exist 

  

 

### Usage

`$ stop`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
instance-id |  |  | The source instanceId (stop all instances if instance-id is not provided) 
name |  |  | The source's name 
namespace |  |  | The source's namespace 
tenant |  |  | The source's tenant 



------------

## <em>update</em>

>bdocs-tab:example Update a Pulsar IO source connector

```bdocs-tab:example_shell
pulsarctl source update
--tenant public
--namespace default
--name update-source
--archive pulsar-io-kafka-2.4.0.nar
--classname org.apache.pulsar.io.kafka.KafkaBytesSource
--destination-topic-name my-topic
--cpu 2
```

>bdocs-tab:example Update a Pulsar IO source connector with source config

```bdocs-tab:example_shell
pulsarctl source create
--source-config "{"publishTopic":"publishTopic", "key":"pulsar"}"
// Other source parameters
```

>bdocs-tab:example Update a Pulsar IO source connector with resource

```bdocs-tab:example_shell
pulsarctl source create
--ram 5656565656
--disk 8080808080808080
--cpu 5.0
// Other source parameters
```

>bdocs-tab:example Update a Pulsar IO source connector with parallelism

```bdocs-tab:example_shell
pulsarctl source create
--parallelism 1
// Other source parameters
```

>bdocs-tab:example Update a Pulsar IO source connector with schema type

```bdocs-tab:example_shell
pulsarctl source create
--schema-type schema.STRING
// Other source parameters
```


### Used For
 

 Update a Pulsar IO source connector. 

  
### Required Permission
 

 This command requires namespace function permissions. 

  
### Output
 
 //normal output 

 Updated (the name of a Pulsar Source) successfully 

  
 //source doesn't exist 

 code: 404 reason: Source (the name of a Pulsar Source) doesn't exist 

  

 

### Usage

`$ update`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
archive | a |  | The path to the NAR archive for the Source. It also supports url-path [http/https/file (file protocol assumes that file already exists on worker host)] from which worker can download the package 
className |  |  | The source's class name if archive is file-url-path (file://) 
cpu |  | 0 | The CPU (in cores) that needs to be allocated per source instance (applicable only to Docker runtime) 
deserialization-classname |  |  | The SerDe classname for the source 
destination-topic-name |  |  | The Pulsar topic to which data is sent 
disk |  | 0 | The disk (in bytes) that need to be allocated per source instance (applicable only to Docker runtime) 
name |  |  | The source's name 
namespace |  |  | The source's namespace 
parallelism |  | 0 | The source's parallelism factor (i.e. the number of source instances to run) 
processing-guarantees |  |  | The processing guarantees (aka delivery semantics) applied to the source 
ram |  | 0 | The RAM (in bytes) that need to be allocated per source instance (applicable only to the process and Docker runtimes) 
schema-type |  |  | The schema type (either a builtin schema like 'avro', 'json', etc.. or custom Schema class name to be used to encode messages emitted from the source 
source-config |  |  | Source config key/values 
source-config-file |  |  | he path to a YAML config file specifying the  
source-type | t |  | The source's connector provider 
tenant |  |  | The source's tenant 




