### Project: Flypack

#### To Do List

- Add a properly logger
- Add a ci-cd pipeline



#### Progress
| Domain  | Postman Collection | Handler | Service | Repository |
| ------------- | ------------- | ------------- | ------------- | ------------- |
| Users | In Progress  | In Progress | In Progress | In Progress|
| People  | In Progress  | In Progress | In Progress | In Progress|
| Roles  | In Progress  | In Progress | In Progress | In Progress|
| Companies  | In Progress  | In Progress | In Progress | In Progress|
| Shippings  | In Progress  | In Progress | In Progress | In Progress|

### Data Provider

| Method  | name | 	API call | 
| ------------- | ------------- | ------------- | 
| getList | GET  | http://localhost:8080/users?sort=["title","ASC"]&range=[0, 24]&filter={"title":"bar"} |
| getOne | GET  | http://localhost:8080/users/123 |
| getMany | GET  | http://localhost:8080/users?filter={"ids":[123,456,789]} |
| getManyReference | GET  | http://localhost:8080/users?filter={"author_id":345} |
| create | POST  | http://localhost:8080/users |
| update | PUT  | http://localhost:8080/users/123 |
| updateMany | Multiple calls to PUT  | http://localhost:8080/users/123 |
| delete | DELETE  | http://localhost:8080/users/123 |
| deleteMany | 	Multiple calls to DELETE  | http://localhost:8080/users/123 |


### Requirements
- go
- docker
- mysql

#### Docker
- docker build -t flypack .
- docker run --rm -it -p 8080:8080 flypack 
#### Docker-Compose
- docker-compose up [--build]

#### Local

- Install nodemon npm install -g nodemon
- Execute with hot reload --exec go run cmd/flypack/main.go --signal SIGTERM