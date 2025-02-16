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

On main.go client files, you can uncomment or comment the lines about the example you want to run.

## Using Evans CLI

First of all, you need to install the Evans CLI, you can download the tar.gz file [here](https://github.com/ktr0731/evans):

```bash
tar -xf evans_file_name.tar.gz
```

You need to move the `evans` to the `/usr/local/bin` directory.

```bash
sudo mv evans /usr/local/bin
```

After that, you need run calculator server:

```bash
make calculator && ./bin/calculator/server
```

You can use the Evans CLI to interact with the gRPC server. To do this, you need to run the following command:

```bash
evans -r repl --host localhost --port 50051
```

You need to select the `calculator.proto` file and the `calculator` service. For that, you need to run the following command on the Evans CLI:

```bash
package calculator
```

You need to select the service:

```bash
service CalculatorService
```

You can call the `Add` method:

```bash
call Add
```

After that, the Evans CLI will ask you to enter the `first_number` and `second_number` values. You can enter the values and press `Enter`.

When you finish, you can exit the Evans CLI:

```bash
exit
```

You can do the same for all services and methods.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.
