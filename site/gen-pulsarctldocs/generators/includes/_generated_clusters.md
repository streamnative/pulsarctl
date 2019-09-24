
------------

# clusters





### Usage

`$ clusters`



------------

## <em>add</em>

>bdocs-tab:example Provisions a new cluster

```bdocs-tab:example_shell
pulsarctl clusters create (cluster-name)
```


### Used For
 

 This command is used for adding the configuration data for a cluster. 

 The configuration data is mainly used for geo-replication between clusters, so please make sure the service urls provided in this command are reachable between clusters. 

 This operation requires Pulsar super-user privileges. 

  
### Required Permission
 

 This command requires super-user permissions. 

  
### Output
 
 //normal output 

 Cluster (cluster-name) added 

  

 

### Usage

`$ add`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
broker-url |  |  | Pulsar cluster broker service url, e.g. pulsar://example.pulsar.io:6650 
broker-url-tls |  |  | Pulsar cluster tls secured broker service url, e.g. pulsar+ssl://example.pulsar.io:6651 
peer-cluster | p | [] | Cluster to be registered as a peer-cluster of this cluster. 
url |  |  | Pulsar cluster web service url, e.g. http://example.pulsar.io:8080 
url-tls |  |  | Pulsar cluster tls secured web service url, e.g. https://example.pulsar.io:8443 



------------

## <em>create-failure-domain</em>

>bdocs-tab:example create the failure domain

```bdocs-tab:example_shell
pulsarctl clusters create-failure-domain (cluster-name) (domain-name)
```

>bdocs-tab:example create the failure domain with brokers

```bdocs-tab:example_shell
pulsarctl clusters create-failure-domain -b (broker-ip):(broker-port) -b (broker-ip):(broker-port) (cluster-name) (domain-name)
```


### Used For
 

 This command is used for creating a failure domain of the (cluster-name). 

  
### Required Permission
 

 This command requires super-user permissions. 

  
### Output
 
 //normal output 

 Create failure domain (domain-name) for cluster (cluster-name) succeed 

  
 //the args need to be specified as (cluster-name) (domain-name) 

 [✖]  need specified two names for cluster and failure domain 

  
 //the specified cluster does not exist in the broker 

 [✖]  code: 412 reason: Cluster <cluster-name> does not exist. 

  

 

### Usage

`$ create-failure-domain`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
brokers | b | [] | Set the failure domain clusters 



------------

## <em>delete</em>

>bdocs-tab:example deleting the cluster named (cluster-name)

```bdocs-tab:example_shell
pulsarctl clusters delete (cluster-name)
```


### Used For
 

 This command is used for deleting an existing cluster. 

  
### Required Permission
 

 This command requires super-user permissions. 

  
### Output
 
 //normal output 

 Cluster (cluster-name) delete successfully. 

  
 //the cluster name is not specified or the cluster name is specified more than one 

 [✖]  only one argument is allowed to be used as a name 

  
 //the specified cluster does not exist in the broker 

 [✖]  code: 412 reason: Cluster <cluster-name> does not exist. 

  

 

### Usage

`$ delete`




------------

## <em>delete-failure-domain</em>

>bdocs-tab:example delete the failure domain

```bdocs-tab:example_shell
pulsarctl clusters delete-failure-domain (cluster-name) (domain-name)
```


### Used For
 

 This command is used for deleting the failure domain (domain-name) of the cluster (cluster-name) 

  
### Required Permission
 

 This command requires super-user permissions. 

  
### Output
 
 //output example 

 Delete failure domain [(domain-name)] for cluster [(cluster-name)] succeed 

  
 //the cluster name and(or) failure domain name is not specified or the name is specified more than one 

 [✖]  need to specified the cluster name and the failure domain name 

  
 //the specified failure domain is not exist 

 code: 404 reason: Domain-name non-existent-failure-domain or cluster standalone does not exist 

  
 //the specified cluster is not exist 

 code: 412 reason: Cluster non-existent-cluster does not exist. 

  

 

