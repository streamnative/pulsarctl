
------------

# namespaces





### Usage

`$ namespaces`



------------

## <em>clear-backlog</em>

>bdocs-tab:example Clear backlog for all topics of the namespace (namespace-name)

```bdocs-tab:example_shell
pulsarctl namespaces clear-backlog (namespace-name)
```

>bdocs-tab:example Clear backlog for all topic of the namespace (namespace-name) with a bundle range <bundle>

```bdocs-tab:example_shell
pulsarctl namespaces clear-backlog --bundle (bundle) (namespace-name)
```

>bdocs-tab:example Clear the specified subscription (subscription-name) backlog for all topics of the namespace (namespace-name)

```bdocs-tab:example_shell
pulsarctl namespaces clear-backlog --subscription (subscription-name) (namespace-name)
```


### Used For
 

 This command is used for clearing backlog for all topics of a namespace. 

  
### Required Permission
 

 This command requires tenant admin permissions. 

  
### Output
 

 

### Usage

`$ clear-backlog`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
bundle | b |  | {start-boundary}_{end-boundary} 
force | f | false | Whether to force clear backlog without prompt 
sub |  |  | subscription name 



------------

## <em>create</em>

>bdocs-tab:example creates a namespace named (namespace-name)

```bdocs-tab:example_shell
pulsarctl namespaces create (namespace-name)
```


### Used For
 

 Creates a new namespace 

  
### Required Permission
 

 This command requires tenant admin permissions. 

  
### Output
 
 //normal output 

 Created (namespace-name) successfully 

  
 //the tenant does not exist 

 [✖]  code: 404 reason: Tenant does not exist 

  
 //you must specify a tenant/namespace name, please check if the tenant/namespace name is provided 

 [✖]  only one argument is allowed to be used as a name 

  
 //the namespace does not exist 

 [✖]  code: 404 reason: Namespace (tenant/namespace) does not exist 

  
 //Invalid number of bundles, please check --bundles value 

 Invalid number of bundles. Number of numBundles has to be in the range of (0, 2^32]. 

  

 

### Usage

`$ create`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
bundles | b | 0 | number of bundles to activate 
clusters | c | [] | List of clusters this namespace will be assigned 



------------

## <em>delete</em>

>bdocs-tab:example Delete a namespace

```bdocs-tab:example_shell
pulsarctl namespaces delete (namespace-name)
```


### Used For
 

 Delete a namespace. The namespace needs to be empty 

  
### Required Permission
 

 This command requires tenant admin permissions. 

  
### Output
 
 //normal output 

 Deleted (namespace-name) successfully 

  
 //you must specify a tenant/namespace name, please check if the tenant/namespace name is provided 

 [✖]  only one argument is allowed to be used as a name 

  
 //the tenant does not exist 

 [✖]  code: 404 reason: Tenant does not exist 

  

 

### Usage

`$ delete`




------------

## <em>delete-anti-affinity-group</em>

>bdocs-tab:example Delete an anti-affinity group of a namespace

```bdocs-tab:example_shell
pulsarctl namespaces delete-anti-affinity-group tenant/namespace
```


### Used For
 

 Delete an anti-affinity group of a namespace 

  
### Required Permission
 

 This command requires tenant admin permissions. 

  
### Output
 
 //normal output 

 Delete the anti-affinity group successfully for [tenant/namespace] 

  
 //you must specify a tenant/namespace name, please check if the tenant/namespace name is provided 

 [✖]  only one argument is allowed to be used as a name 

  
 //the tenant does not exist 

 [✖]  code: 404 reason: Tenant does not exist 

  
 //the namespace does not exist 

 [✖]  code: 404 reason: Namespace (tenant/namespace) does not exist 

  

 

### Usage

`$ delete-anti-affinity-group`




------------

## <em>get-anti-affinity-group</em>

>bdocs-tab:example Get the anti-affinity group of a namespace

```bdocs-tab:example_shell
pulsarctl namespaces get-anti-affinity-group tenant/namespace
```


### Used For
 

 Get the anti-affinity group of a namespace 

  
### Required Permission
 

 This command requires tenant admin permissions. 

  
### Output
 
 //normal output 

 (Anti-affinity group name) 

  
 //you must specify a tenant/namespace name, please check if the tenant/namespace name is provided 

 [✖]  only one argument is allowed to be used as a name 

  
 //the tenant does not exist 

 [✖]  code: 404 reason: Tenant does not exist 

  
 //the namespace does not exist 

 [✖]  code: 404 reason: Namespace (tenant/namespace) does not exist 

  

 

### Usage

`$ get-anti-affinity-group`




------------

## <em>get-anti-affinity-namespaces</em>

>bdocs-tab:example Get the list of namespaces in the same anti-affinity group.

```bdocs-tab:example_shell
pulsarctl namespaces get-anti-affinity-namespaces tenant/namespace
```


### Used For
 

 Get the list of namespaces in the same anti-affinity group. 

  
### Required Permission
 

 This command requires tenant admin permissions. 

  
