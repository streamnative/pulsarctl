
------------

# topics





### Usage

`$ topics`



------------

## <em>bundle-range</em>

>bdocs-tab:example Get namespace bundle range of a topic (topic-name)

```bdocs-tab:example_shell
pulsarctl topic bundle-range (topic-name)
```


### Used For
 

 This command is used for getting namespace bundle range of a topic (partition). 

  
### Required Permission
 

 This command requires super-user permissions. 

  
### Output
 
 //normal output 

 The bundle range of the topic (topic-name) is: (bundle-range) 

  
 //the topic name is not specified 

 [✖]  only one argument is allowed to be used as a name 

  
 //the topic name is not in the format of <tenant>/<namespace>/<topic> or <topic> 

 [✖]  Invalid short topic name '<topic-name>', it should be in the format of <tenant>/<namespace>/<topic> or <topic> 

  
 //the topic name is not in the format of <domain>://<tenant>/<namespace>/<topic> 

 [✖]  Invalid complete topic name '<topic-name>', it should be in the format of <domain>://<tenant>/<namespace>/<topic> 

  
 //the topic name is not in the format of <tenant>/<namespace>/<topic> 

 [✖]  Invalid topic name '<topic-name>', it should be in the format of<tenant>/<namespace>/<topic> 

  
 //the namespace name is not in the format of <tenant>/<namespace> 

 [✖]  The complete name of namespace is invalid. complete name : <namespace-complete-name> 

  
 //the tenant name and(or) namespace name is empty 

 [✖]  Invalid tenant or namespace. [<tenant>/<namespace>] 

  
 //the tenant name contains unsupported special chars. the alphanumeric (a-zA-Z0-9) and the special chars (-=:.%)  is allowed 

 [✖]  Tenant name include unsupported special chars. tenant : [<namespace>] 

  
 //the namespace name contains unsupported special chars. the  alphanumeric (a-zA-Z0-9) and the special chars (-=:.%) is allowed 

 [✖]  Namespace name include unsupported special chars. namespace : [<namespace>] 

  

 

### Usage

`$ bundle-range`




------------

## <em>create</em>

>bdocs-tab:example Create a non-partitioned topic (topic-name)

```bdocs-tab:example_shell
pulsarctl topics create (topic-name) 0
```

>bdocs-tab:example Create a partitioned topic (topic-name) with (partitions-num) partitions

```bdocs-tab:example_shell
pulsarctl topics create (topic-name) (partition-num)
```


### Used For
 

 This command is used for creating topic. 

  
### Required Permission
 

 This command requires namespace admin permissions. 

  
### Output
 
 //normal output 

 Create topic (topic-name) with (partition-num) partitions successfully 

  
 //the topic name and(or) the partitions is not specified 

 [✖]  need to specified the topic name and the partitions 

  
 //the topic has been created 

 [✖]  code: 409 reason: Partitioned topic already exists 

  
 //the topic name is not in the format of <tenant>/<namespace>/<topic> or <topic> 

 [✖]  Invalid short topic name '<topic-name>', it should be in the format of <tenant>/<namespace>/<topic> or <topic> 

  
 //the topic name is not in the format of <domain>://<tenant>/<namespace>/<topic> 

 [✖]  Invalid complete topic name '<topic-name>', it should be in the format of <domain>://<tenant>/<namespace>/<topic> 

  
 //the topic name is not in the format of <tenant>/<namespace>/<topic> 

 [✖]  Invalid topic name '<topic-name>', it should be in the format of<tenant>/<namespace>/<topic> 

  
 //the namespace name is not in the format of <tenant>/<namespace> 

 [✖]  The complete name of namespace is invalid. complete name : <namespace-complete-name> 

  
 //the tenant name and(or) namespace name is empty 

 [✖]  Invalid tenant or namespace. [<tenant>/<namespace>] 

  
 //the tenant name contains unsupported special chars. the alphanumeric (a-zA-Z0-9) and the special chars (-=:.%)  is allowed 

 [✖]  Tenant name include unsupported special chars. tenant : [<namespace>] 

  
 //the namespace name contains unsupported special chars. the  alphanumeric (a-zA-Z0-9) and the special chars (-=:.%) is allowed 

 [✖]  Namespace name include unsupported special chars. namespace : [<namespace>] 

  

 

