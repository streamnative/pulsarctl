
------------

# schemas





### Usage

`$ schemas`



------------

## <em>delete</em>

>bdocs-tab:example Delete the latest schema for a topic

```bdocs-tab:example_shell
pulsarctl schemas delete (topic name)
```


### Used For
 

 Delete the latest schema for a topic 

  
### Required Permission
 

 This command requires namespace admin permissions. 

  
### Output
 
 //normal output 

 Deleted (topic name) successfully 

  
 //you must specify a topic name, please check if the topic name is provided 

 [✖]  only one argument is allowed to be used as a name 

  

 

### Usage

`$ delete`




------------

## <em>get</em>

>bdocs-tab:example Get the schema for a topic

```bdocs-tab:example_shell
pulsarctl schemas get (topic name)
```

>bdocs-tab:example Get the schema for a topic with version

```bdocs-tab:example_shell
pulsarctl schemas get (topic name)
--version 2
```


### Used For
 

 Get the schema for a topic. 

  
### Required Permission
 

 This command requires namespace admin permissions. 

  
### Output
 
 //normal output 

 { 

 "name": "test-schema", 

 "schema": { 

 "type": "record", 

 "name": "Test", 

 "fields": [ 

 { 

 "name": "id", 

 "type": [ 

 "null", 

 "int" 

 ] 

 }, 

 { 

 "name": "name", 

 "type": [ 

 "null", 

 "string" 

 ] 

 } 

 ] 

 }, 

 "type": "AVRO", 

 "properties": {} 

 } 

  
 //HTTP 404 Not Found, please check if the topic name you entered is correct 

 [✖]  code: 404 reason: Not Found 

  
 //you must specify a topic name, please check if the topic name is provided 

 [✖]  only one argument is allowed to be used as a name 

  

 

### Usage

`$ get`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
version |  | 0 | the schema version info 



------------

## <em>upload</em>

>bdocs-tab:example Update the schema for a topic

```bdocs-tab:example_shell
pulsarctl schemas upload
(topic name)
--filename (the file path of schema)
```


### Used For
 

 Update the schema for a topic 

  
### Required Permission
 

 This command requires namespace admin permissions. 

  
### Output
 
 //normal output 

 Upload (topic name) successfully 

  
 //you must specify a topic name, please check if the topic name is provided 

 [✖]  only one argument is allowed to be used as a name 

  
 //no such file or directory 

 [✖]  open (file path): no such file or directory 

  

 

### Usage

`$ upload`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
filename | f |  | filename 




