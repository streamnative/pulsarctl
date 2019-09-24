
------------

# sinks





### Usage

`$ sinks`



------------

## <em>create</em>

>bdocs-tab:example Create a Pulsar Sink in cluster mode

```bdocs-tab:example_shell
pulsarctl sink create
--tenant public
--namespace default
--name (the name of Pulsar Sink)
--inputs test-jdbc
--archive connectors/pulsar-io-jdbc-2.4.0.nar
--sink-config-file connectors/mysql-jdbc-sink.yaml
--parallelism 1
```

>bdocs-tab:example Create a Pulsar Sink in cluster mode with pkg URL

```bdocs-tab:example_shell
pulsarctl sink create
--tenant public
--namespace default
--name (the name of Pulsar Sink)
--inputs test-jdbc
--archive file:/http: + connectors/pulsar-io-jdbc-2.4.0.nar
```

>bdocs-tab:example Create a Pulsar Sink in cluster mode with schema type

```bdocs-tab:example_shell
pulsarctl sink create
--schema-type schema.STRING
// Other sink parameters
```

>bdocs-tab:example Create a Pulsar Sink in cluster mode with parallelism

```bdocs-tab:example_shell
pulsarctl sink create
--parallelism 1
// Other sink parameters
```

>bdocs-tab:example Create a Pulsar Sink in cluster mode with resource

```bdocs-tab:example_shell
pulsarctl sink create
--ram 5656565656
--disk 8080808080808080
--cpu 5.0
// Other sink parameters
```

>bdocs-tab:example Create a Pulsar Sink in cluster mode with sink config

```bdocs-tab:example_shell
pulsarctl sink create
--sink-config "{"publishTopic":"publishTopic", "key":"pulsar"}"
// Other sink parameters
```

>bdocs-tab:example Create a Pulsar Sink in cluster mode with processing guarantees

```bdocs-tab:example_shell
pulsarctl sink create
--processing-guarantees EFFECTIVELY_ONCE
// Other sink parameters
```


### Used For
 

 Create a Pulsar IO sink connector to run in a Pulsar cluster. 

  
### Required Permission
 

 This command requires namespace function permissions. 

  
### Output
 
 //normal output 

 Created (the name of a Pulsar Sinks) successfully 

  
 //sink archive not specified, please check --archive arg 

 [✖]  Sink archive not specified 

  
 //Cannot specify both archive and sink-type, please check --archive and --sink-type args 

 [✖]  Cannot specify both archive and sink-type 

  

 

### Usage

`$ create`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
archive |  |  | Path to the archive file for the sink. It also supports url-path [http/https/file (file protocol assumes that file already exists on worker host)] from which worker can download the package. 
auto-ack |  | false | Whether or not the framework will automatically acknowledge messages 
classname |  |  | The sink's class name if archive is file-url-path (file://) 
cpu |  | 0 | The CPU (in cores) that needs to be allocated per sink instance (applicable only to Docker runtime) 
custom-schema-inputs |  |  | The map of input topics to Schema types or class names (as a JSON string) 
custom-serde-inputs |  |  | The map of input topics to SerDe class names (as a JSON string) 
disk |  | 0 | The disk (in bytes) that need to be allocated per sink instance (applicable only to Docker runtime) 
inputs | i |  | The sink's input topic or topics (multiple topics can be specified as a comma-separated list) 
name |  |  | The sink's name 
namespace |  |  | The sink's namespace 
parallelism |  | 0 | The sink's parallelism factor (i.e. the number of sink instances to run) 
processing-guarantees |  |  | The processing guarantees (aka delivery semantics) applied to the sink 
ram |  | 0 | The RAM (in bytes) that need to be allocated per sink instance (applicable only to the process and Docker runtimes) 
retain-ordering |  | false | Sink consumes and sinks messages in order 
sink-config |  |  | User defined configs key/values 
sink-config-file |  |  | The path to a YAML config file specifying the sink's configuration 
sink-type | t |  | The sink's connector provider 
subs-name |  |  | Pulsar source subscription name if user wants a specific subscription-name for input-topic consumer 
tenant |  |  | The sink's tenant 
timeout-ms |  | 0 | The message timeout in milliseconds 
topics-pattern |  |  | TopicsPattern to consume from list of topics under a namespace that match the pattern. [--input] and [--topicsPattern] are mutually exclusive. Add SerDe class name for a pattern in --customSerdeInputs  (supported for java fun only) 



------------

## <em>delete</em>

>bdocs-tab:example Delete a Pulsar IO sink connector

```bdocs-tab:example_shell
pulsarctl sink delete
--tenant public
--namespace default
--name (the name of Pulsar Sink)
```


### Used For
 

 This command is used for deleting a Pulsar IO sink connector. 

  
### Required Permission
 

 This command requires namespace function permissions. 

  
### Output
 
 //normal output 

 Deleted (the name of a Pulsar Sink) successfully 

  
 //sink doesn't exist 

 code: 404 reason: Sink (the name of a Pulsar Sink) doesn't exist 

  

 

