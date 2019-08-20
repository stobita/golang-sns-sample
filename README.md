# golang-sns-sample

Golang REST API like SNS application

## Overview

### Tools

#### Development Environment
* Docker
* DockerCompose

#### RDB
* MySQL

### Libraries

#### http
* https://github.com/gin-gonic/gin

## Operations

### start api server

```
make dev-up
```

### create migration file

```
make migrate-create NAME=user
```

### migration up

```
make migrate up
```
