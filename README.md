# 🌐 httpserve

> A simple HTTP server written in Go that reads data from standard input (stdin) and serves it over HTTP.
> It allows users to transmit data through HTTP requests by piping data into the server via stdin.

## Usage

```console
httpserve -p <PORT> -a <ADDRESS> -f <FILENAME> -c <CONTENT_TYPE>
```

**Options**

- `-p <PORT>`: Set the port for the server (default: 8080)
- `-a <ADDRESS>`: Set the address for the server (default: 0.0.0.0)
- `-f <FILENAME>`: Set the filename header for downloaded content
- `-c <CONTENT_TYPE>`: Set the content-type header (default: application/octet-stream)
- `-user <USERNAME>`: The user that will be used in authentication prompts
- `-password <PASSWORD>`: The password required to access served data

### Examples

1. Accessing served data without filename or content type

```bash
curl http://localhost:8080
```

Retrieves the served data from the default address, which is **`http://localhost:8080`**.

2. Downloading served data with specified filename and content type

```bash
curl -o output.txt -H "Content-Type: text/plain" http://localhost:8080
```

- `-o output.txt`: Saves the served data to a file named output.txt.
- `-H "Content-Type: text/plain"`: Specifies the content type of the requested data.

## Why?

Honestly, why not. Httpserve is designed to receive data from standard input and serve it over HTTP, offering a
pretty simple and flexible method for sharing and transmitting data across network endpoints.
