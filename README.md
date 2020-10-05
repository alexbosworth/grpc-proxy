# gRPC Proxy

Proxy service to relay requests to a remote gRPC endpoint using REST

## Configuration

Environment Variables:

```ini
# Address for gRPC server
REMOTE_SERVER_SOCKET="host:port"

# Desired PORT (optional: defaults to 8080)
PORT="service_port"

# Full path to the TLS public cert file (optional: defaults to no-TLS)
TLS_CERT_FILE_PATH="path/to/cert/file"

# Full path to the TLS private key file (optional: defaults to no-TLS)
TLS_KEY_FILE_PATH="path/to/key/file"
```

## Docker

Build the image:

```shell
docker build -t alexbosworth/grpc-proxy .
```

Publish the image:

```shell
docker push alexbosworth/grpc-proxy
```

Run the image:

```shell
# Run interactively
docker run -it -e REMOTE_SERVER_SOCKET=grpc_server_socket -p 8080:8080 alexbosworth/grpc-proxy

# To run with TLS, set TLS_CERT_FILE_PATH and TLS_KEY_FILE_PATH env vars and
# use -v to give access to the cert files
```