### Output
 
 //normal output 

 (anti-affinity name list) 

  
 //you must specify a tenant/namespace name, please check if the tenant/namespace name is provided 

 [✖]  only one argument is allowed to be used as a name 

  
 //the tenant does not exist 

 [✖]  code: 404 reason: Tenant does not exist 

  
 //the namespace does not exist 

 [✖]  code: 404 reason: Namespace (tenant/namespace) does not exist 

  

 

### Usage

`$ get-anti-affinity-namespaces`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
cluster | c |  | Cluster name 
group | g |  | Anti-affinity group name 
tenant | t |  | tenant is only used for authorization. 
Client has to be admin of any of the tenant to access this api 



------------

## <em>get-backlog-quotas</em>

>bdocs-tab:example Get the backlog quota policy of a namespace

```bdocs-tab:example_shell
pulsarctl namespaces get-backlog-quotas tenant/namespace
```


### Used For
 

 Get the backlog quota policy of a namespace 

  
### Required Permission
 

 This command requires tenant admin permissions. 

  
### Output
 
 //normal output 

 { 

 "destination_storage" : { 

 "limit" : 10737418240, 

 "policy" : "producer_request_hold" 

 } 

 } 

  
 //you must specify a tenant/namespace name, please check if the tenant/namespace name is provided 

 [✖]  only one argument is allowed to be used as a name 

  
 //the tenant does not exist 

 [✖]  code: 404 reason: Tenant does not exist 

  
 //the namespace does not exist 

 [✖]  code: 404 reason: Namespace (tenant/namespace) does not exist 

  

 

### Usage

`$ get-backlog-quotas`




------------

## <em>get-clusters</em>

>bdocs-tab:example Get the replicated clusters of a namespace

```bdocs-tab:example_shell
pulsarctl namespaces get-clusters tenant/namespace
```


### Used For
 

 Get the replicated clusters of a namespace 

  
### Required Permission
 

 This command requires tenant admin permissions. 

  
### Output
 
 //normal output 

 (cluster name) 

  
 //you must specify a tenant/namespace name, please check if the tenant/namespace name is provided 

 [✖]  only one argument is allowed to be used as a name 

  
 //the tenant does not exist 

 [✖]  code: 404 reason: Tenant does not exist 

  
 //the namespace does not exist 

 [✖]  code: 404 reason: Namespace (tenant/namespace) does not exist 

  

 

### Usage

`$ get-clusters`




------------

## <em>get-dispatch-rate</em>

>bdocs-tab:example Get the default message dispatch rate of namespace (namespace-name)

```bdocs-tab:example_shell
pulsarctl namespaces get-dispatch-rate (namespace)
```


### Used For
 

 This command is used for getting the default message dispatch rate of a namespace. 

  
### Required Permission
 

 This command requires tenant admin permissions. 

  
### Output
 
 //normal output 

 { 

 "dispatchThrottlingRateInMsg" : 0, 

 "dispatchThrottlingRateInByte" : 0, 

 "ratePeriodInSecond" : 1 

 } 

  
 //the namespace name is not in the format of <tenant>/<namespace> 

 [✖]  The complete name of namespace is invalid. complete name : <namespace-complete-name> 

  
 //the tenant name and(or) namespace name is empty 

 [✖]  Invalid tenant or namespace. [<tenant>/<namespace>] 

  
 //the tenant name contains unsupported special chars. the alphanumeric (a-zA-Z0-9) and the special chars (-=:.%)  is allowed 

 [✖]  Tenant name include unsupported special chars. tenant : [<namespace>] 

  
 //the namespace name contains unsupported special chars. the  alphanumeric (a-zA-Z0-9) and the special chars (-=:.%) is allowed 

 [✖]  Namespace name include unsupported special chars. namespace : [<namespace>] 

  

 

### Usage

`$ get-dispatch-rate`




------------

## <em>get-message-ttl</em>

>bdocs-tab:example Get message TTL settings of a namespace

```bdocs-tab:example_shell
pulsarctl namespaces get-message-ttl tenant/namespace
```


### Used For
 

 Get message TTL settings of a namespace 

  
### Required Permission
 

 This command requires tenant admin permissions. 

  
### Output
 
 //normal output 

 (ttl-value) 

  
 //you must specify a tenant/namespace name, please check if the tenant/namespace name is provided 

 [✖]  only one argument is allowed to be used as a name 

  
 //the tenant does not exist 

 [✖]  code: 404 reason: Tenant does not exist 

  
 //the namespace does not exist 

 [✖]  code: 404 reason: Namespace (tenant/namespace) does not exist 

  

 

### Usage

`$ get-message-ttl`




------------

## <em>get-persistence</em>

>bdocs-tab:example Get the persistence policy of a namespace

```bdocs-tab:example_shell
pulsarctl namespaces get-persistence tenant/namespace
```


### Used For
 

 Get the persistence policy of a namespace 

  
### Required Permission
 

 This command requires tenant admin permissions. 

  
