
------------

# functions





### Usage

`$ functions`



------------

## <em>create</em>

>bdocs-tab:example Create a Pulsar Function in cluster mode with jar file

```bdocs-tab:example_shell
pulsarctl functions create
--tenant public
--namespace default
--name (the name of Pulsar Functions>)
--inputs test-input-topic
--output persistent://public/default/test-output-topic
--classname org.apache.pulsar.functions.api.examples.ExclamationFunction
--jar /examples/api-examples.jar
```

>bdocs-tab:example Create a Pulsar Function use function config yaml file

```bdocs-tab:example_shell
pulsarctl functions create
--function-config-file (the path of function config yaml file)
--jar (the path of user code jar)
```

>bdocs-tab:example Create a Pulsar Function in cluster mode with pkg URL

```bdocs-tab:example_shell
pulsarctl functions create
--tenant public
--namespace default
--name (the name of Pulsar Functions)
--inputs test-input-topic
--output persistent://public/default/test-output-topic
--classname org.apache.pulsar.functions.api.examples.ExclamationFunction
--jar file:/http: + /examples/api-examples.jar
```

>bdocs-tab:example Create a Pulsar Function in cluster mode with log topic

```bdocs-tab:example_shell
pulsarctl functions create
--log-topic persistent://public/default/test-log-topic
// Other function parameters
```

>bdocs-tab:example Create a Pulsar Function in cluster mode with dead letter topic

```bdocs-tab:example_shell
pulsarctl functions create
--dead-letter-topic persistent://public/default/test-dead-letter-topic
--max-message-retries 10
// Other function parameters
```

>bdocs-tab:example Create a Pulsar Function in cluster mode with auto ack

```bdocs-tab:example_shell
pulsarctl functions create
--auto-ack
// Other function parameters
```

>bdocs-tab:example Create a Pulsar Function in cluster mode with FQFN

```bdocs-tab:example_shell
pulsarctl functions create
--fqfn tenant/namespace/name eg:public/default/test-fqfn-function
// Other function parameters
```

>bdocs-tab:example Create a Pulsar Function in cluster mode with topics pattern

```bdocs-tab:example_shell
pulsarctl functions create
--topics-pattern persistent://tenant/ns/topicPattern*
// Other function parameters
```

>bdocs-tab:example Create a Pulsar Function in cluster mode with user config

```bdocs-tab:example_shell
pulsarctl functions create
--user-config "{"publishTopic":"publishTopic", "key":"pulsar"}"
// Other function parameters
```

>bdocs-tab:example Create a Pulsar Function in cluster mode with retain ordering

```bdocs-tab:example_shell
pulsarctl functions create
--retain-ordering
// Other function parameters
```

>bdocs-tab:example Create a Pulsar Function in cluster mode with custom schema for inputs topic

```bdocs-tab:example_shell
pulsarctl functions create
--custom-schema-inputs "{"topic-1":"schema.STRING", "topic-2":"schema.JSON"}"
// Other function parameters
```

>bdocs-tab:example Create a Pulsar Function in cluster mode with schema type for output topic

```bdocs-tab:example_shell
pulsarctl functions create
--schema-type schema.STRING
// Other function parameters
```

>bdocs-tab:example Create a Pulsar Function in cluster mode with parallelism

```bdocs-tab:example_shell
pulsarctl functions create
--parallelism 1
// Other function parameters
```

>bdocs-tab:example Create a Pulsar Function in cluster mode with resource

```bdocs-tab:example_shell
pulsarctl functions create
--ram 5656565656
--disk 8080808080808080
--cpu 5.0
// Other function parameters
```

>bdocs-tab:example Create a Pulsar Function in cluster mode with window functions

```bdocs-tab:example_shell
pulsarctl functions create
--window-length-count 10
--window-length-duration-ms 1000
--sliding-interval-count 3
--sliding-interval-duration-ms 1000
// Other function parameters
```


### Used For
 

 This command is used for creating a new Pulsar Function in cluster mode. 

  
### Required Permission
 

 This command requires super-user permissions. 

  
### Output
 
 //normal output 

 Created (the name of a Pulsar Function) successfully 

  

 

### Usage

