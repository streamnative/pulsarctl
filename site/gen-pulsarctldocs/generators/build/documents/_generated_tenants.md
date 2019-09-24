
------------

# tenants





### Usage

`$ tenants`



------------

## <em>create</em>

>bdocs-tab:example create a tenant named (tenant-name)

```bdocs-tab:example_shell
pulsarctl tenants create (tenant-name)
```


### Used For
 

 This command is used for creating a new tenant. 

  
### Required Permission
 

 This command requires super-user permissions. 

  
### Output
 
 //normal output 

 Create tenant (tenant-name) successfully 

  
 //the tenant name is not specified or the tenant name is specified more than one 

 [✖]  only one argument is allowed to be used as a name 

  
 //the specified tenant has been created 

 [✖]  code: 409 reason: Tenant already exists 

  

 

### Usage

`$ create`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
admin-roles | r | [] | Allowed admins to access the tenant 
allowed-clusters | c | [] | Allowed clusters 



------------

## <em>delete</em>

>bdocs-tab:example delete a tenant named (tenant-name)

```bdocs-tab:example_shell
pulsarctl tenants delete (tenant-name)
```


### Used For
 

 This command is used for deleting an existing tenant. 

  
### Required Permission
 

 This command requires super-user permissions. 

  
### Output
 
 //normal output 

 Delete tenant <tenant-name> successfully 

  
 //the tenant name is not specified or the tenant name is specified more than one 

 [✖]  only one argument is allowed to be used as a name 

  
 //the specified tenant does not exist in the broker 

 [✖]  code: 404 reason: The tenant does not exist 

  
 //there has namespace(s) under the tenant (tenant-name) 

 code: 409 reason: The tenant still has active namespaces 

  

 

### Usage

`$ delete`




------------

## <em>get</em>

>bdocs-tab:example get the configuration of tenant (tenant-name)

```bdocs-tab:example_shell
pulsarctl tenants get (tenant-name)
```


### Used For
 

 This command is used for getting the configuration of a tenant. 

  
### Required Permission
 

 This command requires super-user permissions. 

  
### Output
 
 //normal output 

 { 

 "adminRoles": [], 

 "allowedClusters": [ 

 "standalone" 

 ] 

 } 

  
 //the tenant name is not specified or the tenant name is specified more than one 

 [✖]  only one argument is allowed to be used as a name 

  
 //the specified tenant does not exist in the cluster 

 [✖]  code: 404 reason: Tenant does not exist 

  

 

### Usage

`$ get`




------------

## <em>list</em>

>bdocs-tab:example list all the existing tenants

```bdocs-tab:example_shell
pulsarctl tenants list
```


### Used For
 

 This command is used for listing all the existing tenants. 

  
### Required Permission
 

 This command requires super-user permissions. 

  
### Output
 
 //normal output 

 +-------------+ 

 | TENANT NAME | 

 +-------------+ 

 | public      | 

 | sample      | 

 +-------------+ 

  

 

### Usage

`$ list`




------------

## <em>update</em>

>bdocs-tab:example clear the tenant configuration of a tenant

```bdocs-tab:example_shell
pulsarctl tenant update (tenant-name)
```

>bdocs-tab:example update the admin roles for tenant (tenant-name)

```bdocs-tab:example_shell
pulsarctl tenants update --admin-roles (admin-A)--admin-roles (admin-B) (tenant-name)
```

>bdocs-tab:example update the allowed cluster list for tenant (tenant-name)

```bdocs-tab:example_shell
pulsarctl tenants update --allowed-clusters (cluster-A) --allowed-clusters (cluster-B) (tenant-name)
```


### Used For
 

 This command is used for updating the configuration of a tenant. 

  
### Required Permission
 

 This command requires super-user permissions. 

  
### Output
 
 //normal output 

 Update tenant [%s] successfully 

  
 //the tenant name is not specified or the tenant name is specified more than one 

 [✖]  only one argument is allowed to be used as a name 

  
 //the specified tenant does not exist in 

 [✖]  code: 404 reason: Tenant does not exist 

  
 //the flag --admin-roles or --allowed-clusters are not specified 

 [✖]  the admin roles or the allowed clusters is not specified 

  

 

### Usage

`$ update`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
admin-roles | r | [] | Allowed admins to access the tenant 
allowed-clusters | c | [] | Allowed clusters 