### Output
 
 //normal output 

 { 

 "bookkeeperEnsemble": 1, 

 "bookkeeperWriteQuorum": 1, 

 "bookkeeperAckQuorum": 1, 

 "managedLedgerMaxMarkDeleteRate": 0 

 } 

  
 //you must specify a tenant/namespace name, please check if the tenant/namespace name is provided 

 [✖]  only one argument is allowed to be used as a name 

  
 //the tenant does not exist 

 [✖]  code: 404 reason: Tenant does not exist 

  
 //the namespace does not exist 

 [✖]  code: 404 reason: Namespace (tenant/namespace) does not exist 

  

 

### Usage

`$ get-persistence`




------------

## <em>get-replicator-dispatch-rate</em>

>bdocs-tab:example Get the default replicator message dispatch rate of the namespace (namespace-name)

```bdocs-tab:example_shell
pulsarctl namespaces get-replicator-dispatch-rate (namespace)
```


### Used For
 

 This command is used for getting the default replicator message dispatch rate of a namespace. 

  
### Required Permission
 

 This command requires tenant admin permissions. 

  
### Output
 
 //normal output 

 { 

 "dispatchThrottlingRateInMsg" : 0, 

 "dispatchThrottlingRateInByte" : 0, 

 "ratePeriodInSecond" : 1 

 } 

  
 //the namespace name is not in the format of <tenant>/<namespace> 

 [✖]  The complete name of namespace is invalid. complete name : <namespace-complete-name> 

  
 //the tenant name and(or) namespace name is empty 

 [✖]  Invalid tenant or namespace. [<tenant>/<namespace>] 

  
 //the tenant name contains unsupported special chars. the alphanumeric (a-zA-Z0-9) and the special chars (-=:.%)  is allowed 

 [✖]  Tenant name include unsupported special chars. tenant : [<namespace>] 

  
 //the namespace name contains unsupported special chars. the  alphanumeric (a-zA-Z0-9) and the special chars (-=:.%) is allowed 

 [✖]  Namespace name include unsupported special chars. namespace : [<namespace>] 

  

 

### Usage

`$ get-replicator-dispatch-rate`




------------

## <em>get-retention</em>

>bdocs-tab:example Get the retention policy of a namespace

```bdocs-tab:example_shell
pulsarctl namespaces get-retention tenant/namespace
```


### Used For
 

 Get the retention policy of a namespace 

  
### Required Permission
 

 This command requires tenant admin permissions. 

  
### Output
 
 //normal output 

 { 

 "RetentionTimeInMinutes": 0, 

 "RetentionSizeInMB": 0 

 } 

  
 //you must specify a tenant/namespace name, please check if the tenant/namespace name is provided 

 [✖]  only one argument is allowed to be used as a name 

  
 //the tenant does not exist 

 [✖]  code: 404 reason: Tenant does not exist 

  
 //the namespace does not exist 

 [✖]  code: 404 reason: Namespace (tenant/namespace) does not exist 

  

 

### Usage

`$ get-retention`




------------

## <em>get-subscribe-rate</em>

>bdocs-tab:example Get the default subscribe rate per consumer of a namespace (namespace-name)

```bdocs-tab:example_shell
pulsarctl namespaces get-subscribe-rate (namespace)
```


### Used For
 

 This command is used for getting the default subscribe rate per consumer of a namespace. 

  
### Required Permission
 

 This command requires tenant admin permissions. 

  
### Output
 
 //normal output 

 { 

 "subscribeThrottlingRatePerConsumer" : 0, 

 "ratePeriodInSecond" : 30 

 } 

  
 //the namespace name is not specified 

 [✖]  only one argument is allowed to be used as a name 

  
 //the specified namespace name does not exist 

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

`$ get-subscribe-rate`




------------

## <em>get-subscription-dispatch-rate</em>

>bdocs-tab:example Get the default subscription message dispatch rate of namespace (namespace-name)

```bdocs-tab:example_shell
pulsarctl namespaces get-subscription-dispatch-rate (namespace-name)
```


### Used For
 

 This command is used for getting the default subscription message dispatch rate of a namespace. 

  
### Required Permission
 

 This command requires tenant admin permissions. 

  
### Output
 
 //normal output 

 { 

 "dispatchThrottlingRateInMsg" : 0, 

 "dispatchThrottlingRateInByte" : 0, 

 "ratePeriodInSecond" : 1 

 } 

  
 //the namespace name is not specified 

 [✖]  only one argument is allowed to be used as a name 

  
 //the specified namespace name does not exist 

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

`$ get-subscription-dispatch-rate`




------------

## <em>list</em>

>bdocs-tab:example Get the list of namespaces of a tenant

```bdocs-tab:example_shell
pulsarctl namespaces list (tenant name)
```


### Used For
 

 Get the list of namespaces of a tenant 

  
### Required Permission
 

 This command requires tenant admin permissions. 

  
### Output
 
 //normal output 

 +------------------+ 

 |  NAMESPACE NAME  | 

 +------------------+ 

 | public/default   | 

 | public/functions | 

 +------------------+ 

  
 //you must specify a tenant name, please check if the tenant name is provided 

 [✖]  only one argument is allowed to be used as a name 

  
 //the tenant does not exist 

 [✖]  code: 404 reason: Tenant does not exist 

  

 