### Usage

`$ create`




------------

## <em>delete</em>

>bdocs-tab:example Delete a partitioned topic (topic-name)

```bdocs-tab:example_shell
pulsarctl topics delete (topic-name)
```

>bdocs-tab:example Delete a non-partitioned topic (topic-name)

```bdocs-tab:example_shell
pulsarctl topics delete --non-partitioned <topic-name>
```


### Used For
 

 This command is used for deleting an existing topic. 

  
### Required Permission
 

 This command requires namespace admin permissions. 

  
### Output
 
 //normal output 

 Delete topic (topic-name) successfully 

  
 //the topic name is not specified 

 [✖]  only one argument is allowed to be used as a name 

  
 //the partitioned topic does not exist 

 [✖]  code: 404 reason: Partitioned topic does not exist 

  
 //the non-partitioned topic does not exist 

 [✖]  code: 404 reason: Topic not found 

  
 //the topic name is not in the format of <tenant>/<namespace>/<topic> or <topic> 

 [✖]  Invalid short topic name '<topic-name>', it should be in the format of <tenant>/<namespace>/<topic> or <topic> 

  
 //the topic name is not in the format of <domain>://<tenant>/<namespace>/<topic> 

 [✖]  Invalid complete topic name '<topic-name>', it should be in the format of <domain>://<tenant>/<namespace>/<topic> 

  
 //the topic name is not in the format of <tenant>/<namespace>/<topic> 

 [✖]  Invalid topic name '<topic-name>', it should be in the format of<tenant>/<namespace>/<topic> 

  
 //the namespace name is not in the format of <tenant>/<namespace> 

 [✖]  The complete name of namespace is invalid. complete name : <namespace-complete-name> 

  
 //the tenant name and(or) namespace name is empty 

 [✖]  Invalid tenant or namespace. [<tenant>/<namespace>] 

  
 //the tenant name contains unsupported special chars. the alphanumeric (a-zA-Z0-9) and the special chars (-=:.%)  is allowed 

 [✖]  Tenant name include unsupported special chars. tenant : [<namespace>] 

  
 //the namespace name contains unsupported special chars. the  alphanumeric (a-zA-Z0-9) and the special chars (-=:.%) is allowed 

 [✖]  Namespace name include unsupported special chars. namespace : [<namespace>] 

  

 

### Usage

`$ delete`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
delete-schema | d | false | Delete schema while deleting topic 
force | f | false | Close all producer/consumer/replicator and delete topic forcefully 
non-partitioned | n | false | Delete a non-partitioned topic 



------------

## <em>get</em>

>bdocs-tab:example Get hte metadata of an exist topic (topic-name) metadata

```bdocs-tab:example_shell
pulsarctl topics get (topic-name)
```


### Used For
 

 This command is used for getting the metadata of an exist topic. 

  
### Required Permission
 

 This command requires namespace admin permissions. 

  
### Output
 
 //normal output 

 { 

 "partitions": "(partitions)" 

 } 

  
 //the topic name is not specified 

 [✖]  only one argument is allowed to be used as a name 

  
 //the topic name is not in the format of <tenant>/<namespace>/<topic> or <topic> 

 [✖]  Invalid short topic name '<topic-name>', it should be in the format of <tenant>/<namespace>/<topic> or <topic> 

  
 //the topic name is not in the format of <domain>://<tenant>/<namespace>/<topic> 

 [✖]  Invalid complete topic name '<topic-name>', it should be in the format of <domain>://<tenant>/<namespace>/<topic> 

  
 //the topic name is not in the format of <tenant>/<namespace>/<topic> 

 [✖]  Invalid topic name '<topic-name>', it should be in the format of<tenant>/<namespace>/<topic> 

  
 //the namespace name is not in the format of <tenant>/<namespace> 

 [✖]  The complete name of namespace is invalid. complete name : <namespace-complete-name> 

  
 //the tenant name and(or) namespace name is empty 

 [✖]  Invalid tenant or namespace. [<tenant>/<namespace>] 

  
 //the tenant name contains unsupported special chars. the alphanumeric (a-zA-Z0-9) and the special chars (-=:.%)  is allowed 

 [✖]  Tenant name include unsupported special chars. tenant : [<namespace>] 

  
 //the namespace name contains unsupported special chars. the  alphanumeric (a-zA-Z0-9) and the special chars (-=:.%) is allowed 

 [✖]  Namespace name include unsupported special chars. namespace : [<namespace>] 

  

 

