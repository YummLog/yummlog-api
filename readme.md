# yummlog-api

Exposes REST API for the yummlog-interface

## REST API documentation

The API specs are documented in `openapi.yaml`

## Running the api as docker image

### Building the image
```
docker build --tag yummlog-api .
```

### Running the image
```
docker run --rm -p 3000:3000 yummlog-api
```