### Usage

`$ list`




------------

## <em>messages-encryption</em>

>bdocs-tab:example Enable messages encryption for the namespace (namespace-name)

```bdocs-tab:example_shell
pulsarctl namespaces messages-encryption (namespace-name)
```

>bdocs-tab:example Disable messages encryption for the namespace (namespace-name)

```bdocs-tab:example_shell
pulsarct. namespaces messages-encryption --disable (namespace-name)
```


### Used For
 

 This command is used for enabling or disabling messages encryption for a namespace. 

  
### Required Permission
 

 This command requires tenant admin and a broker needs the read-write operations of the global zookeeper. 

  
### Output
 
 //normal output 

 Enable/Disable message encryption for the namespace (namespace-name) 

  
 //the namespace name is not specified 

 [✖]  only one argument is allowed to be used as a name 

  
 //the specified namespace name does not exist 

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

`$ messages-encryption`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
disable |  | false | Disable messages encryption 



------------

## <em>policies</em>

>bdocs-tab:example Get the configuration policies of a namespace

```bdocs-tab:example_shell
pulsarctl namespaces policies (tenant/namespace)
```


### Used For
 

 Get the configuration policies of a namespace 

  
### Required Permission
 

 This command requires tenant admin permissions. 

  
### Output
 
 //normal output 

 { 

 "AuthPolicies": {}, 

 "ReplicationClusters": null, 

 "Bundles": { 

 "boundaries": [ 

 "0x00000000", 

 "0x40000000", 

 "0x80000000", 

 "0xc0000000", 

 "0xffffffff" 

 ], 

 "numBundles": 4 

 }, 

 "BacklogQuotaMap": null, 

 "TopicDispatchRate": { 

 "standalone": { 

 "DispatchThrottlingRateInMsg": 0, 

 "DispatchThrottlingRateInByte": 0, 

 "RatePeriodInSecond": 1 

 } 

 }, 

 "SubscriptionDispatchRate": { 

 "standalone": { 

 "DispatchThrottlingRateInMsg": 0, 

 "DispatchThrottlingRateInByte": 0, 

 "RatePeriodInSecond": 1 

 } 

 }, 

 "ClusterSubscribeRate": { 

 "standalone": { 

 "SubscribeThrottlingRatePerConsumer": 0, 

 "RatePeriodInSecond": 30 

 } 

 }, 

 "Persistence": { 

 "BookkeeperEnsemble": 0, 

 "BookkeeperWriteQuorum": 0, 

 "BookkeeperAckQuorum": 0, 

 "ManagedLedgerMaxMarkDeleteRate": 0 

 }, 

 "DeduplicationEnabled": false, 

 "LatencyStatsSampleRate": null, 

 "MessageTtlInSeconds": 0, 

 "RetentionPolicies": { 

 "RetentionTimeInMinutes": 0, 

 "RetentionSizeInMB": 0 

 }, 

 "Deleted": false, 

 "AntiAffinityGroup": "", 

 "EncryptionRequired": false, 

 "SubscriptionAuthMode": "", 

 "MaxProducersPerTopic": 0, 

 "MaxConsumersPerTopic": 0, 

 "MaxConsumersPerSubscription": 0, 

 "CompactionThreshold": 0, 

 "OffloadThreshold": 0, 

 "OffloadDeletionLagMs": 0, 

 "SchemaAutoUpdateCompatibilityStrategy": "", 

 "SchemaValidationEnforced": false 

 } 

  
 //you must specify a tenant/namespace name, please check if the tenant/namespace name is provided 

 [✖]  only one argument is allowed to be used as a name 

  
 //the tenant does not exist 

 [✖]  code: 404 reason: Tenant does not exist 

  
 //the namespace does not exist 

 [✖]  code: 404 reason: Namespace (tenant/namespace) does not exist 

  

 

### Usage

`$ policies`




------------

## <em>remove-backlog-quota</em>

>bdocs-tab:example Remove a backlog quota policy from a namespace

```bdocs-tab:example_shell
pulsarctl namespaces remove-backlog-quota tenant/namespace
```


### Used For
 

 Remove a backlog quota policy from a namespace 

  
### Required Permission
 

 This command requires tenant admin permissions. 

  
### Output
 
 //normal output 

 Remove backlog quota successfully for [tenant/namespace] 

  
 //you must specify a tenant/namespace name, please check if the tenant/namespace name is provided 

 [✖]  only one argument is allowed to be used as a name 

  
 //the tenant does not exist 

 [✖]  code: 404 reason: Tenant does not exist 

  
 //the namespace does not exist 

 [✖]  code: 404 reason: Namespace (tenant/namespace) does not exist 

  

 

### Usage

`$ remove-backlog-quota`




------------

## <em>set-anti-affinity-group</em>

>bdocs-tab:example Set the anti-affinity group for a namespace

```bdocs-tab:example_shell
pulsarctl namespaces set-anti-affinity-group tenant/namespace
--group (anti-affinity group name)
```


