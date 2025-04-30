# Hello

A simple client & server that communicate over gRPC.

## Protocol Buffers in Buf Schema Registry

```bash
buf lint proto
buf push proto
```

Now they are found at https://buf.build/wcygan/hello

## Running with Docker

To run both the client and server using Docker Compose:

```bash
# Build and start the containers in detached mode
docker compose up --build -d
```

Once the containers are running:

- The **Go server** will be accessible on port `8080` (e.g., for `grpcurl` or `curl` commands).
- The **Deno client** web interface will be accessible in your browser at [http://localhost:8000](http://localhost:8000).

To stop the containers:

```bash
docker compose down
```