`$ create`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
auto-ack |  | true | Whether or not the framework acknowledges messages automatically 
classname |  |  | The class name of a Pulsar Function 
cpu |  | 0 | The cpu in cores that need to be allocated per function instance(applicable only to docker runtime) 
custom-schema-inputs |  |  | The map of input topics to Schema class names (as a JSON string) 
custom-serde-inputs |  |  | The map of input topics to SerDe class names (as a JSON string) 
dead-letter-topic |  |  | The topic where messages that are not processed successfully are sent to 
disk |  | 0 | The disk in bytes that need to be allocated per function instance(applicable only to docker runtime) 
fqfn |  |  | The Fully Qualified Function Name (FQFN) for the function 
function-config-file |  |  | The path to a YAML config file that specifies the configuration of a Pulsar Function 
go |  |  | Path to the main Go executable binary for the function (if the function is written in Go) 
inputs | i |  | The input topic or topics (multiple topics can be specified as a comma-separated list) of a Pulsar Function 
jar |  |  | Path to the JAR file for the function (if the function is written in Java) It also supports URL path [http/https/file (file protocol assumes that file already exists on worker host)] from which worker can download the package. 
log-topic |  |  | The topic to which the logs of a Pulsar Function are produced 
max-message-retries |  | 0 | How many times should we try to process a message before giving up 
name |  |  | The name of a Pulsar Function 
namespace |  |  | The namespace of a Pulsar Function 
output | o |  | The output topic of a Pulsar Function (If none is specified, no output is written) 
output-serde-classname |  |  | The SerDe class to be used for messages output by the function 
parallelism |  | 0 | The parallelism factor of a Pulsar Function (i.e. the number of function instances to run) 
processing-guarantees |  |  | The processing guarantees (aka delivery semantics) applied to the function 
py |  |  | Path to the main Python file/Python Wheel file for the function (if the function is written in Python) 
ram |  | 0 | The ram in bytes that need to be allocated per function instance(applicable only to process/docker runtime) 
retain-ordering |  | false | Function consumes and processes messages in order 
schema-type | t |  | The builtin schema type or custom schema class name to be used for messages output by the function 
sliding-interval-count |  | 0 | The number of messages after which the window slides 
sliding-interval-duration-ms |  | 0 | The time duration after which the window slides 
subs-name |  |  | Pulsar source subscription name if user wants a specific subscription-name for input-topic consumer 
tenant |  |  | The tenant of a Pulsar Function 
timeout-ms |  | 0 | The message timeout in milliseconds 
topics-pattern |  |  | The topic pattern to consume from list of topics under a namespace that match the pattern. [--input] and [--topic-pattern] are mutually exclusive. Add SerDe class name for a pattern in --custom-serde-inputs (supported for java fun only) 
user-config |  |  | User-defined config key/values 
window-length-count |  | 0 | The number of messages per window 
window-length-duration-ms |  | 0 | The time duration of the window in milliseconds 



------------

## <em>delete</em>

>bdocs-tab:example Delete a Pulsar Function that is running on a Pulsar cluster

```bdocs-tab:example_shell
pulsarctl functions delete
--tenant public
--namespace default
--name (the name of Pulsar Functions)
```

>bdocs-tab:example Delete a Pulsar Function that is running on a Pulsar cluster with instance ID

```bdocs-tab:example_shell
pulsarctl functions delete
--tenant public
--namespace default
--name (the name of Pulsar Functions)
--instance-id 1
```

>bdocs-tab:example Delete a Pulsar Function that is running on a Pulsar cluster with FQFN

```bdocs-tab:example_shell
pulsarctl functions delete
--fqfn tenant/namespace/name [eg: public/default/ExampleFunctions]
```


### Used For
 

 This command is used for delete a Pulsar Function that is running on a Pulsar cluster. 

  
### Required Permission
 

 This command requires super-user permissions. 

  
### Output
 
 //normal output 

 Deleted <the name of a Pulsar Function> successfully 

  
 //You must specify a name for the Pulsar Functions or a FQFN, please check the --name args 

 [✖]  you must specify a name for the function or a Fully Qualified Function Name (FQFN) 

  
 //The name of Pulsar Functions doesn't exist, please check the --name args 

 [✖]  code: 404 reason: Function <your function name> doesn't exist 

  

 