### Usage

`$ get`




------------

## <em>internal-stats</em>

>bdocs-tab:example Get internal stats for an existing non-partitioned-topic (topic-name)

```bdocs-tab:example_shell
pulsarctl topic internal-stats (topic-name)
```

>bdocs-tab:example Get internal stats for a partition of a partitioned topic

```bdocs-tab:example_shell
pulsarctl topic internal-stats --partition (partition) (topic-name)
```


### Used For
 

 This command is used for getting the internal stats for a non-partitioned topic or a partition of a partitioned topic. 

  
### Required Permission
 

 This command requires namespace admin permissions. 

  
### Output
 
 //normal output 

 { 

 "entriesAddedCounter": 0, 

 "numberOfEntries": 0, 

 "totalSize": 0, 

 "currentLedgerEntries": 0, 

 "currentLedgerSize": 0, 

 "lastLedgerCreatedTimestamp": "", 

 "lastLedgerCreationFailureTimestamp": "", 

 "waitingCursorsCount": 0, 

 "pendingAddEntriesCount": 0, 

 "lastConfirmedEntry": "", 

 "state": "", 

 "ledgers": [ 

 { 

 "ledgerId": 0, 

 "entries": 0, 

 "size": 0, 

 "offloaded": false 

 } 

 ], 

 "cursors": {} 

 } 

  
 //the topic name is not specified 

 [✖]  only one argument is allowed to be used as a name 

  
 //the specified topic is not exist or the specified topic is a partitioned topic 

 [✖]  code: 404 reason: Topic not found 

  
 //the topic name is not in the format of <tenant>/<namespace>/<topic> or <topic> 

 [✖]  Invalid short topic name '<topic-name>', it should be in the format of <tenant>/<namespace>/<topic> or <topic> 

  
 //the topic name is not in the format of <domain>://<tenant>/<namespace>/<topic> 

 [✖]  Invalid complete topic name '<topic-name>', it should be in the format of <domain>://<tenant>/<namespace>/<topic> 

  
 //the topic name is not in the format of <tenant>/<namespace>/<topic> 

 [✖]  Invalid topic name '<topic-name>', it should be in the format of<tenant>/<namespace>/<topic> 

  
 //the namespace name is not in the format of <tenant>/<namespace> 

 [✖]  The complete name of namespace is invalid. complete name : <namespace-complete-name> 

  
 //the tenant name and(or) namespace name is empty 

 [✖]  Invalid tenant or namespace. [<tenant>/<namespace>] 

  
 //the tenant name contains unsupported special chars. the alphanumeric (a-zA-Z0-9) and the special chars (-=:.%)  is allowed 

 [✖]  Tenant name include unsupported special chars. tenant : [<namespace>] 

  
 //the namespace name contains unsupported special chars. the  alphanumeric (a-zA-Z0-9) and the special chars (-=:.%) is allowed 

 [✖]  Namespace name include unsupported special chars. namespace : [<namespace>] 

  

 

### Usage

`$ internal-stats`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
partition | p | -1 | The partitioned topic index value 



------------

## <em>last-message-id</em>

>bdocs-tab:example Get the last message id of a topic (persistent-topic-name)

```bdocs-tab:example_shell
pulsarctl topic last-message-id (persistent-topic-name)
```

>bdocs-tab:example Get the last message id of a partition of a partitioned topic (topic-name)

```bdocs-tab:example_shell
pulsarctl topic last-message-id --partition (partition) (topic-name)
```


### Used For
 

 This command is used for getting the last message id of a topic (partition). 

  
### Required Permission
 

 This command requires tenant admin permissions. 

  
