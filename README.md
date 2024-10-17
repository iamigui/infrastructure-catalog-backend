# Infrastructure Catalog Backend

## Install Packages

```Go
go mod install
```

## Run the application

```shell
docker-compose up -d
```

## Monitoring URLs

- Grafana: http://localhost:3000/
  - username: admin
  - password: admin
- Prometheus: http://localhost:9090/
- Mongo-Express: http://localhost:8081/
  - username: admin
  - password: admin
- Jaeger: http://localhost:16686/

## Endpoints

### Get Projects

```shell
curl http://localhost:8000/getProjects
```

### Post a Project

```shell
curl -X POST http://localhost:8000/createProject \
-H "Content-Type: application/json" \
-d '{
  "name": "New Infrastructure Project",
  "description": "This project involves building a new data center.",
  "jsonData": {
    "location": "San Francisco, CA"
  }
}'
```

### Get Project by ID

```shell
curl http://localhost:8000/getProject?id=123
```