### Usage

`$ delete-failure-domain`




------------

## <em>get</em>

>bdocs-tab:example getting the (cluster-name) data

```bdocs-tab:example_shell
pulsarctl clusters get (cluster-name)
```


### Used For
 

 This command is used for getting the cluster data of the specified cluster. 

  
### Required Permission
 

 This command requires super-user permissions. 

  
### Output
 
 //normal output 

 { 

 "serviceUrl": "http://localhost:8080", 

 "serviceUrlTls": "", 

 "brokerServiceUrl": "pulsar://localhost:6650", 

 "brokerServiceUrlTls": "", 

 "peerClusterNames": null 

 } 

  
 //the cluster name is not specified or the cluster name is specified more than one 

 [✖]  only one argument is allowed to be used as a name 

  
 //the specified cluster does not exist in the broker 

 [✖]  code: 412 reason: Cluster <cluster-name> does not exist. 

  

 

### Usage

`$ get`




------------

## <em>get-failure-domain</em>

>bdocs-tab:example getting the broker list in the (cluster-name) cluster failure domain (domain-name)

```bdocs-tab:example_shell
pulsarctl clusters get-failure-domain (cluster-name) (domain-name)
```


### Used For
 

 This command is used for getting the specified failure domain on the specified cluster. 

  
### Required Permission
 

 This command requires super-user permissions. 

  
### Output
 
 //output example 

 { 

 "brokers" : [ 

 "failure-broker-A", 

 "failure-broker-B", 

 ] 

 } 

  
 //the cluster name and(or) failure domain name is not specified or the name is specified more than one 

 [✖]  need to specified the cluster name and the failure domain name 

  
 //the specified cluster does not exist in the broker 

 [✖]  code: 412 reason: Cluster <cluster-name> does not exist. 

  

 

### Usage

`$ get-failure-domain`




------------

## <em>get-peer-clusters</em>

>bdocs-tab:example getting the (cluster-name) peer clusters

```bdocs-tab:example_shell
pulsarctl clusters get-peer-clusters (cluster-name)
```


### Used For
 

 This command is used for getting the peer clusters of the specified cluster. 

  
### Required Permission
 

 This command requires super-user permissions. 

  
### Output
 
 //normal output 

 +-------------------+ 

 |   PEER CLUSTERS   | 

 +-------------------+ 

 | test_peer_cluster | 

 +-------------------+ 

  
 //the cluster name is not specified or the cluster name is specified more than one 

 [✖]  only one argument is allowed to be used as a name 

  
 //the specified cluster does not exist in the broker 

 [✖]  code: 412 reason: Cluster <cluster-name> does not exist. 

  

 

### Usage

`$ get-peer-clusters`




------------

## <em>list</em>

>bdocs-tab:example List the existing clusters

```bdocs-tab:example_shell
pulsarctl clusters create-failure-domain (cluster-name) (domain-name)
```




 This command is used for listing the list of available pulsar clusters.

### Usage

`$ list`




------------

## <em>list-failure-domains</em>

>bdocs-tab:example listing all the failure domains under the specified cluster

```bdocs-tab:example_shell
pulsarctl clusters list-failure-domains (cluster-name)
```


### Used For
 

 This command is used for getting all failure domain under the cluster (cluster-name). 

  
### Required Permission
 

 This command requires super-user permissions. 

  
### Output
 
 //output example 

 { 

 "failure-domain": { 

 "brokers": [ 

 "failure-broker-A", 

 "failure-broker-B" 

 ] 

 } 

 } 

  
 //the cluster name is not specified or the cluster name is specified more than one 

 [✖]  only one argument is allowed to be used as a name 

  
 //the specified cluster does not exist in the broker 

 [✖]  code: 412 reason: Cluster <cluster-name> does not exist. 

  

 

### Usage

`$ list-failure-domains`




------------

## <em>update</em>