### Output
 
 //normal output 

 { 

 "LedgerId": 0, 

 "EntryId": 0, 

 "PartitionedIndex": 0 

 } 

  
 //the topic name is not specified 

 [✖]  only one argument is allowed to be used as a name 

  
 //the topic (persistent-topic-name) does not exist in the cluster 

 [✖]  code: 404 reason: Topic not found 

  
 //the topic (persistent-topic-name) does not a persistent topic 

 [✖]  code: 405 reason: GetLastMessageId on a non-persistent topic is not allowed 

  
 //the topic name is not in the format of <tenant>/<namespace>/<topic> or <topic> 

 [✖]  Invalid short topic name '<topic-name>', it should be in the format of <tenant>/<namespace>/<topic> or <topic> 

  
 //the topic name is not in the format of <domain>://<tenant>/<namespace>/<topic> 

 [✖]  Invalid complete topic name '<topic-name>', it should be in the format of <domain>://<tenant>/<namespace>/<topic> 

  
 //the topic name is not in the format of <tenant>/<namespace>/<topic> 

 [✖]  Invalid topic name '<topic-name>', it should be in the format of<tenant>/<namespace>/<topic> 

  
 //the namespace name is not in the format of <tenant>/<namespace> 

 [✖]  The complete name of namespace is invalid. complete name : <namespace-complete-name> 

  
 //the tenant name and(or) namespace name is empty 

 [✖]  Invalid tenant or namespace. [<tenant>/<namespace>] 

  
 //the tenant name contains unsupported special chars. the alphanumeric (a-zA-Z0-9) and the special chars (-=:.%)  is allowed 

 [✖]  Tenant name include unsupported special chars. tenant : [<namespace>] 

  
 //the namespace name contains unsupported special chars. the  alphanumeric (a-zA-Z0-9) and the special chars (-=:.%) is allowed 

 [✖]  Namespace name include unsupported special chars. namespace : [<namespace>] 

  

 

### Usage

`$ last-message-id`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
partition | p | -1 | The partitioned topic index value 



------------

## <em>list</em>

>bdocs-tab:example List all exist topics under the namespace(tenant/namespace)

```bdocs-tab:example_shell
pulsarctl topics list (tenant/namespace)
```


### Used For
 

 This command is used for listing all exist topics under the specified namespace. 

  
### Required Permission
 

 This command requires admin permissions. 

  
### Output
 
 //normal output 

 +----------------------------------------------------------+---------------+ 

 |                        TOPIC NAME                        | PARTITIONED ? | 

 +----------------------------------------------------------+---------------+ 

 +----------------------------------------------------------+---------------+ 

  
 //the namespace is not specified 

 [✖]  only one argument is allowed to be used as a name 

  
 //the tenant of the namespace is not exist 

 [✖]  code: 404 reason: Tenant does not exist 

  
 //the namespace is not exist 

 [✖]  code: 404 reason: Namespace does not exist 

  
 //the namespace name is not in the format of <tenant>/<namespace> 

 [✖]  The complete name of namespace is invalid. complete name : <namespace-complete-name> 

  
 //the tenant name and(or) namespace name is empty 

 [✖]  Invalid tenant or namespace. [<tenant>/<namespace>] 

  
 //the tenant name contains unsupported special chars. the alphanumeric (a-zA-Z0-9) and the special chars (-=:.%)  is allowed 

 [✖]  Tenant name include unsupported special chars. tenant : [<namespace>] 

  
 //the namespace name contains unsupported special chars. the  alphanumeric (a-zA-Z0-9) and the special chars (-=:.%) is allowed 

 [✖]  Namespace name include unsupported special chars. namespace : [<namespace>] 

  

 

### Usage

`$ list`




------------

## <em>lookup</em>

>bdocs-tab:example Lookup the owner broker of the topic (topic-name)

```bdocs-tab:example_shell
pulsarctl topic lookup (topic-name)
```


### Used For
 

 This command is used for looking up the owner broker of a topic. 

  
### Required Permission
 

 This command does not require permissions. 

  
