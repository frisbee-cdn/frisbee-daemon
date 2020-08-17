# Frisbee Daemon

## Description 

This service defines the Lookup Index Layer an marks the core component of the system. 

It's main responsibility is to build and organize the clustered overlay network based on the Kademlia Algorithm.

## Design

TBD

## Local Development Environment

At this moment the only way to start the network is by manually run the main.go file along with these flags:

- host the hostname gor the peer.
- port the port number for the peer.
- boot ( optional ) this flag should be only specified for peers which will play the role as bootstrap peers.

### Non-Docker

In order to locally run the daemon you first need to start the bootstrap peer.
```
go run main.go -host=<hostname> -port=<port_number> -boot

(e.g)

go run main.go -host=RomaniaHostTimisCounty1Boot -port=5001 -boot
```

After this you can start other peers to join your network.

```

go run main.go -host=<hostname> -port=<port_number>

(e.g)

go run main.go -host=RomaniaHostAradCounty1Boot -port=5002
```

Please note that you will need multiple shell instances opened for each peer.


### Docker

TBD

