# grpc

Implementing simple examples of gRPC in Golang.

## Installation

First, you need to install the protocol buffer compiler (protoc) and the Go plugin for the protocol buffer compiler, you can follow the instructions [here](https://grpc.io/docs/protoc-installation/).

## Configuration

### Generate the gRPC code

To generate the gRPC code, you need to run the following command:

```bash
make greet && make calculator
```

### Generate SSL certificates

To generate the SSL certificates, you need to run the following command:

```bash
cd ssl && chmod +x ssl.sh && ./ssl.sh
```

You must configure the environment variables in the `.env` file. Example:

```.env
ENABLE_TLS=true
CERT_FILE_PATH="ssl/server.crt"
KEY_FILE_PATH="ssl/server.pem"
CERT_CLIENT_FILE_PATH="ssl/ca.crt"
```

If you want to disable TLS, you must set the `ENABLE_TLS` variable to `false`.

## Usage

### Run the server

To run the server, you need to run the following command:

```bash
./bin/greet/server
```

You will do the same for the calculator server if you want to run it:

```bash
./bin/calculator/server
```

### Run the client

To run the client, you need to run the following command:

```bash
./bin/greet/client
```

You will do the same for the calculator client if you want to run it:

```bash
./bin/calculator/client
```

## Attention

On main.go client files, you can uncomment or comment the lines about the example you want to run.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.