### Output
 
 // 

 { 

 "brokerUlr": "", 

 "brokerUrlTls": "", 

 "httpUrl": "", 

 "httpUrlTls": "", 

 } 

  
 //the topic name is not specified 

 [✖]  only one argument is allowed to be used as a name 

  
 //the topic name is not in the format of <tenant>/<namespace>/<topic> or <topic> 

 [✖]  Invalid short topic name '<topic-name>', it should be in the format of <tenant>/<namespace>/<topic> or <topic> 

  
 //the topic name is not in the format of <domain>://<tenant>/<namespace>/<topic> 

 [✖]  Invalid complete topic name '<topic-name>', it should be in the format of <domain>://<tenant>/<namespace>/<topic> 

  
 //the topic name is not in the format of <tenant>/<namespace>/<topic> 

 [✖]  Invalid topic name '<topic-name>', it should be in the format of<tenant>/<namespace>/<topic> 

  
 //the namespace name is not in the format of <tenant>/<namespace> 

 [✖]  The complete name of namespace is invalid. complete name : <namespace-complete-name> 

  
 //the tenant name and(or) namespace name is empty 

 [✖]  Invalid tenant or namespace. [<tenant>/<namespace>] 

  
 //the tenant name contains unsupported special chars. the alphanumeric (a-zA-Z0-9) and the special chars (-=:.%)  is allowed 

 [✖]  Tenant name include unsupported special chars. tenant : [<namespace>] 

  
 //the namespace name contains unsupported special chars. the  alphanumeric (a-zA-Z0-9) and the special chars (-=:.%) is allowed 

 [✖]  Namespace name include unsupported special chars. namespace : [<namespace>] 

  

 

### Usage

`$ lookup`




------------

## <em>stats</em>

>bdocs-tab:example Get the non-partitioned topic (topic-name) stats

```bdocs-tab:example_shell
pulsarctl topic stats (topic-name)
```

>bdocs-tab:example Get the partitioned topic (topic-name) stats

```bdocs-tab:example_shell
pulsarctl topic stats --partitioned-topic (topic-name)
```

>bdocs-tab:example Get the partitioned topic (topic-name) stats and per partition stats

```bdocs-tab:example_shell
pulsarctl topic stats --partitioned-topic --per-partition (topic-name)
```


### Used For
 

 This command is used for getting the stats for an existing topic and its connected producers and consumers. (All the rates are computed over a 1 minute window and are relative the last completed 1 minute period) 

  
### Required Permission
 

 This command requires namespace admin permissions. 

  