### Used For
 

 Set the anti-affinity group for a namespace 

  
### Required Permission
 

 This command requires tenant admin permissions. 

  
### Output
 
 //normal output 

 Set the anti-affinity group: (anti-affinity group name) successfully for <tenant/namespace> 

  
 //you must specify a tenant/namespace name, please check if the tenant/namespace name is provided 

 [✖]  only one argument is allowed to be used as a name 

  
 //the tenant does not exist 

 [✖]  code: 404 reason: Tenant does not exist 

  
 //the namespace does not exist 

 [✖]  code: 404 reason: Namespace (tenant/namespace) does not exist 

  

 

### Usage

`$ set-anti-affinity-group`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
group | g |  | Anti-affinity group name 



------------

## <em>set-backlog-quota</em>

>bdocs-tab:example Set a backlog quota policy for a namespace

```bdocs-tab:example_shell
pulsarctl namespaces set-backlog-quota tenant/namespace
--limit 2G
--policy producer_request_hold
```


### Used For
 

 Set a backlog quota policy for a namespace 

  
### Required Permission
 

 This command requires tenant admin permissions. 

  
### Output
 
 //normal output 

 Set backlog quota successfully for [tenant/namespace] 

  
 //you must specify a tenant/namespace name, please check if the tenant/namespace name is provided 

 [✖]  only one argument is allowed to be used as a name 

  
 //the tenant does not exist 

 [✖]  code: 404 reason: Tenant does not exist 

  
 //the namespace does not exist 

 [✖]  code: 404 reason: Namespace (tenant/namespace) does not exist 

  
 //invalid retention policy type, please check --policy arg 

 invalid retention policy type: (policy type) 

  

 

### Usage

`$ set-backlog-quota`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
limit | l |  | Size limit (eg: 10M, 16G) 
policy | p |  | Retention policy to enforce when the limit is reached.
Valid options are: [producer_request_hold, producer_exception, consumer_backlog_eviction] 



------------

## <em>set-clusters</em>

>bdocs-tab:example Set the replicated clusters for a namespace

```bdocs-tab:example_shell
pulsarctl namespaces set-clusters tenant/namespace --clusters (cluster name)
```


### Used For
 

 Set the replicated clusters for a namespace 

  
### Required Permission
 

 This command requires tenant admin permissions. 

  
### Output
 
 //normal output 

 Set replication clusters successfully for tenant/namespace 

  
 //you must specify a tenant/namespace name, please check if the tenant/namespace name is provided 

 [✖]  only one argument is allowed to be used as a name 

  
 //the tenant does not exist 

 [✖]  code: 404 reason: Tenant does not exist 

  
 //the namespace does not exist 

 [✖]  code: 404 reason: Namespace (tenant/namespace) does not exist 

  
 //Invalid cluster name, please check if your cluster name has the appropriate permissions under the current tenant 

 [✖]  code: 403 reason: Cluster name is not in the list of allowed clusters list for tenant [public] 

  

 

### Usage

`$ set-clusters`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
clusters | c |  | Replication Cluster Ids list (comma separated values) 



------------

## <em>set-deduplication</em>

>bdocs-tab:example Enable or disable deduplication for a namespace

```bdocs-tab:example_shell
pulsarctl namespaces set-deduplication tenant/namespace (--enable)
```


### Used For
 

 Enable or disable deduplication for a namespace 

  
### Required Permission
 

 This command requires tenant admin permissions. 

  
### Output
 
 //normal output 

 Set deduplication is [true or false] successfully for public/default 

  
 //you must specify a tenant/namespace name, please check if the tenant/namespace name is provided 

 [✖]  only one argument is allowed to be used as a name 

  
 //the tenant does not exist 

 [✖]  code: 404 reason: Tenant does not exist 

  
 //the namespace does not exist 

 [✖]  code: 404 reason: Namespace (tenant/namespace) does not exist 

  

 

### Usage

`$ set-deduplication`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
enable | e | false | Enable deduplication 



------------

## <em>set-dispatch-rate</em>

>bdocs-tab:example Set the default message dispatch rate by message of the namespace (namespace-name) to (rate)

```bdocs-tab:example_shell
pulsarctl namespaces set-dispatch-rate --msg-rate (rate) (namespace)
```

>bdocs-tab:example Set the default message dispatch rate by byte of the namespace (namespace-name) to (rate)

```bdocs-tab:example_shell
pulsarctl namespaces set-dispatch-rate --byte-rate (rate) (namespace)
```

>bdocs-tab:example Set the default message dispatch rate by time of the namespace (namespace-name) to (period)

```bdocs-tab:example_shell
pulsarctl namespaces set-dispatch-rate --period (period) (namespace)
```


### Used For
 

 This command is used for setting the default message dispatch rate of a namespace. 

  
### Required Permission
 

 This command requires super-user permissions. 

  
