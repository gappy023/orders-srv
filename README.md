# Order Service

This is the Order service

Generated with

```
micro new microservice/part1/orders-srv --namespace=mu.micro.book --alias=order --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: mu.micro.book.srv.order
- Type: srv
- Alias: order

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./order-srv
```

Build a docker image
```
make docker
```