### Usage

`$ delete`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
fqfn |  |  | The Fully Qualified Function Name (FQFN) for the function 
name |  |  | The name of a Pulsar Function 
namespace |  |  | The namespace of a Pulsar Function 
tenant |  |  | The tenant of a Pulsar Function 



------------

## <em>get</em>

>bdocs-tab:example Fetch information about a Pulsar Function

```bdocs-tab:example_shell
pulsarctl functions get
--tenant public
--namespace default
--name (the name of Pulsar Function)
```

>bdocs-tab:example Fetch information about a Pulsar Function with FQFN

```bdocs-tab:example_shell
pulsarctl functions get
--fqfn tenant/namespace/name [eg: public/default/ExampleFunctions]
```


### Used For
 

 Fetch information about a Pulsar Function 

  
### Required Permission
 

 This command requires super-user permissions. 

  
### Output
 
 //normal output 

 { 

 "tenant": "public", 

 "namespace": "default", 

 "name": "test-functions", 

 "className": "org.apache.pulsar.functions.api.examples.ExclamationFunction", 

 "inputSpecs": { 

 "persistent://public/default/test-topic-1": { 

 "isRegexPattern": false 

 } 

 }, 

 "output": "persistent://public/default/test-topic-2", 

 "processingGuarantees": "ATLEAST_ONCE", 

 "retainOrdering": false, 

 "userConfig": {}, 

 "runtime": "JAVA", 

 "autoAck": true, 

 "parallelism": 1, 

 "resources": { 

 "cpu": 1.0, 

 "ram": 1073741824, 

 "disk": 10737418240 

 }, 

 "cleanupSubscription": true 

 } 

  
 //You must specify a name for the Pulsar Functions or a FQFN, please check the --name args 

 [✖]  you must specify a name for the function or a Fully Qualified Function Name (FQFN) 

  
 //The name of Pulsar Functions doesn't exist, please check the --name args 

 [✖]  code: 404 reason: Function (your function name) doesn't exist 

  

 

### Usage

`$ get`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
fqfn |  |  | The Fully Qualified Function Name (FQFN) for the function 
name |  |  | The name of a Pulsar Function 
namespace |  |  | The namespace of a Pulsar Function 
tenant |  |  | The tenant of a Pulsar Function 



------------

## <em>list</em>

>bdocs-tab:example List all Pulsar Functions running under a specific tenant and namespace

```bdocs-tab:example_shell
pulsarctl functions list
--tenant public
--namespace default
```


### Used For
 

 List all Pulsar Functions running under a specific tenant and namespace. 

  
### Required Permission
 

 This command requires super-user permissions. 

  
### Output
 
 //normal output 

 +--------------------+ 

 |   Function Name    | 

 +--------------------+ 

 | test_function_name | 

 +--------------------+ 

  

 

### Usage

`$ list`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
namespace |  |  | The namespace of a Pulsar Function 
tenant |  |  | The tenant of a Pulsar Function 



------------

## <em>putstate</em>

>bdocs-tab:example Put a key/(string value) pair to the state associated with a Pulsar Function

```bdocs-tab:example_shell
pulsarctl functions putstate
--tenant public
--namespace default
--name (the name of Pulsar Function)
(key name) - (string value)
```

>bdocs-tab:example Put a key/(file path) pair to the state associated with a Pulsar Function

```bdocs-tab:example_shell
pulsarctl functions putstate
--tenant public
--namespace default
--name (the name of Pulsar Function)
(key name) = (file path)
```

>bdocs-tab:example Put a key/value pair to the state associated with a Pulsar Function with FQFN

```bdocs-tab:example_shell
pulsarctl functions putstate
--fqfn tenant/namespace/name [eg: public/default/ExampleFunctions]
(key name) - (string value)
```


### Used For
 

 Put a key/value pair to the state associated with a Pulsar Function. 

  
### Required Permission
 

 This command requires namespace function permissions. 

  