### Output
 
 //normal output 

 Success set the default message dispatch rate of the namespace (namespace-name) to (rate) 

  
 //the namespace name is not in the format of <tenant>/<namespace> 

 [✖]  The complete name of namespace is invalid. complete name : <namespace-complete-name> 

  
 //the tenant name and(or) namespace name is empty 

 [✖]  Invalid tenant or namespace. [<tenant>/<namespace>] 

  
 //the tenant name contains unsupported special chars. the alphanumeric (a-zA-Z0-9) and the special chars (-=:.%)  is allowed 

 [✖]  Tenant name include unsupported special chars. tenant : [<namespace>] 

  
 //the namespace name contains unsupported special chars. the  alphanumeric (a-zA-Z0-9) and the special chars (-=:.%) is allowed 

 [✖]  Namespace name include unsupported special chars. namespace : [<namespace>] 

  

 

### Usage

`$ set-dispatch-rate`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
byte-rate | b | -1 | byte dispatch rate (default -1) 
msg-rate | m | -1 | message dispatch rate (default -1) 
period | p | 1 | dispatch rate period (default 1 second) 



------------

## <em>set-message-ttl</em>

>bdocs-tab:example Set Message TTL for a namespace

```bdocs-tab:example_shell
pulsarctl namespaces set-message-ttl tenant/namespace -ttl 10
```


### Used For
 

 Set Message TTL for a namespace 

  
### Required Permission
 

 This command requires tenant admin permissions. 

  
### Output
 
 //normal output 

 Set message TTL successfully for [tenant/namespace] 

  
 //Invalid value for message TTL, please check -ttl arg 

 code: 412 reason: Invalid value for message TTL 

  
 //you must specify a tenant/namespace name, please check if the tenant/namespace name is provided 

 [✖]  only one argument is allowed to be used as a name 

  
 //the tenant does not exist 

 [✖]  code: 404 reason: Tenant does not exist 

  
 //the namespace does not exist 

 [✖]  code: 404 reason: Namespace (tenant/namespace) does not exist 

  

 

### Usage

`$ set-message-ttl`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
messageTTL | t | 0 | Message TTL in seconds 



------------

## <em>set-persistence</em>

>bdocs-tab:example Set the persistence policy for a namespace

```bdocs-tab:example_shell
pulsarctl namespaces set-persistence tenant/namespace
--ensemble-size 2
--write-quorum-size 2
--ack-quorum-size 2
--ml-mark-delete-max-rate 2.0
```


### Used For
 

 Set the persistence policy for a namespace 

  
### Required Permission
 

 This command requires tenant admin permissions. 

  
### Output
 
 //normal output 

 Set the persistence policies successfully for [tenant/namespace] 

  
 //you must specify a tenant/namespace name, please check if the tenant/namespace name is provided 

 [✖]  only one argument is allowed to be used as a name 

  
 //the tenant does not exist 

 [✖]  code: 404 reason: Tenant does not exist 

  
 //the namespace does not exist 

 [✖]  code: 404 reason: Namespace (tenant/namespace) does not exist 

  
 //Bookkeeper Ensemble >= WriteQuorum >= AckQuoru, please c 

 code: 412 reason: Bookkeeper Ensemble >= WriteQuorum >= AckQuoru 

  

 

### Usage

`$ set-persistence`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
ack-quorum-size | a | 0 | Number of acks (garanteed copies) to wait for each entry 
ensemble-size | e | 0 | Number of bookies to use for a topic 
ml-mark-delete-max-rate | r | 0 | Throttling rate of mark-delete operation (0 means no throttle) 
write-quorum-size | w | 0 | How many writes to make of each entry 



------------

## <em>set-replicator-dispatch-rate</em>

>bdocs-tab:example Set the default replicator message dispatch rate by message of the namespace (namespace-name) to (rate)

```bdocs-tab:example_shell
pulsarctl namespaces set-replicator-dispatch-rate --msg-rate (rate) (namespace)
```

>bdocs-tab:example Set the default replicator message dispatch rate by byte of the namespace (namespace-name) to (rate)

```bdocs-tab:example_shell
pulsarctl namespaces set-replicator-dispatch-rate --byte-rate (rate) (namespace)
```

>bdocs-tab:example Set the default replicator message dispatch rate by time of the namespace (namespace-name) to (period)

```bdocs-tab:example_shell
pulsarctl namespaces set-replicator-dispatch-rate --period (period) (namespace)
```


### Used For
 

 This command is used for setting the default replicator message dispatch rate of a namespace. 

  
### Required Permission
 

 This command requires super-user permissions. 

  
### Output
 
 //normal output 

 Success set the default replicator message dispatch rate of the namespace (namespace-name) to (rate) 

  
 //the namespace name is not specified 

 [✖]  only one argument is allowed to be used as a name 

  
 //the replicator-dispatch-rate is not configured 

 [✖]  code: 404 reason: replicator-Dispatch-rate is not configured for cluster standalone 

  
 //the namespace name is not in the format of <tenant>/<namespace> 

 [✖]  The complete name of namespace is invalid. complete name : <namespace-complete-name> 

  
 //the tenant name and(or) namespace name is empty 

 [✖]  Invalid tenant or namespace. [<tenant>/<namespace>] 

  
 //the tenant name contains unsupported special chars. the alphanumeric (a-zA-Z0-9) and the special chars (-=:.%)  is allowed 

 [✖]  Tenant name include unsupported special chars. tenant : [<namespace>] 

  
 //the namespace name contains unsupported special chars. the  alphanumeric (a-zA-Z0-9) and the special chars (-=:.%) is allowed 

 [✖]  Namespace name include unsupported special chars. namespace : [<namespace>] 

  

 

