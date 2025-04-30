# Server

## Deploy

```bash
go run cmd/main.go
```

## Docker

Build the Docker image (run from the `server` directory):

```bash
docker build -t hello-server .
```

Run the container:

```bash
docker run -p 8080:8080 hello-server
```

## Testing

Request:

```bash
grpcurl \
  -plaintext \
  -d '{"name": "Foo"}' \     
  localhost:8080 \
  hello.v1.GreeterService/SayHello
```

Response:

```json
{
  "message": "Hello Foo!"
}
```

You can also do it directly with cURL:

```bash
curl \
  --header "Content-Type: application/json" \
  --header "Connect-Protocol-Version: 1" \
  --data '{"name": "Foo"}' \
  --request POST \
  http://localhost:8080/hello.v1.GreeterService/SayHello
```