### Output
 
 //normal output 

 Put state (the function state) successfully 

  
 //You must specify a name for the Pulsar Functions or a FQFN, please check the --name args 

 [✖]  you must specify a name for the function or a Fully Qualified Function Name (FQFN) 

  
 //The name of Pulsar Functions doesn't exist, please check the `--name` arg 

 [✖]  code: 404 reason: Function (your function name) doesn't exist 

  
 //The state key and state value not specified, please check your input format 

 [✖]  need to specified the state key and state value 

  
 //The format of the input is incorrect, please check. 

 [✖]  error input format 

  

 

### Usage

`$ putstate`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
fqfn |  |  | The Fully Qualified Function Name (FQFN) for the function 
name |  |  | The name of a Pulsar Function 
namespace |  |  | The namespace of a Pulsar Function 
tenant |  |  | The tenant of a Pulsar Function 



------------

## <em>querystate</em>

>bdocs-tab:example Fetch the current state associated with a Pulsar Function

```bdocs-tab:example_shell
pulsarctl functions querystate
--tenant public
--namespace default
--name (the name of Pulsar Function)
--key (the name of key)
--watch
```

>bdocs-tab:example Fetch a key/value pair from the state associated with a Pulsar Function with FQFN

```bdocs-tab:example_shell
pulsarctl functions querystate
--fqfn tenant/namespace/name [eg: public/default/ExampleFunctions]
--key (the name of key)
--watch
```

>bdocs-tab:example Fetch a key/value pair from the state associated with a Pulsar Function

```bdocs-tab:example_shell
pulsarctl functions querystate
--tenant public
--namespace default
--name (the name of Pulsar Function)
--key (the name of key)
```


### Used For
 

 Fetch a key/value pair from the state associated with a Pulsar Function. 

  
### Required Permission
 

 This command requires namespace function permissions. 

  
### Output
 
 //normal output 

 { 

 "key": "pulsar", 

 "stringValue": "hello", 

 "byteValue": null, 

 "numberValue": 0, 

 "version": 6 

 } 

  
 //You must specify a name for the Pulsar Functions or a FQFN, please check the --name args 

 [✖]  you must specify a name for the function or a Fully Qualified Function Name (FQFN) 

  
 //The name of Pulsar Functions doesn't exist, please check the --name args 

 [✖]  code: 404 reason: Function <your function name> doesn't exist 

  
 //key <the name of key> doesn't exist, please check --key args 

 error: key <the name of key> doesn't exist 

  

 

### Usage

`$ querystate`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
fqfn |  |  | The Fully Qualified Function Name (FQFN) for the function 
key | k |  | key 
name |  |  | The name of a Pulsar Function 
namespace |  |  | The namespace of a Pulsar Function 
tenant |  |  | The tenant of a Pulsar Function 
watch | w | false | Watch for changes in the value associated with a key for a Pulsar Function 



------------

## <em>restart</em>

>bdocs-tab:example Restart function instance

```bdocs-tab:example_shell
pulsarctl functions restart
--tenant public
--namespace default
--name (the name of Pulsar Function)
```

>bdocs-tab:example Restart function instance with instance ID

```bdocs-tab:example_shell
pulsarctl functions restart
--tenant public
--namespace default
--name (the name of Pulsar Function)
--instance-id 1
```

>bdocs-tab:example Restart function instance with FQFN

```bdocs-tab:example_shell
pulsarctl functions restart
--fqfn tenant/namespace/name [eg: public/default/ExampleFunctions]
```


### Used For
 

 This command is used for restarting function instance. 

  
### Required Permission
 

 This command requires super-user permissions. 

  
### Output
 
 //normal output 

 Restarted (the name of a Pulsar Function) successfully 

  
 //You must specify a name for the Pulsar Functions or a FQFN, please check the --name args 

 [✖]  you must specify a name for the function or a Fully Qualified Function Name (FQFN) 

  
 //The name of Pulsar Functions doesn't exist, please check the --name args 

 [✖]  code: 404 reason: Function (your function name) doesn't exist 

  
 //Used an instanceID that does not exist or other impermissible actions 

 [✖]  code: 400 reason: Operation not permitted 

  

 

### Usage