### Usage

`$ set-replicator-dispatch-rate`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
byte-rate | b | -1 | byte dispatch rate (default -1) 
msg-rate | m | -1 | message dispatch rate (default -1) 
period | p | 1 | dispatch rate period (default 1 second) 



------------

## <em>set-retention</em>

>bdocs-tab:example Set the retention policy for a namespace

```bdocs-tab:example_shell
pulsarctl namespaces set-retention tenant/namespace --time 100m
```

>bdocs-tab:example Set the retention policy for a namespace

```bdocs-tab:example_shell
pulsarctl namespaces set-retention tenant/namespace --size 1G
```


### Used For
 

 Set the retention policy for a namespace 

  
### Required Permission
 

 This command requires tenant admin permissions. 

  
### Output
 
 //normal output 

 Set retention successfully for [tenant/namespace] 

  
 //you must specify a tenant/namespace name, please check if the tenant/namespace name is provided 

 [✖]  only one argument is allowed to be used as a name 

  
 //the tenant does not exist 

 [✖]  code: 404 reason: Tenant does not exist 

  
 //the namespace does not exist 

 [✖]  code: 404 reason: Namespace (tenant/namespace) does not exist 

  
 //Retention Quota must exceed configured backlog quota for namespace 

 Retention Quota must exceed configured backlog quota for namespace 

  

 

### Usage

`$ set-retention`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
size |  |  | Retention size limit (eg: 10M, 16G, 3T).
0 or less than 1MB means no retention and -1 means infinite size retention 
time |  |  | Retention time in minutes (or minutes, hours,days,weeks eg: 100m, 3h, 2d, 5w).
0 means no retention and -1 means infinite time retention 



------------

## <em>set-subscribe-rate</em>

>bdocs-tab:example Set the default subscribe rate by subscribe of the namespace (namespace-name) (rate)

```bdocs-tab:example_shell
pulsarctl namespaces set-subscribe-rate --subscribe-rate (rate) (namespace)
```

>bdocs-tab:example Set the default subscribe rate by time of the namespace (namespace-name) (period)

```bdocs-tab:example_shell
pulsarctl namespaces set-subscribe-rate --period (period) (namespace)
```


### Used For
 

 This command is used for setting the default subscribe rate per consumer of a namespace. 

  
### Required Permission
 

 This command requires super-user permissions. 

  
### Output
 
 //normal output 

 Success set the default subscribe rate of the namespace (namespace-name) to (rate) 

  
 //the namespace name is not specified 

 [✖]  only one argument is allowed to be used as a name 

  
 //the specified namespace name does not exist 

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

`$ set-subscribe-rate`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
period | p | 30 | dispatch rate period (default 30 second) 
subscribe-rate | m | -1 | message dispatch rate (default -1) 



------------

## <em>set-subscription-auth-mode</em>

>bdocs-tab:example Set the default subscription auth mode (mode) of the namespace (namespace-name)

```bdocs-tab:example_shell
pulsarctl namespaces set-subscription-auth-mode --mode (mode) (namespace-name)
```


### Used For
 

 This command is used for setting the default subscription auth mode of a namespace. 

  
### Required Permission
 

 This command requires tenant admin and a broker needs the read-write operations of the global zookeeper. 

  
### Output
 
 //normal output 

 Successfully set the default subscription auth mode of namespace <namespace-name> to <mode> 

  
 //the namespace name is not specified 

 [✖]  only one argument is allowed to be used as a name 

  
 //the specified namespace name does not exist 

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

`$ set-subscription-auth-mode`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
mode | m |  | Subscription authorization mode of a namespace. (e.g. None, Prefix) 



------------

## <em>set-subscription-dispatch-rate</em>

>bdocs-tab:example Set the default subscription message dispatch rate by message of the namespace (namespace-name) to (rate)

```bdocs-tab:example_shell
pulsarctl namespaces set-subscription-dispatch-rate --msg-rate <rate> <namespace
```

>bdocs-tab:example Set the default subscription message dispatch rate by byte of the namespace (namespace-name) to (rate)

```bdocs-tab:example_shell
pulsarctl namespaces set-subscription-dispatch-rate --byte-rate (rate) (namespace)
```

>bdocs-tab:example Set the default subscriptions message dispatch rate by time of the namespace (namespace-name) to (rate)

```bdocs-tab:example_shell
pulsarctl namespaces set-subscription-dispatch-rate --period (period) (namespace)
```


### Used For
 

 This command is used for setting the default subscription message dispatch rate of a namespace. 

  
### Required Permission
 

 This command requires super-user permissions. 

  
