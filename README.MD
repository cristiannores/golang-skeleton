# API GOLANG SKELETON

## Kafka Example
- ..soon
## Mongo Example 
### File structure mongo files

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

### Run 

```shell
go mod tidy
go run main.go
```

#### For test use case please execute this command  in the terminal or import in your postman collection

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