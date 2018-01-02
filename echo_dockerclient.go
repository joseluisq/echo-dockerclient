// Tiny Go Docker Client middleware for Echo.
// This middleware stores a Docker Client instance 
// on `echo.Context` via `echo.Context.Set("docker-client", *docker.Client)`
// 
// Usage:
// 	e := echo.New()
// 	e.Use(m.DockerClientFromEnv())
// 
// Get the client instance in controller:
// 	client := context.Get("docker-client").(*docker.Client)
// 

package dockerc

import (
	"github.com/fsouza/go-dockerclient"
	"github.com/labstack/echo"
)

const contextDataKey = "docker-client"

// DockerClientFromEnv middleware binds a Docker Client instance
// from Docker's default logic for the environment variables:
// DOCKER_HOST, DOCKER_TLS_VERIFY and DOCKER_CERT_PATH.
func DockerClientFromEnv() echo.MiddlewareFunc {
	return storeDockerClient(docker.NewClientFromEnv())
}

// DockerClient middleware binds a Docker Client instance
// with the given server endpoint. It will use the latest
// remote API version available in the server.
func DockerClient(endpoint string) echo.MiddlewareFunc {
	return storeDockerClient(docker.NewClient(endpoint))
}

// storeDockerClient stores a Docker Client instance on echo.Context
func storeDockerClient(client *docker.Client, err error) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err != nil {
				c.Error(err)
			} else {
				c.Set(contextDataKey, client)
			}

			return next(c)
		}
	}

}