`$ restart`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
fqfn |  |  | The Fully Qualified Function Name (FQFN) for the function 
instance-id |  |  | The function instanceId (restart all instances if instance-id is not provided) 
name |  |  | The name of a Pulsar Function 
namespace |  |  | The namespace of a Pulsar Function 
tenant |  |  | The tenant of a Pulsar Function 



------------

## <em>start</em>

>bdocs-tab:example Starts a stopped function instance

```bdocs-tab:example_shell
pulsarctl functions start
--tenant public
--namespace default
--name (the name of Pulsar Function)
```

>bdocs-tab:example Starts a stopped function instance with instance ID

```bdocs-tab:example_shell
pulsarctl functions start
--tenant public
--namespace default
--name (the name of Pulsar Function)
--instance-id 1
```

>bdocs-tab:example Starts a stopped function instance with FQFN

```bdocs-tab:example_shell
pulsarctl functions start
--fqfn tenant/namespace/name [eg: public/default/ExampleFunctions]
```


### Used For
 

 This command is used for starting a stopped function instance. 

  
### Required Permission
 

 This command requires super-user permissions. 

  
### Output
 
 //normal output 

 Started <the name of a Pulsar Function> successfully 

  
 //You must specify a name for the Pulsar Functions or a FQFN, please check the --name args 

 [✖]  you must specify a name for the function or a Fully Qualified Function Name (FQFN) 

  
 //The name of Pulsar Functions doesn't exist, please check the --name args 

 [✖]  code: 404 reason: Function <your function name> doesn't exist 

  
 //Used an instanceID that does not exist or other impermissible actions 

 [✖]  code: 400 reason: Operation not permitted 

  

 

### Usage

`$ start`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
fqfn |  |  | The Fully Qualified Function Name (FQFN) for the function 
instance-id |  |  | The function instanceId (start all instances if instance-id is not provided) 
name |  |  | The name of a Pulsar Function 
namespace |  |  | The namespace of a Pulsar Function 
tenant |  |  | The tenant of a Pulsar Function 



------------

## <em>stats</em>

>bdocs-tab:example Get the current stats of a Pulsar Function

```bdocs-tab:example_shell
pulsarctl functions stats
--tenant public
--namespace default
--name (the name of Pulsar Function)
```

>bdocs-tab:example Get the current stats of a Pulsar Function with FQFN

```bdocs-tab:example_shell
pulsarctl functions stats
--fqfn tenant/namespace/name [eg: public/default/ExampleFunctions]
```


### Used For
 

 Get the current stats of a Pulsar Function. 

  
### Required Permission
 

 This command requires namespace function permissions. 

  
### Output
 
 //normal output 

 { 

 "receivedTotal": 0, 

 "processedSuccessfullyTotal": 0, 

 "systemExceptionsTotal": 0, 

 "userExceptionsTotal": 0, 

 "avgProcessLatency": 0, 

 "lastInvocation": 0, 

 "oneMin": { 

 "receivedTotal": 0, 

 "processedSuccessfullyTotal": 0, 

 "systemExceptionsTotal": 0, 

 "userExceptionsTotal": 0, 

 "avgProcessLatency": 0 

 }, 

 "instances": [ 

 { 

 "receivedTotal": 0, 

 "processedSuccessfullyTotal": 0, 

 "systemExceptionsTotal": 0, 

 "userExceptionsTotal": 0, 

 "avgProcessLatency": 0, 

 "instanceId": 0, 

 "metrics": { 

 "oneMin": { 

 "receivedTotal": 0, 

 "processedSuccessfullyTotal": 0, 

 "systemExceptionsTotal": 0, 

 "userExceptionsTotal": 0, 

 "avgProcessLatency": 0 

 }, 

 "lastInvocation": 0, 

 "userMetrics": {}, 

 "receivedTotal": 0, 

 "processedSuccessfullyTotal": 0, 

 "systemExceptionsTotal": 0, 

 "userExceptionsTotal": 0, 

 "avgProcessLatency": 0 

 } 

 } 

 ], 

 "instanceId": 0, 

 "metrics": { 

 "oneMin": { 

 "receivedTotal": 0, 

 "processedSuccessfullyTotal": 0, 

 "systemExceptionsTotal": 0, 

 "userExceptionsTotal": 0, 

 "avgProcessLatency": 0 

 }, 

 "lastInvocation": 0, 

 "userMetrics": null, 

 "receivedTotal": 0, 

 "processedSuccessfullyTotal": 0, 

 "systemExceptionsTotal": 0, 

 "userExceptionsTotal": 0, 

 "avgProcessLatency": 0 

 } 

 } 

  
 //You must specify a name for the Pulsar Functions or a FQFN, please check the --name args 

 [✖]  you must specify a name for the function or a Fully Qualified Function Name (FQFN) 

  
 //The name of Pulsar Functions doesn't exist, please check the --name args 

 [✖]  code: 404 reason: Function (your function name) doesn't exist 

  

 

