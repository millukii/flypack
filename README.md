### Flypack


### Instrucciones

#### Docker
- docker build -t flypack .
- docker run --rm -it -p 8080:8080 flypack 
#### Docker-Compose
- docker-compose up [--build]

#### Local

- Instalar nodemon en forma local con npm install -g nodemon
- Ejecutar con nodemon --exec go run cmd/flypack/main.go --signal SIGTERM