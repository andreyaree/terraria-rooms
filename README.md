# WARNING: Still in development! Not finished yet

## What's it?
Terraria-Rooms is a tool for managing, connecting multiple terraria game servers through a single proxy.

## Features
The program provides the following features:
- TCP packet forwarding
- Metrics
  - Active connections
  - Total connections
  - Received
  - Sent
- Lightweight
- Blacklist
- Easy-to-configure
- Global Chat (*Soon*!)

## Build
Install the Go programming language and run the following command in the project root directory:
```bash 
go build ./cmd/terraria-rooms
```

## Run
```bash
./terraria-rooms
```
The program looks for a configuration file next to the binary and creates a default one if it doesn't exist. 

