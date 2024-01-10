# marqeta-client

## How to generate Go Client Code

### fern

Link: https://github.com/fern-api/fern?tab=readme-ov-file#starting-from-openapi

Install:
```shell
npm install -g fern-api
```

Init:
```shell
fern init --openapi yaml/CoreAPI.yaml
```

which prints
```
? Login required. Continue? Yes
Logged in!
Created organization livefire-2015
Created new API: ./fern
```

Generate:
```shell
fern generate
```

### openapi-generator

Link: https://openapi-generator.tech/docs/configuration/

Install:
```shell
brew install openapi-generator
```

Generate:
```shell
openapi-generator generate -i https://raw.githubusercontent.com/marqeta/marqeta-openapi/main/yaml/CoreAPI.yaml  -g go -o marqeta-coreapi-go-client
```

or, with more options

```shell
openapi-generator generate -g go -i CoreAPI.yaml -o marqeta-coreapi-go-client-v3-0-11 --additional-properties=packageName=mqt-core-go,withGoMod=false
```

or, using a config.json file

```shell
openapi-generator generate -g go -i yaml/CoreAPI.yaml -o coreapi -c config.json
```

### openapi-generator-cli

Link: https://github.com/OpenAPITools/openapi-generator-cli#installation

Install:
```shell
npm install -g @openapitools/openapi-generator-cli
```

Generate
```shell
openapi-generator-cli generate -g go -i CoreAPI.yaml -o marqeta-coreapi-go-client-v3-0-11
```

### go-swagger

Link: https://github.com/go-swagger/go-swagger

It does NOT work for OAS3.0

## How to invoke mqtsvc

See package `mqtsvc`.
