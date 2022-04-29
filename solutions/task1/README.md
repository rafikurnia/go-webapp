# go-webapp

[![Go](https://github.com/rafikurnia/go-webapp/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/rafikurnia/go-webapp/actions/workflows/go.yml)

A simple Hello World web application written in Go and uses [gin-gonic](https://gin-gonic.com/) web framework.


## Prerequisites

- Go 1.18.1 (Optional, only if you want to run the application on local machine or perform unit test)
- Docker 19.03.0+


## Getting Started

After you have all prerequisites install, to run the application, you can follow these steps:

1. Run the application using docker-compose via makefile:
   ```bash
   make compose-up
   ```

2. The application will run on the container and listen on port `8080`, and it is mapped to your
   local machine on port 80. Thus, make sure that the port is not occupied.

   After the application is running, you can access it using your web browser on the following link:
   ```bash
   http://localhost/
   ```

   Or, using curl:
   ```bash
   curl -v http://localhost/
   ```

3. When you are finished using the application, you can shut down and remove all resources using the
   following command:
   ```bash
   make compose-down
   ```


## Running unit test locally

To run the test locally on your machine, you can invoke the following command:
```bash
make go-test
```


## Configure Application Parameters

You can update some parameters of the application inside [./config.env](./config.env) file. Table
below explains each parameter in the file:

| Name                          | Description                                                                               |
| ----------------------------- | ----------------------------------------------------------------------------------------- |
| APP_NAME                      | The name of the application. It will be used as the docker image and container name.      |
| VERSION                       | The version of the application. It will be used as a tag for the docker image.            |
| APP_PORT                      | The port number where the application will listen to incoming connections.                |
| HOST_PORT                     | The port number to access the application from host machine.                              |
| GIN_MODE                      | to inform gin-gonic web framework about the environment where the application is running. |
| CPU_LIMIT                     | a maximum number of CPU that the container can use.                                       |
| MEM_LIMIT                     | a maximum amount of memory that the container can use.                                    |
| CPU_RESERVED                  | a number of CPU that is reserved for the container to use.                                |
| MEM_RESERVED                  | an amount of memory that is reserved for container to use.                                |
| NUMBER_OF_REQUESTS            | a number of requests to perform on the benchmarking using ApacheBench.                    |
| NUMBER_OF_CONCURRENT_REQUESTS | A number of multiple requests to make at a time on the benchmarking using ApacheBench.    |


## Makefile help

This project is shipped with Makefile to help on automating the a chain of executions of certain
actions. Table below shows a list of available makefile targets and their descriptoin.

| Target                        | Description                                                              |
| ----------------------------- | ------------------------------------------------------------------------ |
| help                          | Show help menu                                                           |
| go-build                      | Build Go code                                                            |
| go-test                       | Test Go code                                                             |
| go-run                        | Run Go code                                                              |
| go-rm                         | Delete the app binary file                                               |
| docker-build                  | Build an image from a Dockerfile with tag APP_NAME:VERSION               |
| docker-build-ne               | Build the image only if it does not exist                                |
| docker-run                    | Build the image if it does not exists and start container from the image |
| docker-rm                     | Delete the image                                                         |
| compose-up                    | Build the image if not exist and start containers using docker-compose   |
| compose-down                  | Stop and remove containers and networks created by docker-compose        |
| benchmark                     | Run the containers and perform benchmark using apachebench               |
| clean                         | Remove all resources possibly made by this makefile                      |

You can invoke any of the command above by calling `make <target>`.


## Author

[Rafi Kurnia Putra](https://github.com/rafikurnia)


## License

MIT License. See [LICENSE](./LICENSE) for full details.