### Usage

`$ stats`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
fqfn |  |  | The Fully Qualified Function Name (FQFN) for the function 
instance-id |  |  | The function instanceId (Get-stats of all instances if instance-id is not provided) 
name |  |  | The name of a Pulsar Function 
namespace |  |  | The namespace of a Pulsar Function 
tenant |  |  | The tenant of a Pulsar Function 



------------

## <em>status</em>

>bdocs-tab:example Check the current status of a Pulsar Function

```bdocs-tab:example_shell
pulsarctl functions status
--tenant public
--namespace default
--name (the name of Pulsar Function)
```

>bdocs-tab:example Check the current status of a Pulsar Function with FQFN

```bdocs-tab:example_shell
pulsarctl functions status
--fqfn tenant/namespace/name [eg: public/default/ExampleFunctions]
```


### Used For
 

 Check the current status of a Pulsar Function. 

  
### Required Permission
 

 This command requires namespace function permissions. 

  
### Output
 
 //normal output 

 { 

 "numInstances": 1, 

 "numRunning": 1, 

 "instances": [ 

 { 

 "instanceId": 0, 

 "status": { 

 "running": true, 

 "error": "", 

 "numRestarts": 0, 

 "numReceived": 0, 

 "numSuccessfullyProcessed": 0, 

 "numUserExceptions": 0, 

 "latestUserExceptions": [], 

 "numSystemExceptions": 0, 

 "latestSystemExceptions": [], 

 "averageLatency": 0, 

 "lastInvocationTime": 0, 

 "workerId": "c-standalone-fw-127.0.0.1-8080" 

 } 

 } 

 ] 

 } 

  
 //You must specify a name for the Pulsar Functions or a FQFN, please check the --name args 

 [✖]  you must specify a name for the function or a Fully Qualified Function Name (FQFN) 

  
 //The name of Pulsar Functions doesn't exist, please check the --name args 

 [✖]  code: 404 reason: Function (your function name) doesn't exist 

  
 //Used an instanceID that does not exist or other impermissible actions 

 [✖]  code: 400 reason: Operation not permitted 

  

 

### Usage

`$ status`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
fqfn |  |  | The Fully Qualified Function Name (FQFN) for the function 
instance-id |  |  | The function instanceId (Get-status of all instances if instance-id is not provided) 
name |  |  | The name of a Pulsar Function 
namespace |  |  | The namespace of a Pulsar Function 
tenant |  |  | The tenant of a Pulsar Function 



------------

## <em>stop</em>

>bdocs-tab:example Stops function instance

```bdocs-tab:example_shell
pulsarctl functions stop
--tenant public
--namespace default
--name (the name of Pulsar Function)
```

>bdocs-tab:example Stops function instance with instance ID

```bdocs-tab:example_shell
pulsarctl functions stop
--tenant public
--namespace default
--name (the name of Pulsar Function)
--instance-id 1
```

>bdocs-tab:example Stops function instance with FQFN

```bdocs-tab:example_shell
pulsarctl functions stop
--fqfn tenant/namespace/name [eg: public/default/ExampleFunctions]
```


### Used For
 

 This command is used for stopping function instance. 

  
### Required Permission
 

 This command requires super-user permissions. 

  
