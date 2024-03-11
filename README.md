## Purpose

This simple App generates random numbers on a TCP connection.

## How to run it

1. Run "go run main.go" in your machine.

2. Or run it as docker - "run --publish 17777:17777 tcp-random-app"
(you can build it or get dockerfile from "docker pull gitmarut/tcp-random-app:v1.1")

3. Or you can deploy it in Kubernetes from the yaml file given in same directory.

For methods 1 & 2, from another terminal, use nc to connect it.
"nc localhost 17777"
On each press of return key(enter) at client, you get a random number.
It will ignore any text entered except "STOP". It will gracefully terminate connection of entering "STOP"
On pressing ctrl+C at client, it will error EOF and close connection to client.

For method 3 - EDITNEEDED