### Output
 
 //normal output 

 Success set the default subscription message dispatch rate of the namespace (namespace-name) to (rate) 

  
 //the namespace name is not specified 

 [✖]  only one argument is allowed to be used as a name 

  
 //the specified namespace name does not exist 

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

`$ set-subscription-dispatch-rate`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
byte-rate | b | -1 | byte dispatch rate (default -1) 
msg-rate | m | -1 | message dispatch rate (default -1) 
period | p | 1 | dispatch rate period (default 1 second) 



------------

## <em>split-bundle</em>

>bdocs-tab:example Split a namespace-bundle from the current serving broker

```bdocs-tab:example_shell
pulsarctl namespaces split-bundle tenant/namespace --bundle ({start-boundary}_{end-boundary})
```

>bdocs-tab:example Split a namespace-bundle from the current serving broker

```bdocs-tab:example_shell
pulsarctl namespaces split-bundle tenant/namespace
--bundle ({start-boundary}_{end-boundary})
--unload
```


### Used For
 

 Split a namespace-bundle from the current serving broker 

  
### Required Permission
 

 This command requires tenant admin permissions. 

  
### Output
 
 //normal output 

 Split a namespace bundle: ({start-boundary}_{end-boundary}) successfully 

  
 //you must specify a tenant/namespace name, please check if the tenant/namespace name is provided 

 [✖]  only one argument is allowed to be used as a name 

  
 //the tenant does not exist 

 [✖]  code: 404 reason: Tenant does not exist 

  
 //the namespace does not exist 

 [✖]  code: 404 reason: Namespace (tenant/namespace) does not exist 

  
 //Please check if there is an active topic under the current split bundle. 

 [✖]  code: 412 reason: Failed to find ownership for ServiceUnit:public/default/(bundle range) 

  

 

### Usage

`$ split-bundle`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
bundle | b |  | {start-boundary}_{end-boundary} 
unload | u | false | Unload newly split bundles after splitting old bundle 



------------

## <em>topics</em>

>bdocs-tab:example Get the list of topics for a namespace

```bdocs-tab:example_shell
pulsarctl namespaces topics (tenant/namespace)
```


### Used For
 

 Get the list of topics for a namespace 

  
### Required Permission
 

 This command requires namespace admin permissions. 

  
### Output
 
 //normal output 

 +-------------+ 

 | TOPICS NAME | 

 +-------------+ 

 +-------------+ 

  
 //you must specify a tenant/namespace name, please check if the tenant/namespace name is provided 

 [✖]  only one argument is allowed to be used as a name 

  
 //the tenant does not exist 

 [✖]  code: 404 reason: Tenant does not exist 

  
 //the namespace does not exist 

 [✖]  code: 404 reason: Namespace (tenant/namespace) does not exist 

  

 

### Usage

`$ topics`




------------

## <em>unload</em>

>bdocs-tab:example Unload a namespace from the current serving broker

```bdocs-tab:example_shell
pulsarctl namespaces unload tenant/namespace
```

>bdocs-tab:example Unload a namespace with bundle from the current serving broker

```bdocs-tab:example_shell
pulsarctl namespaces unload tenant/namespace --bundle ({start-boundary}_{end-boundary})
```


### Used For
 

 Unload a namespace from the current serving broker 

  
### Required Permission
 

 This command requires tenant admin permissions. 

  
### Output
 
 //normal output 

 Unload namespace (tenant/namespace) (with bundle ({start-boundary}_{end-boundary})) successfully 

  
 //you must specify a tenant/namespace name, please check if the tenant/namespace name is provided 

 [✖]  only one argument is allowed to be used as a name 

  
 //the tenant does not exist 

 [✖]  code: 404 reason: Tenant does not exist 

  
 //the namespace does not exist 

 [✖]  code: 404 reason: Namespace (tenant/namespace) does not exist 

  

 

### Usage

`$ unload`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
bundle | b |  | {start-boundary}_{end-boundary}(e.g. 0x00000000_0xffffffff) 



------------

## <em>unsubscribe</em>

>bdocs-tab:example Unsubscribe the specified subscription <subscription-name> for all topic of the namespace (namespace-name)

```bdocs-tab:example_shell
pulsarctl namespaces unsubscribe (namespace-name) (subscription-name)
```

>bdocs-tab:example Unsubscribe the specified subscription (subscription-name) for all topic of the namespace (namespace-name) with bundle range <bundle>

```bdocs-tab:example_shell
pulsarctl namespaces unsubscribe --bundle (bundle) (namespace-name) (subscription-name)
```


### Used For
 

 This command is used for unsubscribing the specified subscription for all topics of a namespace. 

  
### Required Permission
 

 This command requires tenant admin permissions. 

  
### Output
 
 //normal output 

 Successfully unsubscribe the subscription (subscription-name) for all topics of the namespace (namespace-name) 

  
 //the namespace name is not specified or the subscription name is not specified 

 [✖]  need two arguments apply to the command 

  
 //the specified namespace name does not exist 

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

`$ unsubscribe`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
bundle | b |  | {start_boundary}_{end_boundary} 