### Usage

`$ delete`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
name |  |  | The sink's name 
namespace |  |  | The sink's namespace 
tenant |  |  | The sink's tenant 



------------

## <em>get</em>

>bdocs-tab:example Get the information about a Pulsar IO sink connector

```bdocs-tab:example_shell
pulsarctl sink get
--tenant public
--namespace default
--name (the name of Pulsar Sink)
```


### Used For
 

 Get the information about a Pulsar IO sink connector 

  
### Required Permission
 

 This command requires namespace function permissions. 

  
### Output
 
 //normal output 

 { 

 "tenant": "public", 

 "namespace": "default", 

 "name": "mysql-jdbc-sink", 

 "className": "org.apache.pulsar.io.jdbc.JdbcAutoSchemaSink", 

 "inputSpecs": { 

 "test-jdbc": { 

 "isRegexPattern": false 

 } 

 }, 

 "configs": { 

 "password": "jdbc", 

 "jdbcUrl": "jdbc:mysql://127.0.0.1:3306/test_jdbc", 

 "userName": "root", 

 "tableName": "test_jdbc" 

 }, 

 "parallelism": 1, 

 "processingGuarantees": "ATLEAST_ONCE", 

 "retainOrdering": false, 

 "autoAck": true 

 } 

  
 //sink doesn't exist 

 code: 404 reason: Sink (the name of a Pulsar Sink) doesn't exist 

  

 

### Usage

`$ get`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
name |  |  | The sink's name 
namespace |  |  | The sink's namespace 
tenant |  |  | The sink's tenant 



------------

## <em>list</em>

>bdocs-tab:example Get the list of all the running Pulsar IO sink connectors

```bdocs-tab:example_shell
pulsarctl sink list
--tenant public
--namespace default
```


### Used For
 

 Get the list of all the running Pulsar IO sink connectors 

  
### Required Permission
 

 This command requires namespace function permissions. 

  
### Output
 
 //normal output 

 +--------------------+ 

 |   Sink Name    | 

 +--------------------+ 

 | test_sink_name | 

 +--------------------+ 

  

 

### Usage

`$ list`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
namespace |  |  | The sink's namespace 
tenant |  |  | The sink's tenant 



------------

## <em>restart</em>

>bdocs-tab:example Restart sink instance

```bdocs-tab:example_shell
pulsarctl sink restart
--tenant public
--namespace default
--name (the name of Pulsar Sink)
```

>bdocs-tab:example Restart sink instance with instance ID

```bdocs-tab:example_shell
pulsarctl sink restart
--tenant public
--namespace default
--name (the name of Pulsar Sink)
--instance-id 1
```


### Used For
 

 Restart sink instance. 

  
### Required Permission
 

 This command requires namespace function permissions. 

  
### Output
 
 //normal output 

 Restarted (the name of a Pulsar Sink) successfully 

  
 //sink doesn't exist 

 code: 404 reason: Sink (the name of a Pulsar Sink) doesn't exist 

  

 

### Usage

`$ restart`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
instance-id |  |  | The sink instanceId (stop all instances if instance-id is not provided) 
name |  |  | The sink's name 
namespace |  |  | The sink's namespace 
tenant |  |  | The sink's tenant 



------------

## <em>start</em>

>bdocs-tab:example Start sink instance

```bdocs-tab:example_shell
pulsarctl sink start
--tenant public
--namespace default
--name (the name of Pulsar Sink)
```

>bdocs-tab:example Starts a stopped sink instance with instance ID

```bdocs-tab:example_shell
pulsarctl sink start
--tenant public
--namespace default
--name (the name of Pulsar Sink)
--instance-id 1
```


### Used For
 

 This command is used for starting a stopped sink instance. 

  
### Required Permission
 

 This command requires namespace function permissions. 

  
### Output
 
 //normal output 

 Started (the name of a Pulsar Sink) successfully 

  
 //sink doesn't exist 

 code: 404 reason: Sink (the name of a Pulsar Sink) doesn't exist 

  

 

### Usage

`$ start`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
instance-id |  |  | The sink instanceId (stop all instances if instance-id is not provided) 
name |  |  | The sink's name 
namespace |  |  | The sink's namespace 
tenant |  |  | The sink's tenant 



------------

## <em>status</em>

>bdocs-tab:example Get the current status of a Pulsar Sink

```bdocs-tab:example_shell
pulsarctl sink status
--tenant public
--namespace default
--name (the name of Pulsar Sink)
```


### Used For
 

 Get the current status of a Pulsar Sink. 

  
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

 "numReadFromPulsar" : 0, 

 "numSystemExceptions" : 0, 

 "latestSystemExceptions" : [ ], 

 "numSinkExceptions" : 0, 

 "latestSinkExceptions" : [ ], 

 "numWrittenToSink" : 0, 

 "lastReceivedTime" : 0, 

 "workerId" : "c-standalone-fw-tengdeMBP.lan-8080" 

 } 

 } ] 

 } 

  
 //Update contains no change 

 [✖]  code: 400 reason: Update contains no change 

  
 //The name of Pulsar Sink doesn't exist, please check the --name args 

 [✖]  code: 404 reason: Sink (your sink name) doesn't exist 

  

 

