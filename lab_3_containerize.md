# Containerize the application

In this section, you will containerize the application using Docker. Docker is a platform that enables developers to develop, ship, and run applications in containers. Containers allow developers to package up an application with all of the parts it needs, such as libraries and other dependencies, and ship it all out as one package.

## Instructions

1. MUST create a Dockerfile for the Go application API
1. MUST create a Dockerfile for the consumer application
1. MUST use multi-stage builds for the Go application API and consumer application
1. MUST specific Go version to minimum *new latest version* e.g. `v1.22.0` atleast.
1. use distroless image or scratch image for the final image.
1. SHOULD use the `docker-compose` to run the application and use Dockerfile for the build stage.
1. MUST create command push the image in makefile to push the image to the registry.
for github use `ghcr.io/<username>/<repository>:<tag>` as the image name. for gitlab use `registry.gitlab.com/<username>/<repository>:<tag>`

## Expected Output

- Dockerfile for the Go application API
- Dockerfile for the consumer application
- docker-compose.yml
- Makefile with the following commands:
	- `make build` to build the application
	- `make run` to run the application
	- `make push` to push the image to the registry
- the registry url for the image.