### Output
 
 //normal output 

 Stopped (the name of a Pulsar Function) successfully 

  
 //You must specify a name for the Pulsar Functions or a FQFN, please check the --name args 

 [✖]  you must specify a name for the function or a Fully Qualified Function Name (FQFN) 

  
 //The name of Pulsar Functions doesn't exist, please check the --name args 

 [✖]  code: 404 reason: Function (your function name) doesn't exist 

  
 //Used an instanceID that does not exist or other impermissible actions 

 [✖]  code: 400 reason: Operation not permitted 

  

 

### Usage

`$ stop`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
fqfn |  |  | The Fully Qualified Function Name (FQFN) for the function 
instance-id |  |  | The function instanceId (stop all instances if instance-id is not provided) 
name |  |  | The name of a Pulsar Function 
namespace |  |  | The namespace of a Pulsar Function 
tenant |  |  | The tenant of a Pulsar Function 



------------

## <em>trigger</em>

>bdocs-tab:example Trigger the specified Pulsar Function with a supplied value

```bdocs-tab:example_shell
pulsarctl functions trigger
--tenant public
--namespace default
--name (the name of Pulsar Function)
--topic (the name of input topic)
--trigger-value "hello pulsar"
```

>bdocs-tab:example Trigger the specified Pulsar Function with a supplied value

```bdocs-tab:example_shell
pulsarctl functions trigger
--fqfn tenant/namespace/name [eg: public/default/ExampleFunctions]
--topic (the name of input topic)
--trigger-value "hello pulsar"
```

>bdocs-tab:example Trigger the specified Pulsar Function with a supplied value

```bdocs-tab:example_shell
pulsarctl functions trigger
--tenant public
--namespace default
--name (the name of Pulsar Function)
--topic (the name of input topic)
--trigger-file (the path of trigger file)
```


### Used For
 

 Trigger the specified Pulsar Function with a supplied value. 

  
### Required Permission
 

 This command requires namespace function permissions. 

  
### Output
 
 //You must specify a name for the Pulsar Functions or a FQFN, please check the --name args 

 [✖]  you must specify a name for the function or a Fully Qualified Function Name (FQFN) 

  
 //The name of Pulsar Functions doesn't exist, please check the --name args 

 [✖]  code: 404 reason: Function (your function name) doesn't exist 

  
 //Used an instanceID that does not exist or other impermissible actions 

 [✖]  code: 400 reason: Operation not permitted 

  
 //Function in trigger function has unidentified topic 

 [✖]  code: 400 reason: Function in trigger function has unidentified topic 

  
 //Request Timed Out 

 [✖]  code: 408 reason: Request Timed Out 

  

 

### Usage

`$ trigger`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
fqfn |  |  | The Fully Qualified Function Name (FQFN) for the function 
name |  |  | The name of a Pulsar Function 
namespace |  |  | The namespace of a Pulsar Function 
tenant |  |  | The tenant of a Pulsar Function 
topic |  |  | The specific topic name that the function consumes from that you want to inject the data to 
trigger-file |  |  | The path to the file that contains the data with which you want to trigger the function 
trigger-value |  |  | The value with which you want to trigger the function 



------------

## <em>update</em>

>bdocs-tab:example Change the output topic of a Pulsar Function

```bdocs-tab:example_shell
pulsarctl functions update
--tenant public
--namespace default
--name update-function
--output test-output-topic
```

>bdocs-tab:example Update a Pulsar Function using a function config yaml file

```bdocs-tab:example_shell
pulsarctl functions update
--function-config-file (the path of function config yaml file)
--jar (the path of user code jar)
```

>bdocs-tab:example Change the log topic of a Pulsar Function

```bdocs-tab:example_shell
pulsarctl functions update
--log-topic persistent://public/default/test-log-topic
// Other function parameters
```

>bdocs-tab:example Change the dead letter topic of a Pulsar Function

```bdocs-tab:example_shell
pulsarctl functions update
--dead-letter-topic persistent://public/default/test-dead-letter-topic
--max-message-retries 10
// Other function parameters
```

>bdocs-tab:example Update the user configs of a Pulsar Function

```bdocs-tab:example_shell
pulsarctl functions update
--user-config "{"publishTopic":"publishTopic", "key":"pulsar"}"
// Other function parameters
```

>bdocs-tab:example Change the schemas of the input topics for a Pulsar Function

