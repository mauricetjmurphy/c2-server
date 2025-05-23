# C2 Server Project

## Overview

The implant is a lightweight agent that is deployed on target devices, which communicates with the listener using HTTP and JSON.

## Features

- Secure communication between the server and implants.
- Lightweight implant designed to minimize resource usage on target devices.
- Configurable beaconing interval and task execution.
- Task management and execution, including shell command execution.

## Prerequisites

- Docker

## Setup

### Clone the Repository

```
git clone https://github.com/yourusername/c2-server-project.git
cd c2-server-project
```

### Build the Project

Use the provided Makefile to build the binaries for the implant and listener.

```
make build
```


For cross-compiling to different operating systems:

```
make build_implant GOOS=windows GOARCH=amd64
make build_listener GOOS=windows GOARCH=amd64
```


### Configuration

Edit the configuration file located at config/config.json:

```
{
  "host": "http://127.0.0.1",
  "port": "4000",
  "uri": "/tasks",
  "timeout": 10,
  "interval": 60
}
```


- `host`: The host address of the listener service.
- `port`: The port on which the listener service runs.
- `uri`: The endpoint for tasks.
- `timeout`: Timeout for HTTP requests (in seconds).
- `interval`: Interval between beaconing requests (in seconds).