### Usage

`$ status`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
instance-id |  |  | The sink instanceId (stop all instances if instance-id is not provided) 
name |  |  | The sink's name 
namespace |  |  | The sink's namespace 
tenant |  |  | The sink's tenant 



------------

## <em>stop</em>

>bdocs-tab:example Stops function instance

```bdocs-tab:example_shell
pulsarctl sink stop
--tenant public
--namespace default
--name (the name of Pulsar Sink)
```

>bdocs-tab:example Stops function instance with instance ID

```bdocs-tab:example_shell
pulsarctl sink stop
--tenant public
--namespace default
--name (the name of Pulsar Sink)
--instance-id 1
```


### Used For
 

 This command is used for stopping sink instance. 

  
### Required Permission
 

 This command requires namespace function permissions. 

  
### Output
 
 //normal output 

 Stopped (the name of a Pulsar Sink) successfully 

  
 //sink doesn't exist 

 code: 404 reason: Sink (the name of a Pulsar Sink) doesn't exist 

  

 

### Usage

`$ stop`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
instance-id |  |  | The sink instanceId (stop all instances if instance-id is not provided) 
name |  |  | The sink's name 
namespace |  |  | The sink's namespace 
tenant |  |  | The sink's tenant 



------------

## <em>update</em>

>bdocs-tab:example Update a Pulsar IO sink connector

```bdocs-tab:example_shell
pulsarctl sink update
--tenant public
--namespace default
--name update-source
--archive pulsar-io-kafka-2.4.0.nar
--classname org.apache.pulsar.io.kafka.KafkaBytesSource
--destination-topic-name my-topic
--cpu 2
```

>bdocs-tab:example Update a Pulsar IO sink connector with sink config

```bdocs-tab:example_shell
pulsarctl sink create
--sink-config "{"publishTopic":"publishTopic", "key":"pulsar"}"
// Other sink parameters
```

>bdocs-tab:example Update a Pulsar IO sink connector with resource

```bdocs-tab:example_shell
pulsarctl sink create
--ram 5656565656
--disk 8080808080808080
--cpu 5.0
// Other sink parameters
```

>bdocs-tab:example Update a Pulsar IO sink connector with parallelism

```bdocs-tab:example_shell
pulsarctl sink create
--parallelism 1
// Other sink parameters
```

>bdocs-tab:example Update a Pulsar IO sink connector with schema type

```bdocs-tab:example_shell
pulsarctl sink create
--schema-type schema.STRING
// Other sink parameters
```


### Used For
 

 Update a Pulsar IO sink connector. 

  
### Required Permission
 

 This command requires namespace function permissions. 

  
### Output
 
 //normal output 

 Updated (the name of a Pulsar Sink) successfully 

  
 //sink doesn't exist 

 code: 404 reason: Sink (the name of a Pulsar Sink) doesn't exist 

  

 

### Usage

`$ update`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
archive |  |  | Path to the archive file for the sink. It also supports url-path [http/https/file (file protocol assumes that file already exists on worker host)] from which worker can download the package. 
auto-ack |  | false | Whether or not the framework will automatically acknowledge messages 
classname |  |  | The sink's class name if archive is file-url-path (file://) 
cpu |  | 0 | The CPU (in cores) that needs to be allocated per sink instance (applicable only to Docker runtime) 
custom-schema-inputs |  |  | The map of input topics to Schema types or class names (as a JSON string) 
custom-serde-inputs |  |  | The map of input topics to SerDe class names (as a JSON string) 
disk |  | 0 | The disk (in bytes) that need to be allocated per sink instance (applicable only to Docker runtime) 
inputs | i |  | The sink's input topic or topics (multiple topics can be specified as a comma-separated list) 
name |  |  | The sink's name 
namespace |  |  | The sink's namespace 
parallelism |  | 0 | The sink's parallelism factor (i.e. the number of sink instances to run) 
processing-guarantees |  |  | The processing guarantees (aka delivery semantics) applied to the sink 
ram |  | 0 | The RAM (in bytes) that need to be allocated per sink instance (applicable only to the process and Docker runtimes) 
retain-ordering |  | false | Sink consumes and sinks messages in order 
sink-config |  |  | User defined configs key/values 
sink-config-file |  |  | The path to a YAML config file specifying the sink's configuration 
sink-type | t |  | The sink's connector provider 
subs-name |  |  | Pulsar source subscription name if user wants a specific subscription-name for input-topic consumer 
tenant |  |  | The sink's tenant 
timeout-ms |  | 0 | The message timeout in milliseconds 
topics-pattern |  |  | TopicsPattern to consume from list of topics under a namespace that match the pattern. [--input] and [--topicsPattern] are mutually exclusive. Add SerDe class name for a pattern in --customSerdeInputs  (supported for java fun only) 