```bdocs-tab:example_shell
pulsarctl functions update
--custom-schema-inputs "{"topic-1":"schema.STRING", "topic-2":"schema.JSON"}"
// Other function parameters
```

>bdocs-tab:example Change the schema type of the input topic for a Pulsar Function

```bdocs-tab:example_shell
pulsarctl functions update
--schema-type schema.STRING
// Other function parameters
```

>bdocs-tab:example Change the parallelism of a Pulsar Function

```bdocs-tab:example_shell
pulsarctl functions update
--parallelism 1
// Other function parameters
```

>bdocs-tab:example Change the resource usage for a Pulsar Function

```bdocs-tab:example_shell
pulsarctl functions update
--ram 5656565656
--disk 8080808080808080
--cpu 5.0
// Other function parameters
```

>bdocs-tab:example Update the window configurations for a Pulsar Function

```bdocs-tab:example_shell
pulsarctl functions update
--window-length-count 10
--window-length-duration-ms 1000
--sliding-interval-count 3
--sliding-interval-duration-ms 1000
// Other function parameters
```


### Used For
 

 Update a Pulsar Function that has been deployed to a Pulsar cluster. 

  
### Required Permission
 

 This command requires super-user permissions. 

  
### Output
 
 //normal output 

 Updated (the name of a Pulsar Function) successfully 

  
 //Update contains no change 

 [✖]  code: 400 reason: Update contains no change 

  
 //The name of Pulsar Functions doesn't exist, please check the --name args 

 [✖]  code: 404 reason: Function (your function name) doesn't exist 

  

 

### Usage

`$ update`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
classname |  |  | The class name of a Pulsar Function 
cpu |  | 0 | The cpu in cores that need to be allocated per function instance(applicable only to docker runtime) 
custom-schema-inputs |  |  | The map of input topics to Schema class names (as a JSON string) 
custom-serde-inputs |  |  | The map of input topics to SerDe class names (as a JSON string) 
dead-letter-topic |  |  | The topic where messages that are not processed successfully are sent to 
disk |  | 0 | The disk in bytes that need to be allocated per function instance(applicable only to docker runtime) 
fqfn |  |  | The Fully Qualified Function Name (FQFN) for the function 
function-config-file |  |  | The path to a YAML config file that specifies the configuration of a Pulsar Function 
go |  |  | Path to the main Go executable binary for the function (if the function is written in Go) 
inputs |  |  | The input topic or topics (multiple topics can be specified as a comma-separated list) of a Pulsar Function 
jar |  |  | Path to the JAR file for the function (if the function is written in Java). It also supports URL path [http/https/file (file protocol assumes that file already exists on worker host)] from which worker can download the package. 
log-topic |  |  | The topic to which the logs of a Pulsar Function are produced 
max-message-retries |  | 0 | How many times should we try to process a message before giving up 
name |  |  | The name of a Pulsar Function 
namespace |  |  | The namespace of a Pulsar Function 
output | o |  | The output topic of a Pulsar Function (If none is specified, no output is written) 
output-serde-classname |  |  | The SerDe class to be used for messages output by the function 
parallelism |  | 0 | The parallelism factor of a Pulsar Function (i.e. the number of function instances to run) 
py |  |  | Path to the main Python file/Python Wheel file for the function (if the function is written in Python) 
ram |  | 0 | The ram in bytes that need to be allocated per function instance(applicable only to process/docker runtime) 
schema-type | t |  | The builtin schema type or custom schema class name to be used for messages output by the function 
sliding-interval-count |  | 0 | The number of messages after which the window slides 
sliding-interval-duration-ms |  | 0 | The time duration after which the window slides 
tenant |  |  | The tenant of a Pulsar Function 
timeout-ms |  | 0 | The message timeout in milliseconds 
topics-pattern |  |  | The topic pattern to consume from list of topics under a namespace that match the pattern. [--input] and [--topic-pattern] are mutually exclusive. Add SerDe class name for a pattern in --custom-serde-inputs (supported for java fun only) 
update-auth-data |  | false | Whether or not to update the auth data 
user-config |  |  | User-defined config key/values 
window-length-count |  | 0 | The number of messages per window 
window-length-duration-ms |  | 0 | The time duration of the window in milliseconds 




