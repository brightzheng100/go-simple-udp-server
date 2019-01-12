# Simple UDP Server Written By Go

It's just a simple demo app to demonstrate how to write a simple UDP Server by using Golang.

## Get Started

### Start It Up

```
$ go run *.go
UDP Server listening on: :8888

```

Or specify the `-address`:
```
$ go run *.go -address=0.0.0.0:5678
UDP Server listening on: 0.0.0.0:5678

```

### Play With It

Assuming the UDP Server is listening on local machine's `0.0.0.0:5678`):
```
$ go run *.go -address=0.0.0.0:5678
UDP Server listening on: 0.0.0.0:5678

```

Open another console to play with it.

UDP Testing Client Console, we can simply use `nc` to test it.
We can key in any string and then hit `Enter`, in client side it will echo back and print exactly what we keyed in.
For example, if we key in `A` then `Enter`, and then `B` then `Enter`, we shall see something like:
```
$ nc -u 127.0.0.1 5678
A
A
B
B

```

If you look at the Console for UDP Server, you shall see some logs like:
```
$ go run *.go -address=0.0.0.0:5678
UDP Server listening on: 0.0.0.0:5678
Packet Received: [2] bytes from [127.0.0.1:51959]
Packet Written: [2] bytes to [127.0.0.1:51959]
Packet Received: [2] bytes from [127.0.0.1:51959]
Packet Written: [2] bytes to [127.0.0.1:51959]

```

### Dockerize It

1. Build

```
$ sed 's/${PORT}/5678/g' Dockerfile > Dockerfile.build \
    && docker build --rm -t itstarting/go-simple-udp-server -f Dockerfile.build . \
    && rm Dockerfile.build
```

> Note: as there is a known issue that `ARG` and `ENDPOINT` don't work well (see [here](https://github.com/moby/moby/issues/18492)), have to do some trick like above.


2. Check

```
$ docker images
REPOSITORY                                      TAG                 IMAGE ID            CREATED             SIZE
itstarting/go-simple-udp-server                 latest              5d43f270df0c        6 minutes ago       2.07MB
```

> Note: yeah, it's just ~2MB!


3. Run

```
$ docker run -p 5678:5678/udp itstarting/go-simple-udp-server
```

## Ref

- Good read (actually the code was borrowed from it): https://ops.tips/blog/udp-client-and-server-in-go/