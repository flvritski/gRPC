# Read me

This repo makes use of a `docker-compose.yml` file which consists in 2 services:

- mongo
- mongo-express

There are multiple gRPC endpoints created, for 3 different projects:

- greet
- calculator
- blog

All of the gRPC endpoints are intedended to make an example of the way that gRPC concept has underneath.

We have 3 main types:

- unary (client request->server response)
- server streaming (multiple requests from the client -> stream response from the server)
- client streaming (stream client request -> multiple responses from the server)
- bi-directional streaming (stream request -> stream response)

We also make sure that the connection is secured, using SSL certs.
You can make use of the `ssl.sh` script that's inside the `ssl/` folder. This will generate for you

### ca.key: Certificate Authority private key file (this shouldn't be shared in real-life)

### ca.crt: Certificate Authority trust certificate (this should be shared with users in real-life)

### server.key: Server private key, password protected (this shouldn't be shared)

### server.csr: Server certificate signing request (this should be shared with the CA owner)

### server.crt: Server certificate signed by the CA (this would be sent back by the CA owner) - keep on server

### server.pem: Conversion of server.key into a format gRPC likes (this shouldn't be shared)

## How to run the blog gRPC?

Open 3 zsh terminals inside the root folder \_/>

1. (1st terminal) In the root folder, run `$ make blog`
2. (1st terminal) Run `$ ./bin/blog/server`
3. (2nd terminal) Run `$ cd blog/ && docker-compose up`
4. (3rd terminal) Run `$ ./bin/blog/client`

### Also, in case of any erros (missing packages or etc), you can also run `go mod tidy`.
