# API GOLANG SKELETON ( Clean architecture )

### Config file 

The config file is located in ``` shared/utils/config/config.json```

```javascript
/**
 Some considerations :

 - For run local use "localhost" or use "host.docker.internal" for docker run in mongoUri
 or replace for your own database implementation

 - Port must be the same in docker run command with the specified in this config file
 */
````

````json
{
  "currentStage": "local",
  "developers": ["Cristian Nores"],
  "url": "golang-skeleton",
  "port": "8888",
  "mongoUri" :"mongodb://host.docker.internal:27017"
}
````



### Functionality folder structure

- ```domain/```
  - ```useCases/```  ( all uses cases )
  - ```entities/``` ( all business entities )
- ```interfaces/```
  - ```controllers/``` ( all controllers )
  - ```inputs/``` ( all inputs handlers )
  - ```gateways/```  ( all repository layers *optional )
  - ```outputs/``` ( all outputs methods )

## Kafka Example
- ..soon
## Mongo Task Crud Example 
### File structure mongo module files

- infrastructure 
  - database 
    - mongo 
      - client   ( Client of Mongo )
      - drivers
        - models     ( models for mongo with validations )
        - repository ( Repository layer for Mongo )

### Pre-requisites

- Mongo Docker Container UP with Local Default port

```shell
docker run --name some-mongo -d mongo:tag
```

- Port 8888 available in local environment or change this in config file and curl methods

### Run Locally

```shell
go mod tidy
go run main.go
```

### Run with docker

```shell
docker build -t golang-skeleton . 
docker run -p 8888:8888 golang-skeleton 
```

#### For run use case please execute this command  in the terminal or import in your postman collection

- Add task
``` shell
 
curl --location --request POST 'http://localhost:8888/task' \
--header 'Content-Type: application/json' \
--data-raw '{
"title" : "task title",
"author" : "task author",
"tags" : [ "tag test ","tag test"]
}' 
```
- Get task by title

```shell
curl --location --request GET 'http://localhost:8888/task/task%20title'
```
- Find all task

```shell
curl --location --request GET 'http://localhost:8888/task'
```

- Delete task by title
```shell
curl --location --request DELETE 'http://localhost:8888/task/task%20title'
```