# echo-dockerclient [![Build Status](https://travis-ci.org/joseluisq/echo-dockerclient.svg?branch=master)](https://travis-ci.org/joseluisq/echo-dockerclient)

> Tiny [Go Docker Client](https://github.com/fsouza/go-dockerclient) middleware for [Echo](echo.labstack.com).

## Install

```sh
go get github.com/joseluisq/echo-dockerclient
```

Or via [Dep](https://golang.github.io/dep/)

```sh
dep ensure -add github.com/joseluisq/echo-dockerclient
```

## Usage

```go
package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "github.com/joseluisq/echo-dockerclient"
)

func main() {
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // 1. First, add the Docker Client middleware to Echo server

    // Option 1:
    // Docker Client instance with the given server endpoint
    e.Use(dockerc.DockerClient("unix:///var/run/docker.sock"))

    // Option 2:
    // Or if you're using `docker-machine` or another application that exports 
    // the Docker environment variables like DOCKER_HOST, DOCKER_TLS_VERIFY, DOCKER_CERT_PATH 
    e.Use(dockerc.DockerClientFromEnv())

    // 2. Finally, get the Docker Client instance in controller
    e.GET("/images", func (c echo.Context) error {
        client := c.Get("docker-client").(*docker.Client)
        images, err := client.ListImages(docker.ListImagesOptions{All: false})
        return c.JSON(200, images)
    })
}
```

For more details, check out the [echo-dockerclient](https://github.com/fsouza/go-dockerclient) package and the [Docker Engine API](https://docs.docker.com/develop/sdk/) documentation. 

## Contributions

[Pull requests](https://github.com/joseluisq/redel/pulls) and [issues](https://github.com/joseluisq/redel/issues) are very appreciated.

## License
MIT license

© 2018 [José Luis Quintana](https://git.io/joseluisq)