>bdocs-tab:example updating the web service url of the (cluster-name)

```bdocs-tab:example_shell
pulsarctl clusters update --url http://example:8080 (cluster-name)
```

>bdocs-tab:example updating the tls secured web service url of the (cluster-name)

```bdocs-tab:example_shell
pulsarctl clusters update --url-tls https://example:8080 (cluster-name)
```

>bdocs-tab:example updating the broker service url of the (cluster-name)

```bdocs-tab:example_shell
pulsarctl clusters update --broker-url pulsar://example:6650 (cluster-name)
```

>bdocs-tab:example updating the tls secured web service url of the (cluster-name)

```bdocs-tab:example_shell
pulsarctl clusters update --broker-url-tls pulsar+ssl://example:6650 (cluster-name)
```

>bdocs-tab:example registered as a peer-cluster of the (cluster-name) clusters

```bdocs-tab:example_shell
pulsarctl clusters update -p (cluster-a) -p (cluster-b) (cluster)
```


### Used For
 

 This command is used for updating the cluster data of the specified cluster. 

  
### Required Permission
 

 This command requires super-user permissions. 

  
### Output
 
 //normal output 

 Cluster (cluster-name) updated 

  
 //the cluster name is not specified or the cluster name is specified more than one 

 [✖]  only one argument is allowed to be used as a name 

  
 //the specified cluster does not exist in the broker 

 [✖]  code: 412 reason: Cluster <cluster-name> does not exist. 

  

 

### Usage

`$ update`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
broker-url |  |  | Pulsar cluster broker service url, e.g. pulsar://example.pulsar.io:6650 
broker-url-tls |  |  | Pulsar cluster tls secured broker service url, e.g. pulsar+ssl://example.pulsar.io:6651 
peer-cluster | p | [] | Cluster to be registered as a peer-cluster of this cluster. 
url |  |  | Pulsar cluster web service url, e.g. http://example.pulsar.io:8080 
url-tls |  |  | Pulsar cluster tls secured web service url, e.g. https://example.pulsar.io:8443 



------------

## <em>update-failure-domain</em>

>bdocs-tab:example update the failure domain

```bdocs-tab:example_shell
pulsarctl clusters update-failure-domain (cluster-name) (domain-name)
```

>bdocs-tab:example update the failure domain with brokers

```bdocs-tab:example_shell
pulsarctl clusters update-failure-domain --broker-list <cluster-A> --broker-list (cluster-B) (cluster-name) (domain-name)
```


### Used For
 

 This command is used for updating a failure domain of the (cluster-name). 

  
### Required Permission
 

 This command requires super-user permissions. 

  
### Output
 
 //normal output 

 Update failure domain (domain-name) for cluster (cluster-name) succeed 

  
 //the args need to be specified as (cluster-name) (domain-name) 

 [✖]  need specified two names for cluster and failure domain 

  
 //the specified cluster does not exist in the broker 

 [✖]  code: 412 reason: Cluster <cluster-name> does not exist. 

  

 

### Usage

`$ update-failure-domain`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
broker-list | b | [] | Set the failure domain clusters 



------------

## <em>update-peer-clusters</em>

>bdocs-tab:example updating the <cluster-name> peer clusters

```bdocs-tab:example_shell
pulsarctl clusters update-peer-clusters -p cluster-a -p cluster-b (cluster-name)
```


### Used For
 

 This command is used for updating peer clusters. 

  
### Required Permission
 

 This command requires super-user permissions. 

  
### Output
 
 //output example 

 (cluster-name) peer clusters updated 

  
 //the cluster name is not specified or the cluster name is specified more than one 

 [✖]  only one argument is allowed to be used as a name 

  
 //the specified cluster does not exist in the broker 

 [✖]  code: 412 reason: Cluster <cluster-name> does not exist. 

  

 

### Usage

`$ update-peer-clusters`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
peer-cluster | p | [] | Cluster to be registered as a peer-cluster of this cluster 