### Output
 
 //Get the non-partitioned topic stats 

 { 

 "msgRateIn": 0, 

 "msgRateOut": 0, 

 "msgThroughputIn": 0, 

 "msgThroughputOut": 0, 

 "averageMsgSize": 0, 

 "storageSize": 0, 

 "publishers": [], 

 "subscriptions": {}, 

 "replication": {}, 

 "deduplicationStatus": "Disabled" 

 } 

  
 //Get the partitioned topic stats 

 { 

 "msgRateIn": 0, 

 "msgRateOut": 0, 

 "msgThroughputIn": 0, 

 "msgThroughputOut": 0, 

 "averageMsgSize": 0, 

 "storageSize": 0, 

 "publishers": [], 

 "subscriptions": {}, 

 "replication": {}, 

 "deduplicationStatus": "", 

 "metadata": { 

 "partitions": 1 

 }, 

 "partitions": {} 

 } 

  
 //Get the partitioned topic stats and per partition topic stats 

 { 

 "msgRateIn": 0, 

 "msgRateOut": 0, 

 "msgThroughputIn": 0, 

 "msgThroughputOut": 0, 

 "averageMsgSize": 0, 

 "storageSize": 0, 

 "publishers": [], 

 "subscriptions": {}, 

 "replication": {}, 

 "deduplicationStatus": "", 

 "metadata": { 

 "partitions": 1 

 }, 

 "partitions": { 

 "<topic-name>": { 

 "msgRateIn": 0, 

 "msgRateOut": 0, 

 "msgThroughputIn": 0, 

 "msgThroughputOut": 0, 

 "averageMsgSize": 0, 

 "storageSize": 0, 

 "publishers": [], 

 "subscriptions": {}, 

 "replication": {}, 

 "deduplicationStatus": "" 

 } 

 } 

 } 

  
 //the topic name is not specified 

 [✖]  only one argument is allowed to be used as a name 

  
 //the specified topic does not exist or the specified topic is a partitioned-topic and you don't specified --partitioned-topic or the specified topic is a non-partitioned topic and you specified --partitioned-topic 

 code: 404 reason: Topic not found 

  
 //the topic name is not in the format of <tenant>/<namespace>/<topic> or <topic> 

 [✖]  Invalid short topic name '<topic-name>', it should be in the format of <tenant>/<namespace>/<topic> or <topic> 

  
 //the topic name is not in the format of <domain>://<tenant>/<namespace>/<topic> 

 [✖]  Invalid complete topic name '<topic-name>', it should be in the format of <domain>://<tenant>/<namespace>/<topic> 

  
 //the topic name is not in the format of <tenant>/<namespace>/<topic> 

 [✖]  Invalid topic name '<topic-name>', it should be in the format of<tenant>/<namespace>/<topic> 

  
 //the namespace name is not in the format of <tenant>/<namespace> 

 [✖]  The complete name of namespace is invalid. complete name : <namespace-complete-name> 

  
 //the tenant name and(or) namespace name is empty 

 [✖]  Invalid tenant or namespace. [<tenant>/<namespace>] 

  
 //the tenant name contains unsupported special chars. the alphanumeric (a-zA-Z0-9) and the special chars (-=:.%)  is allowed 

 [✖]  Tenant name include unsupported special chars. tenant : [<namespace>] 

  
 //the namespace name contains unsupported special chars. the  alphanumeric (a-zA-Z0-9) and the special chars (-=:.%) is allowed 

 [✖]  Namespace name include unsupported special chars. namespace : [<namespace>] 

  

 

### Usage

`$ stats`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
partitioned-topic | p | false | Get the partitioned topic stats 
per-partition |  | false | Get the per partition topic stats 



------------

## <em>update</em>

>bdocs-tab:example 

```bdocs-tab:example_shell
pulsarctl topics update (topic-name) (partition-num)
```


### Used For
 

 This command is used for updating the partition number of an exist topic. 

  
### Required Permission
 

 This command requires namespace admin permissions. 

  
### Output
 
 //normal output 

 Update topic (topic-name) with (partition-num) partitions successfully 

  
 //the topic name and(or) the partitions is not specified 

 [✖]  need to specified the topic name and the partitions 

  
 //the partitions number is invalid 

 [✖]  invalid partition number '<number>' 

  
 //the topic is not exist 

 [✖]  code: 409 reason: Topic is not partitioned topic 

  
 //the topic name is not in the format of <tenant>/<namespace>/<topic> or <topic> 

 [✖]  Invalid short topic name '<topic-name>', it should be in the format of <tenant>/<namespace>/<topic> or <topic> 

  
 //the topic name is not in the format of <domain>://<tenant>/<namespace>/<topic> 

 [✖]  Invalid complete topic name '<topic-name>', it should be in the format of <domain>://<tenant>/<namespace>/<topic> 

  
 //the topic name is not in the format of <tenant>/<namespace>/<topic> 

 [✖]  Invalid topic name '<topic-name>', it should be in the format of<tenant>/<namespace>/<topic> 

  
 //the namespace name is not in the format of <tenant>/<namespace> 

 [✖]  The complete name of namespace is invalid. complete name : <namespace-complete-name> 

  
 //the tenant name and(or) namespace name is empty 

 [✖]  Invalid tenant or namespace. [<tenant>/<namespace>] 

  
 //the tenant name contains unsupported special chars. the alphanumeric (a-zA-Z0-9) and the special chars (-=:.%)  is allowed 

 [✖]  Tenant name include unsupported special chars. tenant : [<namespace>] 

  
 //the namespace name contains unsupported special chars. the  alphanumeric (a-zA-Z0-9) and the special chars (-=:.%) is allowed 

 [✖]  Namespace name include unsupported special chars. namespace : [<namespace>] 

  

 

### Usage

`